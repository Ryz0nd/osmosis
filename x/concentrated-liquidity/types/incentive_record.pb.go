// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentratedliquidity/v1beta1/incentive_record.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// IncentiveRecord is the high-level struct we use to deal with an independent
// incentive being distributed on a pool. Note that PoolId, Denom, and MinUptime
// are included in the key so we avoid storing them in state, hence the
// distinction between IncentiveRecord and IncentiveRecordBody.
type IncentiveRecord struct {
	// incentive_id is the id uniquely identifying this incentive record.
	IncentiveId uint64 `protobuf:"varint,1,opt,name=incentive_id,json=incentiveId,proto3" json:"incentive_id,omitempty" yaml:"incentive_id"`
	PoolId      uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// incentive record body holds necessary
	IncentiveRecordBody IncentiveRecordBody `protobuf:"bytes,4,opt,name=incentive_record_body,json=incentiveRecordBody,proto3" json:"incentive_record_body" yaml:"incentive_record_body"`
	// min_uptime is the minimum uptime required for liquidity to qualify for this
	// incentive. It should be always be one of the supported uptimes in
	// types.SupportedUptimes
	MinUptime time.Duration `protobuf:"bytes,5,opt,name=min_uptime,json=minUptime,proto3,stdduration" json:"min_uptime" yaml:"min_uptime"`
}

func (m *IncentiveRecord) Reset()         { *m = IncentiveRecord{} }
func (m *IncentiveRecord) String() string { return proto.CompactTextString(m) }
func (*IncentiveRecord) ProtoMessage()    {}
func (*IncentiveRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef31b586e827443, []int{0}
}
func (m *IncentiveRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveRecord.Merge(m, src)
}
func (m *IncentiveRecord) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveRecord.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveRecord proto.InternalMessageInfo

func (m *IncentiveRecord) GetIncentiveId() uint64 {
	if m != nil {
		return m.IncentiveId
	}
	return 0
}

func (m *IncentiveRecord) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *IncentiveRecord) GetIncentiveRecordBody() IncentiveRecordBody {
	if m != nil {
		return m.IncentiveRecordBody
	}
	return IncentiveRecordBody{}
}

func (m *IncentiveRecord) GetMinUptime() time.Duration {
	if m != nil {
		return m.MinUptime
	}
	return 0
}

// IncentiveRecordBody represents the body stored in state for each individual
// record.
type IncentiveRecordBody struct {
	// remaining_coin is the total amount of incentives to be distributed
	RemainingCoin types.DecCoin `protobuf:"bytes,1,opt,name=remaining_coin,json=remainingCoin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoin" json:"remaining_coin" yaml:"remaining_coins"`
	// emission_rate is the incentive emission rate per second
	EmissionRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=emission_rate,json=emissionRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"emission_rate" yaml:"emission_rate"`
	// start_time is the time when the incentive starts distributing
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
}

func (m *IncentiveRecordBody) Reset()         { *m = IncentiveRecordBody{} }
func (m *IncentiveRecordBody) String() string { return proto.CompactTextString(m) }
func (*IncentiveRecordBody) ProtoMessage()    {}
func (*IncentiveRecordBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef31b586e827443, []int{1}
}
func (m *IncentiveRecordBody) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveRecordBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveRecordBody.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveRecordBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveRecordBody.Merge(m, src)
}
func (m *IncentiveRecordBody) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveRecordBody) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveRecordBody.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveRecordBody proto.InternalMessageInfo

func (m *IncentiveRecordBody) GetRemainingCoin() types.DecCoin {
	if m != nil {
		return m.RemainingCoin
	}
	return types.DecCoin{}
}

func (m *IncentiveRecordBody) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*IncentiveRecord)(nil), "osmosis.concentratedliquidity.v1beta1.IncentiveRecord")
	proto.RegisterType((*IncentiveRecordBody)(nil), "osmosis.concentratedliquidity.v1beta1.IncentiveRecordBody")
}

func init() {
	proto.RegisterFile("osmosis/concentratedliquidity/v1beta1/incentive_record.proto", fileDescriptor_bef31b586e827443)
}

var fileDescriptor_bef31b586e827443 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x4b, 0x29, 0xea, 0xa6, 0x05, 0xe1, 0x14, 0x9a, 0x86, 0x62, 0x57, 0x16, 0x48, 0x15,
	0x52, 0xd6, 0x6a, 0x7a, 0x0b, 0x9c, 0xdc, 0x5c, 0x22, 0xf5, 0x64, 0x81, 0x40, 0x48, 0xc8, 0xac,
	0xbd, 0x8b, 0xbb, 0x6a, 0xec, 0x0d, 0xde, 0x4d, 0x84, 0xdf, 0xa2, 0x48, 0x1c, 0x78, 0x06, 0x9e,
	0xa4, 0xc7, 0x9e, 0x10, 0xe2, 0x90, 0xa2, 0x44, 0xbc, 0x40, 0x9e, 0x00, 0xed, 0x7a, 0x9d, 0xbf,
	0xf6, 0xc0, 0xc9, 0xfe, 0x76, 0xe6, 0x9b, 0x6f, 0xe6, 0xdb, 0x59, 0xf0, 0x8a, 0xf1, 0x84, 0x71,
	0xca, 0xdd, 0x88, 0xa5, 0x11, 0x49, 0x45, 0x86, 0x04, 0xc1, 0x3d, 0xfa, 0x79, 0x40, 0x31, 0x15,
	0xb9, 0x3b, 0x3c, 0x0a, 0x89, 0x40, 0x47, 0x2e, 0x55, 0x41, 0x3a, 0x24, 0x41, 0x46, 0x22, 0x96,
	0x61, 0xd8, 0xcf, 0x98, 0x60, 0xe6, 0x73, 0xcd, 0x86, 0xb7, 0xb2, 0xa1, 0x66, 0x37, 0xf6, 0x22,
	0x95, 0x17, 0x28, 0x92, 0x5b, 0x80, 0xa2, 0x42, 0x63, 0x27, 0x66, 0x31, 0x2b, 0xce, 0xe5, 0x9f,
	0x3e, 0xb5, 0x63, 0xc6, 0xe2, 0x1e, 0x71, 0x15, 0x0a, 0x07, 0x9f, 0x5c, 0x41, 0x13, 0xc2, 0x05,
	0x4a, 0xfa, 0x3a, 0xc1, 0x5a, 0x4d, 0xc0, 0x83, 0x0c, 0x09, 0xca, 0xd2, 0x32, 0x5e, 0x88, 0xb8,
	0x21, 0xe2, 0x64, 0x36, 0x44, 0xc4, 0xa8, 0x8e, 0x3b, 0x3f, 0xd7, 0xc0, 0x83, 0x6e, 0x39, 0x93,
	0xaf, 0x46, 0x32, 0xdb, 0x60, 0x6b, 0x3e, 0x26, 0xc5, 0x75, 0xe3, 0xc0, 0x38, 0x5c, 0xf7, 0x76,
	0xa7, 0x23, 0xbb, 0x96, 0xa3, 0xa4, 0xd7, 0x76, 0x16, 0xa3, 0x8e, 0x5f, 0x9d, 0xc1, 0x2e, 0x36,
	0x77, 0xc1, 0xbd, 0x3e, 0x63, 0x3d, 0x49, 0x5b, 0x93, 0x34, 0x7f, 0x43, 0xc2, 0x2e, 0x36, 0xbf,
	0x19, 0xe0, 0xd1, 0xaa, 0x79, 0x41, 0xc8, 0x70, 0x5e, 0x5f, 0x3f, 0x30, 0x0e, 0xab, 0xad, 0x36,
	0xfc, 0x2f, 0x0b, 0xe1, 0x4a, 0xb3, 0x1e, 0xc3, 0xb9, 0xf7, 0xec, 0x72, 0x64, 0x57, 0xa6, 0x23,
	0x7b, 0x7f, 0xb5, 0xbd, 0x05, 0x19, 0xc7, 0xaf, 0xd1, 0x9b, 0x54, 0xf3, 0x2d, 0x00, 0x09, 0x4d,
	0x83, 0x41, 0x5f, 0x1a, 0x5b, 0xbf, 0xab, 0x5a, 0xd9, 0x83, 0x85, 0xa9, 0xb0, 0x34, 0x15, 0x76,
	0xb4, 0xa9, 0xde, 0x53, 0xad, 0xf4, 0xb0, 0x50, 0x9a, 0x53, 0x9d, 0xef, 0xd7, 0xb6, 0xe1, 0x6f,
	0x26, 0x34, 0x7d, 0x53, 0xe0, 0xbf, 0x6b, 0xa0, 0x76, 0x4b, 0xaf, 0xe6, 0x57, 0x03, 0xdc, 0xcf,
	0x48, 0x82, 0x68, 0x4a, 0xd3, 0x38, 0x90, 0x37, 0xa1, 0xfc, 0xad, 0xb6, 0xf6, 0xa1, 0xde, 0x07,
	0x79, 0x55, 0xb3, 0x71, 0x3b, 0x24, 0x3a, 0x61, 0x34, 0xf5, 0x4e, 0xb5, 0xf0, 0xe3, 0x42, 0x78,
	0xb9, 0x02, 0x77, 0x7e, 0x5c, 0xdb, 0x2f, 0x62, 0x2a, 0xce, 0x06, 0x21, 0x8c, 0x58, 0xa2, 0x37,
	0x4b, 0x7f, 0x9a, 0x1c, 0x9f, 0xbb, 0x22, 0xef, 0x13, 0x5e, 0x56, 0xf3, 0xb7, 0x67, 0x7c, 0x09,
	0xcd, 0x8f, 0x60, 0x9b, 0x24, 0x94, 0x73, 0xca, 0xd2, 0x40, 0xda, 0xae, 0xae, 0x6e, 0xd3, 0x7b,
	0x29, 0x35, 0x7f, 0x8f, 0xec, 0x27, 0x45, 0x1d, 0x8e, 0xcf, 0x21, 0x65, 0x6e, 0x82, 0xc4, 0x19,
	0x3c, 0x25, 0x31, 0x8a, 0xf2, 0x0e, 0x89, 0xa6, 0x23, 0x7b, 0xa7, 0x68, 0x69, 0xa9, 0x82, 0xe3,
	0x6f, 0x95, 0xd8, 0x47, 0x82, 0x98, 0xef, 0x00, 0xe0, 0x02, 0x65, 0x22, 0x50, 0x36, 0xdf, 0x51,
	0x03, 0x37, 0x6e, 0xd8, 0xfc, 0xba, 0x5c, 0xee, 0x55, 0x9f, 0xe7, 0x5c, 0xe7, 0x42, 0xf9, 0xac,
	0x0e, 0x64, 0xba, 0xf7, 0xe1, 0x72, 0x6c, 0x19, 0x57, 0x63, 0xcb, 0xf8, 0x33, 0xb6, 0x8c, 0x8b,
	0x89, 0x55, 0xb9, 0x9a, 0x58, 0x95, 0x5f, 0x13, 0xab, 0xf2, 0xfe, 0x64, 0xc1, 0x10, 0xbd, 0x5b,
	0xcd, 0x1e, 0x0a, 0x79, 0x09, 0xdc, 0xe1, 0xb0, 0x75, 0xec, 0x7e, 0x59, 0x7a, 0xf0, 0xcd, 0xf9,
	0x8b, 0x57, 0x8e, 0x85, 0x1b, 0xaa, 0xb9, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc3,
	0xbd, 0xdd, 0x1f, 0x04, 0x00, 0x00,
}

func (m *IncentiveRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MinUptime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MinUptime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintIncentiveRecord(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.IncentiveRecordBody.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentiveRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PoolId != 0 {
		i = encodeVarintIncentiveRecord(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.IncentiveId != 0 {
		i = encodeVarintIncentiveRecord(dAtA, i, uint64(m.IncentiveId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveRecordBody) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveRecordBody) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveRecordBody) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintIncentiveRecord(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	{
		size := m.EmissionRate.Size()
		i -= size
		if _, err := m.EmissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentiveRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.RemainingCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentiveRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintIncentiveRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentiveRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncentiveRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IncentiveId != 0 {
		n += 1 + sovIncentiveRecord(uint64(m.IncentiveId))
	}
	if m.PoolId != 0 {
		n += 1 + sovIncentiveRecord(uint64(m.PoolId))
	}
	l = m.IncentiveRecordBody.Size()
	n += 1 + l + sovIncentiveRecord(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MinUptime)
	n += 1 + l + sovIncentiveRecord(uint64(l))
	return n
}

func (m *IncentiveRecordBody) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RemainingCoin.Size()
	n += 1 + l + sovIncentiveRecord(uint64(l))
	l = m.EmissionRate.Size()
	n += 1 + l + sovIncentiveRecord(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovIncentiveRecord(uint64(l))
	return n
}

func sovIncentiveRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentiveRecord(x uint64) (n int) {
	return sovIncentiveRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncentiveRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentiveRecord
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
			return fmt.Errorf("proto: IncentiveRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveId", wireType)
			}
			m.IncentiveId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IncentiveId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveRecordBody", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
				return ErrInvalidLengthIncentiveRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentiveRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncentiveRecordBody.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUptime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
				return ErrInvalidLengthIncentiveRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentiveRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MinUptime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentiveRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentiveRecord
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
func (m *IncentiveRecordBody) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentiveRecord
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
			return fmt.Errorf("proto: IncentiveRecordBody: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveRecordBody: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
				return ErrInvalidLengthIncentiveRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentiveRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
				return ErrInvalidLengthIncentiveRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentiveRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EmissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentiveRecord
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
				return ErrInvalidLengthIncentiveRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentiveRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentiveRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentiveRecord
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
func skipIncentiveRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentiveRecord
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
					return 0, ErrIntOverflowIncentiveRecord
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
					return 0, ErrIntOverflowIncentiveRecord
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
				return 0, ErrInvalidLengthIncentiveRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentiveRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentiveRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentiveRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentiveRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentiveRecord = fmt.Errorf("proto: unexpected end of group")
)
