// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/metadata.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MetadataMatcher struct {
	Filter               string                         `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Path                 []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value                *ValueMatcher                  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea39646b3de596af, []int{0}
}

func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher.Unmarshal(m, b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher.Size(m)
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
	return xxx_messageInfo_MetadataMatcher_PathSegment.Unmarshal(m, b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Size(m)
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
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
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0xff, 0x93, 0xf4, 0xeb, 0x3f, 0xc5, 0x0f, 0x02, 0x62, 0xa9, 0xa0, 0xe9, 0xc7, 0xa2,
	0x6e, 0x66, 0xa4, 0x45, 0x17, 0x5d, 0x8e, 0x1b, 0x37, 0x85, 0x10, 0xc1, 0xfd, 0xd5, 0x8e, 0x6d,
	0xb0, 0x9d, 0x89, 0xc9, 0x6d, 0x30, 0x6f, 0x20, 0x2e, 0x5d, 0xfa, 0x3e, 0x3e, 0x8b, 0xaf, 0x20,
	0x5d, 0x49, 0x92, 0x29, 0x48, 0x89, 0xba, 0x1b, 0xee, 0x9c, 0xf3, 0xbb, 0xe7, 0x70, 0x69, 0x5f,
	0xaa, 0x44, 0xa7, 0x1c, 0xd3, 0x50, 0xf2, 0x25, 0xe0, 0xdd, 0x5c, 0x46, 0x3c, 0x19, 0xf1, 0xa5,
	0x44, 0x98, 0x02, 0x02, 0x0b, 0x23, 0x8d, 0xda, 0x39, 0xc8, 0x55, 0x2c, 0x53, 0x31, 0xa3, 0x62,
	0xc9, 0xa8, 0xdd, 0x29, 0x37, 0x27, 0xb0, 0x58, 0xc9, 0xc2, 0xd9, 0xee, 0xac, 0xa6, 0x21, 0x70,
	0x50, 0x4a, 0x23, 0x60, 0xa0, 0x55, 0xcc, 0x13, 0x19, 0xc5, 0x81, 0x56, 0x81, 0x9a, 0x19, 0xc9,
	0x61, 0x02, 0x8b, 0x60, 0x0a, 0x28, 0xf9, 0xe6, 0x51, 0x7c, 0x74, 0x3f, 0x2c, 0xba, 0x37, 0x31,
	0x41, 0x26, 0x05, 0xde, 0x39, 0xa1, 0xb5, 0xfb, 0x60, 0x81, 0x32, 0x6a, 0x11, 0x97, 0x0c, 0xfe,
	0x8b, 0xfa, 0x5a, 0x54, 0x22, 0xcb, 0x25, 0xbe, 0x19, 0x3b, 0x1e, 0xad, 0x84, 0x80, 0xf3, 0x96,
	0xe5, 0xda, 0x83, 0xe6, 0x70, 0xc8, 0x4a, 0x93, 0xb3, 0x2d, 0x2c, 0xf3, 0x00, 0xe7, 0xd7, 0x72,
	0xb6, 0x94, 0x0a, 0x45, 0x63, 0x2d, 0xaa, 0xaf, 0xc4, 0x6a, 0x10, 0x3f, 0x27, 0x39, 0x97, 0xb4,
	0x9a, 0x37, 0x6a, 0xd9, 0x2e, 0x19, 0x34, 0x87, 0xbd, 0x1f, 0x90, 0x37, 0x99, 0xc6, 0xf0, 0x72,
	0xc6, 0x0b, 0xb1, 0xf6, 0x89, 0x5f, 0x78, 0xdb, 0x8f, 0xb4, 0xf9, 0x6d, 0x87, 0x73, 0x44, 0xed,
	0x07, 0x99, 0x6e, 0x75, 0xb8, 0xfa, 0xe7, 0x67, 0xd3, 0xf1, 0xf9, 0xdb, 0xfb, 0xf3, 0xf1, 0x19,
	0x2d, 0xdb, 0xf3, 0x5b, 0xee, 0x5d, 0x5a, 0x8f, 0x0d, 0xde, 0xfe, 0x14, 0x64, 0x7c, 0x9a, 0x61,
	0xfa, 0xb4, 0xfb, 0x37, 0x46, 0x5c, 0xd0, 0x5e, 0xa0, 0x8b, 0x7d, 0x61, 0xa4, 0x9f, 0xd2, 0xf2,
	0x8a, 0x62, 0x67, 0xe3, 0xf3, 0xb2, 0xfb, 0x78, 0xe4, 0xb6, 0x96, 0x1f, 0x6a, 0xf4, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0x41, 0xdc, 0xc2, 0x46, 0x02, 0x00, 0x00,
}
