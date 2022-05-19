package service

import (
	"context"
	v1 "kratos-realworld/api/user/service/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	return &v1.LoginReply{
		User: &v1.LoginReply_User{
			Username: "Jason",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.RegisterReply, err error) {
	return &v1.RegisterReply{}, nil
}
