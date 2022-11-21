package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//string 转 int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//获取当前时间
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
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
	return tmp.Unix()
}

//时间戳转时间
func StrToTime(date int64) string {
	return time.Unix(date, 0).Format("2006-01-02 15:04:05")
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
	p := GetSlice()
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

//校验密码长度
func CheckPasswordLever(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("密码长度必须大于9位")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("密码需要数字")
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("密码需要小写字母")
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("密码需要大写字母")
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("密码需要特殊符号")
	}
	return nil
}

//判断数据是否在数组中
func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

//返回数组切片数据
func GetSlice() map[string]interface{} {
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	return params
}
