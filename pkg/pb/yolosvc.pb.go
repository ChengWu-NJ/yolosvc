// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: yolosvc.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JpgBytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcID   string `protobuf:"bytes,1,opt,name=srcID,proto3" json:"srcID,omitempty"`
	SrcTs   int64  `protobuf:"varint,2,opt,name=srcTs,proto3" json:"srcTs,omitempty"` // unixnano timestamp
	JpgData []byte `protobuf:"bytes,3,opt,name=jpgData,proto3" json:"jpgData,omitempty"`
}

func (x *JpgBytes) Reset() {
	*x = JpgBytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yolosvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JpgBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JpgBytes) ProtoMessage() {}

func (x *JpgBytes) ProtoReflect() protoreflect.Message {
	mi := &file_yolosvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JpgBytes.ProtoReflect.Descriptor instead.
func (*JpgBytes) Descriptor() ([]byte, []int) {
	return file_yolosvc_proto_rawDescGZIP(), []int{0}
}

func (x *JpgBytes) GetSrcID() string {
	if x != nil {
		return x.SrcID
	}
	return ""
}

func (x *JpgBytes) GetSrcTs() int64 {
	if x != nil {
		return x.SrcTs
	}
	return 0
}

func (x *JpgBytes) GetJpgData() []byte {
	if x != nil {
		return x.JpgData
	}
	return nil
}

type HealthzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Htime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=htime,proto3" json:"htime,omitempty"`
}

func (x *HealthzResponse) Reset() {
	*x = HealthzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yolosvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthzResponse) ProtoMessage() {}

func (x *HealthzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yolosvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthzResponse.ProtoReflect.Descriptor instead.
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return file_yolosvc_proto_rawDescGZIP(), []int{1}
}

func (x *HealthzResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *HealthzResponse) GetHtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Htime
	}
	return nil
}

var File_yolosvc_proto protoreflect.FileDescriptor

var file_yolosvc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x79, 0x6f, 0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x77, 0x75, 0x2e, 0x79, 0x6f, 0x6c, 0x6f, 0x73, 0x76, 0x63,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x50, 0x0a, 0x08, 0x4a, 0x70, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x72, 0x63, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x72, 0x63, 0x54, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x72, 0x63, 0x54, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x70, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6a, 0x70, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x59, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x68, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x68, 0x74, 0x69, 0x6d, 0x65, 0x32, 0xbe, 0x02, 0x0a,
	0x09, 0x4f, 0x62, 0x6a, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x64, 0x0a, 0x0c, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x4f, 0x6e, 0x65, 0x4a, 0x70, 0x67, 0x12, 0x1c, 0x2e, 0x63, 0x68, 0x65,
	0x6e, 0x67, 0x77, 0x75, 0x2e, 0x79, 0x6f, 0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x4a, 0x70, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x6e, 0x67,
	0x77, 0x75, 0x2e, 0x79, 0x6f, 0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x70,
	0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x6e, 0x65, 0x6a, 0x70, 0x67,
	0x12, 0x6e, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x4a, 0x70, 0x67, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x77, 0x75, 0x2e, 0x79, 0x6f,
	0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x70, 0x67, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x77, 0x75, 0x2e, 0x79, 0x6f, 0x6c, 0x6f,
	0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x70, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6a, 0x70, 0x67, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x5b, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x77, 0x75, 0x2e, 0x79, 0x6f,
	0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x42, 0x25, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x65, 0x6e,
	0x67, 0x57, 0x75, 0x2d, 0x4e, 0x4a, 0x2f, 0x79, 0x6f, 0x6c, 0x6f, 0x73, 0x76, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yolosvc_proto_rawDescOnce sync.Once
	file_yolosvc_proto_rawDescData = file_yolosvc_proto_rawDesc
)

func file_yolosvc_proto_rawDescGZIP() []byte {
	file_yolosvc_proto_rawDescOnce.Do(func() {
		file_yolosvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_yolosvc_proto_rawDescData)
	})
	return file_yolosvc_proto_rawDescData
}

var file_yolosvc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yolosvc_proto_goTypes = []interface{}{
	(*JpgBytes)(nil),              // 0: chengwu.yolosvc.v1.JpgBytes
	(*HealthzResponse)(nil),       // 1: chengwu.yolosvc.v1.HealthzResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_yolosvc_proto_depIdxs = []int32{
	2, // 0: chengwu.yolosvc.v1.HealthzResponse.htime:type_name -> google.protobuf.Timestamp
	0, // 1: chengwu.yolosvc.v1.ObjDetect.DetectOneJpg:input_type -> chengwu.yolosvc.v1.JpgBytes
	0, // 2: chengwu.yolosvc.v1.ObjDetect.DetectJpgStream:input_type -> chengwu.yolosvc.v1.JpgBytes
	3, // 3: chengwu.yolosvc.v1.ObjDetect.Healthz:input_type -> google.protobuf.Empty
	0, // 4: chengwu.yolosvc.v1.ObjDetect.DetectOneJpg:output_type -> chengwu.yolosvc.v1.JpgBytes
	0, // 5: chengwu.yolosvc.v1.ObjDetect.DetectJpgStream:output_type -> chengwu.yolosvc.v1.JpgBytes
	1, // 6: chengwu.yolosvc.v1.ObjDetect.Healthz:output_type -> chengwu.yolosvc.v1.HealthzResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yolosvc_proto_init() }
func file_yolosvc_proto_init() {
	if File_yolosvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yolosvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JpgBytes); i {
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
		file_yolosvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthzResponse); i {
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
			RawDescriptor: file_yolosvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yolosvc_proto_goTypes,
		DependencyIndexes: file_yolosvc_proto_depIdxs,
		MessageInfos:      file_yolosvc_proto_msgTypes,
	}.Build()
	File_yolosvc_proto = out.File
	file_yolosvc_proto_rawDesc = nil
	file_yolosvc_proto_goTypes = nil
	file_yolosvc_proto_depIdxs = nil
}
