package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/TMAuth/internal/model"
	proto "github.com/eugenshima/TMAuth/proto"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	srv AuthServiceInterface
	proto.UnimplementedAuthServer
}

func NewAuthHandler(srv AuthServiceInterface) *AuthHandler {
	return &AuthHandler{srv: srv}
}

type AuthServiceInterface interface {
	GetProfileByLogin(ctx context.Context, login string) (*model.FullAuthModel, error)
}

func (hnd *AuthHandler) Login(ctx context.Context, req *proto.LoginRequest) (resp *proto.LoginResponse, err error) {
	requestLogin := &model.ProfileLogin{
		Login:    req.Auth.Login,
		Password: req.Auth.Password,
	}
	AuthResult, err := hnd.srv.GetProfileByLogin(ctx, requestLogin.Login)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Login": AuthResult.Login}).Errorf("GetProfileByLogin: %v", err)
		return nil, fmt.Errorf("GetProfileByLogin: %w", err)
	}
	//resp.Login = AuthResult.Login
	return resp, nil
}
