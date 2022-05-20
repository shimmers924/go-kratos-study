package data

import (
	"context"
	"kratos-realworld/internal/biz"

	// "github.com/go-kratos/kratos/v2/internal/context"
	"github.com/go-kratos/kratos/v2/log"
)

func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	// ur.data.db.Create()  //操作数据库
	return nil
}
func (ur *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	return u, nil
}
func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, nil
}
func (ur *userRepo) ListByIds(ctx context.Context, u *biz.User) (*biz.User, error) {
	return nil, nil
}
func (ur *userRepo) ListAll(ctx context.Context, u *biz.User) (*biz.User, error) {
	return nil, nil
}

// 作用为实现biz层userrepo内所有接口
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	// return &userRepo{
	// 	data: data,
	// 	log:  log.NewHelper(logger),
	// }
	return nil
}
