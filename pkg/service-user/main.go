package main

import (
	"service-user/handler"
	pb "service-user/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "service-user"
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
	pb.RegisterServiceUserHandler(srv.Server(), new(handler.ServiceUser))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
