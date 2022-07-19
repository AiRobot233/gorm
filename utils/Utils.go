package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
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

//获取盐
func GetSalt(str string) string {
	r := rand.Int()
	s := strconv.Itoa(r)
	timeUnixNano := time.Now().UnixNano()
	timeS := strconv.FormatInt(timeUnixNano, 10)
	m5 := Md5(str + timeS + s)
	return m5[0:5]
}

//输出分页
func P(data interface{}, count int64) map[string]interface{} {
	var p map[string]interface{}     //定义map
	p = make(map[string]interface{}) //初始化map
	p["list"] = data
	p["total"] = count
	return p
}

//输出错误或正常数据
func R(err *gorm.DB, data interface{}) (bool, interface{}) {
	if err.Error != nil {
		return false, err.Error.Error()
	} else {
		return true, data
	}
}

//json转map
func JSONMethod(content interface{}) []map[string]interface{} {
	var name []map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
