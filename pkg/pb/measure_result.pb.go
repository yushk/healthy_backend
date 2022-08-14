// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: measure_result.proto

package pb

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

type MeasureResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Userid string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid" bson:"userid"`
}

func (x *MeasureResult) Reset() {
	*x = MeasureResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasureResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasureResult) ProtoMessage() {}

func (x *MeasureResult) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasureResult.ProtoReflect.Descriptor instead.
func (*MeasureResult) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{0}
}

func (x *MeasureResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MeasureResult) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type CreateMeasureResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *MeasureResult `protobuf:"bytes,1,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *CreateMeasureResultRequest) Reset() {
	*x = CreateMeasureResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMeasureResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeasureResultRequest) ProtoMessage() {}

func (x *CreateMeasureResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeasureResultRequest.ProtoReflect.Descriptor instead.
func (*CreateMeasureResultRequest) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMeasureResultRequest) GetData() *MeasureResult {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMeasureResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *GetMeasureResultRequest) Reset() {
	*x = GetMeasureResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeasureResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasureResultRequest) ProtoMessage() {}

func (x *GetMeasureResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasureResultRequest.ProtoReflect.Descriptor instead.
func (*GetMeasureResultRequest) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{2}
}

func (x *GetMeasureResultRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateMeasureResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
	Data *MeasureResult `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *UpdateMeasureResultRequest) Reset() {
	*x = UpdateMeasureResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMeasureResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMeasureResultRequest) ProtoMessage() {}

func (x *UpdateMeasureResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMeasureResultRequest.ProtoReflect.Descriptor instead.
func (*UpdateMeasureResultRequest) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMeasureResultRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMeasureResultRequest) GetData() *MeasureResult {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteMeasureResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"id"`
}

func (x *DeleteMeasureResultRequest) Reset() {
	*x = DeleteMeasureResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMeasureResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMeasureResultRequest) ProtoMessage() {}

func (x *DeleteMeasureResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMeasureResultRequest.ProtoReflect.Descriptor instead.
func (*DeleteMeasureResultRequest) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMeasureResultRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMeasureResultsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query" bson:"query"`
	Skip  int64  `protobuf:"varint,2,opt,name=skip,proto3" json:"skip" bson:"skip"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit" bson:"limit"`
}

func (x *GetMeasureResultsRequest) Reset() {
	*x = GetMeasureResultsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeasureResultsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasureResultsRequest) ProtoMessage() {}

func (x *GetMeasureResultsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasureResultsRequest.ProtoReflect.Descriptor instead.
func (*GetMeasureResultsRequest) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{5}
}

func (x *GetMeasureResultsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetMeasureResultsRequest) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetMeasureResultsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMeasureResultsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64            `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count" bson:"total_count"`
	Items      []*MeasureResult `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GetMeasureResultsReply) Reset() {
	*x = GetMeasureResultsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_measure_result_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeasureResultsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeasureResultsReply) ProtoMessage() {}

func (x *GetMeasureResultsReply) ProtoReflect() protoreflect.Message {
	mi := &file_measure_result_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeasureResultsReply.ProtoReflect.Descriptor instead.
func (*GetMeasureResultsReply) Descriptor() ([]byte, []int) {
	return file_measure_result_proto_rawDescGZIP(), []int{6}
}

func (x *GetMeasureResultsReply) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetMeasureResultsReply) GetItems() []*MeasureResult {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_measure_result_proto protoreflect.FileDescriptor

var file_measure_result_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x37, 0x0a, 0x0d, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x91, 0x03, 0x0a, 0x14, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x4a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_measure_result_proto_rawDescOnce sync.Once
	file_measure_result_proto_rawDescData = file_measure_result_proto_rawDesc
)

func file_measure_result_proto_rawDescGZIP() []byte {
	file_measure_result_proto_rawDescOnce.Do(func() {
		file_measure_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_measure_result_proto_rawDescData)
	})
	return file_measure_result_proto_rawDescData
}

var file_measure_result_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_measure_result_proto_goTypes = []interface{}{
	(*MeasureResult)(nil),              // 0: pb.MeasureResult
	(*CreateMeasureResultRequest)(nil), // 1: pb.CreateMeasureResultRequest
	(*GetMeasureResultRequest)(nil),    // 2: pb.GetMeasureResultRequest
	(*UpdateMeasureResultRequest)(nil), // 3: pb.UpdateMeasureResultRequest
	(*DeleteMeasureResultRequest)(nil), // 4: pb.DeleteMeasureResultRequest
	(*GetMeasureResultsRequest)(nil),   // 5: pb.GetMeasureResultsRequest
	(*GetMeasureResultsReply)(nil),     // 6: pb.GetMeasureResultsReply
}
var file_measure_result_proto_depIdxs = []int32{
	0, // 0: pb.CreateMeasureResultRequest.data:type_name -> pb.MeasureResult
	0, // 1: pb.UpdateMeasureResultRequest.data:type_name -> pb.MeasureResult
	0, // 2: pb.GetMeasureResultsReply.items:type_name -> pb.MeasureResult
	1, // 3: pb.MeasureResultManager.CreateMeasureResult:input_type -> pb.CreateMeasureResultRequest
	2, // 4: pb.MeasureResultManager.GetMeasureResult:input_type -> pb.GetMeasureResultRequest
	3, // 5: pb.MeasureResultManager.UpdateMeasureResult:input_type -> pb.UpdateMeasureResultRequest
	4, // 6: pb.MeasureResultManager.DeleteMeasureResult:input_type -> pb.DeleteMeasureResultRequest
	5, // 7: pb.MeasureResultManager.GetMeasureResults:input_type -> pb.GetMeasureResultsRequest
	0, // 8: pb.MeasureResultManager.CreateMeasureResult:output_type -> pb.MeasureResult
	0, // 9: pb.MeasureResultManager.GetMeasureResult:output_type -> pb.MeasureResult
	0, // 10: pb.MeasureResultManager.UpdateMeasureResult:output_type -> pb.MeasureResult
	0, // 11: pb.MeasureResultManager.DeleteMeasureResult:output_type -> pb.MeasureResult
	6, // 12: pb.MeasureResultManager.GetMeasureResults:output_type -> pb.GetMeasureResultsReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_measure_result_proto_init() }
func file_measure_result_proto_init() {
	if File_measure_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_measure_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasureResult); i {
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
		file_measure_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMeasureResultRequest); i {
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
		file_measure_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeasureResultRequest); i {
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
		file_measure_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMeasureResultRequest); i {
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
		file_measure_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMeasureResultRequest); i {
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
		file_measure_result_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeasureResultsRequest); i {
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
		file_measure_result_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeasureResultsReply); i {
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
			RawDescriptor: file_measure_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_measure_result_proto_goTypes,
		DependencyIndexes: file_measure_result_proto_depIdxs,
		MessageInfos:      file_measure_result_proto_msgTypes,
	}.Build()
	File_measure_result_proto = out.File
	file_measure_result_proto_rawDesc = nil
	file_measure_result_proto_goTypes = nil
	file_measure_result_proto_depIdxs = nil
}
