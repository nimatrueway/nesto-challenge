package dto

import "readcommend/internal/repository/model"

type Genre struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func GenreFromModel(dbGenre model.Genre) Genre {
	return Genre{
		ID:    dbGenre.ID,
		Title: dbGenre.Title,
	}
}
