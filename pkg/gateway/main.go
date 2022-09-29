package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"cyber-srv/gateway/handler"
	pb "cyber-srv/gateway/proto/gateway"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	// self grpc service
	pb.RegisterGatewayHandler(srv.Server(), new(handler.Gateway))

	// imported grpc service
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoints := map[string]string{
		"service-user": "localhost:8080",
	}

	for name, addr := range endpoints {
		fmt.Println("registering: ", name, addr)
		pb.RegisterGatewayGWFromEndpoint(context.TODO(), gwmux, addr, opts)
	}

	srv.Server().Handle(srv.Server().NewHandler(gwmux))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func getCred() (*tls.Certificate, credentials.TransportCredentials) {
	var err error
	key, err := ioutil.ReadFile("./certs/_wildcard.jinyi.test-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := ioutil.ReadFile("./certs/_wildcard.jinyi.test.pem")
	if err != nil {
		log.Fatal(err)
	}

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	// cert pool
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(cert)
	if !ok {
		log.Fatal("failed to parse root certificate")
	}

	tlsCred := credentials.NewTLS(&tls.Config{
		ServerName: "jinyi.test",
		RootCAs:    certPool,
	})

	return &pair, tlsCred
}
