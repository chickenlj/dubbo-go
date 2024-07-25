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

package mapping

import (
	gxset "github.com/dubbogo/gost/container/set"
	"github.com/dubbogo/gost/gof/observer"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

// ServiceNameMapping  is the interface which trys to build the mapping between application-level service and interface-level service.
//
// # Map method will map the service to this application-level service
//
// Get method will return the application-level services
type ServiceNameMapping interface {
	Map(url *common.URL) error
	Get(url *common.URL, listener MappingListener) (*gxset.HashSet, error)
	Remove(url *common.URL) error
}

type ServiceMappingChangeEvent struct {
	observer.BaseEvent
	ServiceKey   string
	ServiceNames *gxset.HashSet
}

// NewServiceMappingChangedEvent will create the ServiceMappingChangeEvent
func NewServiceMappingChangedEvent(serviceKey string, serviceNames *gxset.HashSet) *ServiceMappingChangeEvent {
	return &ServiceMappingChangeEvent{
		BaseEvent: observer.BaseEvent{
			Source:    serviceKey,
			Timestamp: time.Now(),
		},
		ServiceKey:   serviceKey,
		ServiceNames: serviceNames,
	}
}

func (sm *ServiceMappingChangeEvent) GetServiceNames() *gxset.HashSet {
	return sm.ServiceNames
}

func (sm *ServiceMappingChangeEvent) GetServiceKey() string {
	return sm.ServiceKey
}
