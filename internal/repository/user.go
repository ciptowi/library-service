package repository

import (
	"context"
	"gorm.io/gorm"
	"library-sevice/internal/dto"
	"library-sevice/internal/models"
	"strings"
)

type User interface {
	Create(ctx context.Context, data map[string]interface{}) error
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]models.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (models.User, error)
	FindByEmail(ctx context.Context, email *string) (*models.User, error)
	FindByName(ctx context.Context, name *string) (*models.User, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (u *user) Create(ctx context.Context, data map[string]interface{}) error {
	return u.Db.WithContext(ctx).Model(&models.User{}).Create(&data).Error
}

func (u *user) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]models.User, *dto.PaginationInfo, error) {
	var users []models.User
	var count int64

	query := u.Db.WithContext(ctx).Model(&models.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(paginate, count), err
}

func (u *user) FindByID(ctx context.Context, ID uint) (models.User, error) {

	var data models.User
	err := u.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (u *user) FindByEmail(ctx context.Context, email *string) (*models.User, error) {
	conn := u.Db.WithContext(ctx)

	var data models.User
	err := conn.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *user) FindByName(ctx context.Context, name *string) (*models.User, error) {
	conn := u.Db.WithContext(ctx)

	var data models.User
	err := conn.Where("name = ?", name).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *user) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := u.Db.WithContext(ctx).Where("id = ?", ID).Model(&models.User{}).Updates(&data).Error
	return err
}

func (u *user) Delete(ctx context.Context, ID uint) error {

	err := u.Db.WithContext(ctx).Where("id = ?", ID).Delete(&models.User{}).Error
	return err
}
