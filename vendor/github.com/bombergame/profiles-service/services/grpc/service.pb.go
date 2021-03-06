// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package profilesgrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type Credentials struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ProfileID struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileID) Reset()         { *m = ProfileID{} }
func (m *ProfileID) String() string { return proto.CompactTextString(m) }
func (*ProfileID) ProtoMessage()    {}
func (*ProfileID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ProfileID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileID.Unmarshal(m, b)
}
func (m *ProfileID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileID.Marshal(b, m, deterministic)
}
func (m *ProfileID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileID.Merge(m, src)
}
func (m *ProfileID) XXX_Size() int {
	return xxx_messageInfo_ProfileID.Size(m)
}
func (m *ProfileID) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileID.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileID proto.InternalMessageInfo

func (m *ProfileID) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "profilesgrpc.Void")
	proto.RegisterType((*Credentials)(nil), "profilesgrpc.Credentials")
	proto.RegisterType((*ProfileID)(nil), "profilesgrpc.ProfileID")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0xca, 0x4f, 0xcb,
	0xcc, 0x49, 0x2d, 0x4e, 0x2f, 0x2a, 0x48, 0x56, 0x62, 0xe3, 0x62, 0x09, 0xcb, 0xcf, 0x4c, 0x51,
	0x72, 0xe5, 0xe2, 0x76, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcc, 0x29, 0x16, 0x92,
	0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05, 0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10,
	0x39, 0x18, 0x5f, 0x49, 0x91, 0x8b, 0x33, 0x00, 0x62, 0xbc, 0xa7, 0x8b, 0x90, 0x08, 0x17, 0x6b,
	0x59, 0x62, 0x4e, 0x29, 0xc4, 0x04, 0xe6, 0x20, 0x08, 0xc7, 0x68, 0x19, 0x23, 0x17, 0x3f, 0x54,
	0x4d, 0x71, 0x30, 0xc4, 0x65, 0x42, 0x0e, 0x5c, 0xfc, 0x9e, 0x79, 0xc9, 0x50, 0xd1, 0xe0, 0xe4,
	0xfc, 0xa2, 0x54, 0x21, 0x71, 0x3d, 0x64, 0x77, 0xea, 0xc1, 0x4d, 0x95, 0x12, 0x42, 0x95, 0x00,
	0xbb, 0x9e, 0x41, 0x28, 0x90, 0x4b, 0xd2, 0x3d, 0xb5, 0x04, 0xae, 0xca, 0xa9, 0x12, 0xd9, 0x37,
	0x92, 0xa8, 0x5a, 0x90, 0xa4, 0xa4, 0x70, 0x59, 0xa3, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x2f, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x22, 0x50, 0xd7, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfilesServiceClient is the client API for ProfilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilesServiceClient interface {
	IncProfileScore(ctx context.Context, in *ProfileID, opts ...grpc.CallOption) (*Void, error)
	GetProfileIDByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*ProfileID, error)
}

type profilesServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfilesServiceClient(cc *grpc.ClientConn) ProfilesServiceClient {
	return &profilesServiceClient{cc}
}

func (c *profilesServiceClient) IncProfileScore(ctx context.Context, in *ProfileID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/profilesgrpc.ProfilesService/IncProfileScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) GetProfileIDByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*ProfileID, error) {
	out := new(ProfileID)
	err := c.cc.Invoke(ctx, "/profilesgrpc.ProfilesService/GetProfileIDByCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServiceServer is the server API for ProfilesService service.
type ProfilesServiceServer interface {
	IncProfileScore(context.Context, *ProfileID) (*Void, error)
	GetProfileIDByCredentials(context.Context, *Credentials) (*ProfileID, error)
}

func RegisterProfilesServiceServer(s *grpc.Server, srv ProfilesServiceServer) {
	s.RegisterService(&_ProfilesService_serviceDesc, srv)
}

func _ProfilesService_IncProfileScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).IncProfileScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profilesgrpc.ProfilesService/IncProfileScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).IncProfileScore(ctx, req.(*ProfileID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_GetProfileIDByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).GetProfileIDByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profilesgrpc.ProfilesService/GetProfileIDByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).GetProfileIDByCredentials(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfilesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profilesgrpc.ProfilesService",
	HandlerType: (*ProfilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncProfileScore",
			Handler:    _ProfilesService_IncProfileScore_Handler,
		},
		{
			MethodName: "GetProfileIDByCredentials",
			Handler:    _ProfilesService_GetProfileIDByCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
