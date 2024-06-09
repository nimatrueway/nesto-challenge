package model

import "database/sql"

type Size struct {
	ID       int32         `gorm:"primary_key"`
	Title    string        `gorm:"type:text;not null"`
	MinPages sql.NullInt16 `gorm:"type:smallint"`
	MaxPages sql.NullInt16 `gorm:"type:smallint"`
}
