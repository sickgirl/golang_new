package controller

//控制器
import (
	"first_go/common"
	"first_go/db"
	"first_go/dto"
	"first_go/gtp"
	"first_go/model"
	"first_go/response"
	"first_go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}}, "获取用户信息成功")
}

func Question(c *gin.Context) {
	user, _ := c.Get("user")
	uid := user.(model.User).ID
	times, err := model.UseTimes(c, uid)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "请求失败请稍后再试")
		return
	} else if times < 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "请求次数已耗尽")
		return
	}
	sentence := c.PostForm("question")
	reply, err := gtp.Completions(sentence)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "当前请求异常")
		return
	} else {
		response.Success(c, gin.H{"times": times, "answer": reply}, "获取用户信息成功")
		return
	}
}

func Login(c *gin.Context) {
	//获取参数  数据验证  判断手机号是否存在  密码是否争取  发放token  返回结构
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")

		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于等于6位")

		return
	}
	//判断手机号 是否存在
	getDb := db.GetDb()
	var user model.User
	getDb.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常,生成密码失败")
		log.Printf("token generate error : %v", err)
		return
	}
	var times, _ = model.GetTimes(c, user.ID)
	num, err := strconv.Atoi(times)
	if err != nil {
		num = 0
	}

	response.Success(c, gin.H{"token": token, "times": num}, "登陆成功")
}

func Register(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	//数据校验
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码必须大于等于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	getDb := db.GetDb()
	fmt.Printf("%+v", getDb)
	if IsTelephoneExit(getDb, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户手机已注册")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密异常")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	getDb.Create(&newUser)

	response.Success(c, nil, "注册成功")

}

func IsTelephoneExit(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	} else {
		return false
	}
}
