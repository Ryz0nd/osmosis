// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/instantiate_msg.proto

package msg

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// ===================== InstantiateMsg
type InstantiateMsg struct {
	// pool_asset_denoms is the list of asset denoms that are initialized
	// at pool creation time.
	PoolAssetDenoms []string `protobuf:"bytes,1,rep,name=pool_asset_denoms,json=poolAssetDenoms,proto3" json:"pool_asset_denoms,omitempty"`
}

func (m *InstantiateMsg) Reset()         { *m = InstantiateMsg{} }
func (m *InstantiateMsg) String() string { return proto.CompactTextString(m) }
func (*InstantiateMsg) ProtoMessage()    {}
func (*InstantiateMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_834d7406456b7712, []int{0}
}
func (m *InstantiateMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstantiateMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstantiateMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstantiateMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstantiateMsg.Merge(m, src)
}
func (m *InstantiateMsg) XXX_Size() int {
	return m.Size()
}
func (m *InstantiateMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_InstantiateMsg.DiscardUnknown(m)
}

var xxx_messageInfo_InstantiateMsg proto.InternalMessageInfo

func (m *InstantiateMsg) GetPoolAssetDenoms() []string {
	if m != nil {
		return m.PoolAssetDenoms
	}
	return nil
}

func init() {
	proto.RegisterType((*InstantiateMsg)(nil), "osmosis.cosmwasmpool.v1beta1.InstantiateMsg")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/instantiate_msg.proto", fileDescriptor_834d7406456b7712)
}

var fileDescriptor_834d7406456b7712 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc8, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x4f, 0xce, 0x2f, 0xce, 0x2d, 0x4f, 0x2c, 0xce, 0x2d, 0xc8, 0xcf, 0xcf,
	0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0xcf,
	0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x2b, 0xc9, 0x4c, 0x2c, 0x49, 0x8d, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0xea, 0xd4, 0x43, 0xd6, 0xa9, 0x07, 0xd5, 0x29, 0x25,
	0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa8, 0x0f, 0x62, 0x41, 0xf4, 0x48, 0xc9, 0x25, 0x83, 0x35,
	0xe9, 0x27, 0x25, 0x16, 0xa7, 0xc2, 0x2d, 0x49, 0xce, 0xcf, 0xcc, 0x83, 0xc8, 0x2b, 0xd9, 0x70,
	0xf1, 0x79, 0x22, 0x2c, 0xf3, 0x2d, 0x4e, 0x17, 0xd2, 0xe2, 0x12, 0x04, 0x99, 0x1b, 0x9f, 0x58,
	0x5c, 0x9c, 0x5a, 0x12, 0x9f, 0x92, 0x9a, 0x97, 0x9f, 0x5b, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x19, 0xc4, 0x0f, 0x92, 0x70, 0x04, 0x89, 0xbb, 0x80, 0x85, 0x9d, 0xa2, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x21, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39,
	0x3f, 0x57, 0x1f, 0xea, 0x6c, 0xdd, 0x9c, 0xc4, 0xa4, 0x62, 0x18, 0x47, 0xbf, 0xac, 0xcc, 0xc8,
	0x58, 0xbf, 0x02, 0x35, 0x10, 0x60, 0x1c, 0xfd, 0xdc, 0xe2, 0xf4, 0x24, 0x36, 0xb0, 0x03, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x18, 0xe2, 0xd1, 0x30, 0x01, 0x00, 0x00,
}

func (m *InstantiateMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstantiateMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstantiateMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolAssetDenoms) > 0 {
		for iNdEx := len(m.PoolAssetDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PoolAssetDenoms[iNdEx])
			copy(dAtA[i:], m.PoolAssetDenoms[iNdEx])
			i = encodeVarintInstantiateMsg(dAtA, i, uint64(len(m.PoolAssetDenoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintInstantiateMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovInstantiateMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InstantiateMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolAssetDenoms) > 0 {
		for _, s := range m.PoolAssetDenoms {
			l = len(s)
			n += 1 + l + sovInstantiateMsg(uint64(l))
		}
	}
	return n
}

func sovInstantiateMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInstantiateMsg(x uint64) (n int) {
	return sovInstantiateMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InstantiateMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInstantiateMsg
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
			return fmt.Errorf("proto: InstantiateMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstantiateMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssetDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInstantiateMsg
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
				return ErrInvalidLengthInstantiateMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInstantiateMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssetDenoms = append(m.PoolAssetDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInstantiateMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInstantiateMsg
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
func skipInstantiateMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInstantiateMsg
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
					return 0, ErrIntOverflowInstantiateMsg
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
					return 0, ErrIntOverflowInstantiateMsg
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
				return 0, ErrInvalidLengthInstantiateMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInstantiateMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInstantiateMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInstantiateMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInstantiateMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInstantiateMsg = fmt.Errorf("proto: unexpected end of group")
)
