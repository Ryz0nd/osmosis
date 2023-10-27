// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/swap_route.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type SwapAmountInRoute struct {
	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty" yaml:"token_out_denom"`
}

func (m *SwapAmountInRoute) Reset()         { *m = SwapAmountInRoute{} }
func (m *SwapAmountInRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountInRoute) ProtoMessage()    {}
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_cddd97a9a05492a8, []int{0}
}
func (m *SwapAmountInRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountInRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountInRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountInRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountInRoute.Merge(m, src)
}
func (m *SwapAmountInRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountInRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountInRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountInRoute proto.InternalMessageInfo

func (m *SwapAmountInRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *SwapAmountInRoute) GetTokenOutDenom() string {
	if m != nil {
		return m.TokenOutDenom
	}
	return ""
}

type SwapAmountOutRoute struct {
	PoolId       uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	TokenInDenom string `protobuf:"bytes,2,opt,name=token_in_denom,json=tokenInDenom,proto3" json:"token_in_denom,omitempty" yaml:"token_in_denom"`
}

func (m *SwapAmountOutRoute) Reset()         { *m = SwapAmountOutRoute{} }
func (m *SwapAmountOutRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountOutRoute) ProtoMessage()    {}
func (*SwapAmountOutRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_cddd97a9a05492a8, []int{1}
}
func (m *SwapAmountOutRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountOutRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountOutRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountOutRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountOutRoute.Merge(m, src)
}
func (m *SwapAmountOutRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountOutRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountOutRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountOutRoute proto.InternalMessageInfo

func (m *SwapAmountOutRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *SwapAmountOutRoute) GetTokenInDenom() string {
	if m != nil {
		return m.TokenInDenom
	}
	return ""
}

type SwapAmountInSplitRoute struct {
	Pools         []SwapAmountInRoute   `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools" yaml:"pools"`
	TokenInAmount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=token_in_amount,json=tokenInAmount,proto3,customtype=cosmossdk.io/math.Int" json:"token_in_amount" yaml:"token_in_amount"`
}

func (m *SwapAmountInSplitRoute) Reset()         { *m = SwapAmountInSplitRoute{} }
func (m *SwapAmountInSplitRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountInSplitRoute) ProtoMessage()    {}
func (*SwapAmountInSplitRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_cddd97a9a05492a8, []int{2}
}
func (m *SwapAmountInSplitRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountInSplitRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountInSplitRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountInSplitRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountInSplitRoute.Merge(m, src)
}
func (m *SwapAmountInSplitRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountInSplitRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountInSplitRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountInSplitRoute proto.InternalMessageInfo

func (m *SwapAmountInSplitRoute) GetPools() []SwapAmountInRoute {
	if m != nil {
		return m.Pools
	}
	return nil
}

type SwapAmountOutSplitRoute struct {
	Pools          []SwapAmountOutRoute  `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools" yaml:"pools"`
	TokenOutAmount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=token_out_amount,json=tokenOutAmount,proto3,customtype=cosmossdk.io/math.Int" json:"token_out_amount" yaml:"token_out_amount"`
}

func (m *SwapAmountOutSplitRoute) Reset()         { *m = SwapAmountOutSplitRoute{} }
func (m *SwapAmountOutSplitRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountOutSplitRoute) ProtoMessage()    {}
func (*SwapAmountOutSplitRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_cddd97a9a05492a8, []int{3}
}
func (m *SwapAmountOutSplitRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountOutSplitRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountOutSplitRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountOutSplitRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountOutSplitRoute.Merge(m, src)
}
func (m *SwapAmountOutSplitRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountOutSplitRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountOutSplitRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountOutSplitRoute proto.InternalMessageInfo

func (m *SwapAmountOutSplitRoute) GetPools() []SwapAmountOutRoute {
	if m != nil {
		return m.Pools
	}
	return nil
}

func init() {
	proto.RegisterType((*SwapAmountInRoute)(nil), "osmosis.poolmanager.v1beta1.SwapAmountInRoute")
	proto.RegisterType((*SwapAmountOutRoute)(nil), "osmosis.poolmanager.v1beta1.SwapAmountOutRoute")
	proto.RegisterType((*SwapAmountInSplitRoute)(nil), "osmosis.poolmanager.v1beta1.SwapAmountInSplitRoute")
	proto.RegisterType((*SwapAmountOutSplitRoute)(nil), "osmosis.poolmanager.v1beta1.SwapAmountOutSplitRoute")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/swap_route.proto", fileDescriptor_cddd97a9a05492a8)
}

var fileDescriptor_cddd97a9a05492a8 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcb, 0xaa, 0xd3, 0x40,
	0x1c, 0xc6, 0x33, 0x5e, 0x8e, 0x38, 0x1e, 0xab, 0x86, 0x73, 0xa9, 0x47, 0x48, 0x4a, 0x56, 0x05,
	0x75, 0x86, 0xd6, 0x45, 0xc5, 0x8d, 0x18, 0xdc, 0x64, 0x55, 0x4c, 0x77, 0x75, 0x11, 0x26, 0x4d,
	0x48, 0x43, 0x93, 0x99, 0xd0, 0x99, 0xb4, 0x76, 0x2b, 0x3e, 0x80, 0x8f, 0xd5, 0x65, 0x37, 0x82,
	0x74, 0x11, 0xa4, 0x7d, 0x83, 0x3c, 0x81, 0xe4, 0x66, 0x93, 0x0a, 0x55, 0xce, 0x6e, 0x2e, 0xff,
	0x6f, 0xfe, 0xbf, 0xef, 0x9b, 0x19, 0xf8, 0x8a, 0xf1, 0x90, 0x71, 0x9f, 0xe3, 0x88, 0xb1, 0x20,
	0x24, 0x94, 0x78, 0xee, 0x1c, 0x2f, 0x7a, 0xb6, 0x2b, 0x48, 0x0f, 0xf3, 0x25, 0x89, 0xac, 0x39,
	0x8b, 0x85, 0x8b, 0xa2, 0x39, 0x13, 0x4c, 0x7e, 0x51, 0x56, 0xa3, 0x5a, 0x35, 0x2a, 0xab, 0x6f,
	0x2e, 0x3c, 0xe6, 0xb1, 0xbc, 0x0e, 0x67, 0xa3, 0x42, 0xa2, 0x7d, 0x03, 0xf0, 0xd9, 0x68, 0x49,
	0xa2, 0x0f, 0x21, 0x8b, 0xa9, 0x30, 0xa8, 0x99, 0x1d, 0x27, 0xbf, 0x84, 0x0f, 0xb2, 0x23, 0x2c,
	0xdf, 0x69, 0x83, 0x0e, 0xe8, 0xde, 0xd3, 0xe5, 0x34, 0x51, 0x5b, 0x2b, 0x12, 0x06, 0xef, 0xb4,
	0x72, 0x43, 0x33, 0xcf, 0xb2, 0x91, 0xe1, 0xc8, 0x3a, 0x7c, 0x22, 0xd8, 0xcc, 0xa5, 0x16, 0x8b,
	0x85, 0xe5, 0xb8, 0x94, 0x85, 0xed, 0x3b, 0x1d, 0xd0, 0x7d, 0xa8, 0xdf, 0xa4, 0x89, 0x7a, 0x55,
	0x88, 0x8e, 0x0a, 0x34, 0xf3, 0x71, 0xbe, 0x32, 0x8c, 0xc5, 0xc7, 0x7c, 0xfe, 0x15, 0x40, 0xf9,
	0x80, 0x31, 0x8c, 0xc5, 0x2d, 0x38, 0xde, 0xc3, 0x56, 0xd1, 0xc6, 0xa7, 0x0d, 0x8c, 0xe7, 0x69,
	0xa2, 0x5e, 0xd6, 0x31, 0xaa, 0x7d, 0xcd, 0x3c, 0xcf, 0x17, 0x0c, 0x5a, 0x40, 0xfc, 0x00, 0xf0,
	0xaa, 0x9e, 0xc5, 0x28, 0x0a, 0xfc, 0x12, 0x64, 0x0c, 0xef, 0x67, 0x5d, 0x78, 0x1b, 0x74, 0xee,
	0x76, 0x1f, 0xf5, 0x11, 0x3a, 0x91, 0x34, 0xfa, 0x2b, 0x4f, 0xfd, 0x62, 0x9d, 0xa8, 0x52, 0x9a,
	0xa8, 0xe7, 0x07, 0x74, 0xae, 0x99, 0xc5, 0x91, 0xb2, 0x55, 0xe5, 0xe7, 0x53, 0x8b, 0xe4, 0xb2,
	0x12, 0x7c, 0x90, 0xa9, 0xb6, 0x89, 0x7a, 0x39, 0xc9, 0xbb, 0x71, 0x67, 0x86, 0x7c, 0x86, 0x43,
	0x22, 0xa6, 0xc8, 0xa0, 0xe2, 0x38, 0xdc, 0x3f, 0xea, 0x2a, 0x5c, 0x83, 0x16, 0x10, 0xda, 0x16,
	0xc0, 0xeb, 0x46, 0xb8, 0x35, 0x63, 0x9f, 0x9b, 0xc6, 0xf0, 0x7f, 0x1a, 0xab, 0x6e, 0xe8, 0xb4,
	0x33, 0x1b, 0x3e, 0x3d, 0x5c, 0x7c, 0xc3, 0xda, 0xdb, 0x7f, 0x59, 0xbb, 0x3e, 0x7e, 0x37, 0x95,
	0xb7, 0x56, 0xf5, 0x70, 0x0a, 0x10, 0xfd, 0xd3, 0x7a, 0xa7, 0x80, 0xcd, 0x4e, 0x01, 0xbf, 0x76,
	0x0a, 0xf8, 0xbe, 0x57, 0xa4, 0xcd, 0x5e, 0x91, 0x7e, 0xee, 0x15, 0x69, 0x3c, 0xf0, 0x7c, 0x31,
	0x8d, 0x6d, 0x34, 0x61, 0x21, 0x2e, 0x5d, 0xbd, 0x0e, 0x88, 0xcd, 0xab, 0x09, 0x5e, 0xf4, 0x7b,
	0xf8, 0x4b, 0xe3, 0x67, 0x89, 0x55, 0xe4, 0x72, 0xfb, 0x2c, 0xff, 0x1a, 0x6f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0xb8, 0x73, 0x26, 0x7d, 0x03, 0x00, 0x00,
}

func (m *SwapAmountInRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountInRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountInRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOutDenom) > 0 {
		i -= len(m.TokenOutDenom)
		copy(dAtA[i:], m.TokenOutDenom)
		i = encodeVarintSwapRoute(dAtA, i, uint64(len(m.TokenOutDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintSwapRoute(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SwapAmountOutRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountOutRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountOutRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenInDenom) > 0 {
		i -= len(m.TokenInDenom)
		copy(dAtA[i:], m.TokenInDenom)
		i = encodeVarintSwapRoute(dAtA, i, uint64(len(m.TokenInDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintSwapRoute(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SwapAmountInSplitRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountInSplitRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountInSplitRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenInAmount.Size()
		i -= size
		if _, err := m.TokenInAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwapRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwapRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SwapAmountOutSplitRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountOutSplitRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountOutSplitRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenOutAmount.Size()
		i -= size
		if _, err := m.TokenOutAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSwapRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSwapRoute(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSwapRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovSwapRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SwapAmountInRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovSwapRoute(uint64(m.PoolId))
	}
	l = len(m.TokenOutDenom)
	if l > 0 {
		n += 1 + l + sovSwapRoute(uint64(l))
	}
	return n
}

func (m *SwapAmountOutRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovSwapRoute(uint64(m.PoolId))
	}
	l = len(m.TokenInDenom)
	if l > 0 {
		n += 1 + l + sovSwapRoute(uint64(l))
	}
	return n
}

func (m *SwapAmountInSplitRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovSwapRoute(uint64(l))
		}
	}
	l = m.TokenInAmount.Size()
	n += 1 + l + sovSwapRoute(uint64(l))
	return n
}

func (m *SwapAmountOutSplitRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovSwapRoute(uint64(l))
		}
	}
	l = m.TokenOutAmount.Size()
	n += 1 + l + sovSwapRoute(uint64(l))
	return n
}

func sovSwapRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSwapRoute(x uint64) (n int) {
	return sovSwapRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SwapAmountInRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwapRoute
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
			return fmt.Errorf("proto: SwapAmountInRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountInRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOutDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwapRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwapRoute
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
func (m *SwapAmountOutRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwapRoute
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
			return fmt.Errorf("proto: SwapAmountOutRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountOutRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenInDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwapRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwapRoute
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
func (m *SwapAmountInSplitRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwapRoute
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
			return fmt.Errorf("proto: SwapAmountInSplitRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountInSplitRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, SwapAmountInRoute{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenInAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwapRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwapRoute
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
func (m *SwapAmountOutSplitRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSwapRoute
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
			return fmt.Errorf("proto: SwapAmountOutSplitRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountOutSplitRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, SwapAmountOutRoute{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSwapRoute
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
				return ErrInvalidLengthSwapRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSwapRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOutAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSwapRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSwapRoute
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
func skipSwapRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSwapRoute
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
					return 0, ErrIntOverflowSwapRoute
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
					return 0, ErrIntOverflowSwapRoute
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
				return 0, ErrInvalidLengthSwapRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSwapRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSwapRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSwapRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSwapRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSwapRoute = fmt.Errorf("proto: unexpected end of group")
)
