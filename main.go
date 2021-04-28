package main

import (
	_ "database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name      string `gorm:"type:vachar(20); not null "`
	Telephone string `gorm:"type:vachar(110); not null;unique "`
	Password  string `gorm:"type:size 255; not null;unique "`
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/user/auth/register", func(c *gin.Context) {

		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		//数据校验
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须大于等于6位"})
		}
		if len(name) == 0 {
			name = RandomString(10)
		}

		c.JSON(200, gin.H{
			"msg":       "用户注册成功",
			"name":      name,
			"password":  password,
			"telephone": telephone,
		})
	})
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnnmQWERTYUIOPASDFGHJKLZXCVBNNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		letters[i] = letters[rand.Intn(len(letters))]
	}
	return string(letters)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "go_test"
	username := "root"
	password := "root"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("connect mysql failed err:" + err.Error())
	}
	return db
}
