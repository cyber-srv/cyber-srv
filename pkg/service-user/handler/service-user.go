package handler

import (
	"context"

	pb "service-user/proto"
)

type ServiceUser struct{}

func (s *ServiceUser) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest, res *pb.GetUserInfoResponse) error {
	res.User = &pb.User{
		Id:    req.Id,
		Name:  "test_user",
		Email: "test@email.com",
		Phone: "177xxxx6396",
	}

	return nil
}
