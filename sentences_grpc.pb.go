// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sentences

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SentencesClient is the client API for Sentences service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentencesClient interface {
	GetAll(ctx context.Context, in *Sentences, opts ...grpc.CallOption) (*Sentences, error)
	Get(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error)
	GetResponse(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error)
	Add(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error)
	AddSentenceGroup(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
	DeleteSentenceGroup(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
	Moderate(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
	Used(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
	VoteGood(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
	VoteBad(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error)
}

type sentencesClient struct {
	cc grpc.ClientConnInterface
}

func NewSentencesClient(cc grpc.ClientConnInterface) SentencesClient {
	return &sentencesClient{cc}
}

func (c *sentencesClient) GetAll(ctx context.Context, in *Sentences, opts ...grpc.CallOption) (*Sentences, error) {
	out := new(Sentences)
	err := c.cc.Invoke(ctx, "/sentences.sentences/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) Get(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) GetResponse(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/GetResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) Add(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) AddSentenceGroup(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/AddSentenceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) DeleteSentenceGroup(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/DeleteSentenceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) Moderate(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/Moderate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) Used(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/Used", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) VoteGood(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/VoteGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sentencesClient) VoteBad(ctx context.Context, in *SentenceGroup, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/sentences.sentences/VoteBad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SentencesServer is the server API for Sentences service.
// All implementations must embed UnimplementedSentencesServer
// for forward compatibility
type SentencesServer interface {
	GetAll(context.Context, *Sentences) (*Sentences, error)
	Get(context.Context, *Sentence) (*Sentence, error)
	GetResponse(context.Context, *Sentence) (*Sentence, error)
	Add(context.Context, *Sentence) (*Sentence, error)
	AddSentenceGroup(context.Context, *SentenceGroup) (*Sentence, error)
	DeleteSentenceGroup(context.Context, *SentenceGroup) (*Sentence, error)
	Moderate(context.Context, *SentenceGroup) (*Sentence, error)
	Used(context.Context, *SentenceGroup) (*Sentence, error)
	VoteGood(context.Context, *SentenceGroup) (*Sentence, error)
	VoteBad(context.Context, *SentenceGroup) (*Sentence, error)
	mustEmbedUnimplementedSentencesServer()
}

// UnimplementedSentencesServer must be embedded to have forward compatible implementations.
type UnimplementedSentencesServer struct {
}

func (UnimplementedSentencesServer) GetAll(context.Context, *Sentences) (*Sentences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSentencesServer) Get(context.Context, *Sentence) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSentencesServer) GetResponse(context.Context, *Sentence) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResponse not implemented")
}
func (UnimplementedSentencesServer) Add(context.Context, *Sentence) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedSentencesServer) AddSentenceGroup(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSentenceGroup not implemented")
}
func (UnimplementedSentencesServer) DeleteSentenceGroup(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSentenceGroup not implemented")
}
func (UnimplementedSentencesServer) Moderate(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Moderate not implemented")
}
func (UnimplementedSentencesServer) Used(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Used not implemented")
}
func (UnimplementedSentencesServer) VoteGood(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteGood not implemented")
}
func (UnimplementedSentencesServer) VoteBad(context.Context, *SentenceGroup) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteBad not implemented")
}
func (UnimplementedSentencesServer) mustEmbedUnimplementedSentencesServer() {}

// UnsafeSentencesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentencesServer will
// result in compilation errors.
type UnsafeSentencesServer interface {
	mustEmbedUnimplementedSentencesServer()
}

func RegisterSentencesServer(s grpc.ServiceRegistrar, srv SentencesServer) {
	s.RegisterService(&Sentences_ServiceDesc, srv)
}

func _Sentences_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentences)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).GetAll(ctx, req.(*Sentences))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).Get(ctx, req.(*Sentence))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/GetResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).GetResponse(ctx, req.(*Sentence))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).Add(ctx, req.(*Sentence))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_AddSentenceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).AddSentenceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/AddSentenceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).AddSentenceGroup(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_DeleteSentenceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).DeleteSentenceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/DeleteSentenceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).DeleteSentenceGroup(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_Moderate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).Moderate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/Moderate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).Moderate(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_Used_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).Used(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/Used",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).Used(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_VoteGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).VoteGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/VoteGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).VoteGood(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sentences_VoteBad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SentencesServer).VoteBad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentences.sentences/VoteBad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SentencesServer).VoteBad(ctx, req.(*SentenceGroup))
	}
	return interceptor(ctx, in, info, handler)
}

// Sentences_ServiceDesc is the grpc.ServiceDesc for Sentences service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sentences_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sentences.sentences",
	HandlerType: (*SentencesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Sentences_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Sentences_Get_Handler,
		},
		{
			MethodName: "GetResponse",
			Handler:    _Sentences_GetResponse_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Sentences_Add_Handler,
		},
		{
			MethodName: "AddSentenceGroup",
			Handler:    _Sentences_AddSentenceGroup_Handler,
		},
		{
			MethodName: "DeleteSentenceGroup",
			Handler:    _Sentences_DeleteSentenceGroup_Handler,
		},
		{
			MethodName: "Moderate",
			Handler:    _Sentences_Moderate_Handler,
		},
		{
			MethodName: "Used",
			Handler:    _Sentences_Used_Handler,
		},
		{
			MethodName: "VoteGood",
			Handler:    _Sentences_VoteGood_Handler,
		},
		{
			MethodName: "VoteBad",
			Handler:    _Sentences_VoteBad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentences.proto",
}