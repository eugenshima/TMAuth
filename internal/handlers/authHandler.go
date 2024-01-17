package handlers

import (
	"context"

	"github.com/eugenshima/TMAuth/internal/model"
	proto "github.com/eugenshima/TMAuth/proto"
)

type AuthHandler struct {
	srv AuthServiceInterface
	proto.UnimplementedAuthServer
}

func NewAuthHandler(srv AuthServiceInterface) *AuthHandler {
	return &AuthHandler{srv: srv}
}

type AuthServiceInterface interface {
	GetProfileByLogin(ctx context.Context) (*model.FullAuthModel, error)
}

func (hnd *AuthHandler) GetProfileByLogin(ctx context.Context)
