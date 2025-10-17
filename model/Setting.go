package model

import "gin/utils"

type Setting struct {
	Id           int              `gorm:"column:id" json:"id,omitempty"`                     //是否可空:NO
	SettingKey   string           `gorm:"column:setting_key;uniqueIndex" json:"setting_key"` //是否可空:NO key
	SettingValue string           `gorm:"column:setting_value" json:"setting_value"`         //是否可空:NO value
	SettingType  int              `gorm:"column:setting_type" json:"setting_type"`           //是否可空:NO value
	CreatedAt    *utils.LocalTime `gorm:"column:created_at" json:"created_at"`               //是否可空:NO
	UpdatedAt    *utils.LocalTime `gorm:"column:updated_at" json:"updated_at"`               //是否可空:NO
}

func (*Setting) TableName() string {
	return "setting"
}
