package model

type Genre struct {
	ID    int32  `gorm:"primary_key"`
	Title string `gorm:"type:text;not null"`
}
