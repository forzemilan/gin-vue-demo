package middleware

import (
	"ginessential/common"
	"ginessential/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 鉴权中间价
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 验证不通过则抛弃本次请求
			ctx.Abort()
			return
		}
		// 前缀Bearer 占了7位，有一个空格
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 验证不通过则抛弃本次请求
			ctx.Abort()
			return
		}
		// 验证通过后，获取claims中的userID
		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)
		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// 验证不通过则抛弃本次请求
			ctx.Abort()
			return
		}
		// 用户存在，将用户信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
