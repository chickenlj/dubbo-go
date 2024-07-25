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

package local

import (
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	common_meta "dubbo.apache.org/dubbo-go/v3/common/metadata"
	triple_api "dubbo.apache.org/dubbo-go/v3/metadata/triple_api/proto"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/dubbogo/gost/log/logger"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

type MetadataServiceProxyFactory interface {
	GetProxy(urls []*common.URL) common_meta.MetadataService
	GetProxyV2(urls []*common.URL) triple_api.MetadataServiceV2
}

type BaseMetadataServiceProxyFactory struct {
}

func NewBaseMetadataServiceProxyFactory() *BaseMetadataServiceProxyFactory {
	return &BaseMetadataServiceProxyFactory{}
}

func (b *BaseMetadataServiceProxyFactory) GetProxy(urls []*common.URL) common_meta.MetadataService {
	if len(urls) == 0 {
		logger.Errorf("metadata service urls not found")
		return nil
	}

	u := urls[0]
	p := extension.GetProtocol(u.Protocol)
	invoker := p.Refer(u)
	if invoker == nil { // can't connect instance
		return nil
	}
	return &MetadataServiceProxy{
		Invoker: invoker,
	}
}

func (b *BaseMetadataServiceProxyFactory) GetProxyV2(urls []*common.URL) triple_api.MetadataServiceV2 {
	invoker, ok := getInvoker(urls)
	if !ok {
		return nil
	}
	return &MetadataServiceProxyV2{
		Invoker: invoker,
	}
}

func getInvoker(urls []*common.URL) (protocol.Invoker, bool) {
	if len(urls) == 0 {
		logger.Errorf("metadata service urls not found")
		return nil, false
	}

	u := urls[0]
	p := extension.GetProtocol(u.Protocol)
	invoker := p.Refer(u)
	if invoker == nil { // can't connect instance
		return nil, false
	}
	return invoker, true
}
