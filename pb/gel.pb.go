// protoc --go_out=plugins=grpc:. gel.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.1
// source: gel.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameters []string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
	Offset     int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_gel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_gel_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetParameters() []string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *Message) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Message `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_gel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_gel_proto_rawDescGZIP(), []int{1}
}

func (x *Logs) GetLogs() []*Message {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamp.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Numbers  map[string]int64     `protobuf:"bytes,2,rep,name=numbers,proto3" json:"numbers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Instants map[string]float64   `protobuf:"bytes,3,rep,name=instants,proto3" json:"instants,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Logs     map[string]*Logs     `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_gel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_gel_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetTs() *timestamp.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Record) GetNumbers() map[string]int64 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

func (x *Record) GetInstants() map[string]float64 {
	if x != nil {
		return x.Instants
	}
	return nil
}

func (x *Record) GetLogs() map[string]*Logs {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_gel_proto protoreflect.FileDescriptor

var file_gel_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x27, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x83, 0x03, 0x0a, 0x06, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73,
	0x12, 0x31, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3b, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x40, 0x0a, 0x0a, 0x47, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x0a, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gel_proto_rawDescOnce sync.Once
	file_gel_proto_rawDescData = file_gel_proto_rawDesc
)

func file_gel_proto_rawDescGZIP() []byte {
	file_gel_proto_rawDescOnce.Do(func() {
		file_gel_proto_rawDescData = protoimpl.X.CompressGZIP(file_gel_proto_rawDescData)
	})
	return file_gel_proto_rawDescData
}

var file_gel_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gel_proto_goTypes = []interface{}{
	(*Message)(nil),             // 0: pb.Message
	(*Logs)(nil),                // 1: pb.Logs
	(*Record)(nil),              // 2: pb.Record
	nil,                         // 3: pb.Record.NumbersEntry
	nil,                         // 4: pb.Record.InstantsEntry
	nil,                         // 5: pb.Record.LogsEntry
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_gel_proto_depIdxs = []int32{
	0, // 0: pb.Logs.logs:type_name -> pb.Message
	6, // 1: pb.Record.ts:type_name -> google.protobuf.Timestamp
	3, // 2: pb.Record.numbers:type_name -> pb.Record.NumbersEntry
	4, // 3: pb.Record.instants:type_name -> pb.Record.InstantsEntry
	5, // 4: pb.Record.logs:type_name -> pb.Record.LogsEntry
	1, // 5: pb.Record.LogsEntry.value:type_name -> pb.Logs
	2, // 6: pb.GelService.SyncRecord:input_type -> pb.Record
	7, // 7: pb.GelService.SyncRecord:output_type -> google.protobuf.Empty
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gel_proto_init() }
func file_gel_proto_init() {
	if File_gel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_gel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
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
		file_gel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_gel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gel_proto_goTypes,
		DependencyIndexes: file_gel_proto_depIdxs,
		MessageInfos:      file_gel_proto_msgTypes,
	}.Build()
	File_gel_proto = out.File
	file_gel_proto_rawDesc = nil
	file_gel_proto_goTypes = nil
	file_gel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GelServiceClient is the client API for GelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GelServiceClient interface {
	SyncRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGelServiceClient(cc grpc.ClientConnInterface) GelServiceClient {
	return &gelServiceClient{cc}
}

func (c *gelServiceClient) SyncRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.GelService/SyncRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GelServiceServer is the server API for GelService service.
type GelServiceServer interface {
	SyncRecord(context.Context, *Record) (*empty.Empty, error)
}

// UnimplementedGelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGelServiceServer struct {
}

func (*UnimplementedGelServiceServer) SyncRecord(context.Context, *Record) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRecord not implemented")
}

func RegisterGelServiceServer(s *grpc.Server, srv GelServiceServer) {
	s.RegisterService(&_GelService_serviceDesc, srv)
}

func _GelService_SyncRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GelServiceServer).SyncRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GelService/SyncRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GelServiceServer).SyncRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

var _GelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GelService",
	HandlerType: (*GelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncRecord",
			Handler:    _GelService_SyncRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gel.proto",
}
