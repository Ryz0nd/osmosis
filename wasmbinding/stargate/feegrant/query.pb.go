// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargate/feegrant/query.proto

package feegrant

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant"
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

// QueryAllowanceResponse is the response type for the Query/Allowance RPC
// method.
type QueryAllowanceResponse struct {
	// allowance is a allowance granted for grantee by granter.
	Allowance *feegrant.Grant `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *QueryAllowanceResponse) Reset()         { *m = QueryAllowanceResponse{} }
func (m *QueryAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowanceResponse) ProtoMessage()    {}
func (*QueryAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c82dd920cbd5059c, []int{0}
}
func (m *QueryAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowanceResponse.Merge(m, src)
}
func (m *QueryAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowanceResponse proto.InternalMessageInfo

func (m *QueryAllowanceResponse) GetAllowance() *feegrant.Grant {
	if m != nil {
		return m.Allowance
	}
	return nil
}

// QueryAllowancesResponse is the response type for the Query/Allowances RPC
// method.
type QueryAllowancesResponse struct {
	// allowances are allowance's granted for grantee by granter.
	Allowances []*feegrant.Grant `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllowancesResponse) Reset()         { *m = QueryAllowancesResponse{} }
func (m *QueryAllowancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllowancesResponse) ProtoMessage()    {}
func (*QueryAllowancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c82dd920cbd5059c, []int{1}
}
func (m *QueryAllowancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllowancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllowancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllowancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllowancesResponse.Merge(m, src)
}
func (m *QueryAllowancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllowancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllowancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllowancesResponse proto.InternalMessageInfo

func (m *QueryAllowancesResponse) GetAllowances() []*feegrant.Grant {
	if m != nil {
		return m.Allowances
	}
	return nil
}

func (m *QueryAllowancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllowanceResponse)(nil), "stargate.feegrant.QueryAllowanceResponse")
	proto.RegisterType((*QueryAllowancesResponse)(nil), "stargate.feegrant.QueryAllowancesResponse")
}

func init() { proto.RegisterFile("stargate/feegrant/query.proto", fileDescriptor_c82dd920cbd5059c) }

var fileDescriptor_c82dd920cbd5059c = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xeb, 0xef, 0x93, 0x90, 0x30, 0x13, 0x1d, 0x20, 0xaa, 0x84, 0x55, 0x75, 0x00, 0x84,
	0x84, 0x4d, 0x60, 0x45, 0x48, 0x65, 0xe9, 0x0a, 0x19, 0x18, 0x60, 0x3a, 0x07, 0x63, 0x2c, 0x25,
	0x76, 0x88, 0xdd, 0x56, 0xbc, 0x05, 0xcf, 0xc0, 0xd3, 0x30, 0x76, 0x64, 0x44, 0xc9, 0x8b, 0xa0,
	0x24, 0xb5, 0x53, 0xc4, 0xc2, 0xe8, 0xbb, 0xdf, 0xff, 0xa7, 0xbb, 0x33, 0x3e, 0xb0, 0x0e, 0x4a,
	0x09, 0x4e, 0xb0, 0x27, 0x21, 0x64, 0x09, 0xda, 0xb1, 0x97, 0xb9, 0x28, 0x5f, 0x69, 0x51, 0x1a,
	0x67, 0x86, 0xbb, 0xbe, 0x4d, 0x7d, 0x7b, 0x74, 0x98, 0x1a, 0x9b, 0x1b, 0xdb, 0xf3, 0x8b, 0x98,
	0x0b, 0x07, 0x71, 0x28, 0x74, 0xd1, 0xd1, 0xc9, 0x9a, 0xe3, 0x60, 0x45, 0xe7, 0x0c, 0x64, 0x01,
	0x52, 0x69, 0x70, 0xca, 0xe8, 0x8e, 0x9d, 0xdc, 0xe1, 0xbd, 0xdb, 0x86, 0x98, 0x66, 0x99, 0x59,
	0x82, 0x4e, 0x45, 0x22, 0x6c, 0x61, 0xb4, 0x15, 0xc3, 0x4b, 0xbc, 0x0d, 0xbe, 0x18, 0xa1, 0x31,
	0x3a, 0xde, 0x39, 0x27, 0xb4, 0x33, 0x87, 0x91, 0xe8, 0xda, 0x4b, 0x67, 0xcd, 0x2b, 0xe9, 0x03,
	0x93, 0x77, 0x84, 0xf7, 0x7f, 0x8a, 0x6d, 0x30, 0x5f, 0x61, 0x1c, 0x40, 0x1b, 0xa1, 0xf1, 0xff,
	0x3f, 0xa8, 0x37, 0x12, 0xc3, 0x19, 0xc6, 0xfd, 0x1e, 0xd1, 0xbf, 0x76, 0xb4, 0x23, 0x9f, 0x6f,
	0x96, 0xa6, 0xdd, 0x21, 0xbd, 0xe1, 0x06, 0x64, 0x58, 0x2b, 0xd9, 0x88, 0x5e, 0x3f, 0x7c, 0x54,
	0x04, 0xad, 0x2a, 0x82, 0xbe, 0x2a, 0x82, 0xde, 0x6a, 0x32, 0x58, 0xd5, 0x64, 0xf0, 0x59, 0x93,
	0xc1, 0xfd, 0x54, 0x2a, 0xf7, 0x3c, 0xe7, 0x34, 0x35, 0x39, 0x6b, 0xbd, 0xca, 0x9e, 0x66, 0xc0,
	0xad, 0x7f, 0xb0, 0x45, 0x7c, 0xc6, 0x96, 0x60, 0x73, 0xae, 0xf4, 0xa3, 0xd2, 0x92, 0xfd, 0xfa,
	0x4c, 0xbe, 0xd5, 0x1e, 0xf8, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x31, 0x32, 0xad, 0xe8,
	0x01, 0x00, 0x00,
}

func (m *QueryAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllowancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllowancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllowancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Allowances) > 0 {
		for iNdEx := len(m.Allowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllowancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for _, e := range m.Allowances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &feegrant.Grant{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllowancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllowancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllowancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allowances = append(m.Allowances, &feegrant.Grant{})
			if err := m.Allowances[len(m.Allowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
