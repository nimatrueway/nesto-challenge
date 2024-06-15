package service

import (
	"context"
	"readcommend/internal/repository/model"
	"readcommend/internal/service/dto"

	"github.com/pkg/errors"
)

type BookRepository interface {
	GetBooks(ctx context.Context, authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]model.Book, error)
	GetAuthors(ctx context.Context, search string, limit int) ([]model.Author, error)
	GetGenres(ctx context.Context) ([]model.Genre, error)
	GetSizes(ctx context.Context) ([]model.Size, error)
	GetEras(ctx context.Context) ([]model.Era, error)
}

type BookServiceImpl struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) *BookServiceImpl {
	return &BookServiceImpl{repo: repo}
}

func (s *BookServiceImpl) GetBooks(ctx context.Context, authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]dto.Book, error) {
	// Future work: verify that authors and genres are valid

	books, err := s.repo.GetBooks(ctx, authors, genres, minPages, maxPages, minYear, maxYear, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get books")
	}

	result := make([]dto.Book, len(books))
	for i, book := range books {
		result[i] = dto.BookFromModel(book)
	}

	return result, nil
}

func (s *BookServiceImpl) GetAuthors(ctx context.Context, search string, limit int) ([]dto.Author, error) {
	authors, err := s.repo.GetAuthors(ctx, search, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get authors")
	}

	result := make([]dto.Author, len(authors))
	for i, author := range authors {
		result[i] = dto.AuthorFromModel(author)
	}

	return result, nil
}

func (s *BookServiceImpl) GetGenres(ctx context.Context) ([]dto.Genre, error) {
	genres, err := s.repo.GetGenres(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get genres")
	}

	result := make([]dto.Genre, len(genres))
	for i, genre := range genres {
		result[i] = dto.GenreFromModel(genre)
	}

	return result, nil
}

func (s *BookServiceImpl) GetSizes(ctx context.Context) ([]dto.Size, error) {
	sizes, err := s.repo.GetSizes(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get sizes")
	}

	result := make([]dto.Size, len(sizes))
	for i, size := range sizes {
		result[i] = dto.SizeFromModel(size)
	}

	return result, nil
}

func (s *BookServiceImpl) GetEras(ctx context.Context) ([]dto.Era, error) {
	eras, err := s.repo.GetEras(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get eras")
	}

	result := make([]dto.Era, len(eras))
	for i, era := range eras {
		result[i] = dto.EraFromModel(era)
	}

	return result, nil
}
