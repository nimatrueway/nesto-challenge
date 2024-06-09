package model

import "database/sql"

type Era struct {
	ID      int32         `gorm:"primary_key"`
	Title   string        `gorm:"type:text;not null"`
	MinYear sql.NullInt16 `gorm:"type:smallint"`
	MaxYear sql.NullInt16 `gorm:"type:smallint"`
}
