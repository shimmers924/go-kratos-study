package service

import (
	v1 "kratos-realworld/api/user/service/v1"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc  *biz.UserUsecase
	log *log.Logger
}

func NewRealworldService(uc *biz.UserUsecase, logger log.Logger) *RealWorldService {
	// return &RealworldService{uc: uc, log: log.NewHelper(logger)}
	return nil
}
