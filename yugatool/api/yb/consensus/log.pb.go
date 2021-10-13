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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: yb/consensus/log.proto

package consensus

import (
	common "github.com/yugabyte/yb-tools/yugatool/api/yb/common"
	util "github.com/yugabyte/yb-tools/yugatool/api/yb/util"
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

// Types of log entries.
type LogEntryTypePB int32

const (
	LogEntryTypePB_UNKNOWN   LogEntryTypePB = 0
	LogEntryTypePB_REPLICATE LogEntryTypePB = 1
	// Roll over active log segment before writing the next entry if active segment is not empty.
	LogEntryTypePB_ROLLOVER_MARKER LogEntryTypePB = 998
	// Serve the purpose of making sure that all entries up to the FLUSH_MARKER entry are flushed.
	LogEntryTypePB_FLUSH_MARKER LogEntryTypePB = 999
)

// Enum value maps for LogEntryTypePB.
var (
	LogEntryTypePB_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "REPLICATE",
		998: "ROLLOVER_MARKER",
		999: "FLUSH_MARKER",
	}
	LogEntryTypePB_value = map[string]int32{
		"UNKNOWN":         0,
		"REPLICATE":       1,
		"ROLLOVER_MARKER": 998,
		"FLUSH_MARKER":    999,
	}
)

func (x LogEntryTypePB) Enum() *LogEntryTypePB {
	p := new(LogEntryTypePB)
	*p = x
	return p
}

func (x LogEntryTypePB) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogEntryTypePB) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_consensus_log_proto_enumTypes[0].Descriptor()
}

func (LogEntryTypePB) Type() protoreflect.EnumType {
	return &file_yb_consensus_log_proto_enumTypes[0]
}

func (x LogEntryTypePB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LogEntryTypePB) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LogEntryTypePB(num)
	return nil
}

// Deprecated: Use LogEntryTypePB.Descriptor instead.
func (LogEntryTypePB) EnumDescriptor() ([]byte, []int) {
	return file_yb_consensus_log_proto_rawDescGZIP(), []int{0}
}

// An entry in the WAL/state machine log.
type LogEntryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      *LogEntryTypePB `protobuf:"varint,1,req,name=type,enum=yb.log.LogEntryTypePB" json:"type,omitempty"`
	Replicate *ReplicateMsg   `protobuf:"bytes,2,opt,name=replicate" json:"replicate,omitempty"`
}

func (x *LogEntryPB) Reset() {
	*x = LogEntryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_consensus_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryPB) ProtoMessage() {}

func (x *LogEntryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_consensus_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryPB.ProtoReflect.Descriptor instead.
func (*LogEntryPB) Descriptor() ([]byte, []int) {
	return file_yb_consensus_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntryPB) GetType() LogEntryTypePB {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return LogEntryTypePB_UNKNOWN
}

func (x *LogEntryPB) GetReplicate() *ReplicateMsg {
	if x != nil {
		return x.Replicate
	}
	return nil
}

// A batch of entries in the WAL.
type LogEntryBatchPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry         []*LogEntryPB `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
	CommittedOpId *util.OpIdPB  `protobuf:"bytes,2,opt,name=committed_op_id,json=committedOpId" json:"committed_op_id,omitempty"`
	// Time when this batch was appended, this time is monotonically increase,
	// and restart safe. I.e. first batch after restart will have time greater than or equal to
	// time of last batch added before restart.
	MonoTime *uint64 `protobuf:"varint,3,opt,name=mono_time,json=monoTime" json:"mono_time,omitempty"`
}

func (x *LogEntryBatchPB) Reset() {
	*x = LogEntryBatchPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_consensus_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntryBatchPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntryBatchPB) ProtoMessage() {}

func (x *LogEntryBatchPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_consensus_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntryBatchPB.ProtoReflect.Descriptor instead.
func (*LogEntryBatchPB) Descriptor() ([]byte, []int) {
	return file_yb_consensus_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogEntryBatchPB) GetEntry() []*LogEntryPB {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *LogEntryBatchPB) GetCommittedOpId() *util.OpIdPB {
	if x != nil {
		return x.CommittedOpId
	}
	return nil
}

func (x *LogEntryBatchPB) GetMonoTime() uint64 {
	if x != nil && x.MonoTime != nil {
		return *x.MonoTime
	}
	return 0
}

// A header for a log segment.
type LogSegmentHeaderPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log format major version.
	MajorVersion *uint32 `protobuf:"varint,1,req,name=major_version,json=majorVersion" json:"major_version,omitempty"`
	// Log format minor version.
	MinorVersion *uint32 `protobuf:"varint,2,req,name=minor_version,json=minorVersion" json:"minor_version,omitempty"`
	// The ID of the tablet this WAL segment stores entries for.
	TabletId []byte `protobuf:"bytes,5,req,name=tablet_id,json=tabletId" json:"tablet_id,omitempty"`
	// The tablet-specific sequence number of this WAL segment.
	SequenceNumber *uint64 `protobuf:"varint,6,req,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// Schema used when appending entries to this log, and its version.
	Schema        *common.SchemaPB `protobuf:"bytes,7,req,name=schema" json:"schema,omitempty"`
	SchemaVersion *uint32          `protobuf:"varint,8,opt,name=schema_version,json=schemaVersion" json:"schema_version,omitempty"`
}

func (x *LogSegmentHeaderPB) Reset() {
	*x = LogSegmentHeaderPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_consensus_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSegmentHeaderPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSegmentHeaderPB) ProtoMessage() {}

func (x *LogSegmentHeaderPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_consensus_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSegmentHeaderPB.ProtoReflect.Descriptor instead.
func (*LogSegmentHeaderPB) Descriptor() ([]byte, []int) {
	return file_yb_consensus_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogSegmentHeaderPB) GetMajorVersion() uint32 {
	if x != nil && x.MajorVersion != nil {
		return *x.MajorVersion
	}
	return 0
}

func (x *LogSegmentHeaderPB) GetMinorVersion() uint32 {
	if x != nil && x.MinorVersion != nil {
		return *x.MinorVersion
	}
	return 0
}

func (x *LogSegmentHeaderPB) GetTabletId() []byte {
	if x != nil {
		return x.TabletId
	}
	return nil
}

func (x *LogSegmentHeaderPB) GetSequenceNumber() uint64 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

func (x *LogSegmentHeaderPB) GetSchema() *common.SchemaPB {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *LogSegmentHeaderPB) GetSchemaVersion() uint32 {
	if x != nil && x.SchemaVersion != nil {
		return *x.SchemaVersion
	}
	return 0
}

// A footer for a log segment.
//
// Log segment footers might not be present (e.g. if the server crashed) so they should contain no
// information that cannot be obtained by actually reading the entries in the log.
//
// We use the footer to keep sparse index entries mapping op_id->offset (right now we just keep the
// first entry with an id in the log).
type LogSegmentFooterPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total number of operations inside this segment.
	NumEntries *int64 `protobuf:"varint,1,req,name=num_entries,json=numEntries" json:"num_entries,omitempty"`
	// The minimum and maximum index of a REPLICATE message in this segment.
	// NOTE: because of log truncation, the min/max are not necessarily the first/last!  For example,
	// a log segment may contain entries "1.5, 1.6, 2.3, 3.3" due to multiple term changes.
	//
	// Because it's possible for a segment to have no REPLICATE messages in it, we set the default
	// to -1 for these fields to avoid accidentally reading 0, which might look like a real log index.
	MinReplicateIndex *int64 `protobuf:"varint,2,opt,name=min_replicate_index,json=minReplicateIndex,def=-1" json:"min_replicate_index,omitempty"`
	MaxReplicateIndex *int64 `protobuf:"varint,3,opt,name=max_replicate_index,json=maxReplicateIndex,def=-1" json:"max_replicate_index,omitempty"`
	// The time (microseconds since epoch) when this segment was closed.
	// NOTE: in configurations where --skip_wal_rewrite=false is specified, log segments are rewritten
	// during bootstrap, so this field will be reset to the time of the bootstrap in each log segment
	// on a newly-restarted server, rather than copied over from the old log segments.
	CloseTimestampMicros *int64 `protobuf:"varint,4,opt,name=close_timestamp_micros,json=closeTimestampMicros" json:"close_timestamp_micros,omitempty"`
}

// Default values for LogSegmentFooterPB fields.
const (
	Default_LogSegmentFooterPB_MinReplicateIndex = int64(-1)
	Default_LogSegmentFooterPB_MaxReplicateIndex = int64(-1)
)

func (x *LogSegmentFooterPB) Reset() {
	*x = LogSegmentFooterPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_consensus_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSegmentFooterPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSegmentFooterPB) ProtoMessage() {}

func (x *LogSegmentFooterPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_consensus_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSegmentFooterPB.ProtoReflect.Descriptor instead.
func (*LogSegmentFooterPB) Descriptor() ([]byte, []int) {
	return file_yb_consensus_log_proto_rawDescGZIP(), []int{3}
}

func (x *LogSegmentFooterPB) GetNumEntries() int64 {
	if x != nil && x.NumEntries != nil {
		return *x.NumEntries
	}
	return 0
}

func (x *LogSegmentFooterPB) GetMinReplicateIndex() int64 {
	if x != nil && x.MinReplicateIndex != nil {
		return *x.MinReplicateIndex
	}
	return Default_LogSegmentFooterPB_MinReplicateIndex
}

func (x *LogSegmentFooterPB) GetMaxReplicateIndex() int64 {
	if x != nil && x.MaxReplicateIndex != nil {
		return *x.MaxReplicateIndex
	}
	return Default_LogSegmentFooterPB_MaxReplicateIndex
}

func (x *LogSegmentFooterPB) GetCloseTimestampMicros() int64 {
	if x != nil && x.CloseTimestampMicros != nil {
		return *x.CloseTimestampMicros
	}
	return 0
}

var File_yb_consensus_log_proto protoreflect.FileDescriptor

var file_yb_consensus_log_proto_rawDesc = []byte{
	0x0a, 0x16, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x79, 0x62, 0x2e, 0x6c, 0x6f, 0x67,
	0x1a, 0x16, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x79, 0x62, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x6f, 0x70, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x79, 0x62, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x42, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x79, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67,
	0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0f,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x42, 0x12,
	0x28, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x79, 0x62, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x50, 0x42, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x79, 0x62, 0x2e, 0x4f, 0x70, 0x49, 0x64, 0x50, 0x42, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x6f, 0x6e, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x6d, 0x6f, 0x6e, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x12, 0x4c,
	0x6f, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50,
	0x42, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x08,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x04, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x07, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x79, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x42, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd3,
	0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x6f, 0x6f,
	0x74, 0x65, 0x72, 0x50, 0x42, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x13, 0x6d, 0x61,
	0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x11, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x34,
	0x0a, 0x16, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x2a, 0x55, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x42, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x52, 0x4f, 0x4c, 0x4c, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x4d,
	0x41, 0x52, 0x4b, 0x45, 0x52, 0x10, 0xe6, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x46, 0x4c, 0x55, 0x53,
	0x48, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52, 0x10, 0xe7, 0x07, 0x42, 0x0c, 0x0a, 0x0a, 0x6f,
	0x72, 0x67, 0x2e, 0x79, 0x62, 0x2e, 0x6c, 0x6f, 0x67,
}

var (
	file_yb_consensus_log_proto_rawDescOnce sync.Once
	file_yb_consensus_log_proto_rawDescData = file_yb_consensus_log_proto_rawDesc
)

func file_yb_consensus_log_proto_rawDescGZIP() []byte {
	file_yb_consensus_log_proto_rawDescOnce.Do(func() {
		file_yb_consensus_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_consensus_log_proto_rawDescData)
	})
	return file_yb_consensus_log_proto_rawDescData
}

var file_yb_consensus_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yb_consensus_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yb_consensus_log_proto_goTypes = []interface{}{
	(LogEntryTypePB)(0),        // 0: yb.log.LogEntryTypePB
	(*LogEntryPB)(nil),         // 1: yb.log.LogEntryPB
	(*LogEntryBatchPB)(nil),    // 2: yb.log.LogEntryBatchPB
	(*LogSegmentHeaderPB)(nil), // 3: yb.log.LogSegmentHeaderPB
	(*LogSegmentFooterPB)(nil), // 4: yb.log.LogSegmentFooterPB
	(*ReplicateMsg)(nil),       // 5: yb.consensus.ReplicateMsg
	(*util.OpIdPB)(nil),        // 6: yb.OpIdPB
	(*common.SchemaPB)(nil),    // 7: yb.SchemaPB
}
var file_yb_consensus_log_proto_depIdxs = []int32{
	0, // 0: yb.log.LogEntryPB.type:type_name -> yb.log.LogEntryTypePB
	5, // 1: yb.log.LogEntryPB.replicate:type_name -> yb.consensus.ReplicateMsg
	1, // 2: yb.log.LogEntryBatchPB.entry:type_name -> yb.log.LogEntryPB
	6, // 3: yb.log.LogEntryBatchPB.committed_op_id:type_name -> yb.OpIdPB
	7, // 4: yb.log.LogSegmentHeaderPB.schema:type_name -> yb.SchemaPB
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yb_consensus_log_proto_init() }
func file_yb_consensus_log_proto_init() {
	if File_yb_consensus_log_proto != nil {
		return
	}
	file_yb_consensus_consensus_proto_init()
	file_yb_consensus_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yb_consensus_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryPB); i {
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
		file_yb_consensus_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntryBatchPB); i {
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
		file_yb_consensus_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSegmentHeaderPB); i {
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
		file_yb_consensus_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSegmentFooterPB); i {
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
			RawDescriptor: file_yb_consensus_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_consensus_log_proto_goTypes,
		DependencyIndexes: file_yb_consensus_log_proto_depIdxs,
		EnumInfos:         file_yb_consensus_log_proto_enumTypes,
		MessageInfos:      file_yb_consensus_log_proto_msgTypes,
	}.Build()
	File_yb_consensus_log_proto = out.File
	file_yb_consensus_log_proto_rawDesc = nil
	file_yb_consensus_log_proto_goTypes = nil
	file_yb_consensus_log_proto_depIdxs = nil
}
