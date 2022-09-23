package auth

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"library-sevice/internal/dto"
	"library-sevice/internal/factory"
	"library-sevice/internal/models"
	"library-sevice/internal/repository"
)

type Service interface {
	Login(ctx context.Context, payload *dto.AuthLoginRequest) (string, *models.User, error)
}

type service struct {
	Repository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Login(ctx context.Context, payload *dto.AuthLoginRequest) (string, *models.User, error) {

	data, err := s.Repository.FindByEmail(ctx, &payload.Email)
	if data == nil {
		return "Email not found!", data, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		return "Wrong password!", data, err
	}
	token, err := data.GenerateToken()
	if err != nil {
		return "Generate Failed!", data, err
	}
	return token, data, nil
}
