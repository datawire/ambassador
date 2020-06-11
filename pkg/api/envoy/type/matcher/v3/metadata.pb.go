// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/v3/metadata.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// [#next-major-version: MetadataMatcher should use StructMatcher]
type MetadataMatcher struct {
	// The filter name to retrieve the Struct from the Metadata.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The path to retrieve the Value from the Struct.
	Path []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// The MetadataMatcher is matched if the value retrieved by path is matched to this value.
	Value                *ValueMatcher `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea39646b3de596af, []int{0}
}
func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Metadata.
// Note: Currently it's not supported to retrieve a value from a list in Metadata. This means that
// if the segment key refers to a list, it has to be the last segment in a path.
type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea39646b3de596af, []int{0, 0}
}
func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.v3.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.v3.MetadataMatcher.PathSegment")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3/metadata.proto", fileDescriptor_ea39646b3de596af)
}

var fileDescriptor_ea39646b3de596af = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4b, 0xc3, 0x50,
	0x14, 0xc5, 0xbd, 0x49, 0xff, 0xf9, 0x8a, 0x5a, 0x02, 0x62, 0xa9, 0x18, 0xd3, 0x3f, 0x43, 0x5d,
	0x12, 0x69, 0x71, 0xe9, 0x18, 0x17, 0x11, 0x0a, 0x21, 0x82, 0xfb, 0xd3, 0x3e, 0xdb, 0x60, 0x9b,
	0x17, 0x93, 0xdb, 0x60, 0x36, 0x47, 0x71, 0x74, 0xf4, 0xa3, 0xb8, 0x0b, 0x1d, 0xf5, 0x1b, 0x48,
	0x17, 0xbf, 0x82, 0x74, 0x92, 0x24, 0xaf, 0x20, 0x35, 0xea, 0xf6, 0x78, 0xf7, 0xdc, 0xdf, 0x39,
	0x87, 0x4b, 0x5a, 0xcc, 0x0d, 0x79, 0x64, 0x60, 0xe4, 0x31, 0x63, 0x42, 0xf1, 0x72, 0xc4, 0x7c,
	0x23, 0xec, 0x1a, 0x13, 0x86, 0x74, 0x40, 0x91, 0xea, 0x9e, 0xcf, 0x91, 0x2b, 0xdb, 0x89, 0x4a,
	0x8f, 0x55, 0xba, 0x50, 0xe9, 0x61, 0xb7, 0x56, 0xcf, 0x5e, 0x0e, 0xe9, 0x78, 0xca, 0xd2, 0xcd,
	0xda, 0xde, 0x74, 0xe0, 0x51, 0x83, 0xba, 0x2e, 0x47, 0x8a, 0x0e, 0x77, 0x03, 0x23, 0x40, 0x8a,
	0xd3, 0x40, 0x8c, 0xeb, 0x3f, 0xc6, 0x21, 0xf3, 0x03, 0x87, 0xbb, 0x8e, 0x3b, 0x14, 0x92, 0x9d,
	0x90, 0x8e, 0x9d, 0x01, 0x45, 0x66, 0x2c, 0x1f, 0xe9, 0xa0, 0xf1, 0x21, 0x91, 0xad, 0xbe, 0xc8,
	0xd9, 0x4f, 0xdd, 0x95, 0x7d, 0x52, 0xb8, 0x72, 0xc6, 0xc8, 0xfc, 0x2a, 0x68, 0xd0, 0x5e, 0x37,
	0x8b, 0x0b, 0x33, 0xe7, 0x4b, 0x1a, 0xd8, 0xe2, 0x5b, 0xb1, 0x48, 0xce, 0xa3, 0x38, 0xaa, 0x4a,
	0x9a, 0xdc, 0x2e, 0x77, 0x3a, 0x7a, 0x66, 0x31, 0x7d, 0x05, 0xab, 0x5b, 0x14, 0x47, 0x67, 0x6c,
	0x38, 0x61, 0x2e, 0x9a, 0xa5, 0x85, 0x99, 0x7f, 0x04, 0xa9, 0x04, 0x76, 0x42, 0x52, 0x8e, 0x49,
	0x3e, 0x29, 0x5c, 0x95, 0x35, 0x68, 0x97, 0x3b, 0xcd, 0x5f, 0x90, 0xe7, 0xb1, 0x46, 0xf0, 0x12,
	0xc6, 0x03, 0x48, 0x15, 0xb0, 0xd3, 0xdd, 0xda, 0x0d, 0x29, 0x7f, 0xf3, 0x50, 0x76, 0x89, 0x7c,
	0xcd, 0xa2, 0x95, 0x0e, 0x27, 0x6b, 0x76, 0xfc, 0xdb, 0x3b, 0x7a, 0x7a, 0xb9, 0x57, 0x0f, 0x49,
	0x96, 0xcf, 0x5f, 0xb9, 0x37, 0x49, 0x31, 0x10, 0x78, 0xf9, 0xd3, 0x84, 0xde, 0x41, 0x8c, 0x69,
	0x91, 0xc6, 0xff, 0x18, 0xf3, 0x74, 0x36, 0x57, 0xe1, 0x75, 0xae, 0xc2, 0xfb, 0x5c, 0x85, 0xe7,
	0xbb, 0xd9, 0x5b, 0x41, 0xaa, 0x00, 0x69, 0x3a, 0x3c, 0xcd, 0xe0, 0xf9, 0xfc, 0x36, 0xca, 0xae,
	0x6d, 0x6e, 0x2c, 0x59, 0x56, 0x7c, 0x33, 0x0b, 0x2e, 0x0a, 0xc9, 0xf1, 0xba, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xbc, 0x6d, 0xd5, 0x79, 0x02, 0x00, 0x00,
}

func (m *MetadataMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		for iNdEx := len(m.Path) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Path[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Filter) > 0 {
		i -= len(m.Filter)
		copy(dAtA[i:], m.Filter)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Filter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataMatcher_PathSegment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher_PathSegment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMatcher_PathSegment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Segment != nil {
		{
			size := m.Segment.Size()
			i -= size
			if _, err := m.Segment.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MetadataMatcher_PathSegment_Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMatcher_PathSegment_Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Key)
	copy(dAtA[i:], m.Key)
	i = encodeVarintMetadata(dAtA, i, uint64(len(m.Key)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetadataMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filter)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Path) > 0 {
		for _, e := range m.Path {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Segment != nil {
		n += m.Segment.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment_Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetadataMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: MetadataMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, &MetadataMatcher_PathSegment{})
			if err := m.Path[len(m.Path)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &ValueMatcher{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MetadataMatcher_PathSegment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: PathSegment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathSegment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Segment = &MetadataMatcher_PathSegment_Key{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)