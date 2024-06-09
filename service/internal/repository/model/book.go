package model

type Book struct {
	ID            int32   `gorm:"primary_key"`
	Title         string  `gorm:"type:text;not null"`
	YearPublished int     `gorm:"type:smallint;not null"`
	Rating        float64 `gorm:"type:numeric(3,2);not null"`
	Pages         int     `gorm:"type:smallint;not null"`
	GenreID       int32   `gorm:"type:integer;not null"`
	AuthorID      int32   `gorm:"type:integer;not null"`
	Genre         Genre   `gorm:"foreignKey:GenreID"`
	Author        Author  `gorm:"foreignKey:AuthorID"`
}
