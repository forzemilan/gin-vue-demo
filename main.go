package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//! 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不得少于6位",
			})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		//! 判断手机号是否存在
		//! 创建用户

		//! 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
		})
	})
	panic(r.Run())
}

// RandomString 随机产生10个大小写字母
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
