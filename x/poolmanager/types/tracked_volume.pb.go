// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/tracked_volume.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type TrackedVolume struct {
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *TrackedVolume) Reset()         { *m = TrackedVolume{} }
func (m *TrackedVolume) String() string { return proto.CompactTextString(m) }
func (*TrackedVolume) ProtoMessage()    {}
func (*TrackedVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2e3e91de3baf1a, []int{0}
}
func (m *TrackedVolume) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrackedVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrackedVolume.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrackedVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedVolume.Merge(m, src)
}
func (m *TrackedVolume) XXX_Size() int {
	return m.Size()
}
func (m *TrackedVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedVolume.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedVolume proto.InternalMessageInfo

func (m *TrackedVolume) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackedVolume)(nil), "osmosis.poolmanager.v1beta1.TrackedVolume")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/tracked_volume.proto", fileDescriptor_0a2e3e91de3baf1a)
}

var fileDescriptor_0a2e3e91de3baf1a = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc8, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0xc8, 0xcf, 0xcf, 0xc9, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0x4e, 0x4d, 0x89,
	0x2f, 0xcb, 0xcf, 0x29, 0xcd, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0xea,
	0xd0, 0x43, 0xd2, 0xa1, 0x07, 0xd5, 0x21, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa7, 0x0f,
	0x62, 0x41, 0xb4, 0x48, 0xc9, 0x25, 0x83, 0xf5, 0xe8, 0x27, 0x25, 0x16, 0xa7, 0xc2, 0x0d, 0x4f,
	0xce, 0xcf, 0xcc, 0x83, 0xc8, 0x2b, 0x95, 0x70, 0xf1, 0x86, 0x40, 0xac, 0x0a, 0x03, 0xdb, 0x24,
	0x94, 0xcc, 0xc5, 0x96, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d,
	0x24, 0xa9, 0x07, 0x31, 0x41, 0x0f, 0x64, 0x02, 0xcc, 0x32, 0x3d, 0xe7, 0xfc, 0xcc, 0x3c, 0x27,
	0x83, 0x13, 0xf7, 0xe4, 0x19, 0x56, 0xdd, 0x97, 0xd7, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x5a, 0x07, 0xa1, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xc1, 0x1a, 0x8a, 0x83, 0xa0, 0x46, 0x3b, 0x05, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x05, 0x92, 0x59, 0x50, 0xdf, 0xea, 0xe6, 0x24, 0x26, 0x15, 0xc3,
	0x38, 0xfa, 0x65, 0x65, 0x46, 0xc6, 0xfa, 0x15, 0x28, 0x61, 0x06, 0xb6, 0x21, 0x89, 0x0d, 0xec,
	0x21, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd8, 0xfa, 0x15, 0x57, 0x01, 0x00, 0x00,
}

func (m *TrackedVolume) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrackedVolume) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrackedVolume) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTrackedVolume(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTrackedVolume(dAtA []byte, offset int, v uint64) int {
	offset -= sovTrackedVolume(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TrackedVolume) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTrackedVolume(uint64(l))
		}
	}
	return n
}

func sovTrackedVolume(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTrackedVolume(x uint64) (n int) {
	return sovTrackedVolume(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrackedVolume) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrackedVolume
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
			return fmt.Errorf("proto: TrackedVolume: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrackedVolume: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrackedVolume
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
				return ErrInvalidLengthTrackedVolume
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTrackedVolume
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrackedVolume(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTrackedVolume
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
func skipTrackedVolume(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrackedVolume
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
					return 0, ErrIntOverflowTrackedVolume
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
					return 0, ErrIntOverflowTrackedVolume
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
				return 0, ErrInvalidLengthTrackedVolume
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTrackedVolume
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTrackedVolume
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTrackedVolume        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrackedVolume          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTrackedVolume = fmt.Errorf("proto: unexpected end of group")
)
