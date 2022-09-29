package handler

import (
	"context"
	"fmt"

	pb "cyber-srv/gateway/proto/gateway"
	service_user "cyber-srv/gateway/proto/service-user"
)

type Gateway struct {
}

func (g *Gateway) Ping(ctx context.Context, in *pb.PingReq, res *pb.PingRes) error {
	fmt.Println(in.User)
	res.User = &service_user.User{
		Id:    "x",
		Name:  "x",
		Email: "x",
		Phone: "x",
	}
	return nil
}
