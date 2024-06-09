package dto

import "readcommend/internal/repository/model"

type Era struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	MinYear *int16 `json:"minYear,omitempty"`
	MaxYear *int16 `json:"maxYear,omitempty"`
}

func EraFromModel(dbEra model.Era) Era {
	return Era{
		ID:      dbEra.ID,
		Title:   dbEra.Title,
		MinYear: int16FromModel(dbEra.MinYear),
		MaxYear: int16FromModel(dbEra.MaxYear),
	}
}
