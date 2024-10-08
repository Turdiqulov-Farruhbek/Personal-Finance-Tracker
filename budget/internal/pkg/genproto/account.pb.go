// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: account.proto

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

type AccountCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Currency string `protobuf:"bytes,4,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *AccountCreate) Reset() {
	*x = AccountCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreate) ProtoMessage() {}

func (x *AccountCreate) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreate.ProtoReflect.Descriptor instead.
func (*AccountCreate) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AccountCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountCreate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AccountCreate) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type AccountGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId    string  `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Name      string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type      string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Currency  string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Balance   float32 `protobuf:"fixed32,6,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt string  `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string  `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *AccountGet) Reset() {
	*x = AccountGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGet) ProtoMessage() {}

func (x *AccountGet) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGet.ProtoReflect.Descriptor instead.
func (*AccountGet) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountGet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountGet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AccountGet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountGet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AccountGet) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AccountGet) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *AccountGet) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AccountGet) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AccountUpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Balance  float32 `protobuf:"fixed32,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AccountUpt) Reset() {
	*x = AccountUpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpt) ProtoMessage() {}

func (x *AccountUpt) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpt.ProtoReflect.Descriptor instead.
func (*AccountUpt) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountUpt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountUpt) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AccountUpt) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AccountUpt) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type AccountUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Body *AccountUpt `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *AccountUpdate) Reset() {
	*x = AccountUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdate) ProtoMessage() {}

func (x *AccountUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdate.ProtoReflect.Descriptor instead.
func (*AccountUpdate) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountUpdate) GetBody() *AccountUpt {
	if x != nil {
		return x.Body
	}
	return nil
}

type AccountFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type       string  `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Currency   string  `protobuf:"bytes,3,opt,name=Currency,proto3" json:"Currency,omitempty"`
	BalanceMin float32 `protobuf:"fixed32,4,opt,name=BalanceMin,proto3" json:"BalanceMin,omitempty"`
	BalanceMax float32 `protobuf:"fixed32,5,opt,name=BalanceMax,proto3" json:"BalanceMax,omitempty"`
	Filter     *Filter `protobuf:"bytes,6,opt,name=Filter,proto3" json:"Filter,omitempty"`
	UserId     string  `protobuf:"bytes,7,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *AccountFilter) Reset() {
	*x = AccountFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountFilter) ProtoMessage() {}

func (x *AccountFilter) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountFilter.ProtoReflect.Descriptor instead.
func (*AccountFilter) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *AccountFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountFilter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AccountFilter) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *AccountFilter) GetBalanceMin() float32 {
	if x != nil {
		return x.BalanceMin
	}
	return 0
}

func (x *AccountFilter) GetBalanceMax() float32 {
	if x != nil {
		return x.BalanceMax
	}
	return 0
}

func (x *AccountFilter) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *AccountFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AccounList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts   []*AccountGet `protobuf:"bytes,1,rep,name=Accounts,proto3" json:"Accounts,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	Limit      int32         `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset     int32         `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *AccounList) Reset() {
	*x = AccounList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccounList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccounList) ProtoMessage() {}

func (x *AccounList) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccounList.ProtoReflect.Descriptor instead.
func (*AccounList) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *AccounList) GetAccounts() []*AccountGet {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *AccounList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *AccounList) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AccounList) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xce, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6a, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x4a, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x74, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xd6,
	0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x4d, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d,
	0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x4d, 0x61, 0x78, 0x12, 0x29, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52,
	0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xb2, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73, 0x75, 0x62,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x75,
	0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x18, 0x5a, 0x16,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_account_proto_goTypes = []interface{}{
	(*AccountCreate)(nil), // 0: submodule.AccountCreate
	(*AccountGet)(nil),    // 1: submodule.AccountGet
	(*AccountUpt)(nil),    // 2: submodule.AccountUpt
	(*AccountUpdate)(nil), // 3: submodule.AccountUpdate
	(*AccountFilter)(nil), // 4: submodule.AccountFilter
	(*AccounList)(nil),    // 5: submodule.AccounList
	(*Filter)(nil),        // 6: submodule.Filter
	(*ById)(nil),          // 7: submodule.ById
	(*Void)(nil),          // 8: submodule.Void
}
var file_account_proto_depIdxs = []int32{
	2, // 0: submodule.AccountUpdate.Body:type_name -> submodule.AccountUpt
	6, // 1: submodule.AccountFilter.Filter:type_name -> submodule.Filter
	1, // 2: submodule.AccounList.Accounts:type_name -> submodule.AccountGet
	0, // 3: submodule.AccountService.CreateAccount:input_type -> submodule.AccountCreate
	7, // 4: submodule.AccountService.GetAccount:input_type -> submodule.ById
	3, // 5: submodule.AccountService.UpdateAccount:input_type -> submodule.AccountUpdate
	7, // 6: submodule.AccountService.DeleteAccount:input_type -> submodule.ById
	4, // 7: submodule.AccountService.ListAccounts:input_type -> submodule.AccountFilter
	8, // 8: submodule.AccountService.CreateAccount:output_type -> submodule.Void
	1, // 9: submodule.AccountService.GetAccount:output_type -> submodule.AccountGet
	8, // 10: submodule.AccountService.UpdateAccount:output_type -> submodule.Void
	8, // 11: submodule.AccountService.DeleteAccount:output_type -> submodule.Void
	5, // 12: submodule.AccountService.ListAccounts:output_type -> submodule.AccounList
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreate); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountGet); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpt); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpdate); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountFilter); i {
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
		file_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccounList); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
