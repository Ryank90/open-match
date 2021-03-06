// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf-spec/frontend.proto

package pb

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

type CreatePlayerRequest struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerRequest) Reset()         { *m = CreatePlayerRequest{} }
func (m *CreatePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerRequest) ProtoMessage()    {}
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{0}
}

func (m *CreatePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerRequest.Unmarshal(m, b)
}
func (m *CreatePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerRequest.Marshal(b, m, deterministic)
}
func (m *CreatePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerRequest.Merge(m, src)
}
func (m *CreatePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerRequest.Size(m)
}
func (m *CreatePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerRequest proto.InternalMessageInfo

func (m *CreatePlayerRequest) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type CreatePlayerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerResponse) Reset()         { *m = CreatePlayerResponse{} }
func (m *CreatePlayerResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerResponse) ProtoMessage()    {}
func (*CreatePlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{1}
}

func (m *CreatePlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerResponse.Unmarshal(m, b)
}
func (m *CreatePlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerResponse.Marshal(b, m, deterministic)
}
func (m *CreatePlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerResponse.Merge(m, src)
}
func (m *CreatePlayerResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerResponse.Size(m)
}
func (m *CreatePlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerResponse proto.InternalMessageInfo

type DeletePlayerRequest struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlayerRequest) Reset()         { *m = DeletePlayerRequest{} }
func (m *DeletePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePlayerRequest) ProtoMessage()    {}
func (*DeletePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{2}
}

func (m *DeletePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlayerRequest.Unmarshal(m, b)
}
func (m *DeletePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlayerRequest.Marshal(b, m, deterministic)
}
func (m *DeletePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlayerRequest.Merge(m, src)
}
func (m *DeletePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePlayerRequest.Size(m)
}
func (m *DeletePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlayerRequest proto.InternalMessageInfo

func (m *DeletePlayerRequest) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type DeletePlayerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePlayerResponse) Reset()         { *m = DeletePlayerResponse{} }
func (m *DeletePlayerResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePlayerResponse) ProtoMessage()    {}
func (*DeletePlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{3}
}

func (m *DeletePlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePlayerResponse.Unmarshal(m, b)
}
func (m *DeletePlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePlayerResponse.Marshal(b, m, deterministic)
}
func (m *DeletePlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePlayerResponse.Merge(m, src)
}
func (m *DeletePlayerResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePlayerResponse.Size(m)
}
func (m *DeletePlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePlayerResponse proto.InternalMessageInfo

type GetUpdatesRequest struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpdatesRequest) Reset()         { *m = GetUpdatesRequest{} }
func (m *GetUpdatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetUpdatesRequest) ProtoMessage()    {}
func (*GetUpdatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{4}
}

func (m *GetUpdatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdatesRequest.Unmarshal(m, b)
}
func (m *GetUpdatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdatesRequest.Marshal(b, m, deterministic)
}
func (m *GetUpdatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdatesRequest.Merge(m, src)
}
func (m *GetUpdatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetUpdatesRequest.Size(m)
}
func (m *GetUpdatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdatesRequest proto.InternalMessageInfo

func (m *GetUpdatesRequest) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type GetUpdatesResponse struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpdatesResponse) Reset()         { *m = GetUpdatesResponse{} }
func (m *GetUpdatesResponse) String() string { return proto.CompactTextString(m) }
func (*GetUpdatesResponse) ProtoMessage()    {}
func (*GetUpdatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6805b20a50ffa9ae, []int{5}
}

func (m *GetUpdatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpdatesResponse.Unmarshal(m, b)
}
func (m *GetUpdatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpdatesResponse.Marshal(b, m, deterministic)
}
func (m *GetUpdatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpdatesResponse.Merge(m, src)
}
func (m *GetUpdatesResponse) XXX_Size() int {
	return xxx_messageInfo_GetUpdatesResponse.Size(m)
}
func (m *GetUpdatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpdatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpdatesResponse proto.InternalMessageInfo

func (m *GetUpdatesResponse) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePlayerRequest)(nil), "api.CreatePlayerRequest")
	proto.RegisterType((*CreatePlayerResponse)(nil), "api.CreatePlayerResponse")
	proto.RegisterType((*DeletePlayerRequest)(nil), "api.DeletePlayerRequest")
	proto.RegisterType((*DeletePlayerResponse)(nil), "api.DeletePlayerResponse")
	proto.RegisterType((*GetUpdatesRequest)(nil), "api.GetUpdatesRequest")
	proto.RegisterType((*GetUpdatesResponse)(nil), "api.GetUpdatesResponse")
}

func init() { proto.RegisterFile("api/protobuf-spec/frontend.proto", fileDescriptor_6805b20a50ffa9ae) }

var fileDescriptor_6805b20a50ffa9ae = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x1b, 0xfe, 0x50, 0xfe, 0xac, 0x1e, 0x74, 0x95, 0x5a, 0x73, 0x2a, 0x3d, 0xf5, 0xd2,
	0xac, 0x54, 0xc4, 0x93, 0x8a, 0x56, 0xed, 0xb5, 0x08, 0x5e, 0xbc, 0x6d, 0x92, 0x49, 0x1a, 0xd8,
	0xec, 0x8e, 0xbb, 0x93, 0x83, 0x9f, 0xd5, 0x2f, 0x23, 0x4d, 0x57, 0x4d, 0xe9, 0x1e, 0x14, 0xaf,
	0x6f, 0xde, 0xfb, 0xc1, 0x7b, 0x0c, 0x1b, 0x49, 0xac, 0x04, 0x5a, 0x43, 0x26, 0x6d, 0x8a, 0xa9,
	0x43, 0xc8, 0x44, 0x61, 0x8d, 0x26, 0xd0, 0x79, 0xd2, 0xca, 0xfc, 0x9f, 0xc4, 0x2a, 0x0e, 0xd8,
	0x6a, 0x70, 0x4e, 0x96, 0xe0, 0x36, 0xb6, 0xf1, 0x0d, 0x3b, 0x9a, 0x5b, 0x90, 0x04, 0x4b, 0x25,
	0xdf, 0xc0, 0x3e, 0xc1, 0x6b, 0x03, 0x8e, 0xf8, 0x84, 0xf5, 0xb1, 0x15, 0x86, 0xd1, 0x28, 0x9a,
	0xec, 0xcd, 0x0e, 0x92, 0xaf, 0x9c, 0x37, 0xfa, 0xfb, 0x78, 0xc0, 0x8e, 0xb7, 0x01, 0x0e, 0x8d,
	0x76, 0xb0, 0x06, 0xdf, 0x83, 0x82, 0x3f, 0x81, 0xb7, 0x01, 0x1e, 0x7c, 0xc5, 0x0e, 0x17, 0x40,
	0xcf, 0x98, 0x4b, 0x02, 0xf7, 0x7b, 0xec, 0x35, 0xe3, 0xdd, 0xf8, 0x06, 0xfa, 0xf3, 0xfc, 0xec,
	0x3d, 0x62, 0xff, 0x1f, 0xfd, 0xd4, 0xfc, 0x81, 0xed, 0x77, 0xcb, 0xf3, 0x61, 0x22, 0xb1, 0x4a,
	0x02, 0x83, 0xc6, 0xa7, 0x81, 0x8b, 0x2f, 0xd4, 0x5b, 0x63, 0xba, 0x55, 0x3d, 0x26, 0x30, 0x9f,
	0xc7, 0x04, 0x77, 0xe9, 0xf1, 0x5b, 0xc6, 0xbe, 0xab, 0xf1, 0x41, 0x6b, 0xdd, 0x99, 0x2a, 0x3e,
	0xd9, 0xd1, 0x3f, 0x01, 0x67, 0xd1, 0xdd, 0xe5, 0xcb, 0x45, 0x59, 0xd1, 0xaa, 0x49, 0x93, 0xcc,
	0xd4, 0x62, 0x61, 0x4c, 0xa9, 0x60, 0xae, 0x4c, 0x93, 0x2f, 0x95, 0xa4, 0xc2, 0xd8, 0x5a, 0x18,
	0x04, 0x3d, 0xad, 0x25, 0x65, 0x2b, 0x51, 0x69, 0x02, 0xab, 0xa5, 0x12, 0x98, 0xa6, 0xfd, 0xf6,
	0x9d, 0xce, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xc0, 0x79, 0x30, 0x99, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	// CreatePlayer will put the player  in state storage, and then look
	// through the 'properties' field for the attributes you have defined as
	// indices your matchmaker config.  If the attributes exist and are valid
	// integers, they will be indexed.
	// INPUT: Player message with these fields populated:
	//  - id
	//  - properties
	// OUTPUT: Result message denoting success or failure (and an error if
	// necessary)
	CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error)
	// DeletePlayer removes the player from state storage by doing the
	// following:
	//  1) Delete player from configured indices.  This effectively removes the
	//     player from matchmaking when using recommended MMF patterns.
	//     Everything after this is just cleanup to save stage storage space.
	//  2) 'Lazily' delete the player's state storage record.  This is kicked
	//     off in the background and may take some time to complete.
	//  2) 'Lazily' delete the player's metadata indicies (like, the timestamp when
	//     they called CreatePlayer, and the last time the record was accessed).  This
	//     is also kicked off in the background and may take some time to complete.
	// INPUT: Player message with the 'id' field populated.
	// OUTPUT: Result message denoting success or failure (and an error if
	// necessary)
	DeletePlayer(ctx context.Context, in *DeletePlayerRequest, opts ...grpc.CallOption) (*DeletePlayerResponse, error)
	// GetUpdates streams matchmaking results from Open Match for the
	// provided player ID.
	// INPUT: Player message with the 'id' field populated.
	// OUTPUT: a stream of player objects with one or more of the following
	// fields populated, if an update to that field is seen in state storage:
	//  - 'assignment': string that usually contains game server connection information.
	//  - 'status': string to communicate current matchmaking status to the client.
	//  - 'error': string to pass along error information to the client.
	//
	// During normal operation, the expectation is that the 'assignment' field
	// will be updated by a Backend process calling the 'CreateAssignments' Backend API
	// endpoint.  'Status' and 'Error' are free for developers to use as they see fit.
	// Even if you had multiple players enter a matchmaking request as a group, the
	// Backend API 'CreateAssignments' call will write the results to state
	// storage separately under each player's ID. OM expects you to make all game
	// clients 'GetUpdates' with their own ID from the Frontend API to get
	// their results.
	//
	// NOTE: This call generates a small amount of load on the Frontend API and state
	//  storage while watching the player record for updates. You are expected
	//  to close the stream from your client after receiving your matchmaking
	//  results (or a reasonable timeout), or you will continue to
	//  generate load on OM until you do!
	// NOTE: Just bear in mind that every update will send egress traffic from
	//  Open Match to game clients! Frugality is recommended.
	GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (Frontend_GetUpdatesClient, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error) {
	out := new(CreatePlayerResponse)
	err := c.cc.Invoke(ctx, "/api.Frontend/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) DeletePlayer(ctx context.Context, in *DeletePlayerRequest, opts ...grpc.CallOption) (*DeletePlayerResponse, error) {
	out := new(DeletePlayerResponse)
	err := c.cc.Invoke(ctx, "/api.Frontend/DeletePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontendClient) GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (Frontend_GetUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/api.Frontend/GetUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendGetUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Frontend_GetUpdatesClient interface {
	Recv() (*GetUpdatesResponse, error)
	grpc.ClientStream
}

type frontendGetUpdatesClient struct {
	grpc.ClientStream
}

func (x *frontendGetUpdatesClient) Recv() (*GetUpdatesResponse, error) {
	m := new(GetUpdatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	// CreatePlayer will put the player  in state storage, and then look
	// through the 'properties' field for the attributes you have defined as
	// indices your matchmaker config.  If the attributes exist and are valid
	// integers, they will be indexed.
	// INPUT: Player message with these fields populated:
	//  - id
	//  - properties
	// OUTPUT: Result message denoting success or failure (and an error if
	// necessary)
	CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerResponse, error)
	// DeletePlayer removes the player from state storage by doing the
	// following:
	//  1) Delete player from configured indices.  This effectively removes the
	//     player from matchmaking when using recommended MMF patterns.
	//     Everything after this is just cleanup to save stage storage space.
	//  2) 'Lazily' delete the player's state storage record.  This is kicked
	//     off in the background and may take some time to complete.
	//  2) 'Lazily' delete the player's metadata indicies (like, the timestamp when
	//     they called CreatePlayer, and the last time the record was accessed).  This
	//     is also kicked off in the background and may take some time to complete.
	// INPUT: Player message with the 'id' field populated.
	// OUTPUT: Result message denoting success or failure (and an error if
	// necessary)
	DeletePlayer(context.Context, *DeletePlayerRequest) (*DeletePlayerResponse, error)
	// GetUpdates streams matchmaking results from Open Match for the
	// provided player ID.
	// INPUT: Player message with the 'id' field populated.
	// OUTPUT: a stream of player objects with one or more of the following
	// fields populated, if an update to that field is seen in state storage:
	//  - 'assignment': string that usually contains game server connection information.
	//  - 'status': string to communicate current matchmaking status to the client.
	//  - 'error': string to pass along error information to the client.
	//
	// During normal operation, the expectation is that the 'assignment' field
	// will be updated by a Backend process calling the 'CreateAssignments' Backend API
	// endpoint.  'Status' and 'Error' are free for developers to use as they see fit.
	// Even if you had multiple players enter a matchmaking request as a group, the
	// Backend API 'CreateAssignments' call will write the results to state
	// storage separately under each player's ID. OM expects you to make all game
	// clients 'GetUpdates' with their own ID from the Frontend API to get
	// their results.
	//
	// NOTE: This call generates a small amount of load on the Frontend API and state
	//  storage while watching the player record for updates. You are expected
	//  to close the stream from your client after receiving your matchmaking
	//  results (or a reasonable timeout), or you will continue to
	//  generate load on OM until you do!
	// NOTE: Just bear in mind that every update will send egress traffic from
	//  Open Match to game clients! Frugality is recommended.
	GetUpdates(*GetUpdatesRequest, Frontend_GetUpdatesServer) error
}

// UnimplementedFrontendServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (*UnimplementedFrontendServer) CreatePlayer(ctx context.Context, req *CreatePlayerRequest) (*CreatePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}
func (*UnimplementedFrontendServer) DeletePlayer(ctx context.Context, req *DeletePlayerRequest) (*DeletePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlayer not implemented")
}
func (*UnimplementedFrontendServer) GetUpdates(req *GetUpdatesRequest, srv Frontend_GetUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Frontend/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).CreatePlayer(ctx, req.(*CreatePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_DeletePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).DeletePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Frontend/DeletePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).DeletePlayer(ctx, req.(*DeletePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontend_GetUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUpdatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FrontendServer).GetUpdates(m, &frontendGetUpdatesServer{stream})
}

type Frontend_GetUpdatesServer interface {
	Send(*GetUpdatesResponse) error
	grpc.ServerStream
}

type frontendGetUpdatesServer struct {
	grpc.ServerStream
}

func (x *frontendGetUpdatesServer) Send(m *GetUpdatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayer",
			Handler:    _Frontend_CreatePlayer_Handler,
		},
		{
			MethodName: "DeletePlayer",
			Handler:    _Frontend_DeletePlayer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUpdates",
			Handler:       _Frontend_GetUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/protobuf-spec/frontend.proto",
}
