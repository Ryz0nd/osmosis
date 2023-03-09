// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v2/balancer/balancerPool.proto

// this is a legacy package that requires additional migration logic
// in order to use the correct packge. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package balancer

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	balancer "github.com/osmosis-labs/osmosis/v15/x/gamm/pool-models/balancer"
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

// PoolParams defined the parameters that will be managed by the pool
// governance in the future. This params are not managed by the chain
// governance. Instead they will be managed by the token holders of the pool.
// The pool's token holders are specified in future_pool_governor.
type PoolParams struct {
	SwapFee                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=swap_fee,json=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee" yaml:"swap_fee"`
	ExitFee                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=exit_fee,json=exitFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exit_fee" yaml:"exit_fee"`
	SmoothWeightChangeParams *balancer.SmoothWeightChangeParams     `protobuf:"bytes,3,opt,name=smooth_weight_change_params,json=smoothWeightChangeParams,proto3" json:"smooth_weight_change_params,omitempty" yaml:"smooth_weight_change_params"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_36e6bec4fd4d490e, []int{0}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func (m *PoolParams) GetSmoothWeightChangeParams() *balancer.SmoothWeightChangeParams {
	if m != nil {
		return m.SmoothWeightChangeParams
	}
	return nil
}

type Pool struct {
	Address    string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Id         uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PoolParams PoolParams `protobuf:"bytes,3,opt,name=pool_params,json=poolParams,proto3" json:"pool_params" yaml:"balancer_pool_params"`
	// This string specifies who will govern the pool in the future.
	// Valid forms of this are:
	// {token name},{duration}
	// {duration}
	// where {token name} if specified is the token which determines the
	// governor, and if not specified is the LP token for this pool.duration is
	// a time specified as 0w,1w,2w, etc. which specifies how long the token
	// would need to be locked up to count in governance. 0w means no lockup.
	// TODO: Further improve these docs
	FuturePoolGovernor string `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
	// sum of all LP tokens sent out
	TotalShares types.Coin `protobuf:"bytes,5,opt,name=total_shares,json=totalShares,proto3" json:"total_shares" yaml:"total_shares"`
	// These are assumed to be sorted by denomiation.
	// They contain the pool asset and the information about the weight
	PoolAssets []balancer.PoolAsset `protobuf:"bytes,6,rep,name=pool_assets,json=poolAssets,proto3" json:"pool_assets" yaml:"pool_assets"`
	// sum of all non-normalized pool weights
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=total_weight,json=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_weight" yaml:"total_weight"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_36e6bec4fd4d490e, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PoolParams)(nil), "osmosis.gamm.v2.balancer.PoolParams")
	proto.RegisterType((*Pool)(nil), "osmosis.gamm.v2.balancer.Pool")
}

func init() {
	proto.RegisterFile("osmosis/gamm/v2/balancer/balancerPool.proto", fileDescriptor_36e6bec4fd4d490e)
}

var fileDescriptor_36e6bec4fd4d490e = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0xd3, 0xb4, 0x81, 0x0d, 0x2a, 0xc2, 0xf4, 0xe0, 0xb6, 0x92, 0xb7, 0x32, 0x08, 0x55,
	0x40, 0xd6, 0x6a, 0x80, 0x4b, 0x39, 0x35, 0xe5, 0x47, 0xbd, 0x15, 0x57, 0x08, 0x81, 0x2a, 0x59,
	0xeb, 0x78, 0x6b, 0x1b, 0xec, 0xac, 0xe5, 0xdd, 0xa4, 0xed, 0x1b, 0x70, 0xe4, 0x08, 0xb7, 0x3e,
	0x04, 0x17, 0xde, 0xa0, 0xe2, 0xd4, 0x23, 0xe2, 0x60, 0xa1, 0x96, 0x27, 0xc8, 0x13, 0xa0, 0xfd,
	0x71, 0x9a, 0x56, 0x0d, 0x12, 0xe2, 0xe4, 0x9d, 0x9d, 0x6f, 0xbe, 0x99, 0xf9, 0x66, 0xbc, 0xe0,
	0x01, 0x65, 0x19, 0x65, 0x09, 0x73, 0x23, 0x9c, 0x65, 0xee, 0xb0, 0xe3, 0x06, 0x38, 0xc5, 0xfd,
	0x1e, 0x29, 0xc6, 0x87, 0x6d, 0x4a, 0x53, 0x94, 0x17, 0x94, 0x53, 0xd3, 0xd2, 0x60, 0x24, 0xc0,
	0x68, 0xd8, 0x41, 0x15, 0x66, 0x69, 0xb1, 0x27, 0x5d, 0xbe, 0xc4, 0xb9, 0xca, 0x50, 0x41, 0x4b,
	0x0b, 0x11, 0x8d, 0xa8, 0xba, 0x17, 0x27, 0x7d, 0x6b, 0x47, 0x94, 0x46, 0x29, 0x71, 0xa5, 0x15,
	0x0c, 0xf6, 0xdc, 0x70, 0x50, 0x60, 0x9e, 0xd0, 0xbe, 0xf6, 0xc3, 0xcb, 0x7e, 0x9e, 0x64, 0x84,
	0x71, 0x9c, 0xe5, 0x15, 0x81, 0x4a, 0xe2, 0xe2, 0x01, 0x8f, 0xdd, 0xe1, 0x5a, 0x40, 0x38, 0x5e,
	0x93, 0xc6, 0x25, 0x7f, 0x80, 0x19, 0x19, 0xfb, 0x7b, 0x34, 0xa9, 0x12, 0x3c, 0xbe, 0xd0, 0x78,
	0x4e, 0x69, 0xda, 0xce, 0x68, 0x48, 0x52, 0xf6, 0x37, 0x05, 0x9c, 0xdf, 0x75, 0x00, 0x84, 0xb9,
	0x8d, 0x0b, 0x9c, 0x31, 0x73, 0x17, 0x5c, 0x63, 0xfb, 0x38, 0xf7, 0xf7, 0x08, 0xb1, 0x8c, 0x15,
	0x63, 0xf5, 0x7a, 0x77, 0xe3, 0xb8, 0x84, 0xb5, 0x9f, 0x25, 0xbc, 0x17, 0x25, 0x3c, 0x1e, 0x04,
	0xa8, 0x47, 0x33, 0x2d, 0x87, 0xfe, 0xb4, 0x59, 0xf8, 0xc1, 0xe5, 0x87, 0x39, 0x61, 0xe8, 0x19,
	0xe9, 0x8d, 0x4a, 0x78, 0xf3, 0x10, 0x67, 0xe9, 0xba, 0x53, 0xf1, 0x38, 0x5e, 0x53, 0x1c, 0x5f,
	0x10, 0x22, 0xd8, 0xc9, 0x41, 0xc2, 0x25, 0x7b, 0xfd, 0xff, 0xd8, 0x2b, 0x1e, 0xc7, 0x6b, 0x8a,
	0xa3, 0x60, 0xff, 0x62, 0x80, 0x65, 0x96, 0x51, 0xca, 0x63, 0x7f, 0x9f, 0x24, 0x51, 0xcc, 0xfd,
	0x5e, 0x8c, 0xfb, 0x11, 0xf1, 0x73, 0xd9, 0x9b, 0x35, 0xb3, 0x62, 0xac, 0xb6, 0x3a, 0x08, 0x5d,
	0x9c, 0xb9, 0x12, 0x12, 0xed, 0xc8, 0xc0, 0x37, 0x32, 0x6e, 0x53, 0x86, 0x29, 0x45, 0xba, 0xf7,
	0x8f, 0x4b, 0x68, 0x8c, 0x4a, 0xe8, 0xe8, 0xae, 0xa6, 0x27, 0x70, 0x3c, 0x8b, 0x4d, 0x61, 0x71,
	0xbe, 0x35, 0x40, 0x43, 0xc8, 0x6c, 0x3e, 0x04, 0x4d, 0x1c, 0x86, 0x05, 0x61, 0x4c, 0xeb, 0x6b,
	0x8e, 0x4a, 0x38, 0xaf, 0xb8, 0xb5, 0xc3, 0xf1, 0x2a, 0x88, 0x39, 0x0f, 0xea, 0x49, 0x28, 0xa5,
	0x6a, 0x78, 0xf5, 0x24, 0x34, 0xdf, 0x83, 0x96, 0x18, 0xec, 0xc5, 0x8e, 0xee, 0xa2, 0x69, 0x5b,
	0x8c, 0xce, 0x27, 0xdb, 0xbd, 0x23, 0x94, 0x1e, 0x95, 0x70, 0x59, 0xe5, 0xaa, 0x20, 0xfe, 0x04,
	0x9f, 0xe3, 0x81, 0xfc, 0x7c, 0x15, 0x5e, 0x81, 0x85, 0xbd, 0x01, 0x1f, 0x14, 0x44, 0x41, 0x22,
	0x3a, 0x24, 0x45, 0x9f, 0x16, 0x56, 0x43, 0x96, 0x0d, 0xcf, 0xa9, 0xae, 0x42, 0x39, 0x9e, 0xa9,
	0xae, 0x45, 0x05, 0x2f, 0xf5, 0xa5, 0xf9, 0x16, 0xdc, 0xe0, 0x94, 0xe3, 0xd4, 0x67, 0x31, 0x2e,
	0x08, 0xb3, 0x66, 0x65, 0xfd, 0x8b, 0x48, 0xff, 0x5e, 0x62, 0xb3, 0xc7, 0x03, 0xd9, 0xa4, 0x49,
	0xbf, 0xbb, 0xac, 0x8b, 0xbe, 0xad, 0x32, 0x4d, 0x06, 0x3b, 0x5e, 0x4b, 0x9a, 0x3b, 0xd2, 0x32,
	0x77, 0xb5, 0x32, 0x98, 0x31, 0xc2, 0x99, 0x35, 0xb7, 0x32, 0xb3, 0xda, 0xea, 0xc0, 0xab, 0x67,
	0x2d, 0x6a, 0xda, 0x10, 0xb8, 0xee, 0x92, 0xe6, 0x37, 0x15, 0xff, 0x04, 0x83, 0xd6, 0x42, 0xc2,
	0x98, 0x19, 0x57, 0x85, 0xab, 0xb9, 0x5b, 0x4d, 0xa9, 0xc1, 0xf3, 0x7f, 0x58, 0xde, 0xad, 0x3e,
	0xbf, 0xdc, 0x87, 0xe2, 0xaa, 0xfa, 0x50, 0x4b, 0xb3, 0x7e, 0xeb, 0xe3, 0x11, 0xac, 0x7d, 0x3e,
	0x82, 0xc6, 0xf7, 0xaf, 0xed, 0x59, 0x51, 0xe8, 0x56, 0xf7, 0xf5, 0xf1, 0xa9, 0x6d, 0x9c, 0x9c,
	0xda, 0xc6, 0xaf, 0x53, 0xdb, 0xf8, 0x74, 0x66, 0xd7, 0x4e, 0xce, 0xec, 0xda, 0x8f, 0x33, 0xbb,
	0xf6, 0xee, 0xe9, 0x44, 0x62, 0xdd, 0x69, 0x3b, 0xc5, 0x01, 0xab, 0x0c, 0x77, 0xb8, 0xf6, 0xc4,
	0x3d, 0xa8, 0x5e, 0x42, 0x59, 0xc8, 0xf8, 0x0d, 0x08, 0xe6, 0xe4, 0x03, 0xf0, 0xe8, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0x17, 0xb2, 0xa1, 0x31, 0x05, 0x00, 0x00,
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SmoothWeightChangeParams != nil {
		{
			size, err := m.SmoothWeightChangeParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBalancerPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.ExitFee.Size()
		i -= size
		if _, err := m.ExitFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalancerPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintBalancerPool(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintBalancerPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBalancerPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBalancerPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovBalancerPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapFee.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	l = m.ExitFee.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	if m.SmoothWeightChangeParams != nil {
		l = m.SmoothWeightChangeParams.Size()
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovBalancerPool(uint64(m.Id))
	}
	l = m.PoolParams.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovBalancerPool(uint64(l))
		}
	}
	l = m.TotalWeight.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	return n
}

func sovBalancerPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBalancerPool(x uint64) (n int) {
	return sovBalancerPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalancerPool
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
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExitFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmoothWeightChangeParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SmoothWeightChangeParams == nil {
				m.SmoothWeightChangeParams = &balancer.SmoothWeightChangeParams{}
			}
			if err := m.SmoothWeightChangeParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalancerPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalancerPool
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalancerPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, balancer.PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalancerPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalancerPool
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
func skipBalancerPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalancerPool
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
					return 0, ErrIntOverflowBalancerPool
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
					return 0, ErrIntOverflowBalancerPool
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
				return 0, ErrInvalidLengthBalancerPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBalancerPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBalancerPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBalancerPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalancerPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBalancerPool = fmt.Errorf("proto: unexpected end of group")
)
