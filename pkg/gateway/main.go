package main

import (
	"context"
	"net/http"

	pb "cyber-srv/gateway/proto/gateway"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	service = "gateway"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	/*
		grpc.WithInsecure is deprecated: use WithTransportCredentials and insecure.NewCredentials()
	*/

	endpoints := ":8080"
	err := pb.RegisterGatewayHandlerFromEndpoint(ctx, mux, endpoints, opts)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8081", mux)
	// Register handler
	// pb.RegisterGatewayHandler(srv.Server(), new(handler.Gateway))
	// Run service
	// if err := srv.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
