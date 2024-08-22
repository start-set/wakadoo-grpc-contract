// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/file.proto

package wakadoo_grpc_contract

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

type FileSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Bytes []byte `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *FileSaveRequest) Reset() {
	*x = FileSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSaveRequest) ProtoMessage() {}

func (x *FileSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSaveRequest.ProtoReflect.Descriptor instead.
func (*FileSaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileSaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileSaveRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type FileSaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileSaveResponse) Reset() {
	*x = FileSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSaveResponse) ProtoMessage() {}

func (x *FileSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSaveResponse.ProtoReflect.Descriptor instead.
func (*FileSaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileSaveResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileGetRequest) Reset() {
	*x = FileGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetRequest) ProtoMessage() {}

func (x *FileGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetRequest.ProtoReflect.Descriptor instead.
func (*FileGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{2}
}

func (x *FileGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *FileGetResponse) Reset() {
	*x = FileGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetResponse) ProtoMessage() {}

func (x *FileGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetResponse.ProtoReflect.Descriptor instead.
func (*FileGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{3}
}

func (x *FileGetResponse) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type FileDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FileDeleteRequest) Reset() {
	*x = FileDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteRequest) ProtoMessage() {}

func (x *FileDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteRequest.ProtoReflect.Descriptor instead.
func (*FileDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{4}
}

func (x *FileDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FileDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDeleted bool `protobuf:"varint,1,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
}

func (x *FileDeleteResponse) Reset() {
	*x = FileDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDeleteResponse) ProtoMessage() {}

func (x *FileDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDeleteResponse.ProtoReflect.Descriptor instead.
func (*FileDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_file_proto_rawDescGZIP(), []int{5}
}

func (x *FileDeleteResponse) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

var File_proto_file_proto protoreflect.FileDescriptor

var file_proto_file_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x46,
	0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x19, 0x5a,
	0x17, 0x2e, 0x3b, 0x77, 0x61, 0x6b, 0x61, 0x64, 0x6f, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_file_proto_rawDescOnce sync.Once
	file_proto_file_proto_rawDescData = file_proto_file_proto_rawDesc
)

func file_proto_file_proto_rawDescGZIP() []byte {
	file_proto_file_proto_rawDescOnce.Do(func() {
		file_proto_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_proto_rawDescData)
	})
	return file_proto_file_proto_rawDescData
}

var file_proto_file_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_file_proto_goTypes = []interface{}{
	(*FileSaveRequest)(nil),    // 0: file.FileSaveRequest
	(*FileSaveResponse)(nil),   // 1: file.FileSaveResponse
	(*FileGetRequest)(nil),     // 2: file.FileGetRequest
	(*FileGetResponse)(nil),    // 3: file.FileGetResponse
	(*FileDeleteRequest)(nil),  // 4: file.FileDeleteRequest
	(*FileDeleteResponse)(nil), // 5: file.FileDeleteResponse
}
var file_proto_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_file_proto_init() }
func file_proto_file_proto_init() {
	if File_proto_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSaveRequest); i {
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
		file_proto_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSaveResponse); i {
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
		file_proto_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetRequest); i {
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
		file_proto_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetResponse); i {
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
		file_proto_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDeleteRequest); i {
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
		file_proto_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDeleteResponse); i {
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
			RawDescriptor: file_proto_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_file_proto_goTypes,
		DependencyIndexes: file_proto_file_proto_depIdxs,
		MessageInfos:      file_proto_file_proto_msgTypes,
	}.Build()
	File_proto_file_proto = out.File
	file_proto_file_proto_rawDesc = nil
	file_proto_file_proto_goTypes = nil
	file_proto_file_proto_depIdxs = nil
}
