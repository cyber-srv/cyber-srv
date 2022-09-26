package main

import (
	"cyber-srv/gateway/handler"
	pb "cyber-srv/gateway/proto/gateway"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
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

	// Register handler
	pb.RegisterGatewayHandler(srv.Server(), new(handler.Gateway))
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
