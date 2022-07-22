package model

import (
	"errors"
	"gin/utils"
	"gorm.io/gorm"
)

//
type User struct {
	Id        int             `gorm:"column:id;primaryKey" json:"id"`      //是否可空:NO
	Name      string          `gorm:"column:name" json:"name"`             //是否可空:NO
	Phone     string          `gorm:"column:phone" json:"phone"`           //是否可空:NO
	Password  string          `gorm:"column:password" json:"-"`            //是否可空:NO
	Salt      string          `gorm:"column:salt" json:"-"`                //是否可空:NO
	CreatedAt utils.LocalTime `gorm:"column:created_at" json:"created_at"` //是否可空:NO
	UpdatedAt utils.LocalTime `gorm:"column:updated_at" json:"updated_at"` //是否可空:NO
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"-"`          //是否可空:NO
	Status    int             `gorm:"column:status" json:"status"`         //是否可空:NO
	RoleId    int             `gorm:"column:role_id" json:"role_id"`       //是否可空:NO
}

func (*User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	user := User{}
	result := tx.Model(u).Where("id != ? AND phone = ?", u.Id, u.Phone).First(&user)
	if result.RowsAffected > 0 {
		return errors.New("手机号已存在")
	}
	return
}
