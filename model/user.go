package model

//服务类
import (
	"first_go/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//TimesGptUser 用户次数
const (
	TimesGptUser = "key:uid:times:%s" //服务名:功能:后续参数
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null "`
	Telephone string `gorm:"type:varchar(110); not null;unique "`
	Password  string `gorm:"type:varchar(255); not null;unique "`
}

func UseTimes(c *gin.Context, uid uint) (int64, error) {
	key := fmt.Sprintf(TimesGptUser, uid)
	res := common.Redis.Incr(key)
	return res.Result()
}

func GetTimes(c *gin.Context, uid uint) (string, error) {
	key := fmt.Sprintf(TimesGptUser, uid)
	res := common.Redis.Get(key)
	return res.Result()
}
