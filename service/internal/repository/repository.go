package repository

import (
	"gorm.io/gorm"
	"readcommend/internal/repository/model"
)

type Repository interface {
	GetBooks(authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]model.Book, error)
	GetAuthors() ([]model.Author, error)
	GetGenres() ([]model.Genre, error)
	GetSizes() ([]model.Size, error)
	GetEras() ([]model.Era, error)
}

type PgRepository struct {
	db *gorm.DB
}

func NewPgRepository(db *gorm.DB) *PgRepository {
	return &PgRepository{
		db: db,
	}
}

func (r *PgRepository) GetBooks(authors, genres []int, minPages, maxPages, minYear, maxYear, limit int) ([]model.Book, error) {
	var books []model.Book

	query := r.db
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

	err := query.Find(&books).Error
	return books, err
}

func (r *PgRepository) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *PgRepository) GetGenres() ([]model.Genre, error) {
	var genres []model.Genre
	err := r.db.Find(&genres).Error
	return genres, err
}

func (r *PgRepository) GetSizes() ([]model.Size, error) {
	var sizes []model.Size
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *PgRepository) GetEras() ([]model.Era, error) {
	var eras []model.Era
	err := r.db.Find(&eras).Error
	return eras, err
}
