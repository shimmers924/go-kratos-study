package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	return &v1.LoginReply{
		User: &v1.LoginReply_User{
			Username: "Jason",
		},
	}, nil
}
