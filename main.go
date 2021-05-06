package main

/**
入口文件
*/
import (
	"first_go/common"
	"github.com/gin-gonic/gin"
)

func main() {
	var _ = common.InitDB() // 初始化数据库连接
	r := gin.Default()      // gin 初始化
	CollectRoutes(r)        //通过路由方法 路由到对应接口
	panic(r.Run())          // 监听并在 0.0.0.0:8080 上启动服务
}
