package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseUserStruct struct {
	ID uint `json:"id" gorm:"primary_key;type:notnull"`
	// Birthday  time.Time      `json:"birthday" gorm:"date"`
	// the insert format for postman "2001-05-18T00:00:00+07:00"
	// or use this "2024-05-22T00:00:00Z"
	Birthday  *time.Time     `json:",omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
