// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/networking/v1/request_matchers.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusCodeMatcher_Comparator int32

const (
	// Strict equality.
	StatusCodeMatcher_EQ StatusCodeMatcher_Comparator = 0
	// Greater than or equal to.
	StatusCodeMatcher_GE StatusCodeMatcher_Comparator = 1
	// Less than or equal to.
	StatusCodeMatcher_LE StatusCodeMatcher_Comparator = 2
)

// Enum value maps for StatusCodeMatcher_Comparator.
var (
	StatusCodeMatcher_Comparator_name = map[int32]string{
		0: "EQ",
		1: "GE",
		2: "LE",
	}
	StatusCodeMatcher_Comparator_value = map[string]int32{
		"EQ": 0,
		"GE": 1,
		"LE": 2,
	}
)

func (x StatusCodeMatcher_Comparator) Enum() *StatusCodeMatcher_Comparator {
	p := new(StatusCodeMatcher_Comparator)
	*p = x
	return p
}

func (x StatusCodeMatcher_Comparator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCodeMatcher_Comparator) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_enumTypes[0].Descriptor()
}

func (StatusCodeMatcher_Comparator) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_enumTypes[0]
}

func (x StatusCodeMatcher_Comparator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCodeMatcher_Comparator.Descriptor instead.
func (StatusCodeMatcher_Comparator) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP(), []int{1, 0}
}

// Describes a matcher against HTTP request headers.
type HeaderMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the name of the header in the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specify the value of the header. If the value is absent a request that
	// has the name header will match, regardless of the header’s value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specify whether the header value should be treated as regex.
	Regex bool `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	//
	//If set to true, the result of the match will be inverted. Defaults to false.
	//
	//Examples:
	//
	//- name=foo, invert_match=true: matches if no header named `foo` is present
	//- name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present
	//- name=foo, value=``\d{3}``, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present.
	InvertMatch bool `protobuf:"varint,4,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
}

func (x *HeaderMatcher) Reset() {
	*x = HeaderMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMatcher) ProtoMessage() {}

func (x *HeaderMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMatcher.ProtoReflect.Descriptor instead.
func (*HeaderMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeaderMatcher) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *HeaderMatcher) GetRegex() bool {
	if x != nil {
		return x.Regex
	}
	return false
}

func (x *HeaderMatcher) GetInvertMatch() bool {
	if x != nil {
		return x.InvertMatch
	}
	return false
}

// Describes a matcher against HTTP response status codes.
type StatusCodeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code value to match against.
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// The comparison type used for matching.
	Comparator StatusCodeMatcher_Comparator `protobuf:"varint,2,opt,name=comparator,proto3,enum=networking.mesh.gloo.solo.io.StatusCodeMatcher_Comparator" json:"comparator,omitempty"`
}

func (x *StatusCodeMatcher) Reset() {
	*x = StatusCodeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCodeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCodeMatcher) ProtoMessage() {}

func (x *StatusCodeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCodeMatcher.ProtoReflect.Descriptor instead.
func (*StatusCodeMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP(), []int{1}
}

func (x *StatusCodeMatcher) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *StatusCodeMatcher) GetComparator() StatusCodeMatcher_Comparator {
	if x != nil {
		return x.Comparator
	}
	return StatusCodeMatcher_EQ
}

// Specify HTTP request level match criteria. All specified conditions must be satisfied for a match to occur.
type HttpMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name assigned to a match. The match's name will be
	// concatenated with the parent route's name and will be logged in
	// the access logs for requests matching this route.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specify match criteria against the targeted path.
	Uri *StringMatch `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify a set of headers which requests must match in entirety (all headers must match).
	Headers []*HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	// Specify a set of URL query parameters which requests must match in entirety (all query params must match).
	QueryParameters []*HttpMatcher_QueryParameterMatcher `protobuf:"bytes,6,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	// Specify an HTTP method to match against.
	Method string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *HttpMatcher) Reset() {
	*x = HttpMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMatcher) ProtoMessage() {}

func (x *HttpMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMatcher.ProtoReflect.Descriptor instead.
func (*HttpMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP(), []int{2}
}

func (x *HttpMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HttpMatcher) GetUri() *StringMatch {
	if x != nil {
		return x.Uri
	}
	return nil
}

func (x *HttpMatcher) GetHeaders() []*HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HttpMatcher) GetQueryParameters() []*HttpMatcher_QueryParameterMatcher {
	if x != nil {
		return x.QueryParameters
	}
	return nil
}

func (x *HttpMatcher) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

// Specify match criteria against the target URL's query parameters.
type HttpMatcher_QueryParameterMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the name of a key that must be present in the requested path's query string.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specify the value of the query parameter keyed on `name`.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// If true, treat `value` as a regular expression.
	Regex bool `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *HttpMatcher_QueryParameterMatcher) Reset() {
	*x = HttpMatcher_QueryParameterMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMatcher_QueryParameterMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMatcher_QueryParameterMatcher) ProtoMessage() {}

func (x *HttpMatcher_QueryParameterMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMatcher_QueryParameterMatcher.ProtoReflect.Descriptor instead.
func (*HttpMatcher_QueryParameterMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP(), []int{2, 0}
}

func (x *HttpMatcher_QueryParameterMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HttpMatcher_QueryParameterMatcher) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *HttpMatcher_QueryParameterMatcher) GetRegex() bool {
	if x != nil {
		return x.Regex
	}
	return false
}

var File_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x72, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x5a, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x51,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x45, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x45,
	0x10, 0x02, 0x22, 0x82, 0x03, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x12, 0x45, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x6a, 0x0a, 0x10, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x57,
	0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_goTypes = []interface{}{
	(StatusCodeMatcher_Comparator)(0),         // 0: networking.mesh.gloo.solo.io.StatusCodeMatcher.Comparator
	(*HeaderMatcher)(nil),                     // 1: networking.mesh.gloo.solo.io.HeaderMatcher
	(*StatusCodeMatcher)(nil),                 // 2: networking.mesh.gloo.solo.io.StatusCodeMatcher
	(*HttpMatcher)(nil),                       // 3: networking.mesh.gloo.solo.io.HttpMatcher
	(*HttpMatcher_QueryParameterMatcher)(nil), // 4: networking.mesh.gloo.solo.io.HttpMatcher.QueryParameterMatcher
	(*StringMatch)(nil),                       // 5: networking.mesh.gloo.solo.io.StringMatch
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_depIdxs = []int32{
	0, // 0: networking.mesh.gloo.solo.io.StatusCodeMatcher.comparator:type_name -> networking.mesh.gloo.solo.io.StatusCodeMatcher.Comparator
	5, // 1: networking.mesh.gloo.solo.io.HttpMatcher.uri:type_name -> networking.mesh.gloo.solo.io.StringMatch
	1, // 2: networking.mesh.gloo.solo.io.HttpMatcher.headers:type_name -> networking.mesh.gloo.solo.io.HeaderMatcher
	4, // 3: networking.mesh.gloo.solo.io.HttpMatcher.query_parameters:type_name -> networking.mesh.gloo.solo.io.HttpMatcher.QueryParameterMatcher
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_mesh_api_networking_v1_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMatcher); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCodeMatcher); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMatcher); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMatcher_QueryParameterMatcher); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_request_matchers_proto_depIdxs = nil
}
