package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

// StrToInt string 转 int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// NowTime 获取当前时间
func NowTime() string {
	return time.Now().Format(time.DateTime)
}

// Md5 md5加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// TimeToStr 时间转时间戳
func TimeToStr(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tmp, _ := time.ParseInLocation(time.DateTime, date, loc)
	return tmp.Unix()
}

// StrToTime 时间戳转时间
func StrToTime(date int64) string {
	return time.Unix(date, 0).Format(time.DateTime)
}

// Int64ToStr int64转string
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// GetSalt 获取盐
func GetSalt(str string) string {
	r := rand.Int()
	s := strconv.Itoa(r)
	timeUnixNano := time.Now().UnixNano()
	timeS := strconv.FormatInt(timeUnixNano, 10)
	m5 := Md5(str + timeS + s)
	return m5[0:5]
}

// P 输出分页
func P(data any, count int64) map[string]any {
	p := GetSlice()
	p["list"] = data
	p["total"] = count
	return p
}

// R 输出错误或正常数据
func R(err *gorm.DB, data any) (bool, any) {
	if err.Error != nil {
		return false, err.Error.Error()
	} else {
		return true, data
	}
}

// JSONMethod json转map
func JSONMethod(content any) []map[string]any {
	var name []map[string]any
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

// CheckPasswordLever 校验密码长度
func CheckPasswordLever(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("密码长度必须大于9位")
	}
	num := `[0-9]{1}`
	aZ := `[a-z]{1}`
	AZ := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("密码需要数字")
	}
	if b, err := regexp.MatchString(aZ, ps); !b || err != nil {
		return fmt.Errorf("密码需要小写字母")
	}
	if b, err := regexp.MatchString(AZ, ps); !b || err != nil {
		return fmt.Errorf("密码需要大写字母")
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("密码需要特殊符号")
	}
	return nil
}

// InArray 判断数据是否在数组中
func InArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// GetSlice 返回数组切片数据
func GetSlice() map[string]any {
	var params map[string]any     //声明变量，不分配内存
	params = make(map[string]any) //必可不少，分配内存
	return params
}

// GetEnvData 读取env文件
func GetEnvData(name string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(name)
}

// SetPwd 修改密码操作
func SetPwd(password string, salt string) (bool, string) {
	err := CheckPasswordLever(password) //校验密码安全性
	if err != nil {
		return false, err.Error()
	}
	return true, Md5(password + salt)
}
