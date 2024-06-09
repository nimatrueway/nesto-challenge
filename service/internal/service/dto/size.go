package dto

import "readcommend/internal/repository/model"

type Size struct {
	ID       int32  `json:"id"`
	Title    string `json:"title"`
	MinPages *int16 `json:"minPages,omitempty"`
	MaxPages *int16 `json:"maxPages,omitempty"`
}

func SizeFromModel(dbSize model.Size) Size {
	return Size{
		ID:       dbSize.ID,
		Title:    dbSize.Title,
		MinPages: int16FromModel(dbSize.MinPages),
		MaxPages: int16FromModel(dbSize.MaxPages),
	}
}
