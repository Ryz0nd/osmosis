// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	migration "github.com/osmosis-labs/osmosis/v23/x/gamm/types/migration"
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

// Params holds parameters for the incentives module
type Params struct {
	PoolCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee" yaml:"pool_creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPoolCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PoolCreationFee
	}
	return nil
}

// GenesisState defines the gamm module's genesis state.
type GenesisState struct {
	Pools []*types1.Any `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// will be renamed to next_pool_id in an upcoming version
	NextPoolNumber   uint64                      `protobuf:"varint,2,opt,name=next_pool_number,json=nextPoolNumber,proto3" json:"next_pool_number,omitempty"`
	Params           Params                      `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
	MigrationRecords *migration.MigrationRecords `protobuf:"bytes,4,opt,name=migration_records,json=migrationRecords,proto3" json:"migration_records,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a324eb7f1dd793e, []int{1}
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

func (m *GenesisState) GetPools() []*types1.Any {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *GenesisState) GetNextPoolNumber() uint64 {
	if m != nil {
		return m.NextPoolNumber
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMigrationRecords() *migration.MigrationRecords {
	if m != nil {
		return m.MigrationRecords
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.gamm.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "osmosis.gamm.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/gamm/v1beta1/genesis.proto", fileDescriptor_5a324eb7f1dd793e)
}

var fileDescriptor_5a324eb7f1dd793e = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x36, 0x8d, 0x84, 0x8b, 0xa0, 0xb5, 0x72, 0x70, 0x2b, 0xe4, 0x04, 0x1f, 0x90,
	0x2f, 0xd9, 0xa5, 0xa9, 0xb8, 0xf4, 0x46, 0x2a, 0x81, 0x40, 0x80, 0x2a, 0xf7, 0xc6, 0x25, 0x5a,
	0x3b, 0x53, 0xd7, 0xc2, 0xeb, 0x89, 0x76, 0x37, 0x55, 0x7d, 0xe3, 0x11, 0x90, 0xb8, 0xf3, 0x00,
	0x9c, 0x79, 0x88, 0x8a, 0x53, 0x8f, 0x9c, 0x02, 0x4a, 0xde, 0x80, 0x27, 0x40, 0xfb, 0xc7, 0x08,
	0xd1, 0x9c, 0xec, 0x99, 0xf9, 0xcd, 0xb7, 0x33, 0xdf, 0xf8, 0x31, 0x4a, 0x8e, 0xb2, 0x94, 0xb4,
	0x60, 0x9c, 0xd3, 0xab, 0xa3, 0x0c, 0x14, 0x3b, 0xa2, 0x05, 0xd4, 0x20, 0x4b, 0x49, 0xe6, 0x02,
	0x15, 0x06, 0x7d, 0xc7, 0x10, 0xcd, 0x10, 0xc7, 0x1c, 0xf6, 0x0b, 0x2c, 0xd0, 0x00, 0x54, 0xff,
	0x59, 0xf6, 0xf0, 0xa0, 0x40, 0x2c, 0x2a, 0xa0, 0x26, 0xca, 0x16, 0x17, 0x94, 0xd5, 0x4d, 0x5b,
	0xca, 0x8d, 0xce, 0xd4, 0xf6, 0xd8, 0xc0, 0x95, 0x22, 0x1b, 0xd1, 0x8c, 0x49, 0xf8, 0x3b, 0x44,
	0x8e, 0x65, 0xed, 0xea, 0x8f, 0x37, 0x4e, 0x29, 0x2f, 0x99, 0x80, 0x99, 0x45, 0xe2, 0x2f, 0x9e,
	0xdf, 0x3b, 0x63, 0x82, 0x71, 0x19, 0x7c, 0xf6, 0xfc, 0xfd, 0x39, 0x62, 0x35, 0xcd, 0x05, 0x30,
	0x55, 0x62, 0x3d, 0xbd, 0x00, 0x08, 0xbd, 0xe1, 0x76, 0xb2, 0x3b, 0x3e, 0x20, 0xee, 0x61, 0xfd,
	0x54, 0xbb, 0x0b, 0x39, 0xc5, 0xb2, 0x9e, 0xbc, 0xb9, 0x59, 0x0e, 0x3a, 0xbf, 0x97, 0x83, 0xb0,
	0x61, 0xbc, 0x3a, 0x89, 0xef, 0x28, 0xc4, 0x5f, 0x7f, 0x0e, 0x92, 0xa2, 0x54, 0x97, 0x8b, 0x8c,
	0xe4, 0xc8, 0xdd, 0x06, 0xee, 0x33, 0x92, 0xb3, 0x0f, 0x54, 0x35, 0x73, 0x90, 0x46, 0x4c, 0xa6,
	0x0f, 0x75, 0xff, 0xa9, 0x6b, 0x7f, 0x01, 0x10, 0x7f, 0xdc, 0xf2, 0xef, 0xbf, 0xb4, 0xbe, 0x9e,
	0x2b, 0xa6, 0x20, 0x78, 0xe6, 0xef, 0x68, 0x46, 0xba, 0xc9, 0xfa, 0xc4, 0x5a, 0x47, 0x5a, 0xeb,
	0xc8, 0xf3, 0xba, 0x99, 0xdc, 0xfb, 0xfe, 0x6d, 0xb4, 0x73, 0x86, 0x58, 0xbd, 0x4a, 0x2d, 0x1d,
	0x24, 0xfe, 0x5e, 0x0d, 0xd7, 0x6a, 0x6a, 0xe6, 0xab, 0x17, 0x3c, 0x03, 0x11, 0x6e, 0x0d, 0xbd,
	0xa4, 0x9b, 0x3e, 0xd0, 0x79, 0xcd, 0xbe, 0x33, 0xd9, 0xe0, 0xc4, 0xef, 0xcd, 0x8d, 0x23, 0xe1,
	0xf6, 0xd0, 0x4b, 0x76, 0xc7, 0x8f, 0xc8, 0xa6, 0x43, 0x12, 0xeb, 0xda, 0xa4, 0xab, 0xd7, 0x4f,
	0x5d, 0x47, 0x70, 0xee, 0xef, 0xf3, 0xb2, 0x10, 0x76, 0x79, 0x01, 0x39, 0x8a, 0x99, 0x0c, 0xbb,
	0x46, 0xe6, 0xc9, 0x66, 0x99, 0xb7, 0x2d, 0x9e, 0x5a, 0x3a, 0xdd, 0xe3, 0xff, 0x65, 0x26, 0xaf,
	0x6f, 0x56, 0x91, 0x77, 0xbb, 0x8a, 0xbc, 0x5f, 0xab, 0xc8, 0xfb, 0xb4, 0x8e, 0x3a, 0xb7, 0xeb,
	0xa8, 0xf3, 0x63, 0x1d, 0x75, 0xde, 0x3f, 0xfd, 0xc7, 0x57, 0xa7, 0x3e, 0xaa, 0x58, 0x26, 0xdb,
	0x80, 0x5e, 0x8d, 0x8f, 0xe9, 0xb5, 0x3d, 0xbf, 0x71, 0x39, 0xeb, 0x19, 0x9b, 0x8e, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0xea, 0x3b, 0x1a, 0xc1, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
	if m.MigrationRecords != nil {
		{
			size, err := m.MigrationRecords.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
	dAtA[i] = 0x1a
	if m.NextPoolNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPoolNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextPoolNumber != 0 {
		n += 1 + sovGenesis(uint64(m.NextPoolNumber))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.MigrationRecords != nil {
		l = m.MigrationRecords.Size()
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
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
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
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
			m.Pools = append(m.Pools, &types1.Any{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPoolNumber", wireType)
			}
			m.NextPoolNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPoolNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MigrationRecords", wireType)
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
			if m.MigrationRecords == nil {
				m.MigrationRecords = &migration.MigrationRecords{}
			}
			if err := m.MigrationRecords.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
