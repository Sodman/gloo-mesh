// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: github.com/solo-io/gloo-mesh/api/networking/v1/ratelimit/rate_limit.proto

package ratelimit

import (
	reflect "reflect"
	sync "sync"

	_ "cuelang.org/go/encoding/protobuf/cue"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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

// Configure the Rate-Limit Filter on a Gateway
type GatewayRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ratelimit service to ask about ratelimit decisions. If not provided,
	// defaults to solo.io rate-limiter server.
	RatelimitServerRef *v1.ObjectRef `protobuf:"bytes,1,opt,name=ratelimit_server_ref,json=ratelimitServerRef,proto3" json:"ratelimit_server_ref,omitempty"`
	// Timeout for the ratelimit service to respond. Defaults to 100ms
	RequestTimeout *duration.Duration `protobuf:"bytes,2,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// Defaults to false
	DenyOnFail bool `protobuf:"varint,3,opt,name=deny_on_fail,json=denyOnFail,proto3" json:"deny_on_fail,omitempty"`
}

func (x *GatewayRateLimit) Reset() {
	*x = GatewayRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayRateLimit) ProtoMessage() {}

func (x *GatewayRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayRateLimit.ProtoReflect.Descriptor instead.
func (*GatewayRateLimit) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayRateLimit) GetRatelimitServerRef() *v1.ObjectRef {
	if x != nil {
		return x.RatelimitServerRef
	}
	return nil
}

func (x *GatewayRateLimit) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *GatewayRateLimit) GetDenyOnFail() bool {
	if x != nil {
		return x.DenyOnFail
	}
	return false
}

// The RateLimitClient specifies either a simplified, abstracted rate limiting model that allows configuring
// the ratelimit Actions directly (raw).
// The corresponding server config should be set in the RateLimitConfig.
type RateLimitClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigType:
	//	*RateLimitClient_Raw
	ConfigType isRateLimitClient_ConfigType `protobuf_oneof:"config_type"`
}

func (x *RateLimitClient) Reset() {
	*x = RateLimitClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitClient) ProtoMessage() {}

func (x *RateLimitClient) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitClient.ProtoReflect.Descriptor instead.
func (*RateLimitClient) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescGZIP(), []int{1}
}

func (m *RateLimitClient) GetConfigType() isRateLimitClient_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *RateLimitClient) GetRaw() *RawRateLimit {
	if x, ok := x.GetConfigType().(*RateLimitClient_Raw); ok {
		return x.Raw
	}
	return nil
}

type isRateLimitClient_ConfigType interface {
	isRateLimitClient_ConfigType()
}

type RateLimitClient_Raw struct {
	// Configure the actions and/or set actions that determine how Envoy composes the descriptors
	Raw *RawRateLimit `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

func (*RateLimitClient_Raw) isRateLimitClient_ConfigType() {}

// Use this field if you want to inline the Envoy rate limits.
// Note that this does not configure the rate limit server. If you are running Gloo Mesh, you need to
// specify the server configuration via the appropriate field in the Gloo Mesh `RateLimitConfig` resource.
// If you are running a custom rate limit server you need to configure it yourself.
type RawRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actions specify how the client (Envoy) will compose the descriptors that
	// will be sent to the server to make a rate limiting decision.
	RateLimits []*v1alpha1.RateLimitActions `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
}

func (x *RawRateLimit) Reset() {
	*x = RawRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawRateLimit) ProtoMessage() {}

func (x *RawRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawRateLimit.ProtoReflect.Descriptor instead.
func (*RawRateLimit) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescGZIP(), []int{2}
}

func (x *RawRateLimit) GetRateLimits() []*v1alpha1.RateLimitActions {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

// Rate limit configuration for a Route or TrafficPolicy. Configures rate limits for individual HTTP routes
type RouteRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Labels to the RateLimitServerConfig ref sent to the ratelimit server
	RatelimitServerConfigSelector *v1.ObjectSelector `protobuf:"bytes,1,opt,name=ratelimit_server_config_selector,json=ratelimitServerConfigSelector,proto3" json:"ratelimit_server_config_selector,omitempty"`
	// Types that are assignable to RateLimitConfigType:
	//	*RouteRateLimit_Raw
	//	*RouteRateLimit_RatelimitClientConfigRef
	RateLimitConfigType isRouteRateLimit_RateLimitConfigType `protobuf_oneof:"rate_limit_config_type"`
}

func (x *RouteRateLimit) Reset() {
	*x = RouteRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRateLimit) ProtoMessage() {}

func (x *RouteRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRateLimit.ProtoReflect.Descriptor instead.
func (*RouteRateLimit) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescGZIP(), []int{3}
}

func (x *RouteRateLimit) GetRatelimitServerConfigSelector() *v1.ObjectSelector {
	if x != nil {
		return x.RatelimitServerConfigSelector
	}
	return nil
}

func (m *RouteRateLimit) GetRateLimitConfigType() isRouteRateLimit_RateLimitConfigType {
	if m != nil {
		return m.RateLimitConfigType
	}
	return nil
}

func (x *RouteRateLimit) GetRaw() *RawRateLimit {
	if x, ok := x.GetRateLimitConfigType().(*RouteRateLimit_Raw); ok {
		return x.Raw
	}
	return nil
}

func (x *RouteRateLimit) GetRatelimitClientConfigRef() *v1.ObjectRef {
	if x, ok := x.GetRateLimitConfigType().(*RouteRateLimit_RatelimitClientConfigRef); ok {
		return x.RatelimitClientConfigRef
	}
	return nil
}

type isRouteRateLimit_RateLimitConfigType interface {
	isRouteRateLimit_RateLimitConfigType()
}

type RouteRateLimit_Raw struct {
	// Configure the actions and/or set actions that determine how Envoy composes the descriptors
	Raw *RawRateLimit `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

type RouteRateLimit_RatelimitClientConfigRef struct {
	// Reference to the RateLimitClientConfig that configures the rate limiting model
	RatelimitClientConfigRef *v1.ObjectRef `protobuf:"bytes,4,opt,name=ratelimit_client_config_ref,json=ratelimitClientConfigRef,proto3,oneof"`
}

func (*RouteRateLimit_Raw) isRouteRateLimit_RateLimitConfigType() {}

func (*RouteRateLimit_RatelimitClientConfigRef) isRouteRateLimit_RateLimitConfigType() {}

var File_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2d, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8,
	0x01, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52,
	0x12, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x66, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x65, 0x6e, 0x79, 0x5f,
	0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64,
	0x65, 0x6e, 0x79, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x22, 0x6a, 0x0a, 0x0f, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x03,
	0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x52, 0x61, 0x77, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22,
	0xbf, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x6a, 0x0a, 0x20, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x1d, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x48,
	0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x77, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x5d, 0x0a, 0x1b, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x18, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x42, 0x18, 0x0a, 0x16, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x54, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_goTypes = []interface{}{
	(*GatewayRateLimit)(nil),          // 0: ratelimit.networking.mesh.gloo.solo.io.GatewayRateLimit
	(*RateLimitClient)(nil),           // 1: ratelimit.networking.mesh.gloo.solo.io.RateLimitClient
	(*RawRateLimit)(nil),              // 2: ratelimit.networking.mesh.gloo.solo.io.RawRateLimit
	(*RouteRateLimit)(nil),            // 3: ratelimit.networking.mesh.gloo.solo.io.RouteRateLimit
	(*v1.ObjectRef)(nil),              // 4: core.skv2.solo.io.ObjectRef
	(*duration.Duration)(nil),         // 5: google.protobuf.Duration
	(*v1alpha1.RateLimitActions)(nil), // 6: ratelimit.api.solo.io.RateLimitActions
	(*v1.ObjectSelector)(nil),         // 7: core.skv2.solo.io.ObjectSelector
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_depIdxs = []int32{
	4, // 0: ratelimit.networking.mesh.gloo.solo.io.GatewayRateLimit.ratelimit_server_ref:type_name -> core.skv2.solo.io.ObjectRef
	5, // 1: ratelimit.networking.mesh.gloo.solo.io.GatewayRateLimit.request_timeout:type_name -> google.protobuf.Duration
	2, // 2: ratelimit.networking.mesh.gloo.solo.io.RateLimitClient.raw:type_name -> ratelimit.networking.mesh.gloo.solo.io.RawRateLimit
	6, // 3: ratelimit.networking.mesh.gloo.solo.io.RawRateLimit.rate_limits:type_name -> ratelimit.api.solo.io.RateLimitActions
	7, // 4: ratelimit.networking.mesh.gloo.solo.io.RouteRateLimit.ratelimit_server_config_selector:type_name -> core.skv2.solo.io.ObjectSelector
	2, // 5: ratelimit.networking.mesh.gloo.solo.io.RouteRateLimit.raw:type_name -> ratelimit.networking.mesh.gloo.solo.io.RawRateLimit
	4, // 6: ratelimit.networking.mesh.gloo.solo.io.RouteRateLimit.ratelimit_client_config_ref:type_name -> core.skv2.solo.io.ObjectRef
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayRateLimit); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitClient); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawRateLimit); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRateLimit); i {
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
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RateLimitClient_Raw)(nil),
	}
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RouteRateLimit_Raw)(nil),
		(*RouteRateLimit_RatelimitClientConfigRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_ratelimit_rate_limit_proto_depIdxs = nil
}
