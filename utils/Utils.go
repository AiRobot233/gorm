package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

//获取当前时间
func NowTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

//md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//时间转时间戳
func TimeToStr(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")                     //设置时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	timestamp := tmp.Unix()
	return timestamp
}

//时间戳转时间
func StrToTime(date int64) string {
	tm := time.Unix(date, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//int64转string
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
