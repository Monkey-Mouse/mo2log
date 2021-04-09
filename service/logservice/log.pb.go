// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: log.proto

package logservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LogModel_Level int32

const (
	LogModel_INFO    LogModel_Level = 0
	LogModel_WARNING LogModel_Level = 1
	LogModel_ERROR   LogModel_Level = 2
	LogModel_FATAL   LogModel_Level = 3
)

// Enum value maps for LogModel_Level.
var (
	LogModel_Level_name = map[int32]string{
		0: "INFO",
		1: "WARNING",
		2: "ERROR",
		3: "FATAL",
	}
	LogModel_Level_value = map[string]int32{
		"INFO":    0,
		"WARNING": 1,
		"ERROR":   2,
		"FATAL":   3,
	}
)

func (x LogModel_Level) Enum() *LogModel_Level {
	p := new(LogModel_Level)
	*p = x
	return p
}

func (x LogModel_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogModel_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_log_proto_enumTypes[0].Descriptor()
}

func (LogModel_Level) Type() protoreflect.EnumType {
	return &file_log_proto_enumTypes[0]
}

func (x LogModel_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogModel_Level.Descriptor instead.
func (LogModel_Level) EnumDescriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{6, 0}
}

type ExtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator        []byte `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Operation       int32  `protobuf:"varint,2,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationTarget []byte `protobuf:"bytes,3,opt,name=operation_target,json=operationTarget,proto3" json:"operation_target,omitempty"`
}

func (x *ExtRequest) Reset() {
	*x = ExtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtRequest) ProtoMessage() {}

func (x *ExtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtRequest.ProtoReflect.Descriptor instead.
func (*ExtRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *ExtRequest) GetOperator() []byte {
	if x != nil {
		return x.Operator
	}
	return nil
}

func (x *ExtRequest) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *ExtRequest) GetOperationTarget() []byte {
	if x != nil {
		return x.OperationTarget
	}
	return nil
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query []byte `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{1}
}

func (x *Query) GetQuery() []byte {
	if x != nil {
		return x.Query
	}
	return nil
}

type Num struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Num) Reset() {
	*x = Num{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Num) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Num) ProtoMessage() {}

func (x *Num) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Num.ProtoReflect.Descriptor instead.
func (*Num) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{2}
}

func (x *Num) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId []byte `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{3}
}

func (x *UserID) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   []byte `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pagesize int32  `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetPagesize() int32 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type LogArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*LogModel `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *LogArray) Reset() {
	*x = LogArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogArray) ProtoMessage() {}

func (x *LogArray) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogArray.ProtoReflect.Descriptor instead.
func (*LogArray) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{5}
}

func (x *LogArray) GetLogs() []*LogModel {
	if x != nil {
		return x.Logs
	}
	return nil
}

type LogModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator             []byte         `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Operation            int32          `protobuf:"varint,2,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationTarget      []byte         `protobuf:"bytes,3,opt,name=operation_target,json=operationTarget,proto3" json:"operation_target,omitempty"`
	LogLevel             LogModel_Level `protobuf:"varint,4,opt,name=log_level,json=logLevel,proto3,enum=proto.LogModel_Level" json:"log_level,omitempty"`
	ExtraMessage         string         `protobuf:"bytes,5,opt,name=extra_message,json=extraMessage,proto3" json:"extra_message,omitempty"`
	CreateTime           int64          `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64          `protobuf:"varint,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	OperationTargetOwner []byte         `protobuf:"bytes,8,opt,name=operation_target_owner,json=operationTargetOwner,proto3" json:"operation_target_owner,omitempty"`
	Processed            bool           `protobuf:"varint,9,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (x *LogModel) Reset() {
	*x = LogModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogModel) ProtoMessage() {}

func (x *LogModel) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogModel.ProtoReflect.Descriptor instead.
func (*LogModel) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{6}
}

func (x *LogModel) GetOperator() []byte {
	if x != nil {
		return x.Operator
	}
	return nil
}

func (x *LogModel) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *LogModel) GetOperationTarget() []byte {
	if x != nil {
		return x.OperationTarget
	}
	return nil
}

func (x *LogModel) GetLogLevel() LogModel_Level {
	if x != nil {
		return x.LogLevel
	}
	return LogModel_INFO
}

func (x *LogModel) GetExtraMessage() string {
	if x != nil {
		return x.ExtraMessage
	}
	return ""
}

func (x *LogModel) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *LogModel) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *LogModel) GetOperationTargetOwner() []byte {
	if x != nil {
		return x.OperationTargetOwner
	}
	return nil
}

func (x *LogModel) GetProcessed() bool {
	if x != nil {
		return x.Processed
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{7}
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x1d, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x22, 0x17, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x21, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2f, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x94, 0x03, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x6c,
	0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x34, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x03,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xba, 0x02, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x73, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x4e, 0x75, 0x6d, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x75, 0x6d,
	0x12, 0x24, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x6c, 0x6f,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_log_proto_goTypes = []interface{}{
	(LogModel_Level)(0), // 0: proto.LogModel.Level
	(*ExtRequest)(nil),  // 1: proto.ExtRequest
	(*Query)(nil),       // 2: proto.Query
	(*Num)(nil),         // 3: proto.Num
	(*UserID)(nil),      // 4: proto.UserID
	(*ListRequest)(nil), // 5: proto.ListRequest
	(*LogArray)(nil),    // 6: proto.LogArray
	(*LogModel)(nil),    // 7: proto.LogModel
	(*Empty)(nil),       // 8: proto.Empty
}
var file_log_proto_depIdxs = []int32{
	7, // 0: proto.LogArray.logs:type_name -> proto.LogModel
	0, // 1: proto.LogModel.log_level:type_name -> proto.LogModel.Level
	7, // 2: proto.LogService.Log:input_type -> proto.LogModel
	1, // 3: proto.LogService.Exist:input_type -> proto.ExtRequest
	5, // 4: proto.LogService.GetUserMsgs:input_type -> proto.ListRequest
	4, // 5: proto.LogService.GetUserNewMsgNum:input_type -> proto.UserID
	4, // 6: proto.LogService.Count:input_type -> proto.UserID
	2, // 7: proto.LogService.CountQuery:input_type -> proto.Query
	2, // 8: proto.LogService.Delete:input_type -> proto.Query
	8, // 9: proto.LogService.Log:output_type -> proto.Empty
	7, // 10: proto.LogService.Exist:output_type -> proto.LogModel
	6, // 11: proto.LogService.GetUserMsgs:output_type -> proto.LogArray
	3, // 12: proto.LogService.GetUserNewMsgNum:output_type -> proto.Num
	3, // 13: proto.LogService.Count:output_type -> proto.Num
	3, // 14: proto.LogService.CountQuery:output_type -> proto.Num
	8, // 15: proto.LogService.Delete:output_type -> proto.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtRequest); i {
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
		file_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Num); i {
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
		file_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogArray); i {
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
		file_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogModel); i {
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
		file_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		EnumInfos:         file_log_proto_enumTypes,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	Log(ctx context.Context, in *LogModel, opts ...grpc.CallOption) (*Empty, error)
	Exist(ctx context.Context, in *ExtRequest, opts ...grpc.CallOption) (*LogModel, error)
	GetUserMsgs(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*LogArray, error)
	GetUserNewMsgNum(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Num, error)
	Count(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Num, error)
	CountQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Num, error)
	Delete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Empty, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Log(ctx context.Context, in *LogModel, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LogService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Exist(ctx context.Context, in *ExtRequest, opts ...grpc.CallOption) (*LogModel, error) {
	out := new(LogModel)
	err := c.cc.Invoke(ctx, "/proto.LogService/Exist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetUserMsgs(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*LogArray, error) {
	out := new(LogArray)
	err := c.cc.Invoke(ctx, "/proto.LogService/GetUserMsgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) GetUserNewMsgNum(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Num, error) {
	out := new(Num)
	err := c.cc.Invoke(ctx, "/proto.LogService/GetUserNewMsgNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Count(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Num, error) {
	out := new(Num)
	err := c.cc.Invoke(ctx, "/proto.LogService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) CountQuery(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Num, error) {
	out := new(Num)
	err := c.cc.Invoke(ctx, "/proto.LogService/CountQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Delete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LogService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	Log(context.Context, *LogModel) (*Empty, error)
	Exist(context.Context, *ExtRequest) (*LogModel, error)
	GetUserMsgs(context.Context, *ListRequest) (*LogArray, error)
	GetUserNewMsgNum(context.Context, *UserID) (*Num, error)
	Count(context.Context, *UserID) (*Num, error)
	CountQuery(context.Context, *Query) (*Num, error)
	Delete(context.Context, *Query) (*Empty, error)
}

// UnimplementedLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (*UnimplementedLogServiceServer) Log(context.Context, *LogModel) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (*UnimplementedLogServiceServer) Exist(context.Context, *ExtRequest) (*LogModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (*UnimplementedLogServiceServer) GetUserMsgs(context.Context, *ListRequest) (*LogArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMsgs not implemented")
}
func (*UnimplementedLogServiceServer) GetUserNewMsgNum(context.Context, *UserID) (*Num, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserNewMsgNum not implemented")
}
func (*UnimplementedLogServiceServer) Count(context.Context, *UserID) (*Num, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (*UnimplementedLogServiceServer) CountQuery(context.Context, *Query) (*Num, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountQuery not implemented")
}
func (*UnimplementedLogServiceServer) Delete(context.Context, *Query) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Log(ctx, req.(*LogModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/Exist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Exist(ctx, req.(*ExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetUserMsgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetUserMsgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/GetUserMsgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetUserMsgs(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_GetUserNewMsgNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).GetUserNewMsgNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/GetUserNewMsgNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).GetUserNewMsgNum(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Count(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_CountQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CountQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/CountQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CountQuery(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LogService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Delete(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LogService_Log_Handler,
		},
		{
			MethodName: "Exist",
			Handler:    _LogService_Exist_Handler,
		},
		{
			MethodName: "GetUserMsgs",
			Handler:    _LogService_GetUserMsgs_Handler,
		},
		{
			MethodName: "GetUserNewMsgNum",
			Handler:    _LogService_GetUserNewMsgNum_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _LogService_Count_Handler,
		},
		{
			MethodName: "CountQuery",
			Handler:    _LogService_CountQuery_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
