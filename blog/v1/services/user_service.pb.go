// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: blog/v1/services/user_service.proto

package services

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	resources "github.com/sei-ri/protoapis/blog/v1/resources"
	resources1 "github.com/sei-ri/protoapis/news/v1/resources"
	protobuf "github.com/sei-ri/protoapis/protobuf"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_services_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_services_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_blog_v1_services_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourcsName string `protobuf:"bytes,1,opt,name=resourcs_name,json=resourcsName,proto3" json:"resourcs_name,omitempty"`
	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_services_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_services_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_blog_v1_services_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetResourcsName() string {
	if x != nil {
		return x.ResourcsName
	}
	return ""
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links *protobuf.SelfLinks `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
	// Types that are assignable to User:
	//	*GetUserResponse_Blog
	//	*GetUserResponse_News
	User isGetUserResponse_User `protobuf_oneof:"user"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_services_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_services_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_blog_v1_services_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserResponse) GetLinks() *protobuf.SelfLinks {
	if x != nil {
		return x.Links
	}
	return nil
}

func (m *GetUserResponse) GetUser() isGetUserResponse_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (x *GetUserResponse) GetBlog() *resources.User {
	if x, ok := x.GetUser().(*GetUserResponse_Blog); ok {
		return x.Blog
	}
	return nil
}

func (x *GetUserResponse) GetNews() *resources1.User {
	if x, ok := x.GetUser().(*GetUserResponse_News); ok {
		return x.News
	}
	return nil
}

type isGetUserResponse_User interface {
	isGetUserResponse_User()
}

type GetUserResponse_Blog struct {
	Blog *resources.User `protobuf:"bytes,2,opt,name=blog,proto3,oneof"`
}

type GetUserResponse_News struct {
	News *resources1.User `protobuf:"bytes,3,opt,name=news,proto3,oneof"`
}

func (*GetUserResponse_Blog) isGetUserResponse_User() {}

func (*GetUserResponse_News) isGetUserResponse_User() {}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourcsName string `protobuf:"bytes,1,opt,name=resourcs_name,json=resourcsName,proto3" json:"resourcs_name,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_services_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_services_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_blog_v1_services_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersRequest) GetResourcsName() string {
	if x != nil {
		return x.ResourcsName
	}
	return ""
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links *protobuf.SelfLinks `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
	Users []*GetUserResponse  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_services_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_services_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_blog_v1_services_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListUsersResponse) GetLinks() *protobuf.SelfLinks {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *ListUsersResponse) GetUsers() []*GetUserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_blog_v1_services_user_service_proto protoreflect.FileDescriptor

var file_blog_v1_services_user_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x65, 0x6c, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x77, 0x64,
	0x22, 0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x72, 0x0c,
	0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73,
	0x42, 0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x72, 0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67,
	0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xe1, 0x02,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6c, 0x66,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x76, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x77, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x65, 0x69, 0x2d, 0x72, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_blog_v1_services_user_service_proto_rawDescOnce sync.Once
	file_blog_v1_services_user_service_proto_rawDescData = file_blog_v1_services_user_service_proto_rawDesc
)

func file_blog_v1_services_user_service_proto_rawDescGZIP() []byte {
	file_blog_v1_services_user_service_proto_rawDescOnce.Do(func() {
		file_blog_v1_services_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_services_user_service_proto_rawDescData)
	})
	return file_blog_v1_services_user_service_proto_rawDescData
}

var file_blog_v1_services_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_blog_v1_services_user_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),  // 0: blog.v1.services.CreateUserRequest
	(*GetUserRequest)(nil),     // 1: blog.v1.services.GetUserRequest
	(*GetUserResponse)(nil),    // 2: blog.v1.services.GetUserResponse
	(*ListUsersRequest)(nil),   // 3: blog.v1.services.ListUsersRequest
	(*ListUsersResponse)(nil),  // 4: blog.v1.services.ListUsersResponse
	(*protobuf.SelfLinks)(nil), // 5: protobuf.SelfLinks
	(*resources.User)(nil),     // 6: blog.v1.resources.User
	(*resources1.User)(nil),    // 7: news.v1.resources.User
}
var file_blog_v1_services_user_service_proto_depIdxs = []int32{
	5, // 0: blog.v1.services.GetUserResponse.links:type_name -> protobuf.SelfLinks
	6, // 1: blog.v1.services.GetUserResponse.blog:type_name -> blog.v1.resources.User
	7, // 2: blog.v1.services.GetUserResponse.news:type_name -> news.v1.resources.User
	5, // 3: blog.v1.services.ListUsersResponse.links:type_name -> protobuf.SelfLinks
	2, // 4: blog.v1.services.ListUsersResponse.users:type_name -> blog.v1.services.GetUserResponse
	0, // 5: blog.v1.services.UserService.CreateUser:input_type -> blog.v1.services.CreateUserRequest
	1, // 6: blog.v1.services.UserService.GetUser:input_type -> blog.v1.services.GetUserRequest
	3, // 7: blog.v1.services.UserService.ListUsers:input_type -> blog.v1.services.ListUsersRequest
	5, // 8: blog.v1.services.UserService.CreateUser:output_type -> protobuf.SelfLinks
	2, // 9: blog.v1.services.UserService.GetUser:output_type -> blog.v1.services.GetUserResponse
	4, // 10: blog.v1.services.UserService.ListUsers:output_type -> blog.v1.services.ListUsersResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_blog_v1_services_user_service_proto_init() }
func file_blog_v1_services_user_service_proto_init() {
	if File_blog_v1_services_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_v1_services_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_blog_v1_services_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_blog_v1_services_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_blog_v1_services_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_blog_v1_services_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
	file_blog_v1_services_user_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GetUserResponse_Blog)(nil),
		(*GetUserResponse_News)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_v1_services_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_v1_services_user_service_proto_goTypes,
		DependencyIndexes: file_blog_v1_services_user_service_proto_depIdxs,
		MessageInfos:      file_blog_v1_services_user_service_proto_msgTypes,
	}.Build()
	File_blog_v1_services_user_service_proto = out.File
	file_blog_v1_services_user_service_proto_rawDesc = nil
	file_blog_v1_services_user_service_proto_goTypes = nil
	file_blog_v1_services_user_service_proto_depIdxs = nil
}
