// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Connectedness int32

const (
	Connectedness_NotConnected  Connectedness = 0
	Connectedness_Connected     Connectedness = 1
	Connectedness_CanConnect    Connectedness = 2
	Connectedness_CannotConnect Connectedness = 3
	Connectedness_Unknown       Connectedness = 4
	Connectedness_Error         Connectedness = 5
)

var Connectedness_name = map[int32]string{
	0: "NotConnected",
	1: "Connected",
	2: "CanConnect",
	3: "CannotConnect",
	4: "Unknown",
	5: "Error",
}

var Connectedness_value = map[string]int32{
	"NotConnected":  0,
	"Connected":     1,
	"CanConnect":    2,
	"CannotConnect": 3,
	"Unknown":       4,
	"Error":         5,
}

func (x Connectedness) String() string {
	return proto.EnumName(Connectedness_name, int32(x))
}

func (Connectedness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type AddrInfo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Addrs                []string `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrInfo) Reset()         { *m = AddrInfo{} }
func (m *AddrInfo) String() string { return proto.CompactTextString(m) }
func (*AddrInfo) ProtoMessage()    {}
func (*AddrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *AddrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrInfo.Unmarshal(m, b)
}
func (m *AddrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrInfo.Marshal(b, m, deterministic)
}
func (m *AddrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrInfo.Merge(m, src)
}
func (m *AddrInfo) XXX_Size() int {
	return xxx_messageInfo_AddrInfo.Size(m)
}
func (m *AddrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddrInfo proto.InternalMessageInfo

func (m *AddrInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AddrInfo) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type Location struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Latitude             float32  `protobuf:"fixed32,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Location) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type PeerInfo struct {
	AddrInfo             *AddrInfo `protobuf:"bytes,1,opt,name=addrInfo,proto3" json:"addrInfo,omitempty"`
	Location             *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerInfo.Unmarshal(m, b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return xxx_messageInfo_PeerInfo.Size(m)
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetAddrInfo() *AddrInfo {
	if m != nil {
		return m.AddrInfo
	}
	return nil
}

func (m *PeerInfo) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type ListenAddrRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenAddrRequest) Reset()         { *m = ListenAddrRequest{} }
func (m *ListenAddrRequest) String() string { return proto.CompactTextString(m) }
func (*ListenAddrRequest) ProtoMessage()    {}
func (*ListenAddrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *ListenAddrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenAddrRequest.Unmarshal(m, b)
}
func (m *ListenAddrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenAddrRequest.Marshal(b, m, deterministic)
}
func (m *ListenAddrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenAddrRequest.Merge(m, src)
}
func (m *ListenAddrRequest) XXX_Size() int {
	return xxx_messageInfo_ListenAddrRequest.Size(m)
}
func (m *ListenAddrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenAddrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenAddrRequest proto.InternalMessageInfo

type ListenAddrReply struct {
	AddrInfo             *AddrInfo `protobuf:"bytes,1,opt,name=addrInfo,proto3" json:"addrInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListenAddrReply) Reset()         { *m = ListenAddrReply{} }
func (m *ListenAddrReply) String() string { return proto.CompactTextString(m) }
func (*ListenAddrReply) ProtoMessage()    {}
func (*ListenAddrReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *ListenAddrReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenAddrReply.Unmarshal(m, b)
}
func (m *ListenAddrReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenAddrReply.Marshal(b, m, deterministic)
}
func (m *ListenAddrReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenAddrReply.Merge(m, src)
}
func (m *ListenAddrReply) XXX_Size() int {
	return xxx_messageInfo_ListenAddrReply.Size(m)
}
func (m *ListenAddrReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenAddrReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListenAddrReply proto.InternalMessageInfo

func (m *ListenAddrReply) GetAddrInfo() *AddrInfo {
	if m != nil {
		return m.AddrInfo
	}
	return nil
}

type PeersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersRequest) Reset()         { *m = PeersRequest{} }
func (m *PeersRequest) String() string { return proto.CompactTextString(m) }
func (*PeersRequest) ProtoMessage()    {}
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *PeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersRequest.Unmarshal(m, b)
}
func (m *PeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersRequest.Marshal(b, m, deterministic)
}
func (m *PeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersRequest.Merge(m, src)
}
func (m *PeersRequest) XXX_Size() int {
	return xxx_messageInfo_PeersRequest.Size(m)
}
func (m *PeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeersRequest proto.InternalMessageInfo

type PeersReply struct {
	Peers                []*PeerInfo `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PeersReply) Reset()         { *m = PeersReply{} }
func (m *PeersReply) String() string { return proto.CompactTextString(m) }
func (*PeersReply) ProtoMessage()    {}
func (*PeersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *PeersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersReply.Unmarshal(m, b)
}
func (m *PeersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersReply.Marshal(b, m, deterministic)
}
func (m *PeersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersReply.Merge(m, src)
}
func (m *PeersReply) XXX_Size() int {
	return xxx_messageInfo_PeersReply.Size(m)
}
func (m *PeersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersReply.DiscardUnknown(m)
}

var xxx_messageInfo_PeersReply proto.InternalMessageInfo

func (m *PeersReply) GetPeers() []*PeerInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

type FindPeerRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindPeerRequest) Reset()         { *m = FindPeerRequest{} }
func (m *FindPeerRequest) String() string { return proto.CompactTextString(m) }
func (*FindPeerRequest) ProtoMessage()    {}
func (*FindPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *FindPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPeerRequest.Unmarshal(m, b)
}
func (m *FindPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPeerRequest.Marshal(b, m, deterministic)
}
func (m *FindPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPeerRequest.Merge(m, src)
}
func (m *FindPeerRequest) XXX_Size() int {
	return xxx_messageInfo_FindPeerRequest.Size(m)
}
func (m *FindPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindPeerRequest proto.InternalMessageInfo

func (m *FindPeerRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type FindPeerReply struct {
	PeerInfo             *PeerInfo `protobuf:"bytes,1,opt,name=peerInfo,proto3" json:"peerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindPeerReply) Reset()         { *m = FindPeerReply{} }
func (m *FindPeerReply) String() string { return proto.CompactTextString(m) }
func (*FindPeerReply) ProtoMessage()    {}
func (*FindPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *FindPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindPeerReply.Unmarshal(m, b)
}
func (m *FindPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindPeerReply.Marshal(b, m, deterministic)
}
func (m *FindPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindPeerReply.Merge(m, src)
}
func (m *FindPeerReply) XXX_Size() int {
	return xxx_messageInfo_FindPeerReply.Size(m)
}
func (m *FindPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindPeerReply proto.InternalMessageInfo

func (m *FindPeerReply) GetPeerInfo() *PeerInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type ConnectPeerRequest struct {
	PeerInfo             *AddrInfo `protobuf:"bytes,1,opt,name=peerInfo,proto3" json:"peerInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ConnectPeerRequest) Reset()         { *m = ConnectPeerRequest{} }
func (m *ConnectPeerRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectPeerRequest) ProtoMessage()    {}
func (*ConnectPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *ConnectPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectPeerRequest.Unmarshal(m, b)
}
func (m *ConnectPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectPeerRequest.Marshal(b, m, deterministic)
}
func (m *ConnectPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectPeerRequest.Merge(m, src)
}
func (m *ConnectPeerRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectPeerRequest.Size(m)
}
func (m *ConnectPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectPeerRequest proto.InternalMessageInfo

func (m *ConnectPeerRequest) GetPeerInfo() *AddrInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

type ConnectPeerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectPeerReply) Reset()         { *m = ConnectPeerReply{} }
func (m *ConnectPeerReply) String() string { return proto.CompactTextString(m) }
func (*ConnectPeerReply) ProtoMessage()    {}
func (*ConnectPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *ConnectPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectPeerReply.Unmarshal(m, b)
}
func (m *ConnectPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectPeerReply.Marshal(b, m, deterministic)
}
func (m *ConnectPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectPeerReply.Merge(m, src)
}
func (m *ConnectPeerReply) XXX_Size() int {
	return xxx_messageInfo_ConnectPeerReply.Size(m)
}
func (m *ConnectPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectPeerReply proto.InternalMessageInfo

type DisconnectPeerRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectPeerRequest) Reset()         { *m = DisconnectPeerRequest{} }
func (m *DisconnectPeerRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectPeerRequest) ProtoMessage()    {}
func (*DisconnectPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *DisconnectPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectPeerRequest.Unmarshal(m, b)
}
func (m *DisconnectPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectPeerRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectPeerRequest.Merge(m, src)
}
func (m *DisconnectPeerRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectPeerRequest.Size(m)
}
func (m *DisconnectPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectPeerRequest proto.InternalMessageInfo

func (m *DisconnectPeerRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type DisconnectPeerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectPeerReply) Reset()         { *m = DisconnectPeerReply{} }
func (m *DisconnectPeerReply) String() string { return proto.CompactTextString(m) }
func (*DisconnectPeerReply) ProtoMessage()    {}
func (*DisconnectPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *DisconnectPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectPeerReply.Unmarshal(m, b)
}
func (m *DisconnectPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectPeerReply.Marshal(b, m, deterministic)
}
func (m *DisconnectPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectPeerReply.Merge(m, src)
}
func (m *DisconnectPeerReply) XXX_Size() int {
	return xxx_messageInfo_DisconnectPeerReply.Size(m)
}
func (m *DisconnectPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectPeerReply proto.InternalMessageInfo

type ConnectednessRequest struct {
	PeerID               string   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectednessRequest) Reset()         { *m = ConnectednessRequest{} }
func (m *ConnectednessRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectednessRequest) ProtoMessage()    {}
func (*ConnectednessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{13}
}

func (m *ConnectednessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectednessRequest.Unmarshal(m, b)
}
func (m *ConnectednessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectednessRequest.Marshal(b, m, deterministic)
}
func (m *ConnectednessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectednessRequest.Merge(m, src)
}
func (m *ConnectednessRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectednessRequest.Size(m)
}
func (m *ConnectednessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectednessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectednessRequest proto.InternalMessageInfo

func (m *ConnectednessRequest) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type ConnectednessReply struct {
	Connectedness        Connectedness `protobuf:"varint,1,opt,name=connectedness,proto3,enum=rpc.Connectedness" json:"connectedness,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConnectednessReply) Reset()         { *m = ConnectednessReply{} }
func (m *ConnectednessReply) String() string { return proto.CompactTextString(m) }
func (*ConnectednessReply) ProtoMessage()    {}
func (*ConnectednessReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{14}
}

func (m *ConnectednessReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectednessReply.Unmarshal(m, b)
}
func (m *ConnectednessReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectednessReply.Marshal(b, m, deterministic)
}
func (m *ConnectednessReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectednessReply.Merge(m, src)
}
func (m *ConnectednessReply) XXX_Size() int {
	return xxx_messageInfo_ConnectednessReply.Size(m)
}
func (m *ConnectednessReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectednessReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectednessReply proto.InternalMessageInfo

func (m *ConnectednessReply) GetConnectedness() Connectedness {
	if m != nil {
		return m.Connectedness
	}
	return Connectedness_NotConnected
}

func init() {
	proto.RegisterEnum("rpc.Connectedness", Connectedness_name, Connectedness_value)
	proto.RegisterType((*AddrInfo)(nil), "rpc.AddrInfo")
	proto.RegisterType((*Location)(nil), "rpc.Location")
	proto.RegisterType((*PeerInfo)(nil), "rpc.PeerInfo")
	proto.RegisterType((*ListenAddrRequest)(nil), "rpc.ListenAddrRequest")
	proto.RegisterType((*ListenAddrReply)(nil), "rpc.ListenAddrReply")
	proto.RegisterType((*PeersRequest)(nil), "rpc.PeersRequest")
	proto.RegisterType((*PeersReply)(nil), "rpc.PeersReply")
	proto.RegisterType((*FindPeerRequest)(nil), "rpc.FindPeerRequest")
	proto.RegisterType((*FindPeerReply)(nil), "rpc.FindPeerReply")
	proto.RegisterType((*ConnectPeerRequest)(nil), "rpc.ConnectPeerRequest")
	proto.RegisterType((*ConnectPeerReply)(nil), "rpc.ConnectPeerReply")
	proto.RegisterType((*DisconnectPeerRequest)(nil), "rpc.DisconnectPeerRequest")
	proto.RegisterType((*DisconnectPeerReply)(nil), "rpc.DisconnectPeerReply")
	proto.RegisterType((*ConnectednessRequest)(nil), "rpc.ConnectednessRequest")
	proto.RegisterType((*ConnectednessReply)(nil), "rpc.ConnectednessReply")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x97, 0xc4, 0x6e, 0xc9, 0x59, 0xd3, 0xa6, 0x67, 0xed, 0x16, 0xc3, 0x1e, 0x4a, 0x7c,
	0xe9, 0x04, 0xa3, 0x56, 0x11, 0x91, 0xc1, 0x70, 0x6d, 0xc5, 0xc2, 0x28, 0x25, 0xcc, 0x57, 0x31,
	0x26, 0xd7, 0x11, 0x16, 0xee, 0x8d, 0xc9, 0x2d, 0xb3, 0xff, 0x8e, 0x7f, 0xa7, 0x0f, 0x92, 0x5f,
	0xb7, 0x49, 0xdb, 0x21, 0x3e, 0x9e, 0xef, 0xf7, 0x9c, 0xcf, 0x3d, 0x27, 0xe7, 0x10, 0xd0, 0xbc,
	0x38, 0x74, 0xe2, 0x84, 0x71, 0x86, 0x4a, 0x12, 0xfb, 0xf6, 0x2b, 0x50, 0x3f, 0x06, 0x41, 0x32,
	0xa7, 0x3f, 0x18, 0x76, 0x40, 0x9e, 0x4f, 0x4d, 0x69, 0x28, 0x8d, 0x34, 0x57, 0x9e, 0x4f, 0xb1,
	0x0f, 0x2d, 0x2f, 0x08, 0x92, 0xd4, 0x94, 0x87, 0xca, 0x48, 0x73, 0x8b, 0xc0, 0xfe, 0x0a, 0xea,
	0x0d, 0xf3, 0x3d, 0x1e, 0x32, 0x8a, 0x26, 0x1c, 0xf9, 0x6c, 0x45, 0x79, 0xb2, 0x2e, 0xcb, 0xaa,
	0x10, 0x2d, 0x50, 0x23, 0x8f, 0x87, 0x7c, 0x15, 0x10, 0x53, 0x1e, 0x4a, 0x23, 0xd9, 0x15, 0x31,
	0x9e, 0x83, 0x16, 0x31, 0x7a, 0x57, 0x98, 0x4a, 0x6e, 0x6e, 0x04, 0xfb, 0x1b, 0xa8, 0x4b, 0x42,
	0x8a, 0x8e, 0x2e, 0x40, 0xf5, 0xca, 0xee, 0xf2, 0x07, 0x8e, 0xc7, 0xba, 0x93, 0xc4, 0xbe, 0x53,
	0xb5, 0xec, 0x0a, 0x3b, 0x4b, 0x8d, 0xca, 0xb6, 0xf2, 0x07, 0xab, 0xd4, 0xaa, 0x57, 0x57, 0xd8,
	0xf6, 0x09, 0xf4, 0x6e, 0xc2, 0x94, 0x13, 0x9a, 0x61, 0x5c, 0xf2, 0x73, 0x45, 0x52, 0x6e, 0x5f,
	0x42, 0xb7, 0x2e, 0xc6, 0xd1, 0xfa, 0x3f, 0x5e, 0xb7, 0x3b, 0xd0, 0xce, 0x9a, 0x4e, 0x2b, 0xda,
	0x6b, 0x80, 0x32, 0xce, 0x40, 0xcf, 0xa0, 0x15, 0x67, 0x91, 0x29, 0x0d, 0x15, 0x41, 0xa9, 0x86,
	0x74, 0x0b, 0xcf, 0xbe, 0x80, 0xee, 0xa7, 0x90, 0x06, 0x99, 0x5c, 0x52, 0xf0, 0x14, 0x0e, 0x33,
	0x4f, 0x2c, 0xa5, 0x8c, 0xec, 0x0f, 0xa0, 0x6f, 0x52, 0xcb, 0x4e, 0xe3, 0x12, 0xd7, 0xe8, 0x54,
	0xbc, 0x21, 0x6c, 0xfb, 0x0a, 0x70, 0xc2, 0x28, 0x25, 0x3e, 0xaf, 0xbf, 0xf4, 0x18, 0x60, 0x33,
	0xaa, 0x00, 0x20, 0x18, 0x0d, 0x40, 0x1c, 0xad, 0xed, 0x97, 0x30, 0x98, 0x86, 0xa9, 0xbf, 0xcb,
	0x7d, 0x6c, 0x82, 0x01, 0x9c, 0x6c, 0x17, 0x64, 0x1c, 0x07, 0xfa, 0x25, 0x9b, 0x04, 0x94, 0xa4,
	0xe9, 0xbf, 0x30, 0x0b, 0x31, 0x4c, 0x95, 0x9f, 0x7d, 0x8d, 0xf7, 0xa0, 0xfb, 0x75, 0x35, 0x2f,
	0xea, 0x8c, 0x31, 0x9f, 0xa8, 0x99, 0xdf, 0x4c, 0x7c, 0x7e, 0x0f, 0x7a, 0xc3, 0x47, 0x03, 0xda,
	0x0b, 0xc6, 0x85, 0x66, 0x1c, 0xa0, 0x0e, 0xda, 0x26, 0x94, 0xb0, 0x03, 0x30, 0xf1, 0x68, 0xa9,
	0x18, 0x32, 0xf6, 0x40, 0x9f, 0x78, 0x94, 0x8a, 0x1a, 0x43, 0xc1, 0x63, 0x38, 0xfa, 0x42, 0xef,
	0x29, 0x7b, 0xa0, 0xc6, 0x13, 0xd4, 0xa0, 0x35, 0x4b, 0x12, 0x96, 0x18, 0xad, 0xf1, 0x1f, 0x19,
	0x94, 0x05, 0xe1, 0x78, 0x09, 0xb0, 0xb9, 0x3c, 0x3c, 0x2d, 0xae, 0x76, 0xfb, 0x3e, 0xad, 0xfe,
	0x8e, 0x9e, 0x7d, 0xb0, 0x03, 0x7c, 0x01, 0xad, 0xfc, 0xd2, 0xb0, 0x27, 0x36, 0x5e, 0x7d, 0x36,
	0xab, 0x5b, 0x97, 0x8a, 0xf4, 0x77, 0xa0, 0x56, 0xa7, 0x83, 0x05, 0x72, 0xeb, 0xe8, 0x2c, 0xdc,
	0x52, 0x8b, 0xba, 0x2b, 0x38, 0xae, 0x6d, 0x1d, 0xcf, 0xea, 0xdf, 0xb2, 0x5e, 0x3d, 0xd8, 0x35,
	0x0a, 0xc0, 0x67, 0xe8, 0x34, 0x37, 0x8e, 0x56, 0x9e, 0xba, 0xf7, 0x6e, 0x2c, 0x73, 0xaf, 0x57,
	0x90, 0x66, 0xdb, 0x4b, 0x7a, 0xba, 0x67, 0xb1, 0x25, 0xe7, 0x6c, 0x9f, 0x95, 0x63, 0xae, 0xdf,
	0xc2, 0x79, 0xc8, 0x1c, 0x4e, 0x7e, 0xf1, 0x30, 0x22, 0x4e, 0xcc, 0x1e, 0x48, 0x72, 0xe7, 0x71,
	0xe2, 0x50, 0xc2, 0xb3, 0xa2, 0xeb, 0xf6, 0xb2, 0x92, 0x16, 0x84, 0x2f, 0xa5, 0xdf, 0xb2, 0x72,
	0x7b, 0x3b, 0xfb, 0x7e, 0x98, 0xff, 0x3b, 0xdf, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x3c,
	0x12, 0x7e, 0x48, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetClient is the client API for Net service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetClient interface {
	ListenAddr(ctx context.Context, in *ListenAddrRequest, opts ...grpc.CallOption) (*ListenAddrReply, error)
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersReply, error)
	FindPeer(ctx context.Context, in *FindPeerRequest, opts ...grpc.CallOption) (*FindPeerReply, error)
	ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerReply, error)
	DisconnectPeer(ctx context.Context, in *DisconnectPeerRequest, opts ...grpc.CallOption) (*DisconnectPeerReply, error)
	Connectedness(ctx context.Context, in *ConnectednessRequest, opts ...grpc.CallOption) (*ConnectednessReply, error)
}

type netClient struct {
	cc *grpc.ClientConn
}

func NewNetClient(cc *grpc.ClientConn) NetClient {
	return &netClient{cc}
}

func (c *netClient) ListenAddr(ctx context.Context, in *ListenAddrRequest, opts ...grpc.CallOption) (*ListenAddrReply, error) {
	out := new(ListenAddrReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/ListenAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersReply, error) {
	out := new(PeersReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netClient) FindPeer(ctx context.Context, in *FindPeerRequest, opts ...grpc.CallOption) (*FindPeerReply, error) {
	out := new(FindPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/FindPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netClient) ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerReply, error) {
	out := new(ConnectPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/ConnectPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netClient) DisconnectPeer(ctx context.Context, in *DisconnectPeerRequest, opts ...grpc.CallOption) (*DisconnectPeerReply, error) {
	out := new(DisconnectPeerReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/DisconnectPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netClient) Connectedness(ctx context.Context, in *ConnectednessRequest, opts ...grpc.CallOption) (*ConnectednessReply, error) {
	out := new(ConnectednessReply)
	err := c.cc.Invoke(ctx, "/rpc.Net/Connectedness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetServer is the server API for Net service.
type NetServer interface {
	ListenAddr(context.Context, *ListenAddrRequest) (*ListenAddrReply, error)
	Peers(context.Context, *PeersRequest) (*PeersReply, error)
	FindPeer(context.Context, *FindPeerRequest) (*FindPeerReply, error)
	ConnectPeer(context.Context, *ConnectPeerRequest) (*ConnectPeerReply, error)
	DisconnectPeer(context.Context, *DisconnectPeerRequest) (*DisconnectPeerReply, error)
	Connectedness(context.Context, *ConnectednessRequest) (*ConnectednessReply, error)
}

// UnimplementedNetServer can be embedded to have forward compatible implementations.
type UnimplementedNetServer struct {
}

func (*UnimplementedNetServer) ListenAddr(ctx context.Context, req *ListenAddrRequest) (*ListenAddrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenAddr not implemented")
}
func (*UnimplementedNetServer) Peers(ctx context.Context, req *PeersRequest) (*PeersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (*UnimplementedNetServer) FindPeer(ctx context.Context, req *FindPeerRequest) (*FindPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPeer not implemented")
}
func (*UnimplementedNetServer) ConnectPeer(ctx context.Context, req *ConnectPeerRequest) (*ConnectPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectPeer not implemented")
}
func (*UnimplementedNetServer) DisconnectPeer(ctx context.Context, req *DisconnectPeerRequest) (*DisconnectPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectPeer not implemented")
}
func (*UnimplementedNetServer) Connectedness(ctx context.Context, req *ConnectednessRequest) (*ConnectednessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connectedness not implemented")
}

func RegisterNetServer(s *grpc.Server, srv NetServer) {
	s.RegisterService(&_Net_serviceDesc, srv)
}

func _Net_ListenAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).ListenAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/ListenAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).ListenAddr(ctx, req.(*ListenAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Net_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Net_FindPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).FindPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/FindPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).FindPeer(ctx, req.(*FindPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Net_ConnectPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).ConnectPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/ConnectPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).ConnectPeer(ctx, req.(*ConnectPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Net_DisconnectPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).DisconnectPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/DisconnectPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).DisconnectPeer(ctx, req.(*DisconnectPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Net_Connectedness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectednessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetServer).Connectedness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Net/Connectedness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetServer).Connectedness(ctx, req.(*ConnectednessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Net_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Net",
	HandlerType: (*NetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListenAddr",
			Handler:    _Net_ListenAddr_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _Net_Peers_Handler,
		},
		{
			MethodName: "FindPeer",
			Handler:    _Net_FindPeer_Handler,
		},
		{
			MethodName: "ConnectPeer",
			Handler:    _Net_ConnectPeer_Handler,
		},
		{
			MethodName: "DisconnectPeer",
			Handler:    _Net_DisconnectPeer_Handler,
		},
		{
			MethodName: "Connectedness",
			Handler:    _Net_Connectedness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
