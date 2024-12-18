// proto/transmitter.proto

// Версия ProtoBuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: transmitter.proto

// Текущий пакет - указывает пространство имен для сервиса и сообщений. Помогает избегать конфликтов имен.

package pb

import (
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

type TransmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransmitRequest) Reset() {
	*x = TransmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitRequest) ProtoMessage() {}

func (x *TransmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitRequest.ProtoReflect.Descriptor instead.
func (*TransmitRequest) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{0}
}

type TransmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string  `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64 `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransmitResponse) Reset() {
	*x = TransmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transmitter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransmitResponse) ProtoMessage() {}

func (x *TransmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransmitResponse.ProtoReflect.Descriptor instead.
func (*TransmitResponse) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{1}
}

func (x *TransmitResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TransmitResponse) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *TransmitResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_transmitter_proto protoreflect.FileDescriptor

var file_transmitter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x22, 0x11, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x32, 0x56, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x12, 0x47, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transmitter_proto_rawDescOnce sync.Once
	file_transmitter_proto_rawDescData = file_transmitter_proto_rawDesc
)

func file_transmitter_proto_rawDescGZIP() []byte {
	file_transmitter_proto_rawDescOnce.Do(func() {
		file_transmitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_transmitter_proto_rawDescData)
	})
	return file_transmitter_proto_rawDescData
}

var file_transmitter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transmitter_proto_goTypes = []any{
	(*TransmitRequest)(nil),  // 0: transmitter.TransmitRequest
	(*TransmitResponse)(nil), // 1: transmitter.TransmitResponse
}
var file_transmitter_proto_depIdxs = []int32{
	0, // 0: transmitter.Transmitter.Transmit:input_type -> transmitter.TransmitRequest
	1, // 1: transmitter.Transmitter.Transmit:output_type -> transmitter.TransmitResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transmitter_proto_init() }
func file_transmitter_proto_init() {
	if File_transmitter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transmitter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TransmitRequest); i {
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
		file_transmitter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TransmitResponse); i {
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
			RawDescriptor: file_transmitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transmitter_proto_goTypes,
		DependencyIndexes: file_transmitter_proto_depIdxs,
		MessageInfos:      file_transmitter_proto_msgTypes,
	}.Build()
	File_transmitter_proto = out.File
	file_transmitter_proto_rawDesc = nil
	file_transmitter_proto_goTypes = nil
	file_transmitter_proto_depIdxs = nil
}
