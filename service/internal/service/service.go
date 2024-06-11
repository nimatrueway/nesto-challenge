package service

import (
	"readcommend/internal/repository"
	"readcommend/internal/service/dto"
)

type BookService interface {
	GetBooks(authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]dto.Book, error)
	GetAuthors(limit int) ([]dto.Author, error)
	GetGenres() ([]dto.Genre, error)
	GetSizes() ([]dto.Size, error)
	GetEras() ([]dto.Era, error)
}

type BookServiceImpl struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{repo: repo}
}

func (s *BookServiceImpl) GetBooks(authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]dto.Book, error) {
	books, err := s.repo.GetBooks(authors, genres, minPages, maxPages, minYear, maxYear, limit)
	if err != nil {
		return nil, err
	}

	result := make([]dto.Book, len(books))
	for i, book := range books {
		result[i] = dto.BookFromModel(book)
	}

	return result, nil
}

func (s *BookServiceImpl) GetAuthors(limit int) ([]dto.Author, error) {
	authors, err := s.repo.GetAuthors(limit)
	if err != nil {
		return nil, err
	}

	result := make([]dto.Author, len(authors))
	for i, author := range authors {
		result[i] = dto.AuthorFromModel(author)
	}

	return result, nil
}

func (s *BookServiceImpl) GetGenres() ([]dto.Genre, error) {
	genres, err := s.repo.GetGenres()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Genre, len(genres))
	for i, genre := range genres {
		result[i] = dto.GenreFromModel(genre)
	}

	return result, nil
}

func (s *BookServiceImpl) GetSizes() ([]dto.Size, error) {
	sizes, err := s.repo.GetSizes()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Size, len(sizes))
	for i, size := range sizes {
		result[i] = dto.SizeFromModel(size)
	}

	return result, nil
}

func (s *BookServiceImpl) GetEras() ([]dto.Era, error) {
	eras, err := s.repo.GetEras()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Era, len(eras))
	for i, era := range eras {
		result[i] = dto.EraFromModel(era)
	}

	return result, nil
}
