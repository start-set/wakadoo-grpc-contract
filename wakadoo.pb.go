// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: wakadoo.proto

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

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wakadoo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_wakadoo_proto_rawDescGZIP(), []int{0}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Session string   `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wakadoo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_wakadoo_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *SignInResponse) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wakadoo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_wakadoo_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

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
		mi := &file_wakadoo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSaveRequest) ProtoMessage() {}

func (x *FileSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[3]
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
	return file_wakadoo_proto_rawDescGZIP(), []int{3}
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

	ByteCount uint64 `protobuf:"varint,1,opt,name=byteCount,proto3" json:"byteCount,omitempty"`
}

func (x *FileSaveResponse) Reset() {
	*x = FileSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wakadoo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSaveResponse) ProtoMessage() {}

func (x *FileSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[4]
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
	return file_wakadoo_proto_rawDescGZIP(), []int{4}
}

func (x *FileSaveResponse) GetByteCount() uint64 {
	if x != nil {
		return x.ByteCount
	}
	return 0
}

type FileGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileGetRequest) Reset() {
	*x = FileGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wakadoo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetRequest) ProtoMessage() {}

func (x *FileGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[5]
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
	return file_wakadoo_proto_rawDescGZIP(), []int{5}
}

func (x *FileGetRequest) GetName() string {
	if x != nil {
		return x.Name
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
		mi := &file_wakadoo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetResponse) ProtoMessage() {}

func (x *FileGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wakadoo_proto_msgTypes[6]
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
	return file_wakadoo_proto_rawDescGZIP(), []int{6}
}

func (x *FileGetResponse) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_wakadoo_proto protoreflect.FileDescriptor

var file_wakadoo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x61, 0x6b, 0x61, 0x64, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x57, 0x0a, 0x0e,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x79, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x62, 0x79, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x46, 0x69, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x27, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x32, 0x8c, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x45, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x3d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19,
	0x5a, 0x17, 0x2e, 0x3b, 0x77, 0x61, 0x6b, 0x61, 0x64, 0x6f, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wakadoo_proto_rawDescOnce sync.Once
	file_wakadoo_proto_rawDescData = file_wakadoo_proto_rawDesc
)

func file_wakadoo_proto_rawDescGZIP() []byte {
	file_wakadoo_proto_rawDescOnce.Do(func() {
		file_wakadoo_proto_rawDescData = protoimpl.X.CompressGZIP(file_wakadoo_proto_rawDescData)
	})
	return file_wakadoo_proto_rawDescData
}

var file_wakadoo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wakadoo_proto_goTypes = []interface{}{
	(*SignInRequest)(nil),    // 0: services.SignInRequest
	(*SignInResponse)(nil),   // 1: services.SignInResponse
	(*Account)(nil),          // 2: services.Account
	(*FileSaveRequest)(nil),  // 3: services.FileSaveRequest
	(*FileSaveResponse)(nil), // 4: services.FileSaveResponse
	(*FileGetRequest)(nil),   // 5: services.FileGetRequest
	(*FileGetResponse)(nil),  // 6: services.FileGetResponse
}
var file_wakadoo_proto_depIdxs = []int32{
	2, // 0: services.SignInResponse.account:type_name -> services.Account
	3, // 1: services.FileService.Save:input_type -> services.FileSaveRequest
	5, // 2: services.FileService.Get:input_type -> services.FileGetRequest
	0, // 3: services.Auth.SignIn:input_type -> services.SignInRequest
	4, // 4: services.FileService.Save:output_type -> services.FileSaveResponse
	6, // 5: services.FileService.Get:output_type -> services.FileGetResponse
	1, // 6: services.Auth.SignIn:output_type -> services.SignInResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wakadoo_proto_init() }
func file_wakadoo_proto_init() {
	if File_wakadoo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wakadoo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_wakadoo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_wakadoo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_wakadoo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_wakadoo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_wakadoo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_wakadoo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wakadoo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_wakadoo_proto_goTypes,
		DependencyIndexes: file_wakadoo_proto_depIdxs,
		MessageInfos:      file_wakadoo_proto_msgTypes,
	}.Build()
	File_wakadoo_proto = out.File
	file_wakadoo_proto_rawDesc = nil
	file_wakadoo_proto_goTypes = nil
	file_wakadoo_proto_depIdxs = nil
}
