package dto

import (
	"readcommend/internal/repository/model"
)

type Book struct {
	ID            int32   `json:"id"`
	Title         string  `json:"title"`
	YearPublished int     `json:"yearPublished"`
	Rating        float64 `json:"rating"`
	Pages         int     `json:"pages"`
	Genre         Genre   `json:"genre"`
	Author        Author  `json:"author"`
}

func BookFromModel(dbBook model.Book) Book {
	return Book{
		ID:            dbBook.ID,
		Title:         dbBook.Title,
		YearPublished: dbBook.YearPublished,
		Rating:        dbBook.Rating,
		Pages:         dbBook.Pages,
		Genre:         GenreFromModel(dbBook.Genre),
		Author:        AuthorFromModel(dbBook.Author),
	}
}
