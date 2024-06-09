package model

import (
	"database/sql"
)

type Era struct {
	ID      int32         `gorm:"primary_key" json:"id"`
	Title   string        `gorm:"type:text;not null" json:"title"`
	MinYear sql.NullInt16 `gorm:"type:smallint" json:"minYear"`
	MaxYear sql.NullInt16 `gorm:"type:smallint" json:"maxYear"`
}

type Size struct {
	ID       int32         `gorm:"primary_key" json:"id"`
	Title    string        `gorm:"type:text;not null" json:"title"`
	MinPages sql.NullInt16 `gorm:"type:smallint" json:"minPages"`
	MaxPages sql.NullInt16 `gorm:"type:smallint" json:"maxPages"`
}

type Genre struct {
	ID    int32  `gorm:"primary_key" json:"id"`
	Title string `gorm:"type:text;not null" json:"title"`
}

type Author struct {
	ID        int32  `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"type:text;not null" json:"firstName"`
	LastName  string `gorm:"type:text;not null" json:"lastName"`
}

type Book struct {
	ID            int32   `gorm:"primary_key" json:"id"`
	Title         string  `gorm:"type:text;not null" json:"title"`
	YearPublished int     `gorm:"type:smallint;not null" json:"yearPublished"`
	Rating        float64 `gorm:"type:numeric(3,2);not null" json:"rating"`
	Pages         int     `gorm:"type:smallint;not null" json:"pages"`
	GenreID       int32   `gorm:"type:integer;not null" json:"-"`
	AuthorID      int32   `gorm:"type:integer;not null" json:"-"`
	Genre         Genre   `gorm:"foreignKey:GenreID" json:"genre"`
	Author        Author  `gorm:"foreignKey:AuthorID" json:"author"`
}
