package service

import (
	"context"
	"server/model/web"
)

type UserService interface {
	CreateUser(c context.Context, req *web.UserCreateRequest) (*web.UserCreateResponse, error)
	Login(c context.Context, req *web.UserLoginRequest) (*web.UserLoginJWT, error)
}