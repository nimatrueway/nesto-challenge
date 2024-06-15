package repository

import (
	"context"
	"regexp"
	"strings"

	"readcommend/internal/repository/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooks(ctx context.Context, authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]model.Book, error)
	GetAuthors(ctx context.Context, search string, limit int) ([]model.Author, error)
	GetGenres(ctx context.Context) ([]model.Genre, error)
	GetSizes(ctx context.Context) ([]model.Size, error)
	GetEras(ctx context.Context) ([]model.Era, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (r *BookRepositoryImpl) GetBooks(ctx context.Context, authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]model.Book, error) {
	var books []model.Book

	query := r.db.WithContext(ctx)
	query = query.InnerJoins("Genre").InnerJoins("Author")

	if authors != nil {
		query = query.Where("author_id IN (?)", authors)
	}
	if genres != nil {
		query = query.Where("genre_id IN (?)", genres)
	}
	if minPages != 0 {
		query = query.Where("pages >= ?", minPages)
	}
	if maxPages != 0 {
		query = query.Where("pages <= ?", maxPages)
	}
	if minYear != 0 {
		query = query.Where("year_published >= ?", minYear)
	}
	if maxYear != 0 {
		query = query.Where("year_published <= ?", maxYear)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Order("book.rating desc, book.id asc").Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) GetAuthors(ctx context.Context, search string, limit int) ([]model.Author, error) {
	var authors []model.Author

	query := r.db.WithContext(ctx)

	if search != "" {
		allWords := regexp.MustCompile(`\S+`).FindAllString(strings.ToLower(search), -1)
		for i, word := range allWords {
			allWords[i] = word + ":*"
		}
		searchTerm := strings.Join(allWords, " & ")
		query = query.Where("to_tsvector('simple', first_name || ' ' || last_name) @@ to_tsquery(?)", searchTerm)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&authors).Error

	return authors, err
}

func (r *BookRepositoryImpl) GetGenres(ctx context.Context) ([]model.Genre, error) {
	var genres []model.Genre
	query := r.db.WithContext(ctx)
	err := query.Find(&genres).Error
	return genres, err
}

func (r *BookRepositoryImpl) GetSizes(ctx context.Context) ([]model.Size, error) {
	var sizes []model.Size
	query := r.db.WithContext(ctx)
	err := query.Find(&sizes).Error
	return sizes, err
}

func (r *BookRepositoryImpl) GetEras(ctx context.Context) ([]model.Era, error) {
	var eras []model.Era
	query := r.db.WithContext(ctx)
	err := query.Find(&eras).Error
	return eras, err
}
