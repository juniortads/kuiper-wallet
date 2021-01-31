// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pb/transaction.proto

package pb

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

type TransactionType int32

const (
	TransactionType_UNKNOWN_STATUS_TYPE TransactionType = 0
	TransactionType_CASHOUT             TransactionType = 1
	TransactionType_CASHIN              TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "UNKNOWN_STATUS_TYPE",
		1: "CASHOUT",
		2: "CASHIN",
	}
	TransactionType_value = map[string]int32{
		"UNKNOWN_STATUS_TYPE": 0,
		"CASHOUT":             1,
		"CASHIN":              2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_pb_transaction_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_pb_transaction_proto_rawDescGZIP(), []int{0}
}

type CreateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Notes            string            `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	AccountId        string            `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	TrackingId       string            `protobuf:"bytes,4,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	TransactionValue *TransactionValue `protobuf:"bytes,5,opt,name=transactionValue,proto3" json:"transactionValue,omitempty"`
	TransactionType  TransactionType   `protobuf:"varint,6,opt,name=transactionType,proto3,enum=pb.TransactionType" json:"transactionType,omitempty"`
}

func (x *CreateTransactionRequest) Reset() {
	*x = CreateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionRequest) ProtoMessage() {}

func (x *CreateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionRequest.ProtoReflect.Descriptor instead.
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_pb_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTransactionRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *CreateTransactionRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateTransactionRequest) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *CreateTransactionRequest) GetTransactionValue() *TransactionValue {
	if x != nil {
		return x.TransactionValue
	}
	return nil
}

func (x *CreateTransactionRequest) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_UNKNOWN_STATUS_TYPE
}

type CreateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTransactionResponse) Reset() {
	*x = CreateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransactionResponse) ProtoMessage() {}

func (x *CreateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransactionResponse.ProtoReflect.Descriptor instead.
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_pb_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTransactionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TransactionValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency string `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionValue) Reset() {
	*x = TransactionValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionValue) ProtoMessage() {}

func (x *TransactionValue) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionValue.ProtoReflect.Descriptor instead.
func (*TransactionValue) Descriptor() ([]byte, []int) {
	return file_pb_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionValue) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransactionValue) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_pb_transaction_proto protoreflect.FileDescriptor

var file_pb_transaction_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x81, 0x02, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x40, 0x0a,
	0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x3d, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2b,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x43, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x41, 0x53, 0x48, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x41, 0x53, 0x48, 0x49, 0x4e, 0x10, 0x02, 0x32, 0x61, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_transaction_proto_rawDescOnce sync.Once
	file_pb_transaction_proto_rawDescData = file_pb_transaction_proto_rawDesc
)

func file_pb_transaction_proto_rawDescGZIP() []byte {
	file_pb_transaction_proto_rawDescOnce.Do(func() {
		file_pb_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_transaction_proto_rawDescData)
	})
	return file_pb_transaction_proto_rawDescData
}

var file_pb_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_transaction_proto_goTypes = []interface{}{
	(TransactionType)(0),              // 0: pb.TransactionType
	(*CreateTransactionRequest)(nil),  // 1: pb.CreateTransactionRequest
	(*CreateTransactionResponse)(nil), // 2: pb.CreateTransactionResponse
	(*TransactionValue)(nil),          // 3: pb.TransactionValue
}
var file_pb_transaction_proto_depIdxs = []int32{
	3, // 0: pb.CreateTransactionRequest.transactionValue:type_name -> pb.TransactionValue
	0, // 1: pb.CreateTransactionRequest.transactionType:type_name -> pb.TransactionType
	1, // 2: pb.Transaction.CreateTransaction:input_type -> pb.CreateTransactionRequest
	2, // 3: pb.Transaction.CreateTransaction:output_type -> pb.CreateTransactionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_transaction_proto_init() }
func file_pb_transaction_proto_init() {
	if File_pb_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionRequest); i {
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
		file_pb_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransactionResponse); i {
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
		file_pb_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionValue); i {
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
			RawDescriptor: file_pb_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_transaction_proto_goTypes,
		DependencyIndexes: file_pb_transaction_proto_depIdxs,
		EnumInfos:         file_pb_transaction_proto_enumTypes,
		MessageInfos:      file_pb_transaction_proto_msgTypes,
	}.Build()
	File_pb_transaction_proto = out.File
	file_pb_transaction_proto_rawDesc = nil
	file_pb_transaction_proto_goTypes = nil
	file_pb_transaction_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/pb.Transaction/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transaction/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _Transaction_CreateTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/transaction.proto",
}
