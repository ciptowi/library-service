package repository

import (
	"context"
	"gorm.io/gorm"
	"library-sevice/internal/dto"
	"library-sevice/internal/models"
	"strings"
)

type Book interface {
	Create(ctx context.Context, data map[string]interface{}) error
	FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]models.Book, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (models.Book, error)
	FindByTitle(ctx context.Context, title *string) (*models.Book, error)
	FindByWriter(ctx context.Context, writer *string) (*models.Book, error)
	FindByISBN(ctx context.Context, isbn *string) (*models.Book, error)
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
	Delete(ctx context.Context, ID uint) error
}

type book struct {
	Db *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{
		db,
	}
}

func (b *book) Create(ctx context.Context, data map[string]interface{}) error {
	return b.Db.WithContext(ctx).Model(&models.Book{}).Create(&data).Error
}

func (b *book) FindAll(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]models.Book, *dto.PaginationInfo, error) {
	var books []models.Book
	var count int64

	query := b.Db.WithContext(ctx).Model(&models.Book{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&books).Error

	return books, dto.CheckInfoPagination(paginate, count), err
}

func (b *book) FindByID(ctx context.Context, ID uint) (models.Book, error) {

	var data models.Book
	err := b.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}

func (b *book) FindByTitle(ctx context.Context, title *string) (*models.Book, error) {
	conn := b.Db.WithContext(ctx)

	var data models.Book
	err := conn.Where("title = ?", title).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *book) FindByWriter(ctx context.Context, writer *string) (*models.Book, error) {
	conn := b.Db.WithContext(ctx)

	var data models.Book
	err := conn.Where("writer = ?", writer).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *book) FindByISBN(ctx context.Context, isbn *string) (*models.Book, error) {
	conn := b.Db.WithContext(ctx)

	var data models.Book
	err := conn.Where("isbn = ?", isbn).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (b *book) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := b.Db.WithContext(ctx).Where("id = ?", ID).Model(&models.Book{}).Updates(&data).Error
	return err
}

func (b *book) Delete(ctx context.Context, ID uint) error {

	err := b.Db.WithContext(ctx).Where("id = ?", ID).Delete(&models.Book{}).Error
	return err
}
