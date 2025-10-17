package admin

import (
	"gin/model"
	"gin/utils"
	"gorm.io/gorm/clause"
)

func SettingWeb() (bool, any) {
	var settingList []*model.Setting
	result := db.Where("setting_type = ?", 1).Find(&settingList)
	return utils.R(result, settingList)
}

func SettingList() (bool, any) {
	var settingList []*model.Setting
	result := db.Find(&settingList)
	return utils.R(result, settingList)
}

func SettingSave(data []model.Setting) (bool, any) {
	for _, item := range data {
		// 类似于 Laravel 的 updateOrCreate
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "setting_key"}},
			DoUpdates: clause.AssignmentColumns([]string{"setting_type", "setting_value"}),
		}).Create(&item).Error

		if err != nil {
			return false, err.Error()
		}
	}
	return true, nil
}
