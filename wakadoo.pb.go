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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_wakadoo_proto protoreflect.FileDescriptor

var file_wakadoo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x61, 0x6b, 0x61, 0x64, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x75, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x53, 0x61,
	0x76, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x3d, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x13, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xe4, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x69, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x19, 0x5a, 0x17, 0x2e, 0x3b, 0x77, 0x61, 0x6b, 0x61, 0x64, 0x6f, 0x6f, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_wakadoo_proto_goTypes = []interface{}{
	(*FileSaveRequest)(nil),               // 0: file.FileSaveRequest
	(*FileGetRequest)(nil),                // 1: file.FileGetRequest
	(*SignInRequest)(nil),                 // 2: auth.SignInRequest
	(*SessionCreateRequest)(nil),          // 3: session.SessionCreateRequest
	(*SessionGetEmailRequest)(nil),        // 4: session.SessionGetEmailRequest
	(*SessionUpdateActivityRequest)(nil),  // 5: session.SessionUpdateActivityRequest
	(*SessionGetActiveRequest)(nil),       // 6: session.SessionGetActiveRequest
	(*FileSaveResponse)(nil),              // 7: file.FileSaveResponse
	(*FileGetResponse)(nil),               // 8: file.FileGetResponse
	(*SignInResponse)(nil),                // 9: auth.SignInResponse
	(*SessionCreateResponse)(nil),         // 10: session.SessionCreateResponse
	(*SessionGetEmailResponse)(nil),       // 11: session.SessionGetEmailResponse
	(*SessionUpdateActivityResponse)(nil), // 12: session.SessionUpdateActivityResponse
	(*SessionGetActiveResponse)(nil),      // 13: session.SessionGetActiveResponse
}
var file_wakadoo_proto_depIdxs = []int32{
	0,  // 0: services.File.Save:input_type -> file.FileSaveRequest
	1,  // 1: services.File.Get:input_type -> file.FileGetRequest
	2,  // 2: services.Auth.SignIn:input_type -> auth.SignInRequest
	3,  // 3: services.Session.Create:input_type -> session.SessionCreateRequest
	4,  // 4: services.Session.GetEmail:input_type -> session.SessionGetEmailRequest
	5,  // 5: services.Session.UpdateLastActivityTime:input_type -> session.SessionUpdateActivityRequest
	6,  // 6: services.Session.GetActive:input_type -> session.SessionGetActiveRequest
	7,  // 7: services.File.Save:output_type -> file.FileSaveResponse
	8,  // 8: services.File.Get:output_type -> file.FileGetResponse
	9,  // 9: services.Auth.SignIn:output_type -> auth.SignInResponse
	10, // 10: services.Session.Create:output_type -> session.SessionCreateResponse
	11, // 11: services.Session.GetEmail:output_type -> session.SessionGetEmailResponse
	12, // 12: services.Session.UpdateLastActivityTime:output_type -> session.SessionUpdateActivityResponse
	13, // 13: services.Session.GetActive:output_type -> session.SessionGetActiveResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_wakadoo_proto_init() }
func file_wakadoo_proto_init() {
	if File_wakadoo_proto != nil {
		return
	}
	file_proto_auth_proto_init()
	file_proto_file_proto_init()
	file_proto_session_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wakadoo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_wakadoo_proto_goTypes,
		DependencyIndexes: file_wakadoo_proto_depIdxs,
	}.Build()
	File_wakadoo_proto = out.File
	file_wakadoo_proto_rawDesc = nil
	file_wakadoo_proto_goTypes = nil
	file_wakadoo_proto_depIdxs = nil
}