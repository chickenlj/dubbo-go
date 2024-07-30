//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: proto/metadata_service_v2.proto

package triple_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metadata_service_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metadata_service_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_proto_metadata_service_v2_proto_rawDescGZIP(), []int{0}
}

func (x *Revision) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MetadataInfoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App      string                    `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Version  string                    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Services map[string]*ServiceInfoV2 `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetadataInfoV2) Reset() {
	*x = MetadataInfoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metadata_service_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataInfoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataInfoV2) ProtoMessage() {}

func (x *MetadataInfoV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metadata_service_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataInfoV2.ProtoReflect.Descriptor instead.
func (*MetadataInfoV2) Descriptor() ([]byte, []int) {
	return file_proto_metadata_service_v2_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataInfoV2) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *MetadataInfoV2) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MetadataInfoV2) GetServices() map[string]*ServiceInfoV2 {
	if x != nil {
		return x.Services
	}
	return nil
}

type ServiceInfoV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group    string            `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Version  string            `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Protocol string            `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port     int32             `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Path     string            `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Params   map[string]string `protobuf:"bytes,7,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceInfoV2) Reset() {
	*x = ServiceInfoV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metadata_service_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfoV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfoV2) ProtoMessage() {}

func (x *ServiceInfoV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metadata_service_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfoV2.ProtoReflect.Descriptor instead.
func (*ServiceInfoV2) Descriptor() ([]byte, []int) {
	return file_proto_metadata_service_v2_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceInfoV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceInfoV2) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ServiceInfoV2) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceInfoV2) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServiceInfoV2) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServiceInfoV2) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ServiceInfoV2) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_proto_metadata_service_v2_proto protoreflect.FileDescriptor

var file_proto_metadata_service_v2_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf8,
	0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x56,
	0x32, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x65, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa0, 0x02, 0x0a, 0x0d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x4c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x32, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x76, 0x0a, 0x11,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56,
	0x32, 0x12, 0x61, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x56, 0x32, 0x42, 0x5a, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x50, 0x01, 0x5a, 0x3b, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x3b, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_metadata_service_v2_proto_rawDescOnce sync.Once
	file_proto_metadata_service_v2_proto_rawDescData = file_proto_metadata_service_v2_proto_rawDesc
)

func file_proto_metadata_service_v2_proto_rawDescGZIP() []byte {
	file_proto_metadata_service_v2_proto_rawDescOnce.Do(func() {
		file_proto_metadata_service_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metadata_service_v2_proto_rawDescData)
	})
	return file_proto_metadata_service_v2_proto_rawDescData
}

var file_proto_metadata_service_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_metadata_service_v2_proto_goTypes = []interface{}{
	(*Revision)(nil),       // 0: org.apache.dubbo.metadata.Revision
	(*MetadataInfoV2)(nil), // 1: org.apache.dubbo.metadata.MetadataInfoV2
	(*ServiceInfoV2)(nil),  // 2: org.apache.dubbo.metadata.ServiceInfoV2
	nil,                    // 3: org.apache.dubbo.metadata.MetadataInfoV2.ServicesEntry
	nil,                    // 4: org.apache.dubbo.metadata.ServiceInfoV2.ParamsEntry
}
var file_proto_metadata_service_v2_proto_depIdxs = []int32{
	3, // 0: org.apache.dubbo.metadata.MetadataInfoV2.services:type_name -> org.apache.dubbo.metadata.MetadataInfoV2.ServicesEntry
	4, // 1: org.apache.dubbo.metadata.ServiceInfoV2.params:type_name -> org.apache.dubbo.metadata.ServiceInfoV2.ParamsEntry
	2, // 2: org.apache.dubbo.metadata.MetadataInfoV2.ServicesEntry.value:type_name -> org.apache.dubbo.metadata.ServiceInfoV2
	0, // 3: org.apache.dubbo.metadata.MetadataServiceV2.GetMetadataInfo:input_type -> org.apache.dubbo.metadata.Revision
	1, // 4: org.apache.dubbo.metadata.MetadataServiceV2.GetMetadataInfo:output_type -> org.apache.dubbo.metadata.MetadataInfoV2
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_metadata_service_v2_proto_init() }
func file_proto_metadata_service_v2_proto_init() {
	if File_proto_metadata_service_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metadata_service_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_metadata_service_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataInfoV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_metadata_service_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfoV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_metadata_service_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_metadata_service_v2_proto_goTypes,
		DependencyIndexes: file_proto_metadata_service_v2_proto_depIdxs,
		MessageInfos:      file_proto_metadata_service_v2_proto_msgTypes,
	}.Build()
	File_proto_metadata_service_v2_proto = out.File
	file_proto_metadata_service_v2_proto_rawDesc = nil
	file_proto_metadata_service_v2_proto_goTypes = nil
	file_proto_metadata_service_v2_proto_depIdxs = nil
}
