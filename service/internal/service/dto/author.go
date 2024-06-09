package dto

import "readcommend/internal/repository/model"

type Author struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func AuthorFromModel(dbAuthor model.Author) Author {
	return Author{
		ID:        dbAuthor.ID,
		FirstName: dbAuthor.FirstName,
		LastName:  dbAuthor.LastName,
	}
}
