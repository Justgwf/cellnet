// Code generated by protoc-gen-go.
// source: coredef.proto
// DO NOT EDIT!

/*
Package coredef is a generated protocol buffer package.

It is generated from these files:
	coredef.proto

It has these top-level messages:
	SessionAccepted
	SessionConnected
	SessionClosed
	PeerInit
	PeerStart
	PeerStop
	UpstreamACK
	CloseClientACK
	DownstreamACK
	RemoteCallREQ
	RemoteCallACK
	TestEchoACK
*/
package coredef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 一个连接接入
type SessionAccepted struct {
}

func (m *SessionAccepted) Reset()                    { *m = SessionAccepted{} }
func (m *SessionAccepted) String() string            { return proto.CompactTextString(m) }
func (*SessionAccepted) ProtoMessage()               {}
func (*SessionAccepted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 已连接
type SessionConnected struct {
}

func (m *SessionConnected) Reset()                    { *m = SessionConnected{} }
func (m *SessionConnected) String() string            { return proto.CompactTextString(m) }
func (*SessionConnected) ProtoMessage()               {}
func (*SessionConnected) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 连接断开
type SessionClosed struct {
}

func (m *SessionClosed) Reset()                    { *m = SessionClosed{} }
func (m *SessionClosed) String() string            { return proto.CompactTextString(m) }
func (*SessionClosed) ProtoMessage()               {}
func (*SessionClosed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 端初始化
type PeerInit struct {
}

func (m *PeerInit) Reset()                    { *m = PeerInit{} }
func (m *PeerInit) String() string            { return proto.CompactTextString(m) }
func (*PeerInit) ProtoMessage()               {}
func (*PeerInit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 端启动
type PeerStart struct {
}

func (m *PeerStart) Reset()                    { *m = PeerStart{} }
func (m *PeerStart) String() string            { return proto.CompactTextString(m) }
func (*PeerStart) ProtoMessage()               {}
func (*PeerStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 端停止
type PeerStop struct {
}

func (m *PeerStop) Reset()                    { *m = PeerStop{} }
func (m *PeerStop) String() string            { return proto.CompactTextString(m) }
func (*PeerStop) ProtoMessage()               {}
func (*PeerStop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 路由上行数据
// gate -> backend
type UpstreamACK struct {
	MsgID    uint32 `protobuf:"varint,1,opt,name=MsgID" json:"MsgID,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	ClientID int64  `protobuf:"varint,3,opt,name=ClientID" json:"ClientID,omitempty"`
}

func (m *UpstreamACK) Reset()                    { *m = UpstreamACK{} }
func (m *UpstreamACK) String() string            { return proto.CompactTextString(m) }
func (*UpstreamACK) ProtoMessage()               {}
func (*UpstreamACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 关闭客户端
// backend -> gate
type CloseClientACK struct {
	ClientID int64 `protobuf:"varint,1,opt,name=ClientID" json:"ClientID,omitempty"`
}

func (m *CloseClientACK) Reset()                    { *m = CloseClientACK{} }
func (m *CloseClientACK) String() string            { return proto.CompactTextString(m) }
func (*CloseClientACK) ProtoMessage()               {}
func (*CloseClientACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 路由下行数据
// backend -> gate
type DownstreamACK struct {
	MsgID    uint32  `protobuf:"varint,1,opt,name=MsgID" json:"MsgID,omitempty"`
	Data     []byte  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	ClientID []int64 `protobuf:"varint,3,rep,name=ClientID" json:"ClientID,omitempty"`
}

func (m *DownstreamACK) Reset()                    { *m = DownstreamACK{} }
func (m *DownstreamACK) String() string            { return proto.CompactTextString(m) }
func (*DownstreamACK) ProtoMessage()               {}
func (*DownstreamACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 请求
type RemoteCallREQ struct {
	MsgID  uint32 `protobuf:"varint,1,opt,name=MsgID" json:"MsgID,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	CallID int64  `protobuf:"varint,3,opt,name=CallID" json:"CallID,omitempty"`
}

func (m *RemoteCallREQ) Reset()                    { *m = RemoteCallREQ{} }
func (m *RemoteCallREQ) String() string            { return proto.CompactTextString(m) }
func (*RemoteCallREQ) ProtoMessage()               {}
func (*RemoteCallREQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 回应
type RemoteCallACK struct {
	MsgID  uint32 `protobuf:"varint,1,opt,name=MsgID" json:"MsgID,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	CallID int64  `protobuf:"varint,3,opt,name=CallID" json:"CallID,omitempty"`
}

func (m *RemoteCallACK) Reset()                    { *m = RemoteCallACK{} }
func (m *RemoteCallACK) String() string            { return proto.CompactTextString(m) }
func (*RemoteCallACK) ProtoMessage()               {}
func (*RemoteCallACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// ==========================================================
// 测试用消息
// ==========================================================
type TestEchoACK struct {
	Content string `protobuf:"bytes,1,opt,name=Content" json:"Content,omitempty"`
}

func (m *TestEchoACK) Reset()                    { *m = TestEchoACK{} }
func (m *TestEchoACK) String() string            { return proto.CompactTextString(m) }
func (*TestEchoACK) ProtoMessage()               {}
func (*TestEchoACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*SessionAccepted)(nil), "coredef.SessionAccepted")
	proto.RegisterType((*SessionConnected)(nil), "coredef.SessionConnected")
	proto.RegisterType((*SessionClosed)(nil), "coredef.SessionClosed")
	proto.RegisterType((*PeerInit)(nil), "coredef.PeerInit")
	proto.RegisterType((*PeerStart)(nil), "coredef.PeerStart")
	proto.RegisterType((*PeerStop)(nil), "coredef.PeerStop")
	proto.RegisterType((*UpstreamACK)(nil), "coredef.UpstreamACK")
	proto.RegisterType((*CloseClientACK)(nil), "coredef.CloseClientACK")
	proto.RegisterType((*DownstreamACK)(nil), "coredef.DownstreamACK")
	proto.RegisterType((*RemoteCallREQ)(nil), "coredef.RemoteCallREQ")
	proto.RegisterType((*RemoteCallACK)(nil), "coredef.RemoteCallACK")
	proto.RegisterType((*TestEchoACK)(nil), "coredef.TestEchoACK")
}

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x69, 0xd3, 0x74, 0x62, 0xac, 0x2e, 0x22, 0xc1, 0x93, 0xe4, 0xa2, 0x07, 0xf1,
	0xe2, 0x2f, 0x28, 0x49, 0x0f, 0x45, 0x04, 0xbb, 0xb1, 0x3f, 0x20, 0x6e, 0x46, 0x0d, 0x24, 0x3b,
	0x21, 0x3b, 0xe0, 0xdf, 0x77, 0x37, 0xa6, 0x62, 0x05, 0x0f, 0xc5, 0xdb, 0xbc, 0x97, 0x97, 0x8f,
	0x37, 0xec, 0x40, 0xac, 0xa8, 0xc7, 0x0a, 0x5f, 0xef, 0xba, 0x9e, 0x98, 0xc4, 0x6c, 0x94, 0xe9,
	0x19, 0x2c, 0x0a, 0x34, 0xa6, 0x26, 0xbd, 0x54, 0x0a, 0x3b, 0xc6, 0x2a, 0x15, 0x70, 0x3a, 0x5a,
	0x19, 0x69, 0x8d, 0xca, 0x79, 0x0b, 0x88, 0x77, 0x5e, 0x43, 0xc6, 0x1a, 0x00, 0xe1, 0x13, 0x62,
	0xbf, 0xd6, 0x35, 0xa7, 0x11, 0xcc, 0xdd, 0x5c, 0x70, 0xd9, 0xf3, 0xee, 0x43, 0xc1, 0xd4, 0xa5,
	0x05, 0x44, 0xdb, 0xce, 0x70, 0x8f, 0x65, 0xbb, 0xcc, 0x1e, 0xc4, 0x39, 0x4c, 0x1f, 0xcd, 0xdb,
	0x3a, 0x4f, 0xbc, 0x2b, 0xef, 0x26, 0x96, 0xd3, 0xd6, 0x09, 0x21, 0x60, 0x92, 0x97, 0x5c, 0x26,
	0x47, 0xd6, 0x3c, 0x96, 0x93, 0xca, 0xce, 0xe2, 0x12, 0xc2, 0xac, 0xa9, 0x51, 0xb3, 0x0d, 0xfb,
	0xd6, 0xf7, 0x65, 0xa8, 0x46, 0x9d, 0xde, 0xc2, 0xc9, 0xd0, 0xe1, 0x2b, 0xe0, 0xb8, 0x3f, 0xd3,
	0xde, 0xaf, 0xf4, 0x16, 0xe2, 0x9c, 0x3e, 0xf4, 0xff, 0x4b, 0xf8, 0x7b, 0xd8, 0x0d, 0xc4, 0x12,
	0x5b, 0x62, 0xcc, 0xca, 0xa6, 0x91, 0xab, 0xcd, 0x01, 0xd8, 0x0b, 0x08, 0xdc, 0x4f, 0xdf, 0x9b,
	0x05, 0x6a, 0x50, 0xfb, 0xc8, 0xc3, 0x9a, 0xfe, 0x85, 0xbc, 0x86, 0xe8, 0x19, 0x0d, 0xaf, 0xd4,
	0x3b, 0x39, 0x60, 0x02, 0x33, 0xfb, 0xa2, 0x6c, 0x37, 0x18, 0x90, 0x73, 0x69, 0xaf, 0x60, 0x90,
	0x2f, 0xc1, 0x70, 0x15, 0xf7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x87, 0xbf, 0x99, 0x26,
	0x02, 0x00, 0x00,
}
