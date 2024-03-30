// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the evm module's genesis state.
type GenesisState struct {
	Params        Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SmartContract *GenesisSmartContract `protobuf:"bytes,2,opt,name=smartContract,proto3" json:"smartContract,omitempty"`
	Chains        []*GenesisChainInfo   `protobuf:"bytes,3,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36103825ee62903, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSmartContract() *GenesisSmartContract {
	if m != nil {
		return m.SmartContract
	}
	return nil
}

func (m *GenesisState) GetChains() []*GenesisChainInfo {
	if m != nil {
		return m.Chains
	}
	return nil
}

type GenesisChainInfo struct {
	ChainReferenceID  string        `protobuf:"bytes,1,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	ChainID           uint64        `protobuf:"varint,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	BlockHeight       uint64        `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BlockHashAtHeight string        `protobuf:"bytes,4,opt,name=blockHashAtHeight,proto3" json:"blockHashAtHeight,omitempty"`
	MinOnChainBalance string        `protobuf:"bytes,5,opt,name=minOnChainBalance,proto3" json:"minOnChainBalance,omitempty"`
	RelayWeights      *RelayWeights `protobuf:"bytes,6,opt,name=relayWeights,proto3" json:"relayWeights,omitempty"`
}

func (m *GenesisChainInfo) Reset()         { *m = GenesisChainInfo{} }
func (m *GenesisChainInfo) String() string { return proto.CompactTextString(m) }
func (*GenesisChainInfo) ProtoMessage()    {}
func (*GenesisChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36103825ee62903, []int{1}
}
func (m *GenesisChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisChainInfo.Merge(m, src)
}
func (m *GenesisChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *GenesisChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisChainInfo proto.InternalMessageInfo

func (m *GenesisChainInfo) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *GenesisChainInfo) GetChainID() uint64 {
	if m != nil {
		return m.ChainID
	}
	return 0
}

func (m *GenesisChainInfo) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *GenesisChainInfo) GetBlockHashAtHeight() string {
	if m != nil {
		return m.BlockHashAtHeight
	}
	return ""
}

func (m *GenesisChainInfo) GetMinOnChainBalance() string {
	if m != nil {
		return m.MinOnChainBalance
	}
	return ""
}

func (m *GenesisChainInfo) GetRelayWeights() *RelayWeights {
	if m != nil {
		return m.RelayWeights
	}
	return nil
}

type GenesisSmartContract struct {
	AbiJson     string `protobuf:"bytes,1,opt,name=abiJson,proto3" json:"abiJson,omitempty"`
	BytecodeHex string `protobuf:"bytes,2,opt,name=bytecodeHex,proto3" json:"bytecodeHex,omitempty"`
}

func (m *GenesisSmartContract) Reset()         { *m = GenesisSmartContract{} }
func (m *GenesisSmartContract) String() string { return proto.CompactTextString(m) }
func (*GenesisSmartContract) ProtoMessage()    {}
func (*GenesisSmartContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36103825ee62903, []int{2}
}
func (m *GenesisSmartContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSmartContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSmartContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSmartContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSmartContract.Merge(m, src)
}
func (m *GenesisSmartContract) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSmartContract) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSmartContract.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSmartContract proto.InternalMessageInfo

func (m *GenesisSmartContract) GetAbiJson() string {
	if m != nil {
		return m.AbiJson
	}
	return ""
}

func (m *GenesisSmartContract) GetBytecodeHex() string {
	if m != nil {
		return m.BytecodeHex
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "palomachain.paloma.evm.GenesisState")
	proto.RegisterType((*GenesisChainInfo)(nil), "palomachain.paloma.evm.GenesisChainInfo")
	proto.RegisterType((*GenesisSmartContract)(nil), "palomachain.paloma.evm.GenesisSmartContract")
}

func init() {
	proto.RegisterFile("palomachain/paloma/evm/genesis.proto", fileDescriptor_a36103825ee62903)
}

var fileDescriptor_a36103825ee62903 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x93, 0xdb, 0x1a, 0xb9, 0xd3, 0x2b, 0x5c, 0x87, 0x8b, 0x84, 0x2e, 0x62, 0xa9, 0x77,
	0x51, 0x2f, 0x97, 0x04, 0xea, 0xd6, 0x85, 0xb6, 0x05, 0x5b, 0x37, 0xca, 0xb8, 0x10, 0xdc, 0xc8,
	0x24, 0x9e, 0x26, 0xc1, 0x64, 0x26, 0x24, 0x63, 0x6d, 0xdf, 0xc2, 0x47, 0xf0, 0x71, 0xba, 0xec,
	0xd2, 0x95, 0x48, 0xfa, 0x22, 0x92, 0x33, 0x29, 0xa4, 0xb6, 0xc1, 0xdd, 0x9c, 0x73, 0x7e, 0xe7,
	0x3f, 0xe7, 0x8b, 0xdc, 0x66, 0x3c, 0x91, 0x29, 0x0f, 0x22, 0x1e, 0x0b, 0x4f, 0xbf, 0x3d, 0x58,
	0xa5, 0x5e, 0x08, 0x02, 0x8a, 0xb8, 0x70, 0xb3, 0x5c, 0x2a, 0x49, 0x9f, 0x34, 0x28, 0x57, 0xbf,
	0x5d, 0x58, 0xa5, 0xfd, 0x9b, 0x50, 0x86, 0x12, 0x11, 0xaf, 0x7a, 0x69, 0xba, 0xff, 0xac, 0x45,
	0x33, 0xe3, 0x39, 0x4f, 0x6b, 0xc9, 0xfe, 0x5d, 0x0b, 0x94, 0x43, 0xc2, 0x37, 0x9f, 0xbf, 0x43,
	0x1c, 0x46, 0xaa, 0x66, 0x87, 0xa5, 0x49, 0xae, 0xde, 0xe8, 0x82, 0x3e, 0x28, 0xae, 0x80, 0xbe,
	0x24, 0x96, 0x16, 0xb3, 0xcd, 0x81, 0x39, 0xea, 0x8d, 0x1d, 0xf7, 0x7c, 0x81, 0xee, 0x7b, 0xa4,
	0x26, 0xdd, 0xed, 0xef, 0xa7, 0x06, 0xab, 0x73, 0x28, 0x23, 0x8f, 0x8a, 0x94, 0xe7, 0x6a, 0x2a,
	0x85, 0xca, 0x79, 0xa0, 0xec, 0x0b, 0x14, 0xb9, 0x6f, 0x13, 0x39, 0x7c, 0xdd, 0xcc, 0x61, 0xc7,
	0x12, 0xf4, 0x15, 0xb1, 0x30, 0xaf, 0xb0, 0x3b, 0x83, 0xce, 0xa8, 0x37, 0x1e, 0xfd, 0x47, 0x6c,
	0x5a, 0xf9, 0x17, 0x62, 0x29, 0x59, 0x9d, 0x37, 0xfc, 0x79, 0x41, 0xae, 0xff, 0x0d, 0xd2, 0x3b,
	0x72, 0x8d, 0x61, 0x06, 0x4b, 0xc8, 0x41, 0x04, 0xb0, 0x98, 0x61, 0xcb, 0x97, 0xec, 0xc4, 0x4f,
	0x6d, 0xf2, 0x10, 0x7d, 0x8b, 0x19, 0x36, 0xd4, 0x65, 0x07, 0x93, 0x0e, 0x48, 0xcf, 0x4f, 0x64,
	0xf0, 0x75, 0x8e, 0x53, 0xb5, 0x3b, 0x18, 0x6d, 0xba, 0xe8, 0x3d, 0x79, 0xac, 0x4d, 0x5e, 0x44,
	0xaf, 0x55, 0xcd, 0x75, 0xf1, 0xa3, 0xd3, 0x40, 0x45, 0xa7, 0xb1, 0x78, 0x27, 0xb0, 0xce, 0x09,
	0x4f, 0xb8, 0x08, 0xc0, 0x7e, 0xa0, 0xe9, 0x93, 0x00, 0x9d, 0x93, 0x2b, 0x5c, 0xea, 0x47, 0xbd,
	0x53, 0xdb, 0xc2, 0x69, 0xdf, 0xb6, 0x0d, 0x88, 0x35, 0x58, 0x76, 0x94, 0x39, 0x64, 0xe4, 0xe6,
	0xdc, 0x2e, 0xaa, 0xce, 0xb9, 0x1f, 0xbf, 0x2d, 0xa4, 0xa8, 0x87, 0x73, 0x30, 0xb1, 0xf3, 0x8d,
	0x82, 0x40, 0x7e, 0x81, 0x39, 0xac, 0x71, 0x2e, 0x97, 0xac, 0xe9, 0x9a, 0x4c, 0xb7, 0xa5, 0x63,
	0xee, 0x4a, 0xc7, 0xfc, 0x53, 0x3a, 0xe6, 0x8f, 0xbd, 0x63, 0xec, 0xf6, 0x8e, 0xf1, 0x6b, 0xef,
	0x18, 0x9f, 0x9e, 0x87, 0xb1, 0x8a, 0xbe, 0xf9, 0x6e, 0x20, 0x53, 0xef, 0xcc, 0xb1, 0xae, 0xf1,
	0x5c, 0xd5, 0x26, 0x83, 0xc2, 0xb7, 0xf0, 0x4e, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2f,
	0xa8, 0xd0, 0x2c, 0x4e, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.SmartContract != nil {
		{
			size, err := m.SmartContract.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RelayWeights != nil {
		{
			size, err := m.RelayWeights.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.MinOnChainBalance) > 0 {
		i -= len(m.MinOnChainBalance)
		copy(dAtA[i:], m.MinOnChainBalance)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MinOnChainBalance)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BlockHashAtHeight) > 0 {
		i -= len(m.BlockHashAtHeight)
		copy(dAtA[i:], m.BlockHashAtHeight)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BlockHashAtHeight)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.ChainID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisSmartContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSmartContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSmartContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BytecodeHex) > 0 {
		i -= len(m.BytecodeHex)
		copy(dAtA[i:], m.BytecodeHex)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BytecodeHex)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AbiJson) > 0 {
		i -= len(m.AbiJson)
		copy(dAtA[i:], m.AbiJson)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AbiJson)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.SmartContract != nil {
		l = m.SmartContract.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ChainID != 0 {
		n += 1 + sovGenesis(uint64(m.ChainID))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovGenesis(uint64(m.BlockHeight))
	}
	l = len(m.BlockHashAtHeight)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.MinOnChainBalance)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.RelayWeights != nil {
		l = m.RelayWeights.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisSmartContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AbiJson)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.BytecodeHex)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContract", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SmartContract == nil {
				m.SmartContract = &GenesisSmartContract{}
			}
			if err := m.SmartContract.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, &GenesisChainInfo{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHashAtHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHashAtHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinOnChainBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinOnChainBalance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayWeights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RelayWeights == nil {
				m.RelayWeights = &RelayWeights{}
			}
			if err := m.RelayWeights.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisSmartContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisSmartContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSmartContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbiJson", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbiJson = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytecodeHex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytecodeHex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)