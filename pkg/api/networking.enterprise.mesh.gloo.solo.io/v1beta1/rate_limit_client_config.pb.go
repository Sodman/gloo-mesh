// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/rate_limit_client_config.proto

package v1beta1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	ratelimit "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1/ratelimit"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Possible states of a RateLimitClientConfig resource reflected in the status by Gloo Mesh while processing a resource.
type RateLimitClientConfigStatus_State int32

const (
	// Resources are in a Pending state before they have been processed by Gloo Mesh.
	RateLimitClientConfigStatus_PENDING RateLimitClientConfigStatus_State = 0
	// Resources are in a Accepted state when they are valid and have been applied successfully to
	// the Gloo Mesh configuration.
	RateLimitClientConfigStatus_ACCEPTED RateLimitClientConfigStatus_State = 1
	// Resources are in an Invalid state when they contain incorrect configuration parameters,
	// such as missing required values or invalid resource references.
	// An invalid state can also result when a resource's configuration is valid
	// but conflicts with another resource which was accepted in an earlier point in time.
	RateLimitClientConfigStatus_INVALID RateLimitClientConfigStatus_State = 2
	// Resources are in a Failed state when they contain correct configuration parameters,
	// but the server encountered an error trying to synchronize the system to
	// the desired state.
	RateLimitClientConfigStatus_FAILED RateLimitClientConfigStatus_State = 3
)

// Enum value maps for RateLimitClientConfigStatus_State.
var (
	RateLimitClientConfigStatus_State_name = map[int32]string{
		0: "PENDING",
		1: "ACCEPTED",
		2: "INVALID",
		3: "FAILED",
	}
	RateLimitClientConfigStatus_State_value = map[string]int32{
		"PENDING":  0,
		"ACCEPTED": 1,
		"INVALID":  2,
		"FAILED":   3,
	}
)

func (x RateLimitClientConfigStatus_State) Enum() *RateLimitClientConfigStatus_State {
	p := new(RateLimitClientConfigStatus_State)
	*p = x
	return p
}

func (x RateLimitClientConfigStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitClientConfigStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_enumTypes[0].Descriptor()
}

func (RateLimitClientConfigStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_enumTypes[0]
}

func (x RateLimitClientConfigStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitClientConfigStatus_State.Descriptor instead.
func (RateLimitClientConfigStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescGZIP(), []int{1, 0}
}

// RateLimitClientConfig contains the client configuration for the rate limit Actions that determine how Envoy
// composes the descriptors that are sent to the rate limit server to check whether a request should be rate-limited
type RateLimitClientConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The RateLimitClient specifies the ratelimit Actions which the client (Envoy) will use to
	// compose the descriptors that will be sent to the server to make a rate limiting decision.
	RateLimits *ratelimit.RateLimitClient `protobuf:"bytes,1,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
}

func (x *RateLimitClientConfigSpec) Reset() {
	*x = RateLimitClientConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitClientConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitClientConfigSpec) ProtoMessage() {}

func (x *RateLimitClientConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitClientConfigSpec.ProtoReflect.Descriptor instead.
func (*RateLimitClientConfigSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitClientConfigSpec) GetRateLimits() *ratelimit.RateLimitClient {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

// The current status of the `RateLimitClientConfig`.
type RateLimitClientConfigStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the RateLimitClientConfig metadata.
	// If the `observedGeneration` does not match `metadata.generation`,
	// Gloo Mesh has not processed the most recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// Any errors found while processing this generation of the resource.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Any warnings found while processing this generation of the resource.
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// The current state of the RateLimitClientConfig.
	State RateLimitClientConfigStatus_State `protobuf:"varint,4,opt,name=state,proto3,enum=networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigStatus_State" json:"state,omitempty"`
}

func (x *RateLimitClientConfigStatus) Reset() {
	*x = RateLimitClientConfigStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitClientConfigStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitClientConfigStatus) ProtoMessage() {}

func (x *RateLimitClientConfigStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitClientConfigStatus.ProtoReflect.Descriptor instead.
func (*RateLimitClientConfigStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitClientConfigStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *RateLimitClientConfigStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RateLimitClientConfigStatus) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *RateLimitClientConfigStatus) GetState() RateLimitClientConfigStatus_State {
	if x != nil {
		return x.State
	}
	return RateLimitClientConfigStatus_PENDING
}

var File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x19, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x58, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0xa1,
	0x02, 0x0a, 0x1b, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x60, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x03, 0x42, 0x5a, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65,
	0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_goTypes = []interface{}{
	(RateLimitClientConfigStatus_State)(0), // 0: networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigStatus.State
	(*RateLimitClientConfigSpec)(nil),      // 1: networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigSpec
	(*RateLimitClientConfigStatus)(nil),    // 2: networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigStatus
	(*ratelimit.RateLimitClient)(nil),      // 3: ratelimit.networking.mesh.gloo.solo.io.RateLimitClient
}
var file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_depIdxs = []int32{
	3, // 0: networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigSpec.rate_limits:type_name -> ratelimit.networking.mesh.gloo.solo.io.RateLimitClient
	0, // 1: networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigStatus.state:type_name -> networking.enterprise.mesh.gloo.solo.io.RateLimitClientConfigStatus.State
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitClientConfigSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitClientConfigStatus); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_enterprise_networking_v1beta1_rate_limit_client_config_proto_depIdxs = nil
}
