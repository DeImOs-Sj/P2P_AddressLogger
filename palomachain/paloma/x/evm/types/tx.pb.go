// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/palomachain/paloma/x/valset/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDeployNewSmartContractRequest struct {
	Title       string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AbiJSON     string            `protobuf:"bytes,4,opt,name=abiJSON,proto3" json:"abiJSON,omitempty"`
	BytecodeHex string            `protobuf:"bytes,5,opt,name=bytecodeHex,proto3" json:"bytecodeHex,omitempty"`
	Metadata    types.MsgMetadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgDeployNewSmartContractRequest) Reset()         { *m = MsgDeployNewSmartContractRequest{} }
func (m *MsgDeployNewSmartContractRequest) String() string { return proto.CompactTextString(m) }
func (*MsgDeployNewSmartContractRequest) ProtoMessage()    {}
func (*MsgDeployNewSmartContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{0}
}
func (m *MsgDeployNewSmartContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeployNewSmartContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployNewSmartContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeployNewSmartContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployNewSmartContractRequest.Merge(m, src)
}
func (m *MsgDeployNewSmartContractRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeployNewSmartContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployNewSmartContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployNewSmartContractRequest proto.InternalMessageInfo

func (m *MsgDeployNewSmartContractRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetAbiJSON() string {
	if m != nil {
		return m.AbiJSON
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetBytecodeHex() string {
	if m != nil {
		return m.BytecodeHex
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type DeployNewSmartContractResponse struct {
}

func (m *DeployNewSmartContractResponse) Reset()         { *m = DeployNewSmartContractResponse{} }
func (m *DeployNewSmartContractResponse) String() string { return proto.CompactTextString(m) }
func (*DeployNewSmartContractResponse) ProtoMessage()    {}
func (*DeployNewSmartContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{1}
}
func (m *DeployNewSmartContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeployNewSmartContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployNewSmartContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeployNewSmartContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployNewSmartContractResponse.Merge(m, src)
}
func (m *DeployNewSmartContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeployNewSmartContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployNewSmartContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployNewSmartContractResponse proto.InternalMessageInfo

type MsgRemoveSmartContractDeploymentRequest struct {
	SmartContractID  uint64            `protobuf:"varint,2,opt,name=smartContractID,proto3" json:"smartContractID,omitempty"`
	ChainReferenceID string            `protobuf:"bytes,3,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	Metadata         types.MsgMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgRemoveSmartContractDeploymentRequest) Reset() {
	*m = MsgRemoveSmartContractDeploymentRequest{}
}
func (m *MsgRemoveSmartContractDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveSmartContractDeploymentRequest) ProtoMessage()    {}
func (*MsgRemoveSmartContractDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{2}
}
func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.Merge(m, src)
}
func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveSmartContractDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest proto.InternalMessageInfo

func (m *MsgRemoveSmartContractDeploymentRequest) GetSmartContractID() uint64 {
	if m != nil {
		return m.SmartContractID
	}
	return 0
}

func (m *MsgRemoveSmartContractDeploymentRequest) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *MsgRemoveSmartContractDeploymentRequest) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type RemoveSmartContractDeploymentResponse struct {
}

func (m *RemoveSmartContractDeploymentResponse) Reset()         { *m = RemoveSmartContractDeploymentResponse{} }
func (m *RemoveSmartContractDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveSmartContractDeploymentResponse) ProtoMessage()    {}
func (*RemoveSmartContractDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{3}
}
func (m *RemoveSmartContractDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveSmartContractDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveSmartContractDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveSmartContractDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSmartContractDeploymentResponse.Merge(m, src)
}
func (m *RemoveSmartContractDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *RemoveSmartContractDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSmartContractDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSmartContractDeploymentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDeployNewSmartContractRequest)(nil), "palomachain.paloma.evm.MsgDeployNewSmartContractRequest")
	proto.RegisterType((*DeployNewSmartContractResponse)(nil), "palomachain.paloma.evm.DeployNewSmartContractResponse")
	proto.RegisterType((*MsgRemoveSmartContractDeploymentRequest)(nil), "palomachain.paloma.evm.MsgRemoveSmartContractDeploymentRequest")
	proto.RegisterType((*RemoveSmartContractDeploymentResponse)(nil), "palomachain.paloma.evm.RemoveSmartContractDeploymentResponse")
}

func init() { proto.RegisterFile("palomachain/paloma/evm/tx.proto", fileDescriptor_631cfc68eb1fd278) }

var fileDescriptor_631cfc68eb1fd278 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x77, 0xb6, 0xe9, 0xb6, 0x4e, 0x11, 0xcb, 0x50, 0x6a, 0x58, 0x30, 0x0d, 0x0b, 0xb6,
	0x6b, 0x0f, 0x09, 0x56, 0x10, 0x11, 0x44, 0x68, 0xf7, 0xd0, 0x16, 0x52, 0x21, 0x7b, 0xf3, 0x36,
	0x3b, 0xfb, 0x9a, 0x06, 0x76, 0xf2, 0xc6, 0xcc, 0x34, 0xee, 0x5e, 0x3d, 0x7a, 0xf2, 0x1b, 0xf8,
	0x15, 0x3c, 0xf8, 0x21, 0x7a, 0xec, 0xd1, 0x93, 0xc8, 0x2e, 0xe8, 0xb7, 0x10, 0xd9, 0xcc, 0xba,
	0xa4, 0x9a, 0xdd, 0xf6, 0xe0, 0x29, 0xef, 0x9f, 0xe7, 0x49, 0xe6, 0xfd, 0xcd, 0x4c, 0xe8, 0x4e,
	0xca, 0x07, 0x28, 0xb9, 0x38, 0xe7, 0x71, 0xe2, 0x9b, 0xd8, 0x87, 0x5c, 0xfa, 0x7a, 0xe8, 0xa5,
	0x19, 0x6a, 0x64, 0xdb, 0x25, 0x81, 0x67, 0x62, 0x0f, 0x72, 0xd9, 0xdc, 0x8a, 0x30, 0xc2, 0x42,
	0xe2, 0x4f, 0x23, 0xa3, 0x6e, 0xde, 0x17, 0xa8, 0x24, 0x2a, 0x5f, 0xaa, 0xc8, 0xcf, 0x1f, 0x4f,
	0x1f, 0xb3, 0xc6, 0x6e, 0xc5, 0x77, 0x72, 0x3e, 0x50, 0xa0, 0x7d, 0x81, 0x52, 0x62, 0x62, 0x74,
	0xad, 0x5f, 0x84, 0xba, 0x81, 0x8a, 0x3a, 0x90, 0x0e, 0x70, 0x74, 0x06, 0xef, 0xba, 0x92, 0x67,
	0xfa, 0x08, 0x13, 0x9d, 0x71, 0xa1, 0x43, 0x78, 0x7b, 0x01, 0x4a, 0xb3, 0x2d, 0xba, 0xaa, 0x63,
	0x3d, 0x00, 0xbb, 0xee, 0x92, 0xf6, 0x9d, 0xd0, 0x24, 0xcc, 0xa5, 0x1b, 0x7d, 0x50, 0x22, 0x8b,
	0x53, 0x1d, 0x63, 0x62, 0xaf, 0x14, 0xbd, 0x72, 0x89, 0xd9, 0x74, 0x8d, 0xf7, 0xe2, 0xd3, 0xee,
	0xab, 0x33, 0xdb, 0x2a, 0xba, 0x7f, 0xd2, 0xa9, 0xb7, 0x37, 0xd2, 0x20, 0xb0, 0x0f, 0xc7, 0x30,
	0xb4, 0x57, 0x8d, 0xb7, 0x54, 0x62, 0xc7, 0x74, 0x5d, 0x82, 0xe6, 0x7d, 0xae, 0xb9, 0xdd, 0x70,
	0x49, 0x7b, 0xe3, 0x60, 0xd7, 0xab, 0x40, 0x63, 0x66, 0xf2, 0x02, 0x15, 0x05, 0x33, 0xf5, 0xa1,
	0x75, 0xf9, 0x6d, 0xa7, 0x16, 0xce, 0xdd, 0xcf, 0xef, 0xbe, 0xff, 0xf9, 0x79, 0x7f, 0x9e, 0x9e,
	0x5a, 0xeb, 0x64, 0xb3, 0x1e, 0xae, 0x89, 0x0c, 0xb8, 0xc6, 0xac, 0xe5, 0x52, 0x67, 0xd1, 0xf0,
	0x2a, 0xc5, 0x44, 0x41, 0xeb, 0x07, 0xa1, 0x7b, 0x81, 0x8a, 0x42, 0x90, 0x98, 0xc3, 0x35, 0x89,
	0x31, 0x4a, 0x48, 0xe6, 0xa4, 0xda, 0xf4, 0x9e, 0x2a, 0x2b, 0x4e, 0x3a, 0x05, 0x33, 0x2b, 0xfc,
	0xbb, 0xcc, 0xf6, 0xe9, 0x66, 0x31, 0x48, 0x08, 0x6f, 0x20, 0x83, 0x44, 0xc0, 0x49, 0x67, 0x86,
	0xf0, 0x9f, 0xfa, 0x35, 0x16, 0xd6, 0xff, 0x67, 0xd1, 0xe8, 0x42, 0xd2, 0x87, 0xac, 0xb5, 0x47,
	0x1f, 0xde, 0x30, 0xa4, 0x21, 0x72, 0xf0, 0xa5, 0x4e, 0x57, 0x02, 0x15, 0xb1, 0x0f, 0x84, 0x6e,
	0x57, 0xc3, 0x63, 0xcf, 0xbc, 0xea, 0x73, 0xec, 0xdd, 0x74, 0xd8, 0x9a, 0x4f, 0x17, 0x39, 0x97,
	0x6f, 0x13, 0xfb, 0x44, 0xe8, 0x83, 0xa5, 0xcb, 0x67, 0x2f, 0x97, 0xac, 0xe9, 0x36, 0xbb, 0xdb,
	0x7c, 0xb1, 0xe8, 0x05, 0xb7, 0xc2, 0x76, 0x78, 0x74, 0x39, 0x76, 0xc8, 0xd5, 0xd8, 0x21, 0xdf,
	0xc7, 0x0e, 0xf9, 0x38, 0x71, 0x6a, 0x57, 0x13, 0xa7, 0xf6, 0x75, 0xe2, 0xd4, 0x5e, 0x3f, 0x8a,
	0x62, 0x7d, 0x7e, 0xd1, 0xf3, 0x04, 0x4a, 0xbf, 0xe2, 0xe2, 0x0e, 0xcd, 0x2f, 0x62, 0x94, 0x82,
	0xea, 0x35, 0x8a, 0x7b, 0xfb, 0xe4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xe1, 0x2a, 0xdf,
	0x49, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DeployNewSmartContract(ctx context.Context, in *MsgDeployNewSmartContractRequest, opts ...grpc.CallOption) (*DeployNewSmartContractResponse, error)
	RemoveSmartContractDeployment(ctx context.Context, in *MsgRemoveSmartContractDeploymentRequest, opts ...grpc.CallOption) (*RemoveSmartContractDeploymentResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DeployNewSmartContract(ctx context.Context, in *MsgDeployNewSmartContractRequest, opts ...grpc.CallOption) (*DeployNewSmartContractResponse, error) {
	out := new(DeployNewSmartContractResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/DeployNewSmartContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveSmartContractDeployment(ctx context.Context, in *MsgRemoveSmartContractDeploymentRequest, opts ...grpc.CallOption) (*RemoveSmartContractDeploymentResponse, error) {
	out := new(RemoveSmartContractDeploymentResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/RemoveSmartContractDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DeployNewSmartContract(context.Context, *MsgDeployNewSmartContractRequest) (*DeployNewSmartContractResponse, error)
	RemoveSmartContractDeployment(context.Context, *MsgRemoveSmartContractDeploymentRequest) (*RemoveSmartContractDeploymentResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DeployNewSmartContract(ctx context.Context, req *MsgDeployNewSmartContractRequest) (*DeployNewSmartContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployNewSmartContract not implemented")
}
func (*UnimplementedMsgServer) RemoveSmartContractDeployment(ctx context.Context, req *MsgRemoveSmartContractDeploymentRequest) (*RemoveSmartContractDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSmartContractDeployment not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DeployNewSmartContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeployNewSmartContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeployNewSmartContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/DeployNewSmartContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeployNewSmartContract(ctx, req.(*MsgDeployNewSmartContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveSmartContractDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveSmartContractDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveSmartContractDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/RemoveSmartContractDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveSmartContractDeployment(ctx, req.(*MsgRemoveSmartContractDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "palomachain.paloma.evm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployNewSmartContract",
			Handler:    _Msg_DeployNewSmartContract_Handler,
		},
		{
			MethodName: "RemoveSmartContractDeployment",
			Handler:    _Msg_RemoveSmartContractDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "palomachain/paloma/evm/tx.proto",
}

func (m *MsgDeployNewSmartContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployNewSmartContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployNewSmartContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.BytecodeHex) > 0 {
		i -= len(m.BytecodeHex)
		copy(dAtA[i:], m.BytecodeHex)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BytecodeHex)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AbiJSON) > 0 {
		i -= len(m.AbiJSON)
		copy(dAtA[i:], m.AbiJSON)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AbiJSON)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *DeployNewSmartContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployNewSmartContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployNewSmartContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveSmartContractDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveSmartContractDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveSmartContractDeploymentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SmartContractID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SmartContractID))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func (m *RemoveSmartContractDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveSmartContractDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveSmartContractDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDeployNewSmartContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AbiJSON)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BytecodeHex)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *DeployNewSmartContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveSmartContractDeploymentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SmartContractID != 0 {
		n += 1 + sovTx(uint64(m.SmartContractID))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *RemoveSmartContractDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeployNewSmartContractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeployNewSmartContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployNewSmartContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbiJSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbiJSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytecodeHex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytecodeHex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeployNewSmartContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeployNewSmartContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployNewSmartContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveSmartContractDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveSmartContractDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveSmartContractDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractID", wireType)
			}
			m.SmartContractID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SmartContractID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RemoveSmartContractDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveSmartContractDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveSmartContractDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
