package model

type Author struct {
	ID        int32  `gorm:"primary_key"`
	FirstName string `gorm:"type:text;not null"`
	LastName  string `gorm:"type:text;not null"`
}
