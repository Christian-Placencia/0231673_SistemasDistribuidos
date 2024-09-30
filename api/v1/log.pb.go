// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: api/v1/log.proto

package v1

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

// The CommitLog message that will be used in communication
type CommitLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []string `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"` // A list of log entries
}

func (x *CommitLog) Reset() {
	*x = CommitLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitLog) ProtoMessage() {}

func (x *CommitLog) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitLog.ProtoReflect.Descriptor instead.
func (*CommitLog) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *CommitLog) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

// Request to append a new entry to the log
type AddEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry string `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *AddEntryRequest) Reset() {
	*x = AddEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryRequest) ProtoMessage() {}

func (x *AddEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryRequest.ProtoReflect.Descriptor instead.
func (*AddEntryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *AddEntryRequest) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

// Response after adding a new entry
type AddEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddEntryResponse) Reset() {
	*x = AddEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryResponse) ProtoMessage() {}

func (x *AddEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryResponse.ProtoReflect.Descriptor instead.
func (*AddEntryResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{2}
}

func (x *AddEntryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Request to get an entry from the log by index
type GetEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *GetEntryRequest) Reset() {
	*x = GetEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryRequest) ProtoMessage() {}

func (x *GetEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryRequest.ProtoReflect.Descriptor instead.
func (*GetEntryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{3}
}

func (x *GetEntryRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

// Response containing the entry or an error
type GetEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry string `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"` // Empty if there's no error
}

func (x *GetEntryResponse) Reset() {
	*x = GetEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryResponse) ProtoMessage() {}

func (x *GetEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryResponse.ProtoReflect.Descriptor instead.
func (*GetEntryResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{4}
}

func (x *GetEntryResponse) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

func (x *GetEntryResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Message for log records, used in advanced log operations
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // The offset of the record in the log
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`    // The value of the record (as bytes)
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{5}
}

func (x *Record) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Record) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// Request to produce a log record
type ProduceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *ProduceRequest) Reset() {
	*x = ProduceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceRequest) ProtoMessage() {}

func (x *ProduceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceRequest.ProtoReflect.Descriptor instead.
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{6}
}

func (x *ProduceRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

// Response after producing a log record, returning its offset
type ProduceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ProduceResponse) Reset() {
	*x = ProduceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProduceResponse) ProtoMessage() {}

func (x *ProduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProduceResponse.ProtoReflect.Descriptor instead.
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{7}
}

func (x *ProduceResponse) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Request to consume a log record by its offset
type ConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ConsumeRequest) Reset() {
	*x = ConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRequest) ProtoMessage() {}

func (x *ConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRequest.ProtoReflect.Descriptor instead.
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{8}
}

func (x *ConsumeRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Response containing the consumed log record
type ConsumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *ConsumeResponse) Reset() {
	*x = ConsumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_log_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeResponse) ProtoMessage() {}

func (x *ConsumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_log_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeResponse.ProtoReflect.Descriptor instead.
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_log_proto_rawDescGZIP(), []int{9}
}

func (x *ConsumeResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

var File_api_v1_log_proto protoreflect.FileDescriptor

var file_api_v1_log_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x25, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x27,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x2c, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3e,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x36,
	0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x36, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x32, 0x84, 0x01, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x78, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x62, 0x61, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_log_proto_rawDescOnce sync.Once
	file_api_v1_log_proto_rawDescData = file_api_v1_log_proto_rawDesc
)

func file_api_v1_log_proto_rawDescGZIP() []byte {
	file_api_v1_log_proto_rawDescOnce.Do(func() {
		file_api_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_log_proto_rawDescData)
	})
	return file_api_v1_log_proto_rawDescData
}

var file_api_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_log_proto_goTypes = []any{
	(*CommitLog)(nil),        // 0: api.CommitLog
	(*AddEntryRequest)(nil),  // 1: api.AddEntryRequest
	(*AddEntryResponse)(nil), // 2: api.AddEntryResponse
	(*GetEntryRequest)(nil),  // 3: api.GetEntryRequest
	(*GetEntryResponse)(nil), // 4: api.GetEntryResponse
	(*Record)(nil),           // 5: api.Record
	(*ProduceRequest)(nil),   // 6: api.ProduceRequest
	(*ProduceResponse)(nil),  // 7: api.ProduceResponse
	(*ConsumeRequest)(nil),   // 8: api.ConsumeRequest
	(*ConsumeResponse)(nil),  // 9: api.ConsumeResponse
}
var file_api_v1_log_proto_depIdxs = []int32{
	5, // 0: api.ProduceRequest.record:type_name -> api.Record
	5, // 1: api.ConsumeResponse.record:type_name -> api.Record
	1, // 2: api.CommitLogService.AddEntry:input_type -> api.AddEntryRequest
	3, // 3: api.CommitLogService.GetEntry:input_type -> api.GetEntryRequest
	6, // 4: api.LogService.Produce:input_type -> api.ProduceRequest
	8, // 5: api.LogService.Consume:input_type -> api.ConsumeRequest
	2, // 6: api.CommitLogService.AddEntry:output_type -> api.AddEntryResponse
	4, // 7: api.CommitLogService.GetEntry:output_type -> api.GetEntryResponse
	7, // 8: api.LogService.Produce:output_type -> api.ProduceResponse
	9, // 9: api.LogService.Consume:output_type -> api.ConsumeResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_log_proto_init() }
func file_api_v1_log_proto_init() {
	if File_api_v1_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_log_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommitLog); i {
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
		file_api_v1_log_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddEntryRequest); i {
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
		file_api_v1_log_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddEntryResponse); i {
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
		file_api_v1_log_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetEntryRequest); i {
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
		file_api_v1_log_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetEntryResponse); i {
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
		file_api_v1_log_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Record); i {
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
		file_api_v1_log_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ProduceRequest); i {
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
		file_api_v1_log_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ProduceResponse); i {
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
		file_api_v1_log_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ConsumeRequest); i {
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
		file_api_v1_log_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ConsumeResponse); i {
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
			RawDescriptor: file_api_v1_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_v1_log_proto_goTypes,
		DependencyIndexes: file_api_v1_log_proto_depIdxs,
		MessageInfos:      file_api_v1_log_proto_msgTypes,
	}.Build()
	File_api_v1_log_proto = out.File
	file_api_v1_log_proto_rawDesc = nil
	file_api_v1_log_proto_goTypes = nil
	file_api_v1_log_proto_depIdxs = nil
}