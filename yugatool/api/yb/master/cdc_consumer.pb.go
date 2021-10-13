// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// The following only applies to changes made to this file as part of YugaByte development.
//
// Portions Copyright (c) YugaByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.  See the License for the specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: yb/master/cdc_consumer.proto

package master

import (
	common "github.com/yugabyte/yb-tools/yugatool/api/yb/common"
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

type ProducerTabletListPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of tablet ids for a given producer table.
	Tablets []string `protobuf:"bytes,1,rep,name=tablets,proto3" json:"tablets,omitempty"`
}

func (x *ProducerTabletListPB) Reset() {
	*x = ProducerTabletListPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_master_cdc_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProducerTabletListPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProducerTabletListPB) ProtoMessage() {}

func (x *ProducerTabletListPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_master_cdc_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProducerTabletListPB.ProtoReflect.Descriptor instead.
func (*ProducerTabletListPB) Descriptor() ([]byte, []int) {
	return file_yb_master_cdc_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *ProducerTabletListPB) GetTablets() []string {
	if x != nil {
		return x.Tablets
	}
	return nil
}

type StreamEntryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from consumer to producer tablet id.
	ConsumerProducerTabletMap      map[string]*ProducerTabletListPB `protobuf:"bytes,1,rep,name=consumer_producer_tablet_map,json=consumerProducerTabletMap,proto3" json:"consumer_producer_tablet_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ConsumerTableId                string                           `protobuf:"bytes,2,opt,name=consumer_table_id,json=consumerTableId,proto3" json:"consumer_table_id,omitempty"`
	ProducerTableId                string                           `protobuf:"bytes,3,opt,name=producer_table_id,json=producerTableId,proto3" json:"producer_table_id,omitempty"`
	SameNumProducerConsumerTablets bool                             `protobuf:"varint,4,opt,name=same_num_producer_consumer_tablets,json=sameNumProducerConsumerTablets,proto3" json:"same_num_producer_consumer_tablets,omitempty"`
}

func (x *StreamEntryPB) Reset() {
	*x = StreamEntryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_master_cdc_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEntryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEntryPB) ProtoMessage() {}

func (x *StreamEntryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_master_cdc_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEntryPB.ProtoReflect.Descriptor instead.
func (*StreamEntryPB) Descriptor() ([]byte, []int) {
	return file_yb_master_cdc_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *StreamEntryPB) GetConsumerProducerTabletMap() map[string]*ProducerTabletListPB {
	if x != nil {
		return x.ConsumerProducerTabletMap
	}
	return nil
}

func (x *StreamEntryPB) GetConsumerTableId() string {
	if x != nil {
		return x.ConsumerTableId
	}
	return ""
}

func (x *StreamEntryPB) GetProducerTableId() string {
	if x != nil {
		return x.ProducerTableId
	}
	return ""
}

func (x *StreamEntryPB) GetSameNumProducerConsumerTablets() bool {
	if x != nil {
		return x.SameNumProducerConsumerTablets
	}
	return false
}

type ProducerEntryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from stream id to metadata for that stream.
	StreamMap     map[string]*StreamEntryPB `protobuf:"bytes,1,rep,name=stream_map,json=streamMap,proto3" json:"stream_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MasterAddrs   []*common.HostPortPB      `protobuf:"bytes,2,rep,name=master_addrs,json=masterAddrs,proto3" json:"master_addrs,omitempty"`
	TserverAddrs  []*common.HostPortPB      `protobuf:"bytes,3,rep,name=tserver_addrs,json=tserverAddrs,proto3" json:"tserver_addrs,omitempty"`
	DisableStream bool                      `protobuf:"varint,4,opt,name=disable_stream,json=disableStream,proto3" json:"disable_stream,omitempty"` // [default = false] implicit in proto3
}

func (x *ProducerEntryPB) Reset() {
	*x = ProducerEntryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_master_cdc_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProducerEntryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProducerEntryPB) ProtoMessage() {}

func (x *ProducerEntryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_master_cdc_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProducerEntryPB.ProtoReflect.Descriptor instead.
func (*ProducerEntryPB) Descriptor() ([]byte, []int) {
	return file_yb_master_cdc_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *ProducerEntryPB) GetStreamMap() map[string]*StreamEntryPB {
	if x != nil {
		return x.StreamMap
	}
	return nil
}

func (x *ProducerEntryPB) GetMasterAddrs() []*common.HostPortPB {
	if x != nil {
		return x.MasterAddrs
	}
	return nil
}

func (x *ProducerEntryPB) GetTserverAddrs() []*common.HostPortPB {
	if x != nil {
		return x.TserverAddrs
	}
	return nil
}

func (x *ProducerEntryPB) GetDisableStream() bool {
	if x != nil {
		return x.DisableStream
	}
	return false
}

type ConsumerRegistryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from producer universe uuid to metadata for that producer.
	ProducerMap map[string]*ProducerEntryPB `protobuf:"bytes,1,rep,name=producer_map,json=producerMap,proto3" json:"producer_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConsumerRegistryPB) Reset() {
	*x = ConsumerRegistryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_master_cdc_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerRegistryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerRegistryPB) ProtoMessage() {}

func (x *ConsumerRegistryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_master_cdc_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerRegistryPB.ProtoReflect.Descriptor instead.
func (*ConsumerRegistryPB) Descriptor() ([]byte, []int) {
	return file_yb_master_cdc_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *ConsumerRegistryPB) GetProducerMap() map[string]*ProducerEntryPB {
	if x != nil {
		return x.ProducerMap
	}
	return nil
}

var File_yb_master_cdc_consumer_proto protoreflect.FileDescriptor

var file_yb_master_cdc_consumer_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x79, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x64, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x79, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x30, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x42, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x73, 0x22, 0x9c, 0x03, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x78, 0x0a, 0x1c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x79, 0x62,
	0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x22, 0x73, 0x61, 0x6d, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1e, 0x73, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x73, 0x1a, 0x6d, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xc2, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x48, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x62, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70,
	0x12, 0x31, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x50, 0x42, 0x52, 0x0b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x50, 0x42, 0x52, 0x0c, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a,
	0x56, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc3, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x51,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x50, 0x42, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x61,
	0x70, 0x1a, 0x5a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x42, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0f, 0x0a,
	0x0d, 0x6f, 0x72, 0x67, 0x2e, 0x79, 0x62, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yb_master_cdc_consumer_proto_rawDescOnce sync.Once
	file_yb_master_cdc_consumer_proto_rawDescData = file_yb_master_cdc_consumer_proto_rawDesc
)

func file_yb_master_cdc_consumer_proto_rawDescGZIP() []byte {
	file_yb_master_cdc_consumer_proto_rawDescOnce.Do(func() {
		file_yb_master_cdc_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_master_cdc_consumer_proto_rawDescData)
	})
	return file_yb_master_cdc_consumer_proto_rawDescData
}

var file_yb_master_cdc_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yb_master_cdc_consumer_proto_goTypes = []interface{}{
	(*ProducerTabletListPB)(nil), // 0: yb.master.ProducerTabletListPB
	(*StreamEntryPB)(nil),        // 1: yb.master.StreamEntryPB
	(*ProducerEntryPB)(nil),      // 2: yb.master.ProducerEntryPB
	(*ConsumerRegistryPB)(nil),   // 3: yb.master.ConsumerRegistryPB
	nil,                          // 4: yb.master.StreamEntryPB.ConsumerProducerTabletMapEntry
	nil,                          // 5: yb.master.ProducerEntryPB.StreamMapEntry
	nil,                          // 6: yb.master.ConsumerRegistryPB.ProducerMapEntry
	(*common.HostPortPB)(nil),    // 7: yb.HostPortPB
}
var file_yb_master_cdc_consumer_proto_depIdxs = []int32{
	4, // 0: yb.master.StreamEntryPB.consumer_producer_tablet_map:type_name -> yb.master.StreamEntryPB.ConsumerProducerTabletMapEntry
	5, // 1: yb.master.ProducerEntryPB.stream_map:type_name -> yb.master.ProducerEntryPB.StreamMapEntry
	7, // 2: yb.master.ProducerEntryPB.master_addrs:type_name -> yb.HostPortPB
	7, // 3: yb.master.ProducerEntryPB.tserver_addrs:type_name -> yb.HostPortPB
	6, // 4: yb.master.ConsumerRegistryPB.producer_map:type_name -> yb.master.ConsumerRegistryPB.ProducerMapEntry
	0, // 5: yb.master.StreamEntryPB.ConsumerProducerTabletMapEntry.value:type_name -> yb.master.ProducerTabletListPB
	1, // 6: yb.master.ProducerEntryPB.StreamMapEntry.value:type_name -> yb.master.StreamEntryPB
	2, // 7: yb.master.ConsumerRegistryPB.ProducerMapEntry.value:type_name -> yb.master.ProducerEntryPB
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_yb_master_cdc_consumer_proto_init() }
func file_yb_master_cdc_consumer_proto_init() {
	if File_yb_master_cdc_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yb_master_cdc_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProducerTabletListPB); i {
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
		file_yb_master_cdc_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEntryPB); i {
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
		file_yb_master_cdc_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProducerEntryPB); i {
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
		file_yb_master_cdc_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerRegistryPB); i {
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
			RawDescriptor: file_yb_master_cdc_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_master_cdc_consumer_proto_goTypes,
		DependencyIndexes: file_yb_master_cdc_consumer_proto_depIdxs,
		MessageInfos:      file_yb_master_cdc_consumer_proto_msgTypes,
	}.Build()
	File_yb_master_cdc_consumer_proto = out.File
	file_yb_master_cdc_consumer_proto_rawDesc = nil
	file_yb_master_cdc_consumer_proto_goTypes = nil
	file_yb_master_cdc_consumer_proto_depIdxs = nil
}
