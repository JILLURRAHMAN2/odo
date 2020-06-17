// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/advertising_channel_sub_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing the different channel subtypes.
type AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType int32

const (
	// Not specified.
	AdvertisingChannelSubTypeEnum_UNSPECIFIED AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 0
	// Used as a return value only. Represents value unknown in this version.
	AdvertisingChannelSubTypeEnum_UNKNOWN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 1
	// Mobile app campaigns for Search.
	AdvertisingChannelSubTypeEnum_SEARCH_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 2
	// Mobile app campaigns for Display.
	AdvertisingChannelSubTypeEnum_DISPLAY_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 3
	// AdWords express campaigns for search.
	AdvertisingChannelSubTypeEnum_SEARCH_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 4
	// AdWords Express campaigns for display.
	AdvertisingChannelSubTypeEnum_DISPLAY_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 5
	// Smart Shopping campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_SMART_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 6
	// Gmail Ad campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_GMAIL_AD AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 7
	// Smart display campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 8
	// Video Outstream campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_OUTSTREAM AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 9
	// Video TrueView for Action campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_ACTION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 10
	// Video campaigns with non-skippable video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_NON_SKIPPABLE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 11
	// Universal App Campaign.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 12
)

var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_MOBILE_APP",
	3:  "DISPLAY_MOBILE_APP",
	4:  "SEARCH_EXPRESS",
	5:  "DISPLAY_EXPRESS",
	6:  "SHOPPING_SMART_ADS",
	7:  "DISPLAY_GMAIL_AD",
	8:  "DISPLAY_SMART_CAMPAIGN",
	9:  "VIDEO_OUTSTREAM",
	10: "VIDEO_ACTION",
	11: "VIDEO_NON_SKIPPABLE",
	12: "APP_CAMPAIGN",
}
var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"SEARCH_MOBILE_APP":      2,
	"DISPLAY_MOBILE_APP":     3,
	"SEARCH_EXPRESS":         4,
	"DISPLAY_EXPRESS":        5,
	"SHOPPING_SMART_ADS":     6,
	"DISPLAY_GMAIL_AD":       7,
	"DISPLAY_SMART_CAMPAIGN": 8,
	"VIDEO_OUTSTREAM":        9,
	"VIDEO_ACTION":           10,
	"VIDEO_NON_SKIPPABLE":    11,
	"APP_CAMPAIGN":           12,
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) String() string {
	return proto.EnumName(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, int32(x))
}
func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_sub_type_02b63a9ec49568b8, []int{0, 0}
}

// An immutable specialization of an Advertising Channel.
type AdvertisingChannelSubTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertisingChannelSubTypeEnum) Reset()         { *m = AdvertisingChannelSubTypeEnum{} }
func (m *AdvertisingChannelSubTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdvertisingChannelSubTypeEnum) ProtoMessage()    {}
func (*AdvertisingChannelSubTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_sub_type_02b63a9ec49568b8, []int{0}
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Unmarshal(m, b)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdvertisingChannelSubTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.Merge(dst, src)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Size(m)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisingChannelSubTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdvertisingChannelSubTypeEnum)(nil), "google.ads.googleads.v1.enums.AdvertisingChannelSubTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType", AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/advertising_channel_sub_type.proto", fileDescriptor_advertising_channel_sub_type_02b63a9ec49568b8)
}

var fileDescriptor_advertising_channel_sub_type_02b63a9ec49568b8 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x18, 0xc5, 0x69, 0x06, 0x1b, 0xb8, 0x13, 0x33, 0x1e, 0x0c, 0x31, 0x51, 0xa4, 0xed, 0x01, 0x12,
	0x45, 0xdc, 0x85, 0x1b, 0x9c, 0xc4, 0x64, 0xd6, 0x9a, 0xc4, 0xaa, 0xd3, 0xf2, 0x47, 0x95, 0xac,
	0x74, 0x89, 0x42, 0xa4, 0xd6, 0x89, 0xea, 0xb4, 0xd2, 0x5e, 0x87, 0x4b, 0x24, 0x5e, 0x84, 0x1b,
	0xde, 0x83, 0x5b, 0x5e, 0x00, 0x25, 0x5e, 0x3a, 0x6e, 0xca, 0x4d, 0x74, 0x74, 0xbe, 0x9f, 0xcf,
	0x17, 0xf9, 0x18, 0xbc, 0x2f, 0xaa, 0xaa, 0x58, 0xe6, 0x56, 0x9a, 0x29, 0x4b, 0xcb, 0x56, 0x6d,
	0x6d, 0x2b, 0x97, 0x9b, 0x95, 0xb2, 0xd2, 0x6c, 0x9b, 0xaf, 0x9b, 0x52, 0x95, 0xb2, 0x10, 0x37,
	0x5f, 0x53, 0x29, 0xf3, 0xa5, 0x50, 0x9b, 0x85, 0x68, 0x6e, 0xeb, 0xdc, 0xac, 0xd7, 0x55, 0x53,
	0xa1, 0x91, 0x3e, 0x66, 0xa6, 0x99, 0x32, 0x77, 0x09, 0xe6, 0xd6, 0x36, 0xbb, 0x84, 0xf3, 0xd7,
	0xfd, 0x82, 0xba, 0xb4, 0x52, 0x29, 0xab, 0x26, 0x6d, 0xca, 0x4a, 0x2a, 0x7d, 0xf8, 0xf2, 0x97,
	0x01, 0x46, 0xf8, 0x7e, 0x87, 0xa7, 0x57, 0xf0, 0xcd, 0x22, 0xb9, 0xad, 0x73, 0x22, 0x37, 0xab,
	0xcb, 0x1f, 0x06, 0x78, 0xb5, 0x97, 0x40, 0x27, 0x60, 0x38, 0x8d, 0x38, 0x23, 0x1e, 0xfd, 0x40,
	0x89, 0x0f, 0x1f, 0xa0, 0x21, 0x38, 0x9a, 0x46, 0xd7, 0x51, 0xfc, 0x31, 0x82, 0x03, 0xf4, 0x02,
	0x3c, 0xe3, 0x04, 0x4f, 0xbc, 0x2b, 0x11, 0xc6, 0x2e, 0x1d, 0x13, 0x81, 0x19, 0x83, 0x06, 0x3a,
	0x03, 0xc8, 0xa7, 0x9c, 0x8d, 0xf1, 0xe7, 0x7f, 0xfd, 0x03, 0x84, 0xc0, 0xd3, 0x3b, 0x9c, 0x7c,
	0x62, 0x13, 0xc2, 0x39, 0x7c, 0x88, 0x4e, 0xc1, 0x49, 0xcf, 0xf6, 0xe6, 0xa3, 0x36, 0x80, 0x5f,
	0xc5, 0x8c, 0xd1, 0x28, 0x10, 0x3c, 0xc4, 0x93, 0x44, 0x60, 0x9f, 0xc3, 0x43, 0xf4, 0x1c, 0xc0,
	0x1e, 0x0e, 0x42, 0x4c, 0xc7, 0x02, 0xfb, 0xf0, 0x08, 0x9d, 0x83, 0xb3, 0xde, 0xd5, 0xb0, 0x87,
	0x43, 0x86, 0x69, 0x10, 0xc1, 0xc7, 0x6d, 0xfc, 0x8c, 0xfa, 0x24, 0x16, 0xf1, 0x34, 0xe1, 0xc9,
	0x84, 0xe0, 0x10, 0x3e, 0x41, 0x10, 0x1c, 0x6b, 0x13, 0x7b, 0x09, 0x8d, 0x23, 0x08, 0xd0, 0x4b,
	0x70, 0xaa, 0x9d, 0x28, 0x8e, 0x04, 0xbf, 0xa6, 0x8c, 0x61, 0x77, 0x4c, 0xe0, 0xb0, 0x45, 0x31,
	0x63, 0xf7, 0x89, 0xc7, 0xee, 0x9f, 0x01, 0xb8, 0xb8, 0xa9, 0x56, 0xe6, 0x7f, 0x5b, 0x71, 0xdf,
	0xec, 0xbd, 0x52, 0xd6, 0xf6, 0xc2, 0x06, 0x5f, 0xdc, 0xbb, 0x80, 0xa2, 0x5a, 0xa6, 0xb2, 0x30,
	0xab, 0x75, 0x61, 0x15, 0xb9, 0xec, 0x5a, 0xeb, 0x1f, 0x4a, 0x5d, 0xaa, 0x3d, 0xef, 0xe6, 0x5d,
	0xf7, 0xfd, 0x66, 0x1c, 0x04, 0x18, 0x7f, 0x37, 0x46, 0x81, 0x8e, 0xc2, 0x99, 0x32, 0xb5, 0x6c,
	0xd5, 0xcc, 0x36, 0xdb, 0x82, 0xd5, 0xcf, 0x7e, 0x3e, 0xc7, 0x99, 0x9a, 0xef, 0xe6, 0xf3, 0x99,
	0x3d, 0xef, 0xe6, 0xbf, 0x8d, 0x0b, 0x6d, 0x3a, 0x0e, 0xce, 0x94, 0xe3, 0xec, 0x08, 0xc7, 0x99,
	0xd9, 0x8e, 0xd3, 0x31, 0x8b, 0xc3, 0xee, 0xc7, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0x7b, 0x4b, 0xe9, 0xcf, 0x02, 0x00, 0x00,
}