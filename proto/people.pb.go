// Code generated by protoc-gen-go. DO NOT EDIT.
// source: people.proto

package people

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

type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_09461903b56db210, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_09461903b56db210, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Id struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_09461903b56db210, []int{2}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*Id)(nil), "Id")
}

func init() { proto.RegisterFile("people.proto", fileDescriptor_09461903b56db210) }

var fileDescriptor_09461903b56db210 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0xcd, 0x2f,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe3, 0x62, 0x0b, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0x13, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98,
	0x83, 0x40, 0x4c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x66, 0xb0, 0x00, 0x53, 0x66, 0x8a,
	0x12, 0x3b, 0x17, 0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x92, 0x14, 0x17, 0x93, 0x67, 0x8a, 0x90,
	0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x29, 0xc4, 0x14, 0xe6, 0x20, 0x08, 0xc7, 0x28, 0x19, 0x64,
	0x09, 0xc8, 0x52, 0x21, 0x49, 0x2e, 0x4e, 0xf7, 0xd4, 0x12, 0xa8, 0x8d, 0xcc, 0x7a, 0x9e, 0x29,
	0x52, 0xec, 0x7a, 0x50, 0x9e, 0x1c, 0x54, 0x0a, 0xac, 0x8e, 0x4d, 0x0f, 0x6c, 0x2a, 0x5c, 0xd6,
	0x80, 0x11, 0x24, 0xef, 0x98, 0x92, 0x02, 0x95, 0x87, 0x89, 0x4b, 0x41, 0x15, 0x6a, 0x30, 0x26,
	0xb1, 0x81, 0x3d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x7b, 0xf7, 0x1a, 0xe0, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeopleClient is the client API for People service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeopleClient interface {
	GetPerson(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Person, error)
	GetPeople(ctx context.Context, in *Empty, opts ...grpc.CallOption) (People_GetPeopleClient, error)
	AddPeople(ctx context.Context, opts ...grpc.CallOption) (People_AddPeopleClient, error)
}

type peopleClient struct {
	cc *grpc.ClientConn
}

func NewPeopleClient(cc *grpc.ClientConn) PeopleClient {
	return &peopleClient{cc}
}

func (c *peopleClient) GetPerson(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/People/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peopleClient) GetPeople(ctx context.Context, in *Empty, opts ...grpc.CallOption) (People_GetPeopleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_People_serviceDesc.Streams[0], "/People/GetPeople", opts...)
	if err != nil {
		return nil, err
	}
	x := &peopleGetPeopleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type People_GetPeopleClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type peopleGetPeopleClient struct {
	grpc.ClientStream
}

func (x *peopleGetPeopleClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peopleClient) AddPeople(ctx context.Context, opts ...grpc.CallOption) (People_AddPeopleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_People_serviceDesc.Streams[1], "/People/AddPeople", opts...)
	if err != nil {
		return nil, err
	}
	x := &peopleAddPeopleClient{stream}
	return x, nil
}

type People_AddPeopleClient interface {
	Send(*Person) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type peopleAddPeopleClient struct {
	grpc.ClientStream
}

func (x *peopleAddPeopleClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peopleAddPeopleClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeopleServer is the server API for People service.
type PeopleServer interface {
	GetPerson(context.Context, *Id) (*Person, error)
	GetPeople(*Empty, People_GetPeopleServer) error
	AddPeople(People_AddPeopleServer) error
}

// UnimplementedPeopleServer can be embedded to have forward compatible implementations.
type UnimplementedPeopleServer struct {
}

func (*UnimplementedPeopleServer) GetPerson(ctx context.Context, req *Id) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedPeopleServer) GetPeople(req *Empty, srv People_GetPeopleServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPeople not implemented")
}
func (*UnimplementedPeopleServer) AddPeople(srv People_AddPeopleServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPeople not implemented")
}

func RegisterPeopleServer(s *grpc.Server, srv PeopleServer) {
	s.RegisterService(&_People_serviceDesc, srv)
}

func _People_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/People/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServer).GetPerson(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _People_GetPeople_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeopleServer).GetPeople(m, &peopleGetPeopleServer{stream})
}

type People_GetPeopleServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type peopleGetPeopleServer struct {
	grpc.ServerStream
}

func (x *peopleGetPeopleServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

func _People_AddPeople_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeopleServer).AddPeople(&peopleAddPeopleServer{stream})
}

type People_AddPeopleServer interface {
	SendAndClose(*Empty) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type peopleAddPeopleServer struct {
	grpc.ServerStream
}

func (x *peopleAddPeopleServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peopleAddPeopleServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _People_serviceDesc = grpc.ServiceDesc{
	ServiceName: "People",
	HandlerType: (*PeopleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPerson",
			Handler:    _People_GetPerson_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPeople",
			Handler:       _People_GetPeople_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddPeople",
			Handler:       _People_AddPeople_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "people.proto",
}
