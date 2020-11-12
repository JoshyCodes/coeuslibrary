// message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: coeus-library/types/protobuf/collections/collections.proto

package collections

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *ListAllRequest) Reset() {
	*x = ListAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllRequest) ProtoMessage() {}

func (x *ListAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllRequest.ProtoReflect.Descriptor instead.
func (*ListAllRequest) Descriptor() ([]byte, []int) {
	return file_coeus_library_types_protobuf_collections_collections_proto_rawDescGZIP(), []int{0}
}

func (x *ListAllRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

type ListAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *ListAllResponse) Reset() {
	*x = ListAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllResponse) ProtoMessage() {}

func (x *ListAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllResponse.ProtoReflect.Descriptor instead.
func (*ListAllResponse) Descriptor() ([]byte, []int) {
	return file_coeus_library_types_protobuf_collections_collections_proto_rawDescGZIP(), []int{1}
}

func (x *ListAllResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

var File_coeus_library_types_protobuf_collections_collections_proto protoreflect.FileDescriptor

var file_coeus_library_types_protobuf_collections_collections_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x65, 0x75, 0x73, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x25, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x5a, 0x0a, 0x12, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44,
	0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coeus_library_types_protobuf_collections_collections_proto_rawDescOnce sync.Once
	file_coeus_library_types_protobuf_collections_collections_proto_rawDescData = file_coeus_library_types_protobuf_collections_collections_proto_rawDesc
)

func file_coeus_library_types_protobuf_collections_collections_proto_rawDescGZIP() []byte {
	file_coeus_library_types_protobuf_collections_collections_proto_rawDescOnce.Do(func() {
		file_coeus_library_types_protobuf_collections_collections_proto_rawDescData = protoimpl.X.CompressGZIP(file_coeus_library_types_protobuf_collections_collections_proto_rawDescData)
	})
	return file_coeus_library_types_protobuf_collections_collections_proto_rawDescData
}

var file_coeus_library_types_protobuf_collections_collections_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coeus_library_types_protobuf_collections_collections_proto_goTypes = []interface{}{
	(*ListAllRequest)(nil),  // 0: collections.ListAllRequest
	(*ListAllResponse)(nil), // 1: collections.ListAllResponse
}
var file_coeus_library_types_protobuf_collections_collections_proto_depIdxs = []int32{
	0, // 0: collections.CollectionsService.ListAll:input_type -> collections.ListAllRequest
	1, // 1: collections.CollectionsService.ListAll:output_type -> collections.ListAllResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coeus_library_types_protobuf_collections_collections_proto_init() }
func file_coeus_library_types_protobuf_collections_collections_proto_init() {
	if File_coeus_library_types_protobuf_collections_collections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllRequest); i {
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
		file_coeus_library_types_protobuf_collections_collections_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllResponse); i {
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
			RawDescriptor: file_coeus_library_types_protobuf_collections_collections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coeus_library_types_protobuf_collections_collections_proto_goTypes,
		DependencyIndexes: file_coeus_library_types_protobuf_collections_collections_proto_depIdxs,
		MessageInfos:      file_coeus_library_types_protobuf_collections_collections_proto_msgTypes,
	}.Build()
	File_coeus_library_types_protobuf_collections_collections_proto = out.File
	file_coeus_library_types_protobuf_collections_collections_proto_rawDesc = nil
	file_coeus_library_types_protobuf_collections_collections_proto_goTypes = nil
	file_coeus_library_types_protobuf_collections_collections_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectionsServiceClient is the client API for CollectionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectionsServiceClient interface {
	ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllResponse, error)
}

type collectionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionsServiceClient(cc grpc.ClientConnInterface) CollectionsServiceClient {
	return &collectionsServiceClient{cc}
}

func (c *collectionsServiceClient) ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllResponse, error) {
	out := new(ListAllResponse)
	err := c.cc.Invoke(ctx, "/collections.CollectionsService/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionsServiceServer is the server API for CollectionsService service.
type CollectionsServiceServer interface {
	ListAll(context.Context, *ListAllRequest) (*ListAllResponse, error)
}

// UnimplementedCollectionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollectionsServiceServer struct {
}

func (*UnimplementedCollectionsServiceServer) ListAll(context.Context, *ListAllRequest) (*ListAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}

func RegisterCollectionsServiceServer(s *grpc.Server, srv CollectionsServiceServer) {
	s.RegisterService(&_CollectionsService_serviceDesc, srv)
}

func _CollectionsService_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionsServiceServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/collections.CollectionsService/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionsServiceServer).ListAll(ctx, req.(*ListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "collections.CollectionsService",
	HandlerType: (*CollectionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAll",
			Handler:    _CollectionsService_ListAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coeus-library/types/protobuf/collections/collections.proto",
}
