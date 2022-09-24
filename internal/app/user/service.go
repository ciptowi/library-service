package user

import (
	"context"
	"library-sevice/internal/dto"
	"library-sevice/internal/factory"
	"library-sevice/internal/models"
	"library-sevice/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(ctx context.Context, payload dto.CreateUserRequest) (string, error)
	Find(ctx context.Context, payload *dto.SearchGetRequest) ([]models.User, dto.PaginationInfo, error)
	FindByID(ctx context.Context, param *dto.ByIDRequest) (models.User, error)
	UpdateByID(ctx context.Context, ID uint, payload dto.UpdateUserRequest) (string, error)
	DeleteByID(ctx context.Context, param *dto.ByIDRequest) (string, error)
}
type service struct {
	UserRepository repository.User
}

func NewService(f *factory.Factory) *service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Create(ctx context.Context, payload dto.CreateUserRequest) (string, error) {
	var data = make(map[string]interface{})

	//if payload.Name != nil {
	data["name"] = payload.Name
	//}
	//if payload.Email != nil {
	data["email"] = payload.Email
	//}
	//if payload.Password != nil {
	// data["password"] = payload.Password
	//}
	_, e := s.UserRepository.FindByEmail(ctx, &payload.Email)
	if e == nil {
		return "Email already exists!", e
	}
	byteFile, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	data["password"] = string(byteFile)
	err := s.UserRepository.Create(ctx, data)
	if err != nil {
		return "Create User Failed", err
	}
	return "Create User Success", nil
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) ([]models.User, dto.PaginationInfo, error) {

	users, info, _ := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)

	var datas []models.User
	for _, user := range users {
		datas = append(datas, models.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return datas, *info, nil
}

func (s *service) FindByID(ctx context.Context, param *dto.ByIDRequest) (models.User, error) {
	user := new(models.User)

	data, _ := s.UserRepository.FindByID(ctx, param.ID)
	user.ID = data.ID
	user.Name = data.Name
	user.Email = data.Email

	return *user, nil
}

func (s *service) UpdateByID(ctx context.Context, ID uint, payload dto.UpdateUserRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Name != nil {
		data["name"] = payload.Name
	}
	if payload.Email != nil {
		data["email"] = payload.Email
	}
	if payload.Password != nil {
		data["password"] = payload.Password
	}
	err := s.UserRepository.Update(ctx, ID, data)
	if err != nil {
		return "Update User Failed", err
	}
	return "Update User Success", nil
}

func (s *service) DeleteByID(ctx context.Context, param *dto.ByIDRequest) (string, error) {
	err := s.UserRepository.Delete(ctx, param.ID)
	if err != nil {
		return "Delete User Failed", err
	}
	return "Delete User Success", nil
}
