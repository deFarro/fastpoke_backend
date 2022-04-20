package artist

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	ID    string `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
