package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// var (
// 	// ErrUserNotFound is user not found.
// 	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
// )

// 定义模型
type User struct {
	username string
}

// 定义data层userrepo内需要实现的业务逻辑接口
type UserRepo interface {
	//Save(context.Context, *Greeter) (*Greeter, error)
	// Update(context.Context, *Greeter) (*Greeter, error)
	// FindByID(context.Context, int64) (*Greeter, error)
	// ListByHello(context.Context, string) ([]*Greeter, error)
	// ListAll(context.Context) ([]*Greeter, error)

	CreateUser(ctx context.Context, user *User) error
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// 为结构体userusecase赋值
func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateCreateUser(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
	}
	return nil
}
