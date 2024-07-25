/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package servicediscovery

import (
	common_meta "dubbo.apache.org/dubbo-go/v3/common/metadata"
	"dubbo.apache.org/dubbo-go/v3/metadata/service/local"
	"dubbo.apache.org/dubbo-go/v3/metadata/service/remote"
	"encoding/gob"
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"time"
)

import (
	gxset "github.com/dubbogo/gost/container/set"
	"github.com/dubbogo/gost/gof/observer"
	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"dubbo.apache.org/dubbo-go/v3/registry/servicediscovery/store"
	"dubbo.apache.org/dubbo-go/v3/remoting"
)

var (
	metaCache *store.CacheManager
	cacheOnce sync.Once
)

func initCache(app string) {
	gob.Register(&common.MetadataInfo{})
	fileName := constant.DefaultMetaFileName + app
	cache, err := store.NewCacheManager(constant.DefaultMetaCacheName, fileName, time.Minute*10, constant.DefaultEntrySize, true)
	if err != nil {
		logger.Fatal("Failed to create cache [%s],the err is %v", constant.DefaultMetaCacheName, err)
	}
	metaCache = cache
}

// ServiceInstancesChangedListenerImpl The Service Discovery Changed  Event Listener
type ServiceInstancesChangedListenerImpl struct {
	app                string
	serviceNames       *gxset.HashSet
	listeners          map[string]registry.NotifyListener
	serviceUrls        map[string][]*common.URL
	revisionToMetadata map[string]*common.MetadataInfo
	allInstances       map[string][]registry.ServiceInstance
	mutex              sync.Mutex
}

func NewServiceInstancesChangedListener(app string, services *gxset.HashSet) registry.ServiceInstancesChangedListener {
	cacheOnce.Do(func() {
		initCache(app)
	})
	return &ServiceInstancesChangedListenerImpl{
		app:                app,
		serviceNames:       services,
		listeners:          make(map[string]registry.NotifyListener),
		serviceUrls:        make(map[string][]*common.URL),
		revisionToMetadata: make(map[string]*common.MetadataInfo),
		allInstances:       make(map[string][]registry.ServiceInstance),
	}
}

// OnEvent on ServiceInstancesChangedEvent the service instances change event
func (lstn *ServiceInstancesChangedListenerImpl) OnEvent(e observer.Event) error {
	ce, ok := e.(*registry.ServiceInstancesChangedEvent)
	if !ok {
		return nil
	}

	lstn.mutex.Lock()
	defer lstn.mutex.Unlock()

	lstn.allInstances[ce.ServiceName] = ce.Instances
	revisionToInstances := make(map[string][]registry.ServiceInstance)
	newRevisionToMetadata := make(map[string]*common.MetadataInfo)
	localServiceToRevisions := make(map[*common.ServiceInfo]*gxset.HashSet)
	protocolRevisionsToUrls := make(map[string]map[*gxset.HashSet][]*common.URL)
	newServiceURLs := make(map[string][]*common.URL)

	logger.Infof("Received instance notification event of service %s, instance list size %d", ce.ServiceName, len(ce.Instances))

	for _, instances := range lstn.allInstances {
		for _, instance := range instances {
			if instance.GetMetadata() == nil {
				logger.Warnf("Instance metadata is nil: %s", instance.GetHost())
				continue
			}
			revision := instance.GetMetadata()[constant.ExportedServicesRevisionPropertyName]
			if "0" == revision {
				logger.Infof("Find instance without valid service metadata: %s", instance.GetHost())
				continue
			}
			subInstances := revisionToInstances[revision]
			if subInstances == nil {
				subInstances = make([]registry.ServiceInstance, 0, 8)
			}
			revisionToInstances[revision] = append(subInstances, instance)
			metadataInfo := lstn.revisionToMetadata[revision]
			if metadataInfo == nil {
				metadataInfo = GetMetadataInfo(lstn.app, instance, revision)
			}
			instance.SetServiceMetadata(metadataInfo)
			for _, service := range metadataInfo.Services {
				if localServiceToRevisions[service] == nil {
					localServiceToRevisions[service] = gxset.NewSet()
				}
				localServiceToRevisions[service].Add(revision)
			}

			newRevisionToMetadata[revision] = metadataInfo
		}
		lstn.revisionToMetadata = newRevisionToMetadata
		for revision, metadataInfo := range newRevisionToMetadata {
			metaCache.Set(revision, metadataInfo)
		}

		for serviceInfo, revisions := range localServiceToRevisions {
			revisionsToUrls := protocolRevisionsToUrls[serviceInfo.Protocol]
			if revisionsToUrls == nil {
				protocolRevisionsToUrls[serviceInfo.Protocol] = make(map[*gxset.HashSet][]*common.URL)
				revisionsToUrls = protocolRevisionsToUrls[serviceInfo.Protocol]
			}
			urls := revisionsToUrls[revisions]
			if urls != nil {
				newServiceURLs[serviceInfo.GetMatchKey()] = urls
			} else {
				urls = make([]*common.URL, 0, 8)
				for _, v := range revisions.Values() {
					r := v.(string)
					for _, i := range revisionToInstances[r] {
						if i != nil {
							urls = append(urls, i.ToURLs(serviceInfo)...)
						}
					}
				}
				revisionsToUrls[revisions] = urls
				newServiceURLs[serviceInfo.GetMatchKey()] = urls
			}
		}
	}

	lstn.serviceUrls = newServiceURLs
	for key, notifyListener := range lstn.listeners {
		urls := lstn.serviceUrls[key]
		events := make([]*registry.ServiceEvent, 0, len(urls))
		for _, url := range urls {
			events = append(events, &registry.ServiceEvent{
				Action:  remoting.EventTypeAdd,
				Service: url,
			})
		}
		notifyListener.NotifyAll(events, func() {})
	}

	return nil
}

// AddListenerAndNotify add notify listener and notify to listen service event
func (lstn *ServiceInstancesChangedListenerImpl) AddListenerAndNotify(serviceKey string, notify registry.NotifyListener) {
	lstn.listeners[serviceKey] = notify
	urls := lstn.serviceUrls[serviceKey]
	for _, url := range urls {
		notify.Notify(&registry.ServiceEvent{
			Action:  remoting.EventTypeAdd,
			Service: url,
		})
	}
}

// RemoveListener remove notify listener
func (lstn *ServiceInstancesChangedListenerImpl) RemoveListener(serviceKey string) {
	delete(lstn.listeners, serviceKey)
}

// GetServiceNames return all listener service names
func (lstn *ServiceInstancesChangedListenerImpl) GetServiceNames() *gxset.HashSet {
	return lstn.serviceNames
}

// Accept return true if the name is the same
func (lstn *ServiceInstancesChangedListenerImpl) Accept(e observer.Event) bool {
	if ce, ok := e.(*registry.ServiceInstancesChangedEvent); ok {
		return lstn.serviceNames.Contains(ce.ServiceName)
	}
	return false
}

// GetPriority returns -1, it will be the first invoked listener
func (lstn *ServiceInstancesChangedListenerImpl) GetPriority() int {
	return -1
}

// GetEventType returns ServiceInstancesChangedEvent
func (lstn *ServiceInstancesChangedListenerImpl) GetEventType() reflect.Type {
	return reflect.TypeOf(&registry.ServiceInstancesChangedEvent{})
}

// GetMetadataInfo get metadata info when MetadataStorageTypePropertyName is null
func GetMetadataInfo(app string, instance registry.ServiceInstance, revision string) *common.MetadataInfo {
	cacheOnce.Do(func() {
		initCache(app)
	})
	if metadataInfo, ok := metaCache.Get(revision); ok {
		return metadataInfo.(*common.MetadataInfo)
	}

	var metadataStorageType string
	metadataInfo := common.EmptyMetadataInfo
	var err error
	if instance.GetMetadata() == nil {
		metadataStorageType = constant.DefaultMetadataStorageType
	} else {
		metadataStorageType = instance.GetMetadata()[constant.MetadataStorageTypePropertyName]
	}
	if metadataStorageType == constant.RemoteMetadataStorageType {
		remoteMetadataServiceImpl, remoteMetadataErr := remote.GetRemoteMetadataService()
		if remoteMetadataErr == nil {
			revision := instance.GetMetadata()[constant.ExportedServicesRevisionPropertyName]
			metadataInfo, err = remoteMetadataServiceImpl.GetMetadata(instance.GetServiceName(), revision)
		} else {
			err = remoteMetadataErr
		}
	} else {
		proxyFactory := local.NewBaseMetadataServiceProxyFactory()
		metadataService := proxyFactory.GetProxy(buildStandardMetadataServiceURL(instance))
		if metadataService != nil {
			defer destroyInvoker(metadataService)
			metadataInfo, err = metadataService.GetMetadataInfo(revision)
		} else {
			err = errors.New("get remote metadata error please check instance " + instance.GetHost() + " is alive")
		}
	}

	if err != nil {
		logger.Errorf("get metadata of %s failed, %v", instance.GetHost(), err)
	}

	if metadataInfo != common.EmptyMetadataInfo {
		metaCache.Set(revision, metadataInfo)
	}

	return metadataInfo
}

func destroyInvoker(metadataService common_meta.MetadataService) {
	if metadataService == nil {
		return
	}

	proxy := metadataService.(*registry.MetadataServiceProxy)
	if proxy.Invoker == nil {
		return
	}

	proxy.Invoker.Destroy()
}

// buildStandardMetadataServiceURL will use standard format to build the metadata service url.
func buildStandardMetadataServiceURL(ins registry.ServiceInstance) []*common.URL {
	ps := getMetadataServiceUrlParams(ins)
	if ps[constant.ProtocolKey] == "" {
		return nil
	}

	res := make([]*common.URL, 0, len(ps))
	sn := ins.GetServiceName()
	host := ins.GetHost()
	convertedParams := make(map[string][]string, len(ps))
	for k, v := range ps {
		convertedParams[k] = []string{v}
	}
	u := common.NewURLWithOptions(common.WithIp(host),
		common.WithPath(constant.MetadataServiceName),
		common.WithProtocol(ps[constant.ProtocolKey]),
		common.WithPort(ps[constant.PortKey]),
		common.WithParams(convertedParams),
		common.WithParamsValue(constant.GroupKey, sn),
		common.WithParamsValue(constant.InterfaceKey, constant.MetadataServiceName))

	if ps[constant.ProtocolKey] == "tri" {
		u.SetAttribute(constant.ClientInfoKey, "info")
		u.Methods = []string{"getMetadataInfo"}

		metaV := ins.GetMetadata()[constant.MetadataVersion]
		if metaV != "" {
			u.Path = constant.MetadataServiceV2Name
			u.SetParam(constant.VersionKey, "2.0.0")
		} else {
			u.SetParam(constant.SerializationKey, constant.Hessian2Serialization)
		}
	}

	res = append(res, u)

	return res
}

// getMetadataServiceUrlParams this will convert the metadata service url parameters to map structure
// it looks like:
// {"dubbo":{"timeout":"10000","version":"1.0.0","dubbo":"2.0.2","release":"2.7.6","port":"20880"}}
func getMetadataServiceUrlParams(ins registry.ServiceInstance) map[string]string {
	ps := ins.GetMetadata()
	res := make(map[string]string, 2)
	if str, ok := ps[constant.MetadataServiceURLParamsPropertyName]; ok && len(str) > 0 {

		err := json.Unmarshal([]byte(str), &res)
		if err != nil {
			logger.Errorf("could not parse the metadata service url parameters to map", err)
		}
	}
	return res
}

func getExportedServicesRevision(serviceInstance registry.ServiceInstance) string {
	metaData := serviceInstance.GetMetadata()
	return metaData[constant.ExportedServicesRevisionPropertyName]
}

func ConvertURLArrToIntfArr(urls []*common.URL) []interface{} {
	if len(urls) == 0 {
		return []interface{}{}
	}

	res := make([]interface{}, 0, len(urls))
	for _, u := range urls {
		res = append(res, u.String())
	}
	return res
}
