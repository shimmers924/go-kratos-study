package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
// func (s *RealWorldService) SayHello(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
// 	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &v1.LoginReply{Message: "Hello " + g.Hello}, nil
// }
