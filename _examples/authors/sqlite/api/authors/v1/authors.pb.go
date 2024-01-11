// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: authors/v1/authors.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bio  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetBio() *wrapperspb.StringValue {
	if x != nil {
		return x.Bio
	}
	return nil
}

type CreateAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bio  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=bio,proto3" json:"bio,omitempty"`
}

func (x *CreateAuthorRequest) Reset() {
	*x = CreateAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorRequest) ProtoMessage() {}

func (x *CreateAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAuthorRequest) GetBio() *wrapperspb.StringValue {
	if x != nil {
		return x.Bio
	}
	return nil
}

type CreateAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ExecResult `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateAuthorResponse) Reset() {
	*x = CreateAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthorResponse) ProtoMessage() {}

func (x *CreateAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthorResponse.ProtoReflect.Descriptor instead.
func (*CreateAuthorResponse) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthorResponse) GetValue() *ExecResult {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAuthorRequest) Reset() {
	*x = DeleteAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorRequest) ProtoMessage() {}

func (x *DeleteAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAuthorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAuthorResponse) Reset() {
	*x = DeleteAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthorResponse) ProtoMessage() {}

func (x *DeleteAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthorResponse.ProtoReflect.Descriptor instead.
func (*DeleteAuthorResponse) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{4}
}

type GetAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuthorRequest) Reset() {
	*x = GetAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorRequest) ProtoMessage() {}

func (x *GetAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{5}
}

func (x *GetAuthorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorResponse) Reset() {
	*x = GetAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorResponse) ProtoMessage() {}

func (x *GetAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorResponse) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type ListAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAuthorsRequest) Reset() {
	*x = ListAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsRequest) ProtoMessage() {}

func (x *ListAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsRequest.ProtoReflect.Descriptor instead.
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{7}
}

type ListAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Author `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListAuthorsResponse) Reset() {
	*x = ListAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthorsResponse) ProtoMessage() {}

func (x *ListAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthorsResponse.ProtoReflect.Descriptor instead.
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{8}
}

func (x *ListAuthorsResponse) GetList() []*Author {
	if x != nil {
		return x.List
	}
	return nil
}

type ExecResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64 `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
	LastInsertId int64 `protobuf:"varint,2,opt,name=lastInsertId,proto3" json:"lastInsertId,omitempty"`
}

func (x *ExecResult) Reset() {
	*x = ExecResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authors_v1_authors_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResult) ProtoMessage() {}

func (x *ExecResult) ProtoReflect() protoreflect.Message {
	mi := &file_authors_v1_authors_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResult.ProtoReflect.Descriptor instead.
func (*ExecResult) Descriptor() ([]byte, []int) {
	return file_authors_v1_authors_proto_rawDescGZIP(), []int{9}
}

func (x *ExecResult) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *ExecResult) GetLastInsertId() int64 {
	if x != nil {
		return x.LastInsertId
	}
	return 0
}

var File_authors_v1_authors_proto protoreflect.FileDescriptor

var file_authors_v1_authors_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x22, 0x59, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x03, 0x62, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x22, 0x44, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x0a,
	0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x49, 0x64, 0x32, 0xb0, 0x03, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x67, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x62, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x0c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x62, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x08, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0xee, 0x01, 0x92, 0x41, 0xd2, 0x01, 0x12, 0xcf, 0x01, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x83, 0x01, 0x42, 0x6f, 0x69, 0x6c, 0x65,
	0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x2a, 0x2a, 0x73, 0x71, 0x6c, 0x63, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2a, 0x2a, 0x2e, 0x20, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x20, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x75, 0x6e, 0x20, 0x60, 0x62,
	0x75, 0x66, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x60, 0x20, 0x74, 0x6f, 0x20,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x22, 0x39,
	0x0a, 0x09, 0x73, 0x71, 0x6c, 0x63, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x12, 0x2c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x77, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x77, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x65, 0x79, 0x2f,
	0x73, 0x71, 0x6c, 0x63, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x16,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authors_v1_authors_proto_rawDescOnce sync.Once
	file_authors_v1_authors_proto_rawDescData = file_authors_v1_authors_proto_rawDesc
)

func file_authors_v1_authors_proto_rawDescGZIP() []byte {
	file_authors_v1_authors_proto_rawDescOnce.Do(func() {
		file_authors_v1_authors_proto_rawDescData = protoimpl.X.CompressGZIP(file_authors_v1_authors_proto_rawDescData)
	})
	return file_authors_v1_authors_proto_rawDescData
}

var file_authors_v1_authors_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_authors_v1_authors_proto_goTypes = []interface{}{
	(*Author)(nil),                 // 0: authors.v1.Author
	(*CreateAuthorRequest)(nil),    // 1: authors.v1.CreateAuthorRequest
	(*CreateAuthorResponse)(nil),   // 2: authors.v1.CreateAuthorResponse
	(*DeleteAuthorRequest)(nil),    // 3: authors.v1.DeleteAuthorRequest
	(*DeleteAuthorResponse)(nil),   // 4: authors.v1.DeleteAuthorResponse
	(*GetAuthorRequest)(nil),       // 5: authors.v1.GetAuthorRequest
	(*GetAuthorResponse)(nil),      // 6: authors.v1.GetAuthorResponse
	(*ListAuthorsRequest)(nil),     // 7: authors.v1.ListAuthorsRequest
	(*ListAuthorsResponse)(nil),    // 8: authors.v1.ListAuthorsResponse
	(*ExecResult)(nil),             // 9: authors.v1.ExecResult
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
}
var file_authors_v1_authors_proto_depIdxs = []int32{
	10, // 0: authors.v1.Author.bio:type_name -> google.protobuf.StringValue
	10, // 1: authors.v1.CreateAuthorRequest.bio:type_name -> google.protobuf.StringValue
	9,  // 2: authors.v1.CreateAuthorResponse.value:type_name -> authors.v1.ExecResult
	0,  // 3: authors.v1.GetAuthorResponse.author:type_name -> authors.v1.Author
	0,  // 4: authors.v1.ListAuthorsResponse.list:type_name -> authors.v1.Author
	1,  // 5: authors.v1.AuthorsService.CreateAuthor:input_type -> authors.v1.CreateAuthorRequest
	3,  // 6: authors.v1.AuthorsService.DeleteAuthor:input_type -> authors.v1.DeleteAuthorRequest
	5,  // 7: authors.v1.AuthorsService.GetAuthor:input_type -> authors.v1.GetAuthorRequest
	7,  // 8: authors.v1.AuthorsService.ListAuthors:input_type -> authors.v1.ListAuthorsRequest
	2,  // 9: authors.v1.AuthorsService.CreateAuthor:output_type -> authors.v1.CreateAuthorResponse
	4,  // 10: authors.v1.AuthorsService.DeleteAuthor:output_type -> authors.v1.DeleteAuthorResponse
	6,  // 11: authors.v1.AuthorsService.GetAuthor:output_type -> authors.v1.GetAuthorResponse
	8,  // 12: authors.v1.AuthorsService.ListAuthors:output_type -> authors.v1.ListAuthorsResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_authors_v1_authors_proto_init() }
func file_authors_v1_authors_proto_init() {
	if File_authors_v1_authors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authors_v1_authors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_authors_v1_authors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorRequest); i {
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
		file_authors_v1_authors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthorResponse); i {
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
		file_authors_v1_authors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorRequest); i {
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
		file_authors_v1_authors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAuthorResponse); i {
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
		file_authors_v1_authors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorRequest); i {
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
		file_authors_v1_authors_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorResponse); i {
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
		file_authors_v1_authors_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsRequest); i {
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
		file_authors_v1_authors_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthorsResponse); i {
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
		file_authors_v1_authors_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResult); i {
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
			RawDescriptor: file_authors_v1_authors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authors_v1_authors_proto_goTypes,
		DependencyIndexes: file_authors_v1_authors_proto_depIdxs,
		MessageInfos:      file_authors_v1_authors_proto_msgTypes,
	}.Build()
	File_authors_v1_authors_proto = out.File
	file_authors_v1_authors_proto_rawDesc = nil
	file_authors_v1_authors_proto_goTypes = nil
	file_authors_v1_authors_proto_depIdxs = nil
}
