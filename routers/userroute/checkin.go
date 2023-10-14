package userroute

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func checkIn(ctx *gin.Context) {
	// 签到功能
	userIdStr := ctx.GetString("user_id")
	err := handles.CheckInUserToday(userIdStr)
	if err != nil {
		// response json {error:''}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "签到成功"})
}

func userInfo(ctx *gin.Context) {
	userIdStr := ctx.GetString("user_id")
	userInfo, _ := handles.FindUserById(userIdStr)
	ctx.JSON(200, gin.H{"data": userInfo})
}
