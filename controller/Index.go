package controller

import (
	"fmt"
	"gin/model"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func Jwt(context *gin.Context) {
	data := utils.GetSlice()
	data["uid"] = 1
	data["user_name"] = "hhh"
	_, token := utils.GetJwt(data, 21)
	utils.Success(context, token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, data)
}

type Tes struct {
	Sql string `form:"sql" json:"sql"`
}

func Test(c *gin.Context) {
	params := Tes{}
	if err := c.ShouldBindJSON(&params); err == nil {
		// 定义正则表达式，匹配 ```sql 和 ``` 之间的 SQL 语句
		re := regexp.MustCompile("(?i)```sql\\s*(.*?)\\s*```")
		// 提取匹配的 SQL 语句
		matches := re.FindStringSubmatch(params.Sql)
		if len(matches) >= 2 {
			cleanSQL := strings.TrimSpace(matches[1])
			var results []map[string]any // 这里改为切片，支持多行数据
			fmt.Println("sql:" + cleanSQL)
			model.GetDb().Raw(cleanSQL).Scan(&results)
			utils.Success(c, results)
		} else {
			utils.Error(c, "未找到 SQL 语句")
		}
	} else {
		utils.Error(c, err.Error())
	}
}

func SetUnitUser(c *gin.Context) {
	var units []model.Unit
	var db = model.GetDb()
	db.Where("is_unit = 1 AND pid != 5").Find(&units)

	password := "Aa@112233"

	for _, unit := range units {
		// 检查是否已经存在同名的用户
		var count int64
		db.Model(&model.User{}).Where("name = ?", unit.Name).Count(&count) // 明确指定用户表
		if count > 0 {
			continue
		}

		// 循环添加数据
		salt := utils.GetSalt(password)
		addUser := model.User{
			Name:     unit.Name,
			Phone:    "",
			Password: utils.Md5(password + salt),
			Salt:     salt,
			Status:   1,
			RoleId:   2,
			UnitId:   unit.Id,
		}

		// 执行插入操作
		result := db.Create(&addUser)
		if result.Error != nil {
			fmt.Printf("添加用户失败: %s, 错误: %v\n", unit.Name, result.Error)
		} else {
			fmt.Printf("添加成功: %s\n", unit.Name)
		}
	}
}
