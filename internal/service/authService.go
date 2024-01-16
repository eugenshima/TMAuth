package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/TMAuth/internal/model"
)

type ProfileServiceService struct {
	rps ProfileRepositoryInterface
}

func NewProfileServiceService(rps ProfileRepositoryInterface) *ProfileServiceService {
	return &ProfileServiceService{rps: rps}
}

type ProfileRepositoryInterface interface {
	GetProfileByLogin(context.Context, string) (*model.FullAuthModel, error)
}

func (s *ProfileServiceService) GetProfileByLogin(ctx context.Context) error {
	str := "jija"
	auth, err := s.rps.GetProfileByLogin(ctx, str)
	if err != nil {
		return fmt.Errorf("GetProfileByLogin: %w", err)
	}
	fmt.Println(auth, auth.Login)
	return nil
}
