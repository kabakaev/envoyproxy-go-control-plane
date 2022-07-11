// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/access_loggers/open_telemetry/v4alpha/logs_service.proto

package envoy_extensions_access_loggers_open_telemetry_v4alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v4alpha "github.com/kabakaev/envoyproxy-go-control-plane/envoy/extensions/access_loggers/grpc/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the built-in *envoy.access_loggers.open_telemetry*
// :ref:`AccessLog <envoy_v3_api_msg_config.accesslog.v3.AccessLog>`. This configuration will
// populate `opentelemetry.proto.collector.v1.logs.ExportLogsServiceRequest.resource_logs <https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/collector/logs/v1/logs_service.proto>`_.
// OpenTelemetry `Resource <https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/resource/v1/resource.proto>`_
// attributes are filled with Envoy node info. In addition, the request start time is set in the
// dedicated field.
// [#extension: envoy.access_loggers.open_telemetry]
// [#comment:TODO(itamarkam): allow configuration for resource attributes.]
type OpenTelemetryAccessLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [#comment:TODO(itamarkam): add 'filter_state_objects_to_log' to logs.]
	CommonConfig *v4alpha.CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// OpenTelemetry `LogResource <https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/logs/v1/logs.proto>`_
	// fields, following `Envoy access logging formatting <https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log/usage>`_.
	//
	// See 'body' in the LogResource proto for more details.
	// Example: ``body { string_value: "%PROTOCOL%" }``.
	Body *v1.AnyValue `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// See 'attributes' in the LogResource proto for more details.
	// Example: ``attributes { values { key: "user_agent" value { string_value: "%REQ(USER-AGENT)%" } } }``.
	Attributes *v1.KeyValueList `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *OpenTelemetryAccessLogConfig) Reset() {
	*x = OpenTelemetryAccessLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenTelemetryAccessLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenTelemetryAccessLogConfig) ProtoMessage() {}

func (x *OpenTelemetryAccessLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenTelemetryAccessLogConfig.ProtoReflect.Descriptor instead.
func (*OpenTelemetryAccessLogConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescGZIP(), []int{0}
}

func (x *OpenTelemetryAccessLogConfig) GetCommonConfig() *v4alpha.CommonGrpcAccessLogConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

func (x *OpenTelemetryAccessLogConfig) GetBody() *v1.AnyValue {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *OpenTelemetryAccessLogConfig) GetAttributes() *v1.KeyValueList {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto protoreflect.FileDescriptor

var file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6f, 0x70, 0x65,
	0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfc, 0x02, 0x0a, 0x1c, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x76, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x4b, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x5a, 0x9a, 0xc5, 0x88, 0x1e, 0x55, 0x0a, 0x53, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x6a, 0x0a, 0x44, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x10, 0x4c, 0x6f, 0x67, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescOnce sync.Once
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescData = file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDesc
)

func file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescGZIP() []byte {
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescData)
	})
	return file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDescData
}

var file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_goTypes = []interface{}{
	(*OpenTelemetryAccessLogConfig)(nil),      // 0: envoy.extensions.access_loggers.open_telemetry.v4alpha.OpenTelemetryAccessLogConfig
	(*v4alpha.CommonGrpcAccessLogConfig)(nil), // 1: envoy.extensions.access_loggers.grpc.v4alpha.CommonGrpcAccessLogConfig
	(*v1.AnyValue)(nil),                       // 2: opentelemetry.proto.common.v1.AnyValue
	(*v1.KeyValueList)(nil),                   // 3: opentelemetry.proto.common.v1.KeyValueList
}
var file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.access_loggers.open_telemetry.v4alpha.OpenTelemetryAccessLogConfig.common_config:type_name -> envoy.extensions.access_loggers.grpc.v4alpha.CommonGrpcAccessLogConfig
	2, // 1: envoy.extensions.access_loggers.open_telemetry.v4alpha.OpenTelemetryAccessLogConfig.body:type_name -> opentelemetry.proto.common.v1.AnyValue
	3, // 2: envoy.extensions.access_loggers.open_telemetry.v4alpha.OpenTelemetryAccessLogConfig.attributes:type_name -> opentelemetry.proto.common.v1.KeyValueList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_init() }
func file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_init() {
	if File_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenTelemetryAccessLogConfig); i {
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
			RawDescriptor: file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_msgTypes,
	}.Build()
	File_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto = out.File
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_rawDesc = nil
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_goTypes = nil
	file_envoy_extensions_access_loggers_open_telemetry_v4alpha_logs_service_proto_depIdxs = nil
}
