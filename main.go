package main

import (
	"io"
	"micro-server/handler"
	pb "micro-server/proto"

	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	"go-micro.dev/v4/cmd/micro/debug/trace/jaeger"
	// To explicitly declare gRPC transport
	//"github.com/asim/go-micro/plugins/server/grpc/v4"
	// micro.Server(grpc.NewServer()),
)

var (
	service = "micro-server"
)

func main() {
	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(closer io.Closer) {
		_ = closer.Close()
	}(closer)

	// Create service
	// micro.Version(version) - set version
	// micro.Name(service) - set service name
	// micro.Server(grpc.NewServer()), // must come before any other options
	srv := micro.NewService(
		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
	)
	srv.Init()

	// Register handler
	_ = pb.RegisterTransactionPersistenceHandler(srv.Server(), new(handler.TransactionPersistence))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
