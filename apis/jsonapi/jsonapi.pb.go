// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: apis/jsonapi/jsonapi.proto

package jsonapi

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Where specified, a links member can be used to represent links.
type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a string whose value is a URI-reference [RFC3986 Section 4.1] pointing to the link’s target.
	Href string `protobuf:"bytes,1,opt,name=href,proto3" json:"href,omitempty"`
	// a string indicating the link’s relation type.
	Rel *string `protobuf:"bytes,2,opt,name=rel,proto3,oneof" json:"rel,omitempty"`
	// a link to a description document (e.g. OpenAPI or JSON Schema) for the link target.
	Describedby *string `protobuf:"bytes,3,opt,name=describedby,proto3,oneof" json:"describedby,omitempty"`
	// a string which serves as a label for the destination of a link
	// such that it can be used as a human-readable identifier (e.g., a menu entry).
	Title *string `protobuf:"bytes,4,opt,name=title,proto3,oneof" json:"title,omitempty"`
	// a string indicating the media type of the link’s target.
	Type *string `protobuf:"bytes,5,opt,name=type,proto3,oneof" json:"type,omitempty"`
	// a string or an array of strings indicating the language(s) of the link’s target.
	// An array of strings indicates that the link’s target is available in multiple languages.
	Hreflang *string `protobuf:"bytes,6,opt,name=hreflang,proto3,oneof" json:"hreflang,omitempty"`
	// a meta object containing non-standard meta-information about the link.
	Meta map[string]*anypb.Any `protobuf:"bytes,7,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{0}
}

func (x *Links) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Links) GetRel() string {
	if x != nil && x.Rel != nil {
		return *x.Rel
	}
	return ""
}

func (x *Links) GetDescribedby() string {
	if x != nil && x.Describedby != nil {
		return *x.Describedby
	}
	return ""
}

func (x *Links) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Links) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *Links) GetHreflang() string {
	if x != nil && x.Hreflang != nil {
		return *x.Hreflang
	}
	return ""
}

func (x *Links) GetMeta() map[string]*anypb.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data is the primary data for a response.
	Data []*anypb.Any `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// Errors is an array of error objects.
	Errors []*ErrorObject `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetData() []*anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetErrors() []*ErrorObject {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ErrorObjectSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a JSON Pointer [RFC6901] to the value in the request document that caused the error
	// [e.g. "/data" for a primary data object, or "/data/attributes/title" for a specific attribute].
	Pointer string `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
	// a string indicating which URI query parameter caused the error.
	Parameter string `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	// a string indicating the name of a single request header which caused the error.
	Header string `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *ErrorObjectSource) Reset() {
	*x = ErrorObjectSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorObjectSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorObjectSource) ProtoMessage() {}

func (x *ErrorObjectSource) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorObjectSource.ProtoReflect.Descriptor instead.
func (*ErrorObjectSource) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorObjectSource) GetPointer() string {
	if x != nil {
		return x.Pointer
	}
	return ""
}

func (x *ErrorObjectSource) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *ErrorObjectSource) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

type ErrorObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a unique identifier for this particular occurrence of the problem.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// a links object containing references to the source of the error.
	Links *Links `protobuf:"bytes,2,opt,name=links,proto3,oneof" json:"links,omitempty"`
	// the HTTP status code applicable to this problem, expressed as a string value.
	Status uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// an application-specific error code, expressed as a string value.
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	// a short, human-readable summary of the problem
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	// a human-readable explanation specific to this occurrence of the problem. Like title.
	Detail string `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail,omitempty"`
	// an object containing references to the source of the error.
	Source *ErrorObjectSource `protobuf:"bytes,7,opt,name=source,proto3,oneof" json:"source,omitempty"`
	// a meta object containing non-standard meta-information about the error.
	Meta map[string]*anypb.Any `protobuf:"bytes,8,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ErrorObject) Reset() {
	*x = ErrorObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorObject) ProtoMessage() {}

func (x *ErrorObject) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorObject.ProtoReflect.Descriptor instead.
func (*ErrorObject) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ErrorObject) GetLinks() *Links {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *ErrorObject) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ErrorObject) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorObject) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ErrorObject) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *ErrorObject) GetSource() *ErrorObjectSource {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ErrorObject) GetMeta() map[string]*anypb.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ErrorCaller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Line     int32  `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
}

func (x *ErrorCaller) Reset() {
	*x = ErrorCaller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorCaller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorCaller) ProtoMessage() {}

func (x *ErrorCaller) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorCaller.ProtoReflect.Descriptor instead.
func (*ErrorCaller) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorCaller) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *ErrorCaller) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *ErrorCaller) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

// pageInfo is used to indicate whether more edges exist prior or following the set defined by the clients arguments.
type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hasPreviousPage is used to indicate whether more edges exist prior to the set defined by the clients arguments.
	// If the client is paginating with last/before, then the server must return true if prior edges exist, otherwise false.
	// If the client is paginating with first/after, then the client may return true if edges prior to after exist,
	// if it can do so efficiently, otherwise may return false.
	HasPreviousPage bool `protobuf:"varint,1,opt,name=has_previous_page,json=hasPreviousPage,proto3" json:"has_previous_page,omitempty"`
	// hasNextPage is used to indicate whether more edges exist following the set defined by the clients arguments.
	// If the client is paginating with first/after, then the server must return true if further edges exist, otherwise false.
	// If the client is paginating with last/before, then the client may return true if edges further from before exist,
	// if it can do so efficiently, otherwise may return false.
	HasNextPage bool `protobuf:"varint,2,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
	// startCursor is the cursor to the first node in edges. Or the cursor of the representation of the first returned element.
	StartCursor *string `protobuf:"bytes,3,opt,name=start_cursor,json=startCursor,proto3,oneof" json:"start_cursor,omitempty"`
	// endCursor is the cursor to the last node in edges. Or the cursor of the representation of the last returned element.
	EndCursor *string `protobuf:"bytes,4,opt,name=end_cursor,json=endCursor,proto3,oneof" json:"end_cursor,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{5}
}

func (x *PageInfo) GetHasPreviousPage() bool {
	if x != nil {
		return x.HasPreviousPage
	}
	return false
}

func (x *PageInfo) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

func (x *PageInfo) GetStartCursor() string {
	if x != nil && x.StartCursor != nil {
		return *x.StartCursor
	}
	return ""
}

func (x *PageInfo) GetEndCursor() string {
	if x != nil && x.EndCursor != nil {
		return *x.EndCursor
	}
	return ""
}

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// first is the number of items to return from the beginning of the list.
	First int64 `protobuf:"varint,1,opt,name=first,proto3" json:"first,omitempty"`
	// after is the cursor to the first node in edges that should be returned.
	After *string `protobuf:"bytes,2,opt,name=after,proto3,oneof" json:"after,omitempty"`
	// last is the number of items to return from the end of the list.
	Last int64 `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	// before is the cursor to the last node in edges that should be returned.
	Before *string `protobuf:"bytes,4,opt,name=before,proto3,oneof" json:"before,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_jsonapi_jsonapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_apis_jsonapi_jsonapi_proto_rawDescGZIP(), []int{6}
}

func (x *PaginationRequest) GetFirst() int64 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *PaginationRequest) GetAfter() string {
	if x != nil && x.After != nil {
		return *x.After
	}
	return ""
}

func (x *PaginationRequest) GetLast() int64 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *PaginationRequest) GetBefore() string {
	if x != nil && x.Before != nil {
		return *x.Before
	}
	return ""
}

var File_apis_jsonapi_jsonapi_proto protoreflect.FileDescriptor

var file_apis_jsonapi_jsonapi_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x04, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x49, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x92,
	0x41, 0x32, 0x4a, 0x30, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69,
	0x64, 0x6f, 0x63, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x22, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x26, 0x0a, 0x03, 0x72, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x48, 0x00, 0x52, 0x03, 0x72, 0x65, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x92, 0x41, 0x0b, 0x4a, 0x09, 0x22, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x50, 0x49, 0x22, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x64, 0x62, 0x79, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x20, 0x4a, 0x1e, 0x22,
	0x4c, 0x65, 0x61, 0x72, 0x6e, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0x48, 0x02, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x74,
	0x65, 0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0x48, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x08, 0x68, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x4a, 0x07, 0x22, 0x65, 0x6e,
	0x2d, 0x55, 0x53, 0x22, 0x48, 0x04, 0x52, 0x08, 0x68, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x6e, 0x67,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x6c, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x62, 0x79, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x68, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0x67, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x92,
	0x41, 0x0d, 0x4a, 0x0b, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x67, 0x65, 0x22, 0x52,
	0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x92, 0x41, 0x12,
	0x4a, 0x10, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x2e, 0x41, 0x53,
	0x43, 0x22, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x92,
	0x41, 0x11, 0x4a, 0x0f, 0x22, 0x58, 0x2d, 0x53, 0x4f, 0x4d, 0x45, 0x2d, 0x48, 0x45, 0x41, 0x44,
	0x45, 0x52, 0x22, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xdb, 0x04, 0x0a, 0x0b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x92, 0x41, 0x0f, 0x4a, 0x0d, 0x22, 0x42,
	0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0a, 0x92, 0x41, 0x07, 0x4a, 0x05, 0x22, 0x34, 0x30, 0x30, 0x22, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0x92, 0x41, 0x16, 0x4a, 0x14, 0x22, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x22, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x5b, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x45, 0x92, 0x41, 0x42, 0x4a, 0x40, 0x22, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73,
	0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2c, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x28, 0x73, 0x29, 0x20, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x6d, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x55, 0x92, 0x41, 0x52, 0x4a, 0x50, 0x22, 0x41, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x75,
	0x6e, 0x64, 0x65, 0x72, 0x20, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x67, 0x65, 0x20, 0x69,
	0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x2c, 0x20, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x27, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x27, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x65, 0x61, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x27, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x27, 0x22, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3c,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe8, 0x01, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x50, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61,
	0x73, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x3d, 0x22,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x61, 0x47,
	0x56, 0x73, 0x62, 0x47, 0x38, 0x3d, 0x22, 0x48, 0x01, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47,
	0x38, 0x3d, 0x22, 0x48, 0x00, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c,
	0x61, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0x4a, 0x0a, 0x22, 0x61, 0x47, 0x56, 0x73, 0x62,
	0x47, 0x38, 0x3d, 0x22, 0x48, 0x01, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x65, 0x72, 0x75, 0x2d, 0x61, 0x69, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_jsonapi_jsonapi_proto_rawDescOnce sync.Once
	file_apis_jsonapi_jsonapi_proto_rawDescData = file_apis_jsonapi_jsonapi_proto_rawDesc
)

func file_apis_jsonapi_jsonapi_proto_rawDescGZIP() []byte {
	file_apis_jsonapi_jsonapi_proto_rawDescOnce.Do(func() {
		file_apis_jsonapi_jsonapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_jsonapi_jsonapi_proto_rawDescData)
	})
	return file_apis_jsonapi_jsonapi_proto_rawDescData
}

var file_apis_jsonapi_jsonapi_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apis_jsonapi_jsonapi_proto_goTypes = []interface{}{
	(*Links)(nil),             // 0: apis.jsonapi.Links
	(*Response)(nil),          // 1: apis.jsonapi.Response
	(*ErrorObjectSource)(nil), // 2: apis.jsonapi.ErrorObjectSource
	(*ErrorObject)(nil),       // 3: apis.jsonapi.ErrorObject
	(*ErrorCaller)(nil),       // 4: apis.jsonapi.ErrorCaller
	(*PageInfo)(nil),          // 5: apis.jsonapi.PageInfo
	(*PaginationRequest)(nil), // 6: apis.jsonapi.PaginationRequest
	nil,                       // 7: apis.jsonapi.Links.MetaEntry
	nil,                       // 8: apis.jsonapi.ErrorObject.MetaEntry
	(*anypb.Any)(nil),         // 9: google.protobuf.Any
}
var file_apis_jsonapi_jsonapi_proto_depIdxs = []int32{
	7, // 0: apis.jsonapi.Links.meta:type_name -> apis.jsonapi.Links.MetaEntry
	9, // 1: apis.jsonapi.Response.data:type_name -> google.protobuf.Any
	3, // 2: apis.jsonapi.Response.errors:type_name -> apis.jsonapi.ErrorObject
	0, // 3: apis.jsonapi.ErrorObject.links:type_name -> apis.jsonapi.Links
	2, // 4: apis.jsonapi.ErrorObject.source:type_name -> apis.jsonapi.ErrorObjectSource
	8, // 5: apis.jsonapi.ErrorObject.meta:type_name -> apis.jsonapi.ErrorObject.MetaEntry
	9, // 6: apis.jsonapi.Links.MetaEntry.value:type_name -> google.protobuf.Any
	9, // 7: apis.jsonapi.ErrorObject.MetaEntry.value:type_name -> google.protobuf.Any
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_apis_jsonapi_jsonapi_proto_init() }
func file_apis_jsonapi_jsonapi_proto_init() {
	if File_apis_jsonapi_jsonapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_jsonapi_jsonapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorObjectSource); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorObject); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorCaller); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_apis_jsonapi_jsonapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationRequest); i {
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
	file_apis_jsonapi_jsonapi_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_apis_jsonapi_jsonapi_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_apis_jsonapi_jsonapi_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_apis_jsonapi_jsonapi_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_jsonapi_jsonapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_jsonapi_jsonapi_proto_goTypes,
		DependencyIndexes: file_apis_jsonapi_jsonapi_proto_depIdxs,
		MessageInfos:      file_apis_jsonapi_jsonapi_proto_msgTypes,
	}.Build()
	File_apis_jsonapi_jsonapi_proto = out.File
	file_apis_jsonapi_jsonapi_proto_rawDesc = nil
	file_apis_jsonapi_jsonapi_proto_goTypes = nil
	file_apis_jsonapi_jsonapi_proto_depIdxs = nil
}
