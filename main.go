package main

/**
入口文件
*/
import (
	"first_go/common"
	"github.com/gin-gonic/gin"
)

func main() {

	var _ = common.InitDB()
	r := gin.Default()
	CollectRoutes(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
