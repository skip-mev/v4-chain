// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/revshare/revshare.proto

package types

import (
	fmt "fmt"
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

// MarketMapperRevShareDetails specifies any details associated with the market
// mapper revenue share
type MarketMapperRevShareDetails struct {
	// Unix timestamp recorded when the market revenue share expires
	ExpirationTs uint64 `protobuf:"varint,1,opt,name=expiration_ts,json=expirationTs,proto3" json:"expiration_ts,omitempty"`
}

func (m *MarketMapperRevShareDetails) Reset()         { *m = MarketMapperRevShareDetails{} }
func (m *MarketMapperRevShareDetails) String() string { return proto.CompactTextString(m) }
func (*MarketMapperRevShareDetails) ProtoMessage()    {}
func (*MarketMapperRevShareDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b9759663d195798, []int{0}
}
func (m *MarketMapperRevShareDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMapperRevShareDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMapperRevShareDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMapperRevShareDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMapperRevShareDetails.Merge(m, src)
}
func (m *MarketMapperRevShareDetails) XXX_Size() int {
	return m.Size()
}
func (m *MarketMapperRevShareDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMapperRevShareDetails.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMapperRevShareDetails proto.InternalMessageInfo

func (m *MarketMapperRevShareDetails) GetExpirationTs() uint64 {
	if m != nil {
		return m.ExpirationTs
	}
	return 0
}

func init() {
	proto.RegisterType((*MarketMapperRevShareDetails)(nil), "dydxprotocol.revshare.MarketMapperRevShareDetails")
}

func init() {
	proto.RegisterFile("dydxprotocol/revshare/revshare.proto", fileDescriptor_5b9759663d195798)
}

var fileDescriptor_5b9759663d195798 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4a, 0x2d, 0x2b, 0xce, 0x48, 0x2c,
	0x4a, 0x85, 0x33, 0xf4, 0xc0, 0x52, 0x42, 0xa2, 0xc8, 0xaa, 0xf4, 0x60, 0x92, 0x4a, 0x4e, 0x5c,
	0xd2, 0xbe, 0x89, 0x45, 0xd9, 0xa9, 0x25, 0xbe, 0x89, 0x05, 0x05, 0xa9, 0x45, 0x41, 0xa9, 0x65,
	0xc1, 0x20, 0x71, 0x97, 0xd4, 0x92, 0xc4, 0xcc, 0x9c, 0x62, 0x21, 0x65, 0x2e, 0xde, 0xd4, 0x8a,
	0x82, 0xcc, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0xf8, 0x92, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x96, 0x20, 0x1e, 0x84, 0x60, 0x48, 0xb1, 0x53, 0xc8, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e,
	0xcb, 0x31, 0x44, 0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa3,
	0xb8, 0xb2, 0xcc, 0x44, 0x37, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x1f, 0x2e, 0x52, 0x81, 0x70, 0x79,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x58, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xec,
	0xc5, 0x93, 0xd9, 0xdf, 0x00, 0x00, 0x00,
}

func (m *MarketMapperRevShareDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMapperRevShareDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMapperRevShareDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpirationTs != 0 {
		i = encodeVarintRevshare(dAtA, i, uint64(m.ExpirationTs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRevshare(dAtA []byte, offset int, v uint64) int {
	offset -= sovRevshare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketMapperRevShareDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpirationTs != 0 {
		n += 1 + sovRevshare(uint64(m.ExpirationTs))
	}
	return n
}

func sovRevshare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRevshare(x uint64) (n int) {
	return sovRevshare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketMapperRevShareDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRevshare
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
			return fmt.Errorf("proto: MarketMapperRevShareDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMapperRevShareDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTs", wireType)
			}
			m.ExpirationTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevshare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRevshare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRevshare
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
func skipRevshare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRevshare
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
					return 0, ErrIntOverflowRevshare
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
					return 0, ErrIntOverflowRevshare
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
				return 0, ErrInvalidLengthRevshare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRevshare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRevshare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRevshare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRevshare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRevshare = fmt.Errorf("proto: unexpected end of group")
)
