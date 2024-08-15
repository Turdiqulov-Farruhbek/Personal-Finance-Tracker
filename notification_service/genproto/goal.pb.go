// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: goal.proto

package genproto

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

type GoalCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string  `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,3,opt,name=TargetAmount,proto3" json:"TargetAmount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,4,opt,name=CurrentAmount,proto3" json:"CurrentAmount,omitempty"`
	Deadline      string  `protobuf:"bytes,5,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
}

func (x *GoalCreate) Reset() {
	*x = GoalCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalCreate) ProtoMessage() {}

func (x *GoalCreate) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalCreate.ProtoReflect.Descriptor instead.
func (*GoalCreate) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{0}
}

func (x *GoalCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoalCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalCreate) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalCreate) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalCreate) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

type GoalUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,2,opt,name=TargetAmount,proto3" json:"TargetAmount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,3,opt,name=CurrentAmount,proto3" json:"CurrentAmount,omitempty"`
	Deadline      string  `protobuf:"bytes,4,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	Status        string  `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GoalUpt) Reset() {
	*x = GoalUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalUpt) ProtoMessage() {}

func (x *GoalUpt) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalUpt.ProtoReflect.Descriptor instead.
func (*GoalUpt) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{1}
}

func (x *GoalUpt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalUpt) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalUpt) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalUpt) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GoalUpt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GoalUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body *GoalUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *GoalUpdate) Reset() {
	*x = GoalUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalUpdate) ProtoMessage() {}

func (x *GoalUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalUpdate.ProtoReflect.Descriptor instead.
func (*GoalUpdate) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{2}
}

func (x *GoalUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GoalUpdate) GetBody() *GoalUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type GoalGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=TargetAmount,proto3" json:"TargetAmount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=CurrentAmount,proto3" json:"CurrentAmount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt     string  `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *GoalGet) Reset() {
	*x = GoalGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalGet) ProtoMessage() {}

func (x *GoalGet) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalGet.ProtoReflect.Descriptor instead.
func (*GoalGet) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{3}
}

func (x *GoalGet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GoalGet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoalGet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalGet) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalGet) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalGet) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GoalGet) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GoalGet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GoalGet) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GoalFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string  `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Status       string  `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	TargetFrom   float32 `protobuf:"fixed32,4,opt,name=TargetFrom,proto3" json:"TargetFrom,omitempty"`
	TargetTo     float32 `protobuf:"fixed32,5,opt,name=TargetTo,proto3" json:"TargetTo,omitempty"`
	DeadlineFrom string  `protobuf:"bytes,6,opt,name=DeadlineFrom,proto3" json:"DeadlineFrom,omitempty"`
	DeadlineTo   string  `protobuf:"bytes,7,opt,name=DeadlineTo,proto3" json:"DeadlineTo,omitempty"`
	Filter       *Filter `protobuf:"bytes,8,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *GoalFilter) Reset() {
	*x = GoalFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalFilter) ProtoMessage() {}

func (x *GoalFilter) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalFilter.ProtoReflect.Descriptor instead.
func (*GoalFilter) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{4}
}

func (x *GoalFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoalFilter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GoalFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalFilter) GetTargetFrom() float32 {
	if x != nil {
		return x.TargetFrom
	}
	return 0
}

func (x *GoalFilter) GetTargetTo() float32 {
	if x != nil {
		return x.TargetTo
	}
	return 0
}

func (x *GoalFilter) GetDeadlineFrom() string {
	if x != nil {
		return x.DeadlineFrom
	}
	return ""
}

func (x *GoalFilter) GetDeadlineTo() string {
	if x != nil {
		return x.DeadlineTo
	}
	return ""
}

func (x *GoalFilter) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GoalList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goals      []*GoalGet `protobuf:"bytes,1,rep,name=Goals,proto3" json:"Goals,omitempty"`
	TotalCount int32      `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	Limit      int32      `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset     int32      `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *GoalList) Reset() {
	*x = GoalList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalList) ProtoMessage() {}

func (x *GoalList) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalList.ProtoReflect.Descriptor instead.
func (*GoalList) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{5}
}

func (x *GoalList) GetGoals() []*GoalGet {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *GoalList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GoalList) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GoalList) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_goal_proto protoreflect.FileDescriptor

var file_goal_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x47, 0x6f, 0x61, 0x6c, 0x55,
	0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x44, 0x0a, 0x0a, 0x47, 0x6f, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61,
	0x6c, 0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xff, 0x01, 0x0a, 0x07, 0x47,
	0x6f, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xfb, 0x01, 0x0a,
	0x0a, 0x47, 0x6f, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x12,
	0x29, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x47,
	0x6f, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x47, 0x6f, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x05, 0x47, 0x6f, 0x61, 0x6c,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32,
	0x92, 0x02, 0x0a, 0x0b, 0x47, 0x6f, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x15, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x12, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x47, 0x6f, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x13, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x12,
	0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x12, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6f, 0x61,
	0x6c, 0x47, 0x65, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goal_proto_rawDescOnce sync.Once
	file_goal_proto_rawDescData = file_goal_proto_rawDesc
)

func file_goal_proto_rawDescGZIP() []byte {
	file_goal_proto_rawDescOnce.Do(func() {
		file_goal_proto_rawDescData = protoimpl.X.CompressGZIP(file_goal_proto_rawDescData)
	})
	return file_goal_proto_rawDescData
}

var file_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goal_proto_goTypes = []interface{}{
	(*GoalCreate)(nil), // 0: submodule.GoalCreate
	(*GoalUpt)(nil),    // 1: submodule.GoalUpt
	(*GoalUpdate)(nil), // 2: submodule.GoalUpdate
	(*GoalGet)(nil),    // 3: submodule.GoalGet
	(*GoalFilter)(nil), // 4: submodule.GoalFilter
	(*GoalList)(nil),   // 5: submodule.GoalList
	(*Filter)(nil),     // 6: submodule.Filter
	(*ById)(nil),       // 7: submodule.ById
	(*Void)(nil),       // 8: submodule.Void
}
var file_goal_proto_depIdxs = []int32{
	1, // 0: submodule.GoalUpdate.Body:type_name -> submodule.GoalUpt
	6, // 1: submodule.GoalFilter.Filter:type_name -> submodule.Filter
	3, // 2: submodule.GoalList.Goals:type_name -> submodule.GoalGet
	0, // 3: submodule.GoalService.CreateGoal:input_type -> submodule.GoalCreate
	2, // 4: submodule.GoalService.UpdateGoal:input_type -> submodule.GoalUpdate
	7, // 5: submodule.GoalService.DeleteGoal:input_type -> submodule.ById
	4, // 6: submodule.GoalService.ListGoals:input_type -> submodule.GoalFilter
	7, // 7: submodule.GoalService.GetGoal:input_type -> submodule.ById
	8, // 8: submodule.GoalService.CreateGoal:output_type -> submodule.Void
	8, // 9: submodule.GoalService.UpdateGoal:output_type -> submodule.Void
	8, // 10: submodule.GoalService.DeleteGoal:output_type -> submodule.Void
	5, // 11: submodule.GoalService.ListGoals:output_type -> submodule.GoalList
	3, // 12: submodule.GoalService.GetGoal:output_type -> submodule.GoalGet
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_goal_proto_init() }
func file_goal_proto_init() {
	if File_goal_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_goal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalCreate); i {
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
		file_goal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalUpt); i {
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
		file_goal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalUpdate); i {
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
		file_goal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalGet); i {
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
		file_goal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalFilter); i {
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
		file_goal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalList); i {
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
			RawDescriptor: file_goal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goal_proto_goTypes,
		DependencyIndexes: file_goal_proto_depIdxs,
		MessageInfos:      file_goal_proto_msgTypes,
	}.Build()
	File_goal_proto = out.File
	file_goal_proto_rawDesc = nil
	file_goal_proto_goTypes = nil
	file_goal_proto_depIdxs = nil
}