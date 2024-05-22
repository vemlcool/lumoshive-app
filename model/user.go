package model

type User struct {
	Name    string `json:"name" gorm:"type:varchar(256)"`
	Phone   string `json:"phone" gorm:"type:varchar(20)"`
	Address string `json:"address" gorm:"type:varchar(256)"`
	Email   string `json:"email" gorm:"type:varchar(35)"`
	BaseUserStruct
}
