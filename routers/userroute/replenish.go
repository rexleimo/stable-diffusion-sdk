package userroute

import (
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"

	"github.com/gin-gonic/gin"
)

func replenish(ctx *gin.Context) {
	userIdStr := ctx.GetString("user_id")
	var payload models.User
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		// resp json error
		ctx.JSON(400, gin.H{
			"error":      err.Error(),
			"error_code": 40013,
		})
		return
	}

	user, _ := handles.FindUserById(userIdStr)

	if user.IsReplenish == 1 {
		ctx.JSON(400, gin.H{
			"error":      "already replenish",
			"error_code": 40203,
		})
		return
	}

	user.Bonus += user.Bonus + 100
	user.Name = payload.Name
	user.Avatar = payload.Avatar
	user.IsReplenish = 1

	err2 := handles.UpdateUser(user)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": "replenish success"})
}
