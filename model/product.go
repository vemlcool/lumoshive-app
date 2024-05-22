package model

type Product struct {
	Name   string `json:"name" gorm:"type:varchar(256)"`
	Detail string `json:"detail" gorm:"type:varchar(50)"`
	Type   string `json:"type" gorm:"type:varchar(20)"`
	Stock  int    `json:"stock" gorm:"type:smallint"`
	BaseProductStruct
}
