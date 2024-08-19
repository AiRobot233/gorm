package model

import (
	"errors"
	"gin/utils"
	"gorm.io/gorm"
)

type User struct {
	Id        int              `gorm:"column:id;primaryKey" json:"id,omitempty"`      //是否可空:NO
	Name      string           `gorm:"column:name" json:"name,omitempty"`             //是否可空:NO
	Phone     string           `gorm:"column:phone" json:"phone,omitempty"`           //是否可空:NO
	Password  string           `gorm:"column:password" json:"-"`                      //是否可空:NO
	Salt      string           `gorm:"column:salt" json:"-"`                          //是否可空:NO
	CreatedAt *utils.LocalTime `gorm:"column:created_at" json:"created_at,omitempty"` //是否可空:NO
	UpdatedAt *utils.LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"` //是否可空:NO
	DeletedAt gorm.DeletedAt   `gorm:"column:deleted_at" json:"-"`                    //是否可空:NO
	Status    int              `gorm:"column:status" json:"status,omitempty"`         //是否可空:NO
	RoleId    int              `gorm:"column:role_id" json:"role_id,omitempty"`       //是否可空:NO
	UnitId    int              `gorm:"column:unit_id" json:"unit_id"`                 //是否可空:NO
	Role      *Role            `json:"role,omitempty"`
	Unit      *Unit            `json:"unit,omitempty"`
}

func (*User) TableName() string {
	return "user"
}

// BeforeSave 新增修改事件
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	user := User{}
	result := tx.Model(u).Where("id != ? AND `name` = ?", u.Id, u.Name).First(&user)
	if result.RowsAffected > 0 {
		return errors.New("名称已存在")
	}
	return
}

// UserSearch 搜索
type UserSearch struct {
	Keyword string `form:"keyword"`
	RoleId  int    `form:"role_id"`
}

// UserSearchFunc 搜索函数
func UserSearchFunc(params UserSearch) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params.Keyword != "" {
			db.Where("(`phone` LIKE ? OR `name` LIKE ?)", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
		}
		if params.RoleId > 0 {
			db.Where("role_id = ?", params.RoleId)
		}
		return db
	}
}
