// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargate/slashing/query.proto

package slashing

import (
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	types "github.com/cosmos/cosmos-sdk/x/slashing/types"
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

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	Params types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d51489b12a14e470, []int{0}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() types.Params {
	if m != nil {
		return m.Params
	}
	return types.Params{}
}

// QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
type QuerySigningInfoResponse struct {
	// val_signing_info is the signing info of requested val cons address
	ValSigningInfo types.ValidatorSigningInfo `protobuf:"bytes,1,opt,name=val_signing_info,json=valSigningInfo,proto3" json:"val_signing_info"`
}

func (m *QuerySigningInfoResponse) Reset()         { *m = QuerySigningInfoResponse{} }
func (m *QuerySigningInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoResponse) ProtoMessage()    {}
func (*QuerySigningInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d51489b12a14e470, []int{1}
}
func (m *QuerySigningInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoResponse.Merge(m, src)
}
func (m *QuerySigningInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoResponse proto.InternalMessageInfo

func (m *QuerySigningInfoResponse) GetValSigningInfo() types.ValidatorSigningInfo {
	if m != nil {
		return m.ValSigningInfo
	}
	return types.ValidatorSigningInfo{}
}

// QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
type QuerySigningInfosResponse struct {
	// info is the signing info of all validators
	Info       []types.ValidatorSigningInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info"`
	Pagination *query.PageResponse          `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySigningInfosResponse) Reset()         { *m = QuerySigningInfosResponse{} }
func (m *QuerySigningInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosResponse) ProtoMessage()    {}
func (*QuerySigningInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d51489b12a14e470, []int{2}
}
func (m *QuerySigningInfosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosResponse.Merge(m, src)
}
func (m *QuerySigningInfosResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosResponse proto.InternalMessageInfo

func (m *QuerySigningInfosResponse) GetInfo() []types.ValidatorSigningInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *QuerySigningInfosResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsResponse)(nil), "stargate.slashing.QueryParamsResponse")
	proto.RegisterType((*QuerySigningInfoResponse)(nil), "stargate.slashing.QuerySigningInfoResponse")
	proto.RegisterType((*QuerySigningInfosResponse)(nil), "stargate.slashing.QuerySigningInfosResponse")
}

func init() { proto.RegisterFile("stargate/slashing/query.proto", fileDescriptor_d51489b12a14e470) }

var fileDescriptor_d51489b12a14e470 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xdb, 0xef, 0x23, 0x2c, 0xc6, 0xc4, 0x28, 0xba, 0x40, 0x12, 0x8b, 0x61, 0xa1, 0xc6,
	0x84, 0x8e, 0xe8, 0xda, 0x85, 0x6c, 0x88, 0x3b, 0x45, 0xe3, 0x42, 0x63, 0xc8, 0x2d, 0x0c, 0xc3,
	0x24, 0xed, 0x4c, 0xed, 0x0c, 0x35, 0xbc, 0x85, 0x2f, 0xe2, 0x7b, 0xb0, 0x64, 0xe9, 0xca, 0x18,
	0x78, 0x11, 0xd3, 0xf9, 0x03, 0x44, 0xe3, 0xc6, 0xdd, 0xf4, 0xdc, 0x39, 0xbf, 0x73, 0xdb, 0x53,
	0xb4, 0x2f, 0x15, 0x64, 0x14, 0x14, 0xc1, 0x32, 0x06, 0x39, 0x62, 0x9c, 0xe2, 0xe7, 0x31, 0xc9,
	0x26, 0x61, 0x9a, 0x09, 0x25, 0x2a, 0xdb, 0x6e, 0x1c, 0xba, 0x71, 0xed, 0xa4, 0x2f, 0x64, 0x22,
	0x24, 0x8e, 0x40, 0x12, 0x73, 0x17, 0xe7, 0xad, 0x88, 0x28, 0x68, 0xe1, 0x14, 0x28, 0xe3, 0xa0,
	0x98, 0xe0, 0xc6, 0x5e, 0xdb, 0xa5, 0x82, 0x0a, 0x7d, 0xc4, 0xc5, 0xc9, 0xaa, 0x87, 0x96, 0xb0,
	0x4c, 0x74, 0x7e, 0x27, 0x98, 0x7b, 0x8d, 0x3b, 0xb4, 0x73, 0x53, 0xf0, 0xaf, 0x21, 0x83, 0x44,
	0x76, 0x89, 0x4c, 0x05, 0x97, 0xa4, 0x72, 0x81, 0xca, 0xa9, 0x56, 0xaa, 0xfe, 0x81, 0x7f, 0xbc,
	0x71, 0x56, 0x0f, 0x0d, 0x6f, 0xb9, 0x62, 0x68, 0x79, 0xa1, 0x31, 0xb6, 0x4b, 0xd3, 0x8f, 0xba,
	0xd7, 0xb5, 0xa6, 0xc6, 0x04, 0x55, 0x35, 0xf5, 0x96, 0x51, 0xce, 0x38, 0xbd, 0xe2, 0x43, 0xb1,
	0x44, 0x3f, 0xa1, 0xad, 0x1c, 0xe2, 0x9e, 0x34, 0xa3, 0x1e, 0xe3, 0x43, 0x61, 0x43, 0x9a, 0xbf,
	0x86, 0xdc, 0x43, 0xcc, 0x06, 0xa0, 0x44, 0xb6, 0x06, 0xb4, 0x91, 0x9b, 0x39, 0xc4, 0x6b, 0x6a,
	0xe3, 0xcd, 0x47, 0x7b, 0xdf, 0xb3, 0x57, 0xef, 0xd5, 0x41, 0x25, 0x1b, 0xf8, 0xff, 0xaf, 0x81,
	0x1a, 0x50, 0xe9, 0x20, 0xb4, 0x6a, 0xa2, 0xfa, 0x4f, 0xef, 0x7f, 0xe4, 0x70, 0x45, 0x6d, 0xa1,
	0xa9, 0x78, 0xf5, 0x99, 0x28, 0x71, 0x5b, 0x74, 0xd7, 0xac, 0xed, 0xc7, 0xe9, 0x3c, 0xf0, 0x67,
	0xf3, 0xc0, 0xff, 0x9c, 0x07, 0xfe, 0xeb, 0x22, 0xf0, 0x66, 0x8b, 0xc0, 0x7b, 0x5f, 0x04, 0xde,
	0xc3, 0x25, 0x65, 0x6a, 0x34, 0x8e, 0xc2, 0xbe, 0x48, 0xb0, 0xe6, 0x32, 0xd9, 0x8c, 0x21, 0x92,
	0xee, 0x01, 0xe7, 0xad, 0x53, 0xfc, 0x02, 0x32, 0x89, 0x18, 0x1f, 0x14, 0x15, 0xff, 0xf8, 0xcd,
	0xa2, 0xb2, 0x2e, 0xf9, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x5b, 0xdf, 0x4e, 0x82, 0x02,
	0x00, 0x00,
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValSigningInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Info) > 0 {
		for iNdEx := len(m.Info) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Info[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ValSigningInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Info) > 0 {
		for _, e := range m.Info {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySigningInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValSigningInfo", wireType)
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
			if err := m.ValSigningInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySigningInfosResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
			m.Info = append(m.Info, types.ValidatorSigningInfo{})
			if err := m.Info[len(m.Info)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
