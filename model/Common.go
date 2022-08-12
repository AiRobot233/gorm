package model

import (
	"gin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetDb() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Second * 60) //设置链接池超时时间
	return db
}

//分页
func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := utils.StrToInt(page)
		pageSize := utils.StrToInt(pageSize)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
