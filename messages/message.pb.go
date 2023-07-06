// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: message.proto

package messages

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

type Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Matrix) Reset() {
	*x = Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matrix) ProtoMessage() {}

func (x *Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matrix.ProtoReflect.Descriptor instead.
func (*Matrix) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Matrix) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []float32 `protobuf:"fixed32,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetValues() []float32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type DTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Layer1WeightsMatrix []*Matrix `protobuf:"bytes,1,rep,name=layer1_weights_matrix,json=layer1WeightsMatrix,proto3" json:"layer1_weights_matrix,omitempty"`
	Bias1               []float32 `protobuf:"fixed32,2,rep,packed,name=bias1,proto3" json:"bias1,omitempty"`
	Layer2WeightsMatrix []*Matrix `protobuf:"bytes,3,rep,name=layer2_weights_matrix,json=layer2WeightsMatrix,proto3" json:"layer2_weights_matrix,omitempty"`
	Bias2               []float32 `protobuf:"fixed32,4,rep,packed,name=bias2,proto3" json:"bias2,omitempty"`
	Layer3WeightsMatrix []*Matrix `protobuf:"bytes,5,rep,name=layer3_weights_matrix,json=layer3WeightsMatrix,proto3" json:"layer3_weights_matrix,omitempty"`
	Bias3               []float32 `protobuf:"fixed32,6,rep,packed,name=bias3,proto3" json:"bias3,omitempty"`
}

func (x *DTO) Reset() {
	*x = DTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DTO) ProtoMessage() {}

func (x *DTO) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DTO.ProtoReflect.Descriptor instead.
func (*DTO) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *DTO) GetLayer1WeightsMatrix() []*Matrix {
	if x != nil {
		return x.Layer1WeightsMatrix
	}
	return nil
}

func (x *DTO) GetBias1() []float32 {
	if x != nil {
		return x.Bias1
	}
	return nil
}

func (x *DTO) GetLayer2WeightsMatrix() []*Matrix {
	if x != nil {
		return x.Layer2WeightsMatrix
	}
	return nil
}

func (x *DTO) GetBias2() []float32 {
	if x != nil {
		return x.Bias2
	}
	return nil
}

func (x *DTO) GetLayer3WeightsMatrix() []*Matrix {
	if x != nil {
		return x.Layer3WeightsMatrix
	}
	return nil
}

func (x *DTO) GetBias3() []float32 {
	if x != nil {
		return x.Bias3
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x06, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x77,
	0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x1d, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x99, 0x02, 0x0a, 0x03, 0x44, 0x54, 0x4f, 0x12, 0x44, 0x0a,
	0x15, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x5f,
	0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x13,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x61, 0x73, 0x31, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x05, 0x62, 0x69, 0x61, 0x73, 0x31, 0x12, 0x44, 0x0a, 0x15, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x32, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x13, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x69, 0x61, 0x73, 0x32, 0x18, 0x04, 0x20, 0x03, 0x28, 0x02, 0x52, 0x05,
	0x62, 0x69, 0x61, 0x73, 0x32, 0x12, 0x44, 0x0a, 0x15, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x13, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x69, 0x61, 0x73, 0x33, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x05, 0x62, 0x69, 0x61, 0x73,
	0x33, 0x32, 0x37, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x28, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x54, 0x4f, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x54, 0x4f, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(*Matrix)(nil), // 0: messages.Matrix
	(*Row)(nil),    // 1: messages.Row
	(*DTO)(nil),    // 2: messages.DTO
}
var file_message_proto_depIdxs = []int32{
	1, // 0: messages.Matrix.rows:type_name -> messages.Row
	0, // 1: messages.DTO.layer1_weights_matrix:type_name -> messages.Matrix
	0, // 2: messages.DTO.layer2_weights_matrix:type_name -> messages.Matrix
	0, // 3: messages.DTO.layer3_weights_matrix:type_name -> messages.Matrix
	2, // 4: messages.Coordinator.Client:input_type -> messages.DTO
	2, // 5: messages.Coordinator.Client:output_type -> messages.DTO
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matrix); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DTO); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}