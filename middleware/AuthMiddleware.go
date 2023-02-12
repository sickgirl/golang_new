package middleware

import (
	"first_go/common"
	"first_go/db"
	"first_go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取 authorization header
		tokenString := context.GetHeader("Authorization")
		//验证 token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			//基础接受参数校验
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			println(err.Error())
			//解密验证
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
			context.Abort()
			return
		}
		userId := claims.UserId

		println("test111:", userId)
		var user model.User
		db1 := db.GetDb()

		db1.First(&user, userId)
		if user.ID == 0 {
			//用户信息校验
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足3"})
			context.Abort()
			return
		}

		//用户信息存在  则写入上下文
		context.Set("user", user)
		context.Next()

	}
}
