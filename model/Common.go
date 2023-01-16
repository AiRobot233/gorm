package model

import (
	"gin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDb() *gorm.DB {
	dsn := utils.GetEnvData("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlDB, dbErr := db.DB()
	if dbErr != nil {
		panic(err.Error())
	}
	sqlDB.SetConnMaxLifetime(time.Second * 60) //设置链接池超时时间
	return db
}

// Paginate 分页
func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := utils.StrToInt(page)
		pageSize := utils.StrToInt(pageSize)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
