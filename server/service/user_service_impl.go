package service

import (
	"context"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b"
)
type UserServiceImpl struct {
	Repository repository.UserRepository
	timeout time.Duration
}

func NewService(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repository: repository,
		timeout: time.Duration(2) * time.Second,
	}
}

func (s *UserServiceImpl) CreateUser(c context.Context, req *web.UserCreateRequest) (*web.UserCreateResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// HASH PASSWORD
	hashedPassoword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &domain.User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassoword,
	}
	
	r, err := s.Repository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &web.UserCreateResponse{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}

	return res, nil
}

type JWTClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *UserServiceImpl) 	Login(c context.Context, req *web.UserLoginRequest) (*web.UserLoginJWT, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.FindByEmail(ctx, req.Email)
	if err != nil {
		return &web.UserLoginJWT{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &web.UserLoginJWT{}, err	
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return &web.UserLoginJWT{}, err
	}

	return &web.UserLoginJWT{
		AccessToken: ss,
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
	}, nil
}