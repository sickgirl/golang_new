package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

func main() {

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
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	fmt.Println("hello world!")
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
