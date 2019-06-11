// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dnspb.proto

package dnspkg 

import (
	fmt "fmt"
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

type DnsLookupRequest struct {
	HostNames            []string `protobuf:"bytes,1,rep,name=hostNames,proto3" json:"hostNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnsLookupRequest) Reset()         { *m = DnsLookupRequest{} }
func (m *DnsLookupRequest) String() string { return proto.CompactTextString(m) }
func (*DnsLookupRequest) ProtoMessage()    {}
func (*DnsLookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f23846b31913176, []int{0}
}

func (m *DnsLookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsLookupRequest.Unmarshal(m, b)
}
func (m *DnsLookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsLookupRequest.Marshal(b, m, deterministic)
}
func (m *DnsLookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsLookupRequest.Merge(m, src)
}
func (m *DnsLookupRequest) XXX_Size() int {
	return xxx_messageInfo_DnsLookupRequest.Size(m)
}
func (m *DnsLookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsLookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DnsLookupRequest proto.InternalMessageInfo

func (m *DnsLookupRequest) GetHostNames() []string {
	if m != nil {
		return m.HostNames
	}
	return nil
}

type HostNameToIpAddressMapping struct {
	HostName             string   `protobuf:"bytes,1,opt,name=hostName,proto3" json:"hostName,omitempty"`
	Ipv4Addr             string   `protobuf:"bytes,2,opt,name=ipv4Addr,proto3" json:"ipv4Addr,omitempty"`
	Ipv6Addr             string   `protobuf:"bytes,3,opt,name=ipv6Addr,proto3" json:"ipv6Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostNameToIpAddressMapping) Reset()         { *m = HostNameToIpAddressMapping{} }
func (m *HostNameToIpAddressMapping) String() string { return proto.CompactTextString(m) }
func (*HostNameToIpAddressMapping) ProtoMessage()    {}
func (*HostNameToIpAddressMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f23846b31913176, []int{1}
}

func (m *HostNameToIpAddressMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostNameToIpAddressMapping.Unmarshal(m, b)
}
func (m *HostNameToIpAddressMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostNameToIpAddressMapping.Marshal(b, m, deterministic)
}
func (m *HostNameToIpAddressMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostNameToIpAddressMapping.Merge(m, src)
}
func (m *HostNameToIpAddressMapping) XXX_Size() int {
	return xxx_messageInfo_HostNameToIpAddressMapping.Size(m)
}
func (m *HostNameToIpAddressMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_HostNameToIpAddressMapping.DiscardUnknown(m)
}

var xxx_messageInfo_HostNameToIpAddressMapping proto.InternalMessageInfo

func (m *HostNameToIpAddressMapping) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *HostNameToIpAddressMapping) GetIpv4Addr() string {
	if m != nil {
		return m.Ipv4Addr
	}
	return ""
}

func (m *HostNameToIpAddressMapping) GetIpv6Addr() string {
	if m != nil {
		return m.Ipv6Addr
	}
	return ""
}

type DnsLookupResponse struct {
	HostNameMapping      []*HostNameToIpAddressMapping `protobuf:"bytes,1,rep,name=hostNameMapping,proto3" json:"hostNameMapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *DnsLookupResponse) Reset()         { *m = DnsLookupResponse{} }
func (m *DnsLookupResponse) String() string { return proto.CompactTextString(m) }
func (*DnsLookupResponse) ProtoMessage()    {}
func (*DnsLookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f23846b31913176, []int{2}
}

func (m *DnsLookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsLookupResponse.Unmarshal(m, b)
}
func (m *DnsLookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsLookupResponse.Marshal(b, m, deterministic)
}
func (m *DnsLookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsLookupResponse.Merge(m, src)
}
func (m *DnsLookupResponse) XXX_Size() int {
	return xxx_messageInfo_DnsLookupResponse.Size(m)
}
func (m *DnsLookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsLookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DnsLookupResponse proto.InternalMessageInfo

func (m *DnsLookupResponse) GetHostNameMapping() []*HostNameToIpAddressMapping {
	if m != nil {
		return m.HostNameMapping
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsLookupRequest)(nil), "dnsproto.DnsLookupRequest")
	proto.RegisterType((*HostNameToIpAddressMapping)(nil), "dnsproto.HostNameToIpAddressMapping")
	proto.RegisterType((*DnsLookupResponse)(nil), "dnsproto.DnsLookupResponse")
}

func init() { proto.RegisterFile("dnspb.proto", fileDescriptor_9f23846b31913176) }

var fileDescriptor_9f23846b31913176 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0x4d, 0x4b,
	0xc9, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xc9, 0x2b, 0x06, 0xb3, 0x94,
	0x0c, 0xb8, 0x04, 0x5c, 0xf2, 0x8a, 0x7d, 0xf2, 0xf3, 0xb3, 0x4b, 0x0b, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0x33, 0xf2, 0x8b, 0x4b, 0xfc, 0x12, 0x73, 0x53, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0x10, 0x02, 0x4a, 0x05, 0x5c, 0x52, 0x1e, 0x50, 0x4e,
	0x48, 0xbe, 0x67, 0x81, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x6f, 0x62, 0x41, 0x41, 0x66,
	0x5e, 0xba, 0x90, 0x14, 0x17, 0x07, 0x4c, 0xa9, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c,
	0x0f, 0x92, 0xcb, 0x2c, 0x28, 0x33, 0x01, 0xe9, 0x90, 0x60, 0x82, 0xc8, 0xc1, 0xf8, 0x50, 0x39,
	0x33, 0xb0, 0x1c, 0x33, 0x5c, 0x0e, 0xcc, 0x57, 0x4a, 0xe6, 0x12, 0x44, 0x72, 0x63, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1f, 0x17, 0x3f, 0xcc, 0x60, 0xa8, 0xdd, 0x60, 0xa7, 0x72, 0x1b,
	0xa9, 0xe8, 0xc1, 0x3c, 0xa7, 0x87, 0xdb, 0x9d, 0x41, 0xe8, 0x9a, 0x93, 0xd8, 0xc0, 0x5a, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x37, 0xe3, 0x94, 0x8f, 0x29, 0x01, 0x00, 0x00,
}