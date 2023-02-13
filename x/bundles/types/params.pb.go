// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/bundles/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the bundles module parameters.
type Params struct {
	// upload_timeout ...
	UploadTimeout uint64 `protobuf:"varint,1,opt,name=upload_timeout,json=uploadTimeout,proto3" json:"upload_timeout,omitempty"`
	// storage_cost ...
	StorageCost github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=storage_cost,json=storageCost,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"storage_cost"`
	// network_fee ...
	NetworkFee string `protobuf:"bytes,3,opt,name=network_fee,json=networkFee,proto3" json:"network_fee,omitempty"`
	// max_points ...
	MaxPoints uint64 `protobuf:"varint,4,opt,name=max_points,json=maxPoints,proto3" json:"max_points,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfd3a74b72a01aaa, []int{0}
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

func (m *Params) GetUploadTimeout() uint64 {
	if m != nil {
		return m.UploadTimeout
	}
	return 0
}

func (m *Params) GetNetworkFee() string {
	if m != nil {
		return m.NetworkFee
	}
	return ""
}

func (m *Params) GetMaxPoints() uint64 {
	if m != nil {
		return m.MaxPoints
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "kyve.bundles.v1beta1.Params")
}

func init() { proto.RegisterFile("kyve/bundles/v1beta1/params.proto", fileDescriptor_cfd3a74b72a01aaa) }

var fileDescriptor_cfd3a74b72a01aaa = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x1b, 0x1d, 0x83, 0x65, 0xea, 0xa1, 0xec, 0x50, 0x04, 0xb3, 0x29, 0x28, 0x3b, 0x68,
	0xc2, 0xf0, 0x1f, 0x4c, 0xdd, 0x45, 0x90, 0x39, 0x44, 0xd0, 0x4b, 0x49, 0xbb, 0xcf, 0xae, 0x6c,
	0xe9, 0x57, 0x9a, 0x74, 0x6e, 0xff, 0xc2, 0xdf, 0xe4, 0x69, 0xc7, 0x1d, 0xc5, 0xc3, 0x90, 0xed,
	0x8f, 0x48, 0xd3, 0x22, 0x9e, 0x12, 0x1e, 0x9e, 0xef, 0xe5, 0xe5, 0xa5, 0xa7, 0xd3, 0xe5, 0x1c,
	0x44, 0x90, 0x27, 0xe3, 0x19, 0x68, 0x31, 0xef, 0x05, 0x60, 0x64, 0x4f, 0xa4, 0x32, 0x93, 0x4a,
	0xf3, 0x34, 0x43, 0x83, 0x6e, 0xab, 0x50, 0x78, 0xa5, 0xf0, 0x4a, 0x39, 0x6e, 0x45, 0x18, 0xa1,
	0x15, 0x44, 0xf1, 0x2b, 0xdd, 0xb3, 0x4f, 0x42, 0xeb, 0x43, 0x7b, 0xec, 0x9e, 0xd3, 0xa3, 0x3c,
	0x9d, 0xa1, 0x1c, 0xfb, 0x26, 0x56, 0x80, 0xb9, 0xf1, 0x48, 0x87, 0x74, 0x6b, 0xa3, 0xc3, 0x92,
	0x3e, 0x95, 0xd0, 0x7d, 0xa4, 0x07, 0xda, 0x60, 0x26, 0x23, 0xf0, 0x43, 0xd4, 0xc6, 0xdb, 0xeb,
	0x90, 0x6e, 0xa3, 0xcf, 0x57, 0x9b, 0xb6, 0xf3, 0xbd, 0x69, 0x5f, 0x44, 0xb1, 0x99, 0xe4, 0x01,
	0x0f, 0x51, 0x89, 0x10, 0xb5, 0x42, 0x5d, 0x3d, 0x57, 0x7a, 0x3c, 0x15, 0x66, 0x99, 0x82, 0xe6,
	0xb7, 0x10, 0x8e, 0x9a, 0x55, 0xc6, 0x0d, 0x6a, 0xe3, 0xb6, 0x69, 0x33, 0x01, 0xf3, 0x8e, 0xd9,
	0xd4, 0x7f, 0x03, 0xf0, 0xf6, 0x8b, 0xc4, 0x11, 0xad, 0xd0, 0x00, 0xc0, 0x3d, 0xa1, 0x54, 0xc9,
	0x85, 0x9f, 0x62, 0x9c, 0x18, 0xed, 0xd5, 0x6c, 0xad, 0x86, 0x92, 0x8b, 0xa1, 0x05, 0xfd, 0xc1,
	0x6a, 0xcb, 0xc8, 0x7a, 0xcb, 0xc8, 0xcf, 0x96, 0x91, 0x8f, 0x1d, 0x73, 0xd6, 0x3b, 0xe6, 0x7c,
	0xed, 0x98, 0xf3, 0x7a, 0xf9, 0xaf, 0xce, 0xfd, 0xcb, 0xf3, 0xdd, 0x43, 0x99, 0x29, 0xc2, 0x89,
	0x8c, 0x13, 0xb1, 0xf8, 0xdb, 0xd1, 0x16, 0x0b, 0xea, 0x76, 0x93, 0xeb, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x63, 0x93, 0xce, 0x64, 0x01, 0x00, 0x00,
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
	if m.MaxPoints != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPoints))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NetworkFee) > 0 {
		i -= len(m.NetworkFee)
		copy(dAtA[i:], m.NetworkFee)
		i = encodeVarintParams(dAtA, i, uint64(len(m.NetworkFee)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.StorageCost.Size()
		i -= size
		if _, err := m.StorageCost.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.UploadTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UploadTimeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if m.UploadTimeout != 0 {
		n += 1 + sovParams(uint64(m.UploadTimeout))
	}
	l = m.StorageCost.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.NetworkFee)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxPoints != 0 {
		n += 1 + sovParams(uint64(m.MaxPoints))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadTimeout", wireType)
			}
			m.UploadTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UploadTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageCost", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StorageCost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPoints", wireType)
			}
			m.MaxPoints = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPoints |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
