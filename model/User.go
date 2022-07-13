package model

import (
	"gin/utils"
	"gorm.io/gorm"
)

//
type User struct {
	Id        int             `gorm:"column:id" json:"id"`                 //是否可空:NO
	Name      string          `gorm:"column:name" json:"name"`             //是否可空:YES
	Age       int             `gorm:"column:age" json:"age"`               //是否可空:YES
	CreatedAt utils.LocalTime `gorm:"column:created_at" json:"created_at"` //是否可空:YES
	UpdatedAt utils.LocalTime `gorm:"column:updated_at" json:"updated_at"` //是否可空:YES
	Books     []Book          `gorm:"foreignKey:UserId;references:Id" json:"books"`
}

func (*User) TableName() string {
	return "user"
}

func (*User) Db() *gorm.DB {
	return GetDb()
}
