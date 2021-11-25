// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/micro-server.proto

package microserver

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for TransactionPersistence service

func NewTransactionPersistenceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TransactionPersistence service

type TransactionPersistenceService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (TransactionPersistence_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (TransactionPersistence_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (TransactionPersistence_BidiStreamService, error)
}

type transactionPersistenceService struct {
	c    client.Client
	name string
}

func NewTransactionPersistenceService(name string, c client.Client) TransactionPersistenceService {
	return &transactionPersistenceService{
		c:    c,
		name: name,
	}
}

func (c *transactionPersistenceService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "TransactionPersistence.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionPersistenceService) ClientStream(ctx context.Context, opts ...client.CallOption) (TransactionPersistence_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "TransactionPersistence.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &transactionPersistenceServiceClientStream{stream}, nil
}

type TransactionPersistence_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type transactionPersistenceServiceClientStream struct {
	stream client.Stream
}

func (x *transactionPersistenceServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *transactionPersistenceService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (TransactionPersistence_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "TransactionPersistence.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &transactionPersistenceServiceServerStream{stream}, nil
}

type TransactionPersistence_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type transactionPersistenceServiceServerStream struct {
	stream client.Stream
}

func (x *transactionPersistenceServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transactionPersistenceService) BidiStream(ctx context.Context, opts ...client.CallOption) (TransactionPersistence_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "TransactionPersistence.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &transactionPersistenceServiceBidiStream{stream}, nil
}

type TransactionPersistence_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type transactionPersistenceServiceBidiStream struct {
	stream client.Stream
}

func (x *transactionPersistenceServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TransactionPersistence service

type TransactionPersistenceHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, TransactionPersistence_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, TransactionPersistence_ServerStreamStream) error
	BidiStream(context.Context, TransactionPersistence_BidiStreamStream) error
}

func RegisterTransactionPersistenceHandler(s server.Server, hdlr TransactionPersistenceHandler, opts ...server.HandlerOption) error {
	type transactionPersistence interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type TransactionPersistence struct {
		transactionPersistence
	}
	h := &transactionPersistenceHandler{hdlr}
	return s.Handle(s.NewHandler(&TransactionPersistence{h}, opts...))
}

type transactionPersistenceHandler struct {
	TransactionPersistenceHandler
}

func (h *transactionPersistenceHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.TransactionPersistenceHandler.Call(ctx, in, out)
}

func (h *transactionPersistenceHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.TransactionPersistenceHandler.ClientStream(ctx, &transactionPersistenceClientStreamStream{stream})
}

type TransactionPersistence_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type transactionPersistenceClientStreamStream struct {
	stream server.Stream
}

func (x *transactionPersistenceClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *transactionPersistenceHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TransactionPersistenceHandler.ServerStream(ctx, m, &transactionPersistenceServerStreamStream{stream})
}

type TransactionPersistence_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type transactionPersistenceServerStreamStream struct {
	stream server.Stream
}

func (x *transactionPersistenceServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *transactionPersistenceHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.TransactionPersistenceHandler.BidiStream(ctx, &transactionPersistenceBidiStreamStream{stream})
}

type TransactionPersistence_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type transactionPersistenceBidiStreamStream struct {
	stream server.Stream
}

func (x *transactionPersistenceBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *transactionPersistenceBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transactionPersistenceBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transactionPersistenceBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *transactionPersistenceBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}