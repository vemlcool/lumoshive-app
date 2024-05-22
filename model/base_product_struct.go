package model

import (
	"time"
)

type BaseProductStruct struct {
	ID            uint      `json:"id" gorm:"primary_key;type:notnull"`
	LastModified  time.Time `gorm:"autoUpdateTime:milli" json:"last_modified"`
	LastAddStock  time.Time `gorm:"autoCreateTime:milli" json:"last_add_stock"`
	LastTakeStock time.Time `gorm:"autoCreateTime:milli" json:"last_take_stock"`
}
