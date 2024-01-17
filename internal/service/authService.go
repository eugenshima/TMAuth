package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/TMAuth/internal/model"
)

// AuthService represents authentication
type AuthService struct {
	rps AuthRepositoryInterface
}

// NewAuthService creates a new AuthService
func NewAuthService(rps AuthRepositoryInterface) *AuthService {
	return &AuthService{rps: rps}
}

// AuthRepositoryInterface represents a repository interface
type AuthRepositoryInterface interface {
	GetProfileByLogin(context.Context, string) (*model.FullAuthModel, error)
}

// GetProfileByLogin function returns the profile associated with the given login
func (s *AuthService) GetProfileByLogin(ctx context.Context) (*model.FullAuthModel, error) {
	str := "jija"
	auth, err := s.rps.GetProfileByLogin(ctx, str)
	if err != nil {
		return nil, fmt.Errorf("GetProfileByLogin: %w", err)
	}
	fmt.Println(auth, auth.Login)
	return auth, nil
}
