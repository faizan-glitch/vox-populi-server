package models

import (
	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	Question string
	Replies  uint
}
