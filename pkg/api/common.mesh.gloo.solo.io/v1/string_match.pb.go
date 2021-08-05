// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: github.com/solo-io/gloo-mesh/api/common/v1/string_match.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Describes how to match a given string in HTTP headers. Match is case-sensitive.
type StringMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string match type.
	//
	// Types that are assignable to MatchType:
	//	*StringMatch_Exact
	//	*StringMatch_Prefix
	//	*StringMatch_Regex
	//	*StringMatch_Suffix
	MatchType isStringMatch_MatchType `protobuf_oneof:"match_type"`
	//If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no effect for the regex match.
	IgnoreCase bool `protobuf:"varint,5,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
}

func (x *StringMatch) Reset() {
	*x = StringMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatch) ProtoMessage() {}

func (x *StringMatch) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatch.ProtoReflect.Descriptor instead.
func (*StringMatch) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescGZIP(), []int{0}
}

func (m *StringMatch) GetMatchType() isStringMatch_MatchType {
	if m != nil {
		return m.MatchType
	}
	return nil
}

func (x *StringMatch) GetExact() string {
	if x, ok := x.GetMatchType().(*StringMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatch) GetPrefix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatch) GetRegex() string {
	if x, ok := x.GetMatchType().(*StringMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

func (x *StringMatch) GetSuffix() string {
	if x, ok := x.GetMatchType().(*StringMatch_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (x *StringMatch) GetIgnoreCase() bool {
	if x != nil {
		return x.IgnoreCase
	}
	return false
}

type isStringMatch_MatchType interface {
	isStringMatch_MatchType()
}

type StringMatch_Exact struct {
	// Exact string match.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatch_Prefix struct {
	// Prefix-based match.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatch_Regex struct {
	// ECMAscript style regex-based match.
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

type StringMatch_Suffix struct {
	// Suffix-based match.
	Suffix string `protobuf:"bytes,4,opt,name=suffix,proto3,oneof"`
}

func (*StringMatch_Exact) isStringMatch_MatchType() {}

func (*StringMatch_Prefix) isStringMatch_MatchType() {}

func (*StringMatch_Regex) isStringMatch_MatchType() {}

func (*StringMatch_Suffix) isStringMatch_MatchType() {}

var File_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x73, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x42, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_goTypes = []interface{}{
	(*StringMatch)(nil), // 0: common.mesh.gloo.solo.io.StringMatch
}
var file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatch); i {
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
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StringMatch_Exact)(nil),
		(*StringMatch_Prefix)(nil),
		(*StringMatch_Regex)(nil),
		(*StringMatch_Suffix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_common_v1_string_match_proto_depIdxs = nil
}
