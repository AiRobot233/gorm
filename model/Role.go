package model

import (
	"gin/utils"
	"gorm.io/gorm"
)

//
type Role struct {
	Id        int             `gorm:"column:id" json:"id"`                 //是否可空:NO
	Pid       int             `gorm:"column:pid" json:"pid"`               //是否可空:NO 上级id
	Name      string          `gorm:"column:name" json:"name"`             //是否可空:NO 名称
	Rule      string          `gorm:"column:rule" json:"rule"`             //是否可空:NO 权限
	CreatedAt utils.LocalTime `gorm:"column:created_at" json:"created_at"` //是否可空:NO
	UpdatedAt utils.LocalTime `gorm:"column:updated_at" json:"updated_at"` //是否可空:NO
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`          //是否可空:YES
}

func (*Role) TableName() string {
	return "role"
}
