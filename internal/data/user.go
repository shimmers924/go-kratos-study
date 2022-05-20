package data

import (
	"context"
	"kratos-realworld/internal/biz"

	// "github.com/go-kratos/kratos/v2/internal/context"
	"github.com/go-kratos/kratos/v2/log"
)

// 作用为实现biz层userrepo内所有接口
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (r *userRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
// 	return g, nil
// }

// func (r *userRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
// 	return g, nil
// }

// func (r *userRepo) FindByID(context.Context, int64) (*biz.User, error) {
// 	return nil, nil
// }

// func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
// 	return nil, nil
// }

// func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
// 	return nil, nil
// }
func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	// ur.data.db.Create()  //操作数据库
	return nil
}
