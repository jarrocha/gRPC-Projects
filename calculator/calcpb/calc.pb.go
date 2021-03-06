// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Operand int32

const (
	Operand_UNKNOWN Operand = 0
	Operand_SUM     Operand = 1
	Operand_SUB     Operand = 2
	Operand_MUL     Operand = 3
	Operand_DIV     Operand = 4
)

var Operand_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUM",
	2: "SUB",
	3: "MUL",
	4: "DIV",
}

var Operand_value = map[string]int32{
	"UNKNOWN": 0,
	"SUM":     1,
	"SUB":     2,
	"MUL":     3,
	"DIV":     4,
}

func (x Operand) String() string {
	return proto.EnumName(Operand_name, int32(x))
}

func (Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a87762106ce7f1ca, []int{0}
}

type Operation struct {
	Operator             Operand  `protobuf:"varint,1,opt,name=operator,proto3,enum=Operand" json:"operator,omitempty"`
	Number1              float64  `protobuf:"fixed64,2,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float64  `protobuf:"fixed64,3,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87762106ce7f1ca, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetOperator() Operand {
	if m != nil {
		return m.Operator
	}
	return Operand_UNKNOWN
}

func (m *Operation) GetNumber1() float64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Operation) GetNumber2() float64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type OperRequest struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OperRequest) Reset()         { *m = OperRequest{} }
func (m *OperRequest) String() string { return proto.CompactTextString(m) }
func (*OperRequest) ProtoMessage()    {}
func (*OperRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87762106ce7f1ca, []int{1}
}

func (m *OperRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperRequest.Unmarshal(m, b)
}
func (m *OperRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperRequest.Marshal(b, m, deterministic)
}
func (m *OperRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperRequest.Merge(m, src)
}
func (m *OperRequest) XXX_Size() int {
	return xxx_messageInfo_OperRequest.Size(m)
}
func (m *OperRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperRequest proto.InternalMessageInfo

func (m *OperRequest) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type OperRespond struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperRespond) Reset()         { *m = OperRespond{} }
func (m *OperRespond) String() string { return proto.CompactTextString(m) }
func (*OperRespond) ProtoMessage()    {}
func (*OperRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87762106ce7f1ca, []int{2}
}

func (m *OperRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperRespond.Unmarshal(m, b)
}
func (m *OperRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperRespond.Marshal(b, m, deterministic)
}
func (m *OperRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperRespond.Merge(m, src)
}
func (m *OperRespond) XXX_Size() int {
	return xxx_messageInfo_OperRespond.Size(m)
}
func (m *OperRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_OperRespond.DiscardUnknown(m)
}

var xxx_messageInfo_OperRespond proto.InternalMessageInfo

func (m *OperRespond) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterEnum("Operand", Operand_name, Operand_value)
	proto.RegisterType((*Operation)(nil), "Operation")
	proto.RegisterType((*OperRequest)(nil), "OperRequest")
	proto.RegisterType((*OperRespond)(nil), "OperRespond")
}

func init() { proto.RegisterFile("calcpb/calc.proto", fileDescriptor_a87762106ce7f1ca) }

var fileDescriptor_a87762106ce7f1ca = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4f, 0x83, 0x30,
	0x14, 0xc6, 0xd7, 0x61, 0x80, 0x3d, 0x8c, 0xa9, 0x3d, 0x18, 0xe2, 0x69, 0x21, 0x9a, 0xa0, 0x07,
	0x8c, 0xf5, 0xa0, 0xf1, 0x38, 0xbd, 0x18, 0x1d, 0x4b, 0xba, 0xa0, 0x89, 0x37, 0x60, 0xef, 0xb0,
	0x04, 0x29, 0x96, 0xe2, 0xdf, 0x6f, 0x5a, 0xc6, 0xe0, 0xd4, 0xef, 0xeb, 0x2f, 0xef, 0x7d, 0x5f,
	0x1e, 0x9c, 0x97, 0x79, 0x55, 0x36, 0xc5, 0x9d, 0x79, 0x92, 0x46, 0x49, 0x2d, 0x23, 0x84, 0xc5,
	0xa6, 0x41, 0x95, 0xeb, 0xbd, 0xac, 0xd9, 0x15, 0xf8, 0xd2, 0x1a, 0xa9, 0x42, 0xb2, 0x24, 0xf1,
	0x19, 0xf7, 0x13, 0x4b, 0xeb, 0x9d, 0x38, 0x12, 0x16, 0x82, 0x57, 0x77, 0x3f, 0x05, 0xaa, 0xfb,
	0x70, 0xbe, 0x24, 0x31, 0x11, 0x83, 0x1d, 0x09, 0x0f, 0x9d, 0x29, 0xe1, 0xd1, 0x23, 0x04, 0x66,
	0x91, 0xc0, 0xdf, 0x0e, 0x5b, 0xcd, 0xe2, 0x49, 0xaa, 0x4d, 0x0a, 0x38, 0x24, 0xc7, 0x1f, 0x31,
	0xc2, 0xe8, 0x7a, 0x18, 0x6c, 0x1b, 0x59, 0xef, 0xd8, 0x05, 0xb8, 0x02, 0xdb, 0xae, 0xd2, 0x76,
	0x8a, 0x88, 0x83, 0xbb, 0x7d, 0x06, 0xef, 0x50, 0x94, 0x05, 0xe0, 0x65, 0xe9, 0x7b, 0xba, 0xf9,
	0x4a, 0xe9, 0x8c, 0x79, 0xe0, 0x6c, 0xb3, 0x35, 0x25, 0xbd, 0x58, 0xd1, 0xb9, 0x11, 0xeb, 0xec,
	0x83, 0x3a, 0x46, 0xbc, 0xbe, 0x7d, 0xd2, 0x13, 0xfe, 0xd4, 0x47, 0x6c, 0x51, 0xfd, 0xed, 0x4b,
	0x64, 0x37, 0xb0, 0x78, 0xc9, 0xab, 0xb2, 0xab, 0x72, 0x8d, 0xec, 0x34, 0x99, 0xd4, 0xbe, 0x1c,
	0x9c, 0xed, 0x12, 0xcd, 0x56, 0xfe, 0xb7, 0xdb, 0x5f, 0xb4, 0x70, 0xed, 0x35, 0x1f, 0xfe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0x1f, 0x69, 0x2d, 0x62, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperServiceClient is the client API for OperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperServiceClient interface {
	// Unary API
	Calculate(ctx context.Context, in *OperRequest, opts ...grpc.CallOption) (*OperRespond, error)
}

type operServiceClient struct {
	cc *grpc.ClientConn
}

func NewOperServiceClient(cc *grpc.ClientConn) OperServiceClient {
	return &operServiceClient{cc}
}

func (c *operServiceClient) Calculate(ctx context.Context, in *OperRequest, opts ...grpc.CallOption) (*OperRespond, error) {
	out := new(OperRespond)
	err := c.cc.Invoke(ctx, "/OperService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperServiceServer is the server API for OperService service.
type OperServiceServer interface {
	// Unary API
	Calculate(context.Context, *OperRequest) (*OperRespond, error)
}

func RegisterOperServiceServer(s *grpc.Server, srv OperServiceServer) {
	s.RegisterService(&_OperService_serviceDesc, srv)
}

func _OperService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OperService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperServiceServer).Calculate(ctx, req.(*OperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OperService",
	HandlerType: (*OperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _OperService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calcpb/calc.proto",
}
