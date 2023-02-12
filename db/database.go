package db

//链接数据库 及 返回链接
import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	DB, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("connect mysql failed err:" + err.Error())
	}
	//err2 := db.AutoMigrate(&model.User{})
	//if err2 != nil {
	//	panic("create mysql table failed err:" + err2.Error())
	//}
	//DB = db
	return DB
}

func GetDb() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	DB, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("connect mysql failed err:" + err.Error())
	}
	return DB
}
