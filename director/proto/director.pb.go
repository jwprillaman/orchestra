// Code generated by protoc-gen-go. DO NOT EDIT.
// source: director.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Players struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Players) Reset()         { *m = Players{} }
func (m *Players) String() string { return proto.CompactTextString(m) }
func (*Players) ProtoMessage()    {}
func (*Players) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{0}
}
func (m *Players) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Players.Unmarshal(m, b)
}
func (m *Players) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Players.Marshal(b, m, deterministic)
}
func (dst *Players) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Players.Merge(dst, src)
}
func (m *Players) XXX_Size() int {
	return xxx_messageInfo_Players.Size(m)
}
func (m *Players) XXX_DiscardUnknown() {
	xxx_messageInfo_Players.DiscardUnknown(m)
}

var xxx_messageInfo_Players proto.InternalMessageInfo

func (m *Players) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type Songs struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Players              []string `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Songs) Reset()         { *m = Songs{} }
func (m *Songs) String() string { return proto.CompactTextString(m) }
func (*Songs) ProtoMessage()    {}
func (*Songs) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{1}
}
func (m *Songs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Songs.Unmarshal(m, b)
}
func (m *Songs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Songs.Marshal(b, m, deterministic)
}
func (dst *Songs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Songs.Merge(dst, src)
}
func (m *Songs) XXX_Size() int {
	return xxx_messageInfo_Songs.Size(m)
}
func (m *Songs) XXX_DiscardUnknown() {
	xxx_messageInfo_Songs.DiscardUnknown(m)
}

var xxx_messageInfo_Songs proto.InternalMessageInfo

func (m *Songs) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Songs) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

type Filter struct {
	PlayerName           string   `protobuf:"bytes,1,opt,name=playerName,proto3" json:"playerName,omitempty"`
	SongIds              string   `protobuf:"bytes,2,opt,name=songIds,proto3" json:"songIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{2}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *Filter) GetSongIds() string {
	if m != nil {
		return m.SongIds
	}
	return ""
}

type Player struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{3}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PlayerReport struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mem                  uint64   `protobuf:"varint,2,opt,name=mem,proto3" json:"mem,omitempty"`
	SongIds              []int64  `protobuf:"varint,3,rep,packed,name=songIds,proto3" json:"songIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerReport) Reset()         { *m = PlayerReport{} }
func (m *PlayerReport) String() string { return proto.CompactTextString(m) }
func (*PlayerReport) ProtoMessage()    {}
func (*PlayerReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_director_0515fb61c22c641e, []int{5}
}
func (m *PlayerReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerReport.Unmarshal(m, b)
}
func (m *PlayerReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerReport.Marshal(b, m, deterministic)
}
func (dst *PlayerReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerReport.Merge(dst, src)
}
func (m *PlayerReport) XXX_Size() int {
	return xxx_messageInfo_PlayerReport.Size(m)
}
func (m *PlayerReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerReport.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerReport proto.InternalMessageInfo

func (m *PlayerReport) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlayerReport) GetMem() uint64 {
	if m != nil {
		return m.Mem
	}
	return 0
}

func (m *PlayerReport) GetSongIds() []int64 {
	if m != nil {
		return m.SongIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Players)(nil), "proto.Players")
	proto.RegisterType((*Songs)(nil), "proto.Songs")
	proto.RegisterType((*Filter)(nil), "proto.Filter")
	proto.RegisterType((*Player)(nil), "proto.Player")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*PlayerReport)(nil), "proto.PlayerReport")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DirectorClient is the client API for Director service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DirectorClient interface {
	GetPlayers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Players, error)
	GetSongs(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Songs, error)
	RegisterPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Response, error)
	RemovePlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Response, error)
	Report(ctx context.Context, in *PlayerReport, opts ...grpc.CallOption) (*Response, error)
}

type directorClient struct {
	cc *grpc.ClientConn
}

func NewDirectorClient(cc *grpc.ClientConn) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) GetPlayers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Players, error) {
	out := new(Players)
	err := c.cc.Invoke(ctx, "/proto.Director/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) GetSongs(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Songs, error) {
	out := new(Songs)
	err := c.cc.Invoke(ctx, "/proto.Director/GetSongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) RegisterPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Director/RegisterPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) RemovePlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Director/RemovePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) Report(ctx context.Context, in *PlayerReport, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Director/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectorServer is the server API for Director service.
type DirectorServer interface {
	GetPlayers(context.Context, *Filter) (*Players, error)
	GetSongs(context.Context, *Filter) (*Songs, error)
	RegisterPlayer(context.Context, *Player) (*Response, error)
	RemovePlayer(context.Context, *Player) (*Response, error)
	Report(context.Context, *PlayerReport) (*Response, error)
}

func RegisterDirectorServer(s *grpc.Server, srv DirectorServer) {
	s.RegisterService(&_Director_serviceDesc, srv)
}

func _Director_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Director/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).GetPlayers(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_GetSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).GetSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Director/GetSongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).GetSongs(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_RegisterPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).RegisterPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Director/RegisterPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).RegisterPlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_RemovePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).RemovePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Director/RemovePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).RemovePlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Director/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).Report(ctx, req.(*PlayerReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _Director_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayers",
			Handler:    _Director_GetPlayers_Handler,
		},
		{
			MethodName: "GetSongs",
			Handler:    _Director_GetSongs_Handler,
		},
		{
			MethodName: "RegisterPlayer",
			Handler:    _Director_RegisterPlayer_Handler,
		},
		{
			MethodName: "RemovePlayer",
			Handler:    _Director_RemovePlayer_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _Director_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "director.proto",
}

func init() { proto.RegisterFile("director.proto", fileDescriptor_director_0515fb61c22c641e) }

var fileDescriptor_director_0515fb61c22c641e = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x13, 0xd3, 0xa4, 0xe9, 0x90, 0x46, 0x59, 0x3d, 0x84, 0x22, 0x5a, 0xf6, 0x54, 0x0f,
	0x16, 0x69, 0xc1, 0x07, 0x10, 0xb1, 0x78, 0x29, 0xb2, 0x3e, 0x41, 0x4c, 0x86, 0x10, 0x68, 0xb2,
	0x61, 0x77, 0x15, 0x7c, 0x77, 0x0f, 0xb2, 0xff, 0xb0, 0x95, 0x5e, 0x7a, 0xda, 0xf9, 0x66, 0xe6,
	0xfb, 0x98, 0xfd, 0x41, 0x5e, 0xb7, 0x02, 0x2b, 0xc5, 0xc5, 0x72, 0x10, 0x5c, 0x71, 0x12, 0x9b,
	0x87, 0xde, 0xc2, 0xf8, 0x6d, 0x57, 0x7e, 0xa3, 0x90, 0xe4, 0x0a, 0xe2, 0xbe, 0xec, 0x50, 0x16,
	0xe1, 0x3c, 0x5a, 0x4c, 0x98, 0x15, 0x74, 0x0d, 0xf1, 0x3b, 0xef, 0x1b, 0x49, 0x2e, 0x20, 0x6a,
	0x6b, 0x3b, 0x8c, 0x98, 0x2e, 0x49, 0x01, 0xe3, 0xc1, 0x7a, 0x8b, 0x33, 0x63, 0xf1, 0x92, 0x3e,
	0x41, 0xf2, 0xd2, 0xee, 0x14, 0x0a, 0x72, 0x03, 0x60, 0x9b, 0xdb, 0xb2, 0xc3, 0x22, 0x9c, 0x87,
	0x8b, 0x09, 0xdb, 0xeb, 0xe8, 0x0c, 0xc9, 0xfb, 0xe6, 0xb5, 0xd6, 0x19, 0x7a, 0xe8, 0x25, 0xbd,
	0x86, 0xc4, 0x5e, 0x46, 0x08, 0x8c, 0xfa, 0x3f, 0xb7, 0xa9, 0xe9, 0x23, 0xa4, 0x0c, 0xe5, 0xc0,
	0x7b, 0x69, 0x33, 0x3e, 0xab, 0x0a, 0xa5, 0x34, 0x2b, 0x29, 0xf3, 0x52, 0xdf, 0xdc, 0xc9, 0xc6,
	0x25, 0xeb, 0x92, 0x6e, 0x21, 0xb3, 0xa9, 0x0c, 0x07, 0x2e, 0xd4, 0xb1, 0x6c, 0xe3, 0xc2, 0xce,
	0xb8, 0x46, 0x4c, 0x97, 0xfb, 0x57, 0x46, 0xe6, 0xff, 0x5e, 0xae, 0x7e, 0x42, 0x48, 0x9f, 0x1d,
	0x59, 0x72, 0x0f, 0xb0, 0x41, 0xe5, 0x79, 0x4e, 0x2d, 0xe9, 0xa5, 0x25, 0x31, 0xcb, 0x9d, 0x74,
	0x63, 0x1a, 0x90, 0x3b, 0x48, 0x37, 0xa8, 0x2c, 0xdd, 0x7f, 0xcb, 0x99, 0x93, 0x66, 0x48, 0x03,
	0xb2, 0x82, 0x9c, 0x61, 0xd3, 0x4a, 0x85, 0xc2, 0x41, 0x99, 0x1e, 0xc4, 0xcd, 0xce, 0x9d, 0xf4,
	0x50, 0x68, 0x40, 0x1e, 0x20, 0x63, 0xd8, 0xf1, 0x2f, 0x3c, 0xc1, 0x91, 0x38, 0x2c, 0x97, 0x07,
	0xbb, 0xb6, 0x79, 0xc4, 0xf1, 0x91, 0x98, 0xce, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x03, 0xad,
	0x5b, 0x86, 0x5e, 0x02, 0x00, 0x00,
}
