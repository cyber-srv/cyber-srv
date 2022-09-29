package main

import (
	"context"
	service_user "cyber-srv/gateway/proto/service-user"
	"fmt"

	"go-micro.dev/v4"
)

func main() {
	srv := micro.NewService()

	srv.Init()

	cl := service_user.NewServiceUserService("service-user", srv.Client())

	res, err := cl.GetUserInfo(context.TODO(), &service_user.GetUserInfoRequest{
		Id: "123",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
