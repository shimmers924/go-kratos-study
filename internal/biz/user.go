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
	Email        string
	Token        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

//用户登录模型
type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

//定义接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error //接口内方法
	UpdateUser(context.Context, int64) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ListByIds(context.Context, []int) ([]*User, error)
	ListAll(context.Context) ([]*User, error)
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// 把data 层实现注入到biz层
func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

// 为UserUsecase结构体添加一个指针接收器的方法
func (uc *UserUsecase) Login(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
	}
	return nil
}
func (uc *UserUsecase) update(ctx context.Context, u *User) error {
	if err := uc.ur.UpdateUser(ctx, u); err != nil {
	}
	return nil
}
func (uc *UserUsecase) GetUserByEmail(ctx context.Context, u *User) error {
	if err := uc.ur.GetUserByEmail(ctx, u.Email); err != nil {
	}
	return nil
}
func (uc *UserUsecase) ListByIds(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
	}
	return nil
}
func (uc *UserUsecase) ListAll(ctx context.Context, u *User) error {
	if err := uc.ur.CreateUser(ctx, u); err != nil {
	}
	return nil
}
