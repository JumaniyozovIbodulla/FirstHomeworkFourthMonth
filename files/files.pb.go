// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: files.proto

package files

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CreateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IsExists bool   `protobuf:"varint,2,opt,name=isExists,proto3" json:"isExists,omitempty"`
	IsWrote  bool   `protobuf:"varint,3,opt,name=isWrote,proto3" json:"isWrote,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateResponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

func (x *CreateResponse) GetIsWrote() bool {
	if x != nil {
		return x.IsWrote
	}
	return false
}

func (x *CreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_files_proto protoreflect.FileDescriptor

var file_files_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x57, 0x72, 0x6f, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x57, 0x72, 0x6f, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc8, 0x01, 0x0a, 0x05, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x49, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x57, 0x72, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x46, 0x72, 0x6f, 0x6d,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x68, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_files_proto_rawDescOnce sync.Once
	file_files_proto_rawDescData = file_files_proto_rawDesc
)

func file_files_proto_rawDescGZIP() []byte {
	file_files_proto_rawDescOnce.Do(func() {
		file_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_files_proto_rawDescData)
	})
	return file_files_proto_rawDescData
}

var file_files_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_files_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: CreateRequest
	(*CreateResponse)(nil), // 1: CreateResponse
}
var file_files_proto_depIdxs = []int32{
	0, // 0: Files.CreateFile:input_type -> CreateRequest
	0, // 1: Files.IsFileExists:input_type -> CreateRequest
	0, // 2: Files.WroteToFile:input_type -> CreateRequest
	0, // 3: Files.ReadFromFile:input_type -> CreateRequest
	1, // 4: Files.CreateFile:output_type -> CreateResponse
	1, // 5: Files.IsFileExists:output_type -> CreateResponse
	1, // 6: Files.WroteToFile:output_type -> CreateResponse
	1, // 7: Files.ReadFromFile:output_type -> CreateResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_files_proto_init() }
func file_files_proto_init() {
	if File_files_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_files_proto_goTypes,
		DependencyIndexes: file_files_proto_depIdxs,
		MessageInfos:      file_files_proto_msgTypes,
	}.Build()
	File_files_proto = out.File
	file_files_proto_rawDesc = nil
	file_files_proto_goTypes = nil
	file_files_proto_depIdxs = nil
}
