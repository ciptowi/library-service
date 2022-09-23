package book

import (
	"context"
	"library-sevice/internal/dto"
	"library-sevice/internal/factory"
	"library-sevice/internal/models"
	"library-sevice/internal/repository"
)

type Service interface {
	Create(ctx context.Context, payload dto.CreateBookRequest) (string, error)
	Find(ctx context.Context, payload *dto.SearchGetRequest) ([]models.Book, dto.PaginationInfo, error)
	FindByID(ctx context.Context, param *dto.ByIDRequest) (models.Book, error)
	UpdateByID(ctx context.Context, ID uint, payload dto.UpdateBookRequest) (string, error)
	DeleteByID(ctx context.Context, param *dto.ByIDRequest) (string, error)
}
type service struct {
	BookRepository repository.Book
}

func NewService(f *factory.Factory) *service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) Create(ctx context.Context, payload dto.CreateBookRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Title != nil {
		data["title"] = payload.Title
	}
	if payload.Writer != nil {
		data["writer"] = payload.Writer
	}
	if payload.Isbn != nil {
		data["isbn"] = payload.Isbn
	}

	err := s.BookRepository.Create(ctx, data)
	if err != nil {
		return "Create Book Failed", err
	}
	return "Create Book Success", nil
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) ([]models.Book, dto.PaginationInfo, error) {
	books, info, _ := s.BookRepository.FindAll(ctx, payload, &payload.Pagination)
	return books, *info, nil
}

func (s *service) FindByID(ctx context.Context, param *dto.ByIDRequest) (models.Book, error) {
	data, _ := s.BookRepository.FindByID(ctx, param.ID)
	return data, nil
}

func (s *service) UpdateByID(ctx context.Context, ID uint, payload dto.UpdateBookRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Title != nil {
		data["title"] = payload.Title
	}
	if payload.Writer != nil {
		data["writer"] = payload.Writer
	}
	if payload.Isbn != nil {
		data["isbn"] = payload.Isbn
	}
	err := s.BookRepository.Update(ctx, ID, data)
	if err != nil {
		return "Update Book Failed", err
	}
	return "Update Book Success", nil
}

func (s *service) DeleteByID(ctx context.Context, param *dto.ByIDRequest) (string, error) {
	err := s.BookRepository.Delete(ctx, param.ID)
	if err != nil {
		return "Delete Book Failed", err
	}
	return "Delete Book Success", nil
}
