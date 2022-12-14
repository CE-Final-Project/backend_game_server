// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: role_message.proto

package authService

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	CreatedBy uint64                 `protobuf:"varint,4,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	UpdatedBy uint64                 `protobuf:"varint,6,opt,name=UpdatedBy,proto3" json:"UpdatedBy,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetCreatedBy() uint64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *Role) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Role) GetUpdatedBy() uint64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

type CreateRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *CreateRoleReq) Reset() {
	*x = CreateRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleReq) ProtoMessage() {}

func (x *CreateRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleReq.ProtoReflect.Descriptor instead.
func (*CreateRoleReq) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRoleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRes) Reset() {
	*x = CreateRoleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRes) ProtoMessage() {}

func (x *CreateRoleRes) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRes.ProtoReflect.Descriptor instead.
func (*CreateRoleRes) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleRes) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type UpdateRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *UpdateRoleReq) Reset() {
	*x = UpdateRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleReq) ProtoMessage() {}

func (x *UpdateRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleReq.ProtoReflect.Descriptor instead.
func (*UpdateRoleReq) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRoleReq) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateRoleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UpdateRoleRes) Reset() {
	*x = UpdateRoleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRes) ProtoMessage() {}

func (x *UpdateRoleRes) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRes.ProtoReflect.Descriptor instead.
func (*UpdateRoleRes) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRoleRes) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type DeleteRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRoleReq) Reset() {
	*x = DeleteRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleReq) ProtoMessage() {}

func (x *DeleteRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleReq.ProtoReflect.Descriptor instead.
func (*DeleteRoleReq) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRoleReq) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteRoleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoleRes) Reset() {
	*x = DeleteRoleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRes) ProtoMessage() {}

func (x *DeleteRoleRes) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRes.ProtoReflect.Descriptor instead.
func (*DeleteRoleRes) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{6}
}

type SearchRolesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=Search,proto3" json:"Search,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size   int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SearchRolesReq) Reset() {
	*x = SearchRolesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRolesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRolesReq) ProtoMessage() {}

func (x *SearchRolesReq) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRolesReq.ProtoReflect.Descriptor instead.
func (*SearchRolesReq) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{7}
}

func (x *SearchRolesReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *SearchRolesReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRolesReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SearchRolesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64   `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	TotalPages int64   `protobuf:"varint,2,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Page       int64   `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size       int64   `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	HasMore    bool    `protobuf:"varint,5,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	Roles      []*Role `protobuf:"bytes,6,rep,name=Roles,proto3" json:"Roles,omitempty"`
}

func (x *SearchRolesRes) Reset() {
	*x = SearchRolesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRolesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRolesRes) ProtoMessage() {}

func (x *SearchRolesRes) ProtoReflect() protoreflect.Message {
	mi := &file_role_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRolesRes.ProtoReflect.Descriptor instead.
func (*SearchRolesRes) Descriptor() ([]byte, []int) {
	return file_role_message_proto_rawDescGZIP(), []int{8}
}

func (x *SearchRolesRes) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *SearchRolesRes) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *SearchRolesRes) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRolesRes) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchRolesRes) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *SearchRolesRes) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_role_message_proto protoreflect.FileDescriptor

var file_role_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_role_message_proto_rawDescOnce sync.Once
	file_role_message_proto_rawDescData = file_role_message_proto_rawDesc
)

func file_role_message_proto_rawDescGZIP() []byte {
	file_role_message_proto_rawDescOnce.Do(func() {
		file_role_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_message_proto_rawDescData)
	})
	return file_role_message_proto_rawDescData
}

var file_role_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_role_message_proto_goTypes = []interface{}{
	(*Role)(nil),                  // 0: authentication.Role
	(*CreateRoleReq)(nil),         // 1: authentication.CreateRoleReq
	(*CreateRoleRes)(nil),         // 2: authentication.CreateRoleRes
	(*UpdateRoleReq)(nil),         // 3: authentication.UpdateRoleReq
	(*UpdateRoleRes)(nil),         // 4: authentication.UpdateRoleRes
	(*DeleteRoleReq)(nil),         // 5: authentication.DeleteRoleReq
	(*DeleteRoleRes)(nil),         // 6: authentication.DeleteRoleRes
	(*SearchRolesReq)(nil),        // 7: authentication.SearchRolesReq
	(*SearchRolesRes)(nil),        // 8: authentication.SearchRolesRes
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_role_message_proto_depIdxs = []int32{
	9, // 0: authentication.Role.CreatedAt:type_name -> google.protobuf.Timestamp
	9, // 1: authentication.Role.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: authentication.CreateRoleRes.role:type_name -> authentication.Role
	0, // 3: authentication.UpdateRoleRes.role:type_name -> authentication.Role
	0, // 4: authentication.SearchRolesRes.Roles:type_name -> authentication.Role
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_role_message_proto_init() }
func file_role_message_proto_init() {
	if File_role_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_role_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleReq); i {
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
		file_role_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRes); i {
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
		file_role_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleReq); i {
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
		file_role_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRes); i {
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
		file_role_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleReq); i {
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
		file_role_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRes); i {
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
		file_role_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRolesReq); i {
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
		file_role_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRolesRes); i {
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
			RawDescriptor: file_role_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_message_proto_goTypes,
		DependencyIndexes: file_role_message_proto_depIdxs,
		MessageInfos:      file_role_message_proto_msgTypes,
	}.Build()
	File_role_message_proto = out.File
	file_role_message_proto_rawDesc = nil
	file_role_message_proto_goTypes = nil
	file_role_message_proto_depIdxs = nil
}
