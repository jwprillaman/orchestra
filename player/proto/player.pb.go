// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

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

type PlayRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Params               []string `protobuf:"bytes,3,rep,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayRequest) Reset()         { *m = PlayRequest{} }
func (m *PlayRequest) String() string { return proto.CompactTextString(m) }
func (*PlayRequest) ProtoMessage()    {}
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_5a478e18fb3bcbb7, []int{0}
}
func (m *PlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayRequest.Unmarshal(m, b)
}
func (m *PlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayRequest.Marshal(b, m, deterministic)
}
func (dst *PlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayRequest.Merge(dst, src)
}
func (m *PlayRequest) XXX_Size() int {
	return xxx_messageInfo_PlayRequest.Size(m)
}
func (m *PlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayRequest proto.InternalMessageInfo

func (m *PlayRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlayRequest) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

type PlayResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayResponse) Reset()         { *m = PlayResponse{} }
func (m *PlayResponse) String() string { return proto.CompactTextString(m) }
func (*PlayResponse) ProtoMessage()    {}
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_5a478e18fb3bcbb7, []int{1}
}
func (m *PlayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayResponse.Unmarshal(m, b)
}
func (m *PlayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayResponse.Marshal(b, m, deterministic)
}
func (dst *PlayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayResponse.Merge(dst, src)
}
func (m *PlayResponse) XXX_Size() int {
	return xxx_messageInfo_PlayResponse.Size(m)
}
func (m *PlayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlayResponse proto.InternalMessageInfo

func (m *PlayResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PlayResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StopRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_5a478e18fb3bcbb7, []int{2}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (dst *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(dst, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StopResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_player_5a478e18fb3bcbb7, []int{3}
}
func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (dst *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(dst, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func (m *StopResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StopResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PlayRequest)(nil), "proto.PlayRequest")
	proto.RegisterType((*PlayResponse)(nil), "proto.PlayResponse")
	proto.RegisterType((*StopRequest)(nil), "proto.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "proto.StopResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlayerClient is the client API for Player service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerClient interface {
	Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (*PlayResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type playerClient struct {
	cc *grpc.ClientConn
}

func NewPlayerClient(cc *grpc.ClientConn) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (*PlayResponse, error) {
	out := new(PlayResponse)
	err := c.cc.Invoke(ctx, "/proto.Player/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/proto.Player/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServer is the server API for Player service.
type PlayerServer interface {
	Play(context.Context, *PlayRequest) (*PlayResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
}

func RegisterPlayerServer(s *grpc.Server, srv PlayerServer) {
	s.RegisterService(&_Player_serviceDesc, srv)
}

func _Player_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Player/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).Play(ctx, req.(*PlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Player_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Player/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Player_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _Player_Play_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Player_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player.proto",
}

func init() { proto.RegisterFile("player.proto", fileDescriptor_player_5a478e18fb3bcbb7) }

var fileDescriptor_player_5a478e18fb3bcbb7 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x37, 0xc9, 0x5a, 0xdd, 0xd9, 0x22, 0x32, 0x82, 0x04, 0x41, 0x28, 0x39, 0xf5, 0xb4,
	0xa0, 0x9e, 0xf4, 0x11, 0x04, 0xa1, 0xa4, 0x4f, 0x10, 0xdb, 0x20, 0x85, 0xb6, 0x89, 0x4d, 0x7b,
	0xe8, 0xdb, 0x4b, 0x92, 0xb6, 0xd4, 0x93, 0xa7, 0xfc, 0x33, 0xe1, 0x9b, 0xff, 0x9f, 0x81, 0xd4,
	0xb6, 0x6a, 0xd6, 0xc3, 0xc5, 0x0e, 0x66, 0x34, 0x78, 0x15, 0x1e, 0xf1, 0x06, 0xe7, 0xa2, 0x55,
	0xb3, 0xd4, 0x3f, 0x93, 0x76, 0x23, 0x22, 0x1c, 0x7b, 0xd5, 0x69, 0x4e, 0x32, 0x92, 0x9f, 0x64,
	0xd0, 0xf8, 0x00, 0x49, 0xa1, 0x06, 0xd5, 0x39, 0xce, 0x32, 0x96, 0x9f, 0xe4, 0x52, 0x89, 0x0f,
	0x48, 0x23, 0xea, 0xac, 0xe9, 0x9d, 0x46, 0x0e, 0xd7, 0xe5, 0x54, 0x55, 0xda, 0xb9, 0x80, 0xdf,
	0xc8, 0xb5, 0xc4, 0x5b, 0xa0, 0x4d, 0xcd, 0x69, 0x46, 0x72, 0x26, 0x69, 0x53, 0xe3, 0x1d, 0xb0,
	0x4f, 0xf7, 0xcd, 0x59, 0x30, 0xf1, 0x52, 0x3c, 0xc1, 0xb9, 0x1c, 0x8d, 0x5d, 0x63, 0x44, 0x80,
	0xac, 0x80, 0x78, 0x87, 0x34, 0x7e, 0xff, 0x6b, 0xb5, 0x8c, 0xa6, 0xdb, 0xe8, 0x97, 0x1e, 0x92,
	0x22, 0x2c, 0x8e, 0xcf, 0x70, 0xf4, 0x0a, 0x31, 0x9e, 0xe0, 0xb2, 0x5b, 0xfc, 0xf1, 0xfe, 0x4f,
	0x2f, 0xda, 0x88, 0x83, 0x47, 0xbc, 0xf1, 0x86, 0xec, 0x42, 0x6e, 0xc8, 0x3e, 0x99, 0x38, 0x7c,
	0x25, 0xa1, 0xfb, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x17, 0x73, 0x7e, 0x22, 0x6f, 0x01, 0x00,
	0x00,
}
