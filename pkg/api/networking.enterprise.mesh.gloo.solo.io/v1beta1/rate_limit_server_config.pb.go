// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/rate_limit_server_config.proto

package v1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "cuelang.org/go/encoding/protobuf/cue"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1/ratelimit"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
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

// Possible states of a RateLimitServerConfig resource reflected in the status by Gloo Mesh while processing a resource.
type RateLimitServerConfigStatus_State int32

const (
	// Resources are in a Pending state before they have been processed by Gloo Mesh.
	RateLimitServerConfigStatus_PENDING RateLimitServerConfigStatus_State = 0
	// Resources are in a Accepted state when they are valid and have been applied successfully to
	// the Gloo Mesh configuration.
	RateLimitServerConfigStatus_ACCEPTED RateLimitServerConfigStatus_State = 1
	// Resources are in an Invalid state when they contain incorrect configuration parameters,
	// such as missing required values or invalid resource references.
	// An invalid state can also result when a resource's configuration is valid
	// but conflicts with another resource which was accepted in an earlier point in time.
	RateLimitServerConfigStatus_REJECTED RateLimitServerConfigStatus_State = 2
	// Resources are in a Failed state when they contain correct configuration parameters,
	// but the server encountered an error trying to synchronize the system to
	// the desired state.
	RateLimitServerConfigStatus_FAILED RateLimitServerConfigStatus_State = 3
)

// Enum value maps for RateLimitServerConfigStatus_State.
var (
	RateLimitServerConfigStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTED",
		2: "REJECTED",
		3: "FAILED",
	}
	RateLimitServerConfigStatus_State_value = map[string]int32{
		"PENDING":  0,
		"ACCEPTED": 1,
		"REJECTED": 2,
		"FAILED":   3,
	}
)

func (x RateLimitServerConfigStatus_State) Enum() *RateLimitServerConfigStatus_State {
	p := new(RateLimitServerConfigStatus_State)
	*p = x
	return p
}

func (x RateLimitServerConfigStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitServerConfigStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_enumTypes[0].Descriptor()
}

func (RateLimitServerConfigStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_enumTypes[0]
}

func (x RateLimitServerConfigStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitServerConfigStatus_State.Descriptor instead.
func (RateLimitServerConfigStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescGZIP(), []int{1, 0}
}

// A `RateLimitConfig` describes the ratelimit server policy.
type RateLimitServerConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each `RateLimitConfig` is an instance of one specific configuration type.
	// Currently, only raw configuration is supported, but going forward we are planning on adding
	// more high-level configuration formats to support specific use cases.
	//
	// Types that are assignable to ConfigType:
	//	*RateLimitServerConfigSpec_Raw_
	ConfigType isRateLimitServerConfigSpec_ConfigType `protobuf_oneof:"config_type"`
}

func (x *RateLimitServerConfigSpec) Reset() {
	*x = RateLimitServerConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitServerConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitServerConfigSpec) ProtoMessage() {}

func (x *RateLimitServerConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitServerConfigSpec.ProtoReflect.Descriptor instead.
func (*RateLimitServerConfigSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescGZIP(), []int{0}
}

func (m *RateLimitServerConfigSpec) GetConfigType() isRateLimitServerConfigSpec_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *RateLimitServerConfigSpec) GetRaw() *RateLimitServerConfigSpec_Raw {
	if x, ok := x.GetConfigType().(*RateLimitServerConfigSpec_Raw_); ok {
		return x.Raw
	}
	return nil
}

type isRateLimitServerConfigSpec_ConfigType interface {
	isRateLimitServerConfigSpec_ConfigType()
}

type RateLimitServerConfigSpec_Raw_ struct {
	// Define a policy using the raw configuration format used by the server and the client (Envoy).
	Raw *RateLimitServerConfigSpec_Raw `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

func (*RateLimitServerConfigSpec_Raw_) isRateLimitServerConfigSpec_ConfigType() {}

// The current status of the `RateLimitServerConfig`.
type RateLimitServerConfigStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the RateLimitServerConfig metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Any warnings found while processing this generation of the resource.
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// The current state of the RateLimitServerConfig.
	State RateLimitServerConfigStatus_State `protobuf:"varint,4,opt,name=state,proto3,enum=networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigStatus_State" json:"state,omitempty"`
}

func (x *RateLimitServerConfigStatus) Reset() {
	*x = RateLimitServerConfigStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitServerConfigStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitServerConfigStatus) ProtoMessage() {}

func (x *RateLimitServerConfigStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitServerConfigStatus.ProtoReflect.Descriptor instead.
func (*RateLimitServerConfigStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitServerConfigStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RateLimitServerConfigStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RateLimitServerConfigStatus) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *RateLimitServerConfigStatus) GetState() RateLimitServerConfigStatus_State {
	if x != nil {
		return x.State
	}
	return RateLimitServerConfigStatus_PENDING
}

// This object allows users to specify rate limit policies using the raw configuration formats
// used by the server and the client (Envoy). When using this configuration type, it is up to
// the user to ensure that server and client configurations match to implement the desired behavior.
// The server (and the client libraries that are shipped with it) will ensure that there are no
// collisions between raw configurations defined on separate `RateLimitConfig` resources.
type RateLimitServerConfigSpec_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The descriptors that will be applied to the server.
	// {{/* Note: validation of this field disabled because it slows down cue tremendously*/}}
	Descriptors []*v1alpha1.Descriptor `protobuf:"bytes,1,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The set descriptors that will be applied to the server.
	// {{/* Note: validation of this field disabled because it slows down cue tremendously*/}}
	SetDescriptors []*v1alpha1.SetDescriptor `protobuf:"bytes,2,rep,name=set_descriptors,json=setDescriptors,proto3" json:"set_descriptors,omitempty"`
}

func (x *RateLimitServerConfigSpec_Raw) Reset() {
	*x = RateLimitServerConfigSpec_Raw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitServerConfigSpec_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitServerConfigSpec_Raw) ProtoMessage() {}

func (x *RateLimitServerConfigSpec_Raw) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitServerConfigSpec_Raw.ProtoReflect.Descriptor instead.
func (*RateLimitServerConfigSpec_Raw) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RateLimitServerConfigSpec_Raw) GetDescriptors() []*v1alpha1.Descriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *RateLimitServerConfigSpec_Raw) GetSetDescriptors() []*v1alpha1.SetDescriptor {
	if x != nil {
		return x.SetDescriptors
	}
	return nil
}

var File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f,
	0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb0, 0x02, 0x0a, 0x19, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5a,
	0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x52, 0x61, 0x77, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x1a, 0xa7, 0x01, 0x0a, 0x03, 0x52,
	0x61, 0x77, 0x12, 0x4a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x54,
	0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x42, 0x05, 0xea,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xa2, 0x02, 0x0a, 0x1b, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x60, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3c, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x5a, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_goTypes = []interface{}{
	(RateLimitServerConfigStatus_State)(0), // 0: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigStatus.State
	(*RateLimitServerConfigSpec)(nil),      // 1: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec
	(*RateLimitServerConfigStatus)(nil),    // 2: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigStatus
	(*RateLimitServerConfigSpec_Raw)(nil),  // 3: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec.Raw
	(*v1alpha1.Descriptor)(nil),            // 4: ratelimit.api.solo.io.Descriptor
	(*v1alpha1.SetDescriptor)(nil),         // 5: ratelimit.api.solo.io.SetDescriptor
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_depIdxs = []int32{
	3, // 0: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec.raw:type_name -> networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec.Raw
	0, // 1: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigStatus.state:type_name -> networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigStatus.State
	4, // 2: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec.Raw.descriptors:type_name -> ratelimit.api.solo.io.Descriptor
	5, // 3: networking.enterprise.mesh.gloo.solo.io.RateLimitServerConfigSpec.Raw.set_descriptors:type_name -> ratelimit.api.solo.io.SetDescriptor
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitServerConfigSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitServerConfigStatus); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitServerConfigSpec_Raw); i {
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
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RateLimitServerConfigSpec_Raw_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_server_config_proto_depIdxs = nil
}
