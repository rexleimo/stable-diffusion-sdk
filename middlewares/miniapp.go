package middlewares

import (
	"fmt"
	"stable-diffusion-sdk/utils/jwtutils"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiniAppAuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 验证小程序登录信息
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			ctx.JSON(401, gin.H{"code": 401, "error": "未登录"})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(401, gin.H{"code": 401, "error": "未登录"})
			ctx.Abort()
			return
		}
		mc, err := jwtutils.Parse(parts[1])
		if err != nil {
			fmt.Println(err)
			ctx.JSON(401, gin.H{"code": 401, "error": "无效token"})
			ctx.Abort()
			return
		}

		userId := mc.Subject
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
