package main

/**
入口文件
*/
import (
	"first_go/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	var _ = common.InitDB() // 初始化数据库连接
	r := gin.Default()      // gin 初始化
	CollectRoutes(r)        //通过路由方法 路由到对应接口
	port := viper.GetString("sever.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}

}
