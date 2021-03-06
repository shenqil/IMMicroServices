// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: demo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Demo 示例对象
type Demo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`               // 唯一标识
	Code      string                 `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`           // 编号
	Name      string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`           // 名称
	Memo      string                 `protobuf:"bytes,4,opt,name=Memo,proto3" json:"Memo,omitempty"`           // 备注
	Status    int32                  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`      // 状态(1:启用 2:停用)
	Creator   string                 `protobuf:"bytes,6,opt,name=Creator,proto3" json:"Creator,omitempty"`     // 创建者
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"` // 创建时间
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"` // 更新时间
}

func (x *Demo) Reset() {
	*x = Demo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Demo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Demo) ProtoMessage() {}

func (x *Demo) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Demo.ProtoReflect.Descriptor instead.
func (*Demo) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Demo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Demo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Demo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Demo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Demo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Demo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Demo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Demo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// DemoQueryParam 查询条件
type DemoQueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationParam *PaginationParam `protobuf:"bytes,1,opt,name=PaginationParam,proto3" json:"PaginationParam,omitempty"`
	Code            string           `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`             // 编号
	QueryValue      string           `protobuf:"bytes,3,opt,name=QueryValue,proto3" json:"QueryValue,omitempty"` // 查询值
}

func (x *DemoQueryParam) Reset() {
	*x = DemoQueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoQueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoQueryParam) ProtoMessage() {}

func (x *DemoQueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoQueryParam.ProtoReflect.Descriptor instead.
func (*DemoQueryParam) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *DemoQueryParam) GetPaginationParam() *PaginationParam {
	if x != nil {
		return x.PaginationParam
	}
	return nil
}

func (x *DemoQueryParam) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DemoQueryParam) GetQueryValue() string {
	if x != nil {
		return x.QueryValue
	}
	return ""
}

// DemoQueryOptions 示例对象查询可选参数项
type DemoQueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderFields []*OrderField `protobuf:"bytes,1,rep,name=OrderFields,proto3" json:"OrderFields,omitempty"` // 排序字段
}

func (x *DemoQueryOptions) Reset() {
	*x = DemoQueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoQueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoQueryOptions) ProtoMessage() {}

func (x *DemoQueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoQueryOptions.ProtoReflect.Descriptor instead.
func (*DemoQueryOptions) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *DemoQueryOptions) GetOrderFields() []*OrderField {
	if x != nil {
		return x.OrderFields
	}
	return nil
}

type DemoQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params  *DemoQueryParam   `protobuf:"bytes,1,opt,name=Params,proto3" json:"Params,omitempty"`
	Options *DemoQueryOptions `protobuf:"bytes,2,opt,name=Options,proto3" json:"Options,omitempty"`
}

func (x *DemoQueryRequest) Reset() {
	*x = DemoQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoQueryRequest) ProtoMessage() {}

func (x *DemoQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoQueryRequest.ProtoReflect.Descriptor instead.
func (*DemoQueryRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *DemoQueryRequest) GetParams() *DemoQueryParam {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *DemoQueryRequest) GetOptions() *DemoQueryOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type DemoGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DemoGetRequest) Reset() {
	*x = DemoGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoGetRequest) ProtoMessage() {}

func (x *DemoGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoGetRequest.ProtoReflect.Descriptor instead.
func (*DemoGetRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{4}
}

func (x *DemoGetRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DemoDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DemoDeleteRequest) Reset() {
	*x = DemoDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoDeleteRequest) ProtoMessage() {}

func (x *DemoDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoDeleteRequest.ProtoReflect.Descriptor instead.
func (*DemoDeleteRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{5}
}

func (x *DemoDeleteRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DemoUpdateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status uint32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DemoUpdateStatusRequest) Reset() {
	*x = DemoUpdateStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoUpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoUpdateStatusRequest) ProtoMessage() {}

func (x *DemoUpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoUpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*DemoUpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{6}
}

func (x *DemoUpdateStatusRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DemoUpdateStatusRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// DemoQueryResult 示例对象查询结果
type DemoQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*Demo           `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	PageResult *PaginationResult `protobuf:"bytes,2,opt,name=PageResult,proto3" json:"PageResult,omitempty"`
}

func (x *DemoQueryResult) Reset() {
	*x = DemoQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoQueryResult) ProtoMessage() {}

func (x *DemoQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoQueryResult.ProtoReflect.Descriptor instead.
func (*DemoQueryResult) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{7}
}

func (x *DemoQueryResult) GetData() []*Demo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DemoQueryResult) GetPageResult() *PaginationResult {
	if x != nil {
		return x.PageResult
	}
	return nil
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65,
	0x6d, 0x6f, 0x50, 0x42, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x89,
	0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6d, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x43, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x10, 0x44, 0x65,
	0x6d, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x76, 0x0a, 0x10, 0x44, 0x65, 0x6d, 0x6f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x20,
	0x0a, 0x0e, 0x44, 0x65, 0x6d, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6d, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x41, 0x0a, 0x17, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x0f, 0x44, 0x65, 0x6d, 0x6f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xdc, 0x02, 0x0a, 0x0b, 0x44, 0x65,
	0x6d, 0x6f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e,
	0x44, 0x65, 0x6d, 0x6f, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x1a, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x42, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x50, 0x42, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x42, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_demo_proto_goTypes = []interface{}{
	(*Demo)(nil),                    // 0: demoPB.Demo
	(*DemoQueryParam)(nil),          // 1: demoPB.DemoQueryParam
	(*DemoQueryOptions)(nil),        // 2: demoPB.DemoQueryOptions
	(*DemoQueryRequest)(nil),        // 3: demoPB.DemoQueryRequest
	(*DemoGetRequest)(nil),          // 4: demoPB.DemoGetRequest
	(*DemoDeleteRequest)(nil),       // 5: demoPB.DemoDeleteRequest
	(*DemoUpdateStatusRequest)(nil), // 6: demoPB.DemoUpdateStatusRequest
	(*DemoQueryResult)(nil),         // 7: demoPB.DemoQueryResult
	(*timestamppb.Timestamp)(nil),   // 8: google.protobuf.Timestamp
	(*PaginationParam)(nil),         // 9: commonPB.PaginationParam
	(*OrderField)(nil),              // 10: commonPB.OrderField
	(*PaginationResult)(nil),        // 11: commonPB.PaginationResult
	(*IDResult)(nil),                // 12: commonPB.IDResult
	(*StatusResult)(nil),            // 13: commonPB.StatusResult
}
var file_demo_proto_depIdxs = []int32{
	8,  // 0: demoPB.Demo.CreatedAt:type_name -> google.protobuf.Timestamp
	8,  // 1: demoPB.Demo.UpdatedAt:type_name -> google.protobuf.Timestamp
	9,  // 2: demoPB.DemoQueryParam.PaginationParam:type_name -> commonPB.PaginationParam
	10, // 3: demoPB.DemoQueryOptions.OrderFields:type_name -> commonPB.OrderField
	1,  // 4: demoPB.DemoQueryRequest.Params:type_name -> demoPB.DemoQueryParam
	2,  // 5: demoPB.DemoQueryRequest.Options:type_name -> demoPB.DemoQueryOptions
	0,  // 6: demoPB.DemoQueryResult.Data:type_name -> demoPB.Demo
	11, // 7: demoPB.DemoQueryResult.PageResult:type_name -> commonPB.PaginationResult
	3,  // 8: demoPB.DemoGreeter.Query:input_type -> demoPB.DemoQueryRequest
	4,  // 9: demoPB.DemoGreeter.Get:input_type -> demoPB.DemoGetRequest
	0,  // 10: demoPB.DemoGreeter.Create:input_type -> demoPB.Demo
	0,  // 11: demoPB.DemoGreeter.Update:input_type -> demoPB.Demo
	5,  // 12: demoPB.DemoGreeter.Delete:input_type -> demoPB.DemoDeleteRequest
	6,  // 13: demoPB.DemoGreeter.UpdateStatus:input_type -> demoPB.DemoUpdateStatusRequest
	7,  // 14: demoPB.DemoGreeter.Query:output_type -> demoPB.DemoQueryResult
	0,  // 15: demoPB.DemoGreeter.Get:output_type -> demoPB.Demo
	12, // 16: demoPB.DemoGreeter.Create:output_type -> commonPB.IDResult
	13, // 17: demoPB.DemoGreeter.Update:output_type -> commonPB.StatusResult
	13, // 18: demoPB.DemoGreeter.Delete:output_type -> commonPB.StatusResult
	13, // 19: demoPB.DemoGreeter.UpdateStatus:output_type -> commonPB.StatusResult
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Demo); i {
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
		file_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoQueryParam); i {
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
		file_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoQueryOptions); i {
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
		file_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoQueryRequest); i {
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
		file_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoGetRequest); i {
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
		file_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoDeleteRequest); i {
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
		file_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoUpdateStatusRequest); i {
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
		file_demo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoQueryResult); i {
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
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}
