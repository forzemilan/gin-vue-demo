package controller

import (
	"ginessential/common"
	"ginessential/dto"
	"ginessential/model"
	"ginessential/response"
	"ginessential/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Register 用户注册功能
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//! 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//! 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不得少于6位")
		return
	}
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	//! 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	//! 加密密码并创建用户
	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "密码加密失败")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPasswd),
	}
	DB.Create(&newUser)
	//! 返回结果
	response.Success(ctx, nil, "注册成功")
}

// Login 用户登录
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//! 获取数据
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//! 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不得少于6位")
		return
	}
	//! 判断手机号是否存在;存在-->获取用户名
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//! 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//! 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate failed, %v", err)
		return
	}
	//! 返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

// Info 用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

// 判断手机号是否已经存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
