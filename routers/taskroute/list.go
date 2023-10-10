package taskroute

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func getTaskList(ctx *gin.Context) {
	userId, ok := ctx.Get("user_id")
	if ok == false {
		ctx.JSON(400, gin.H{"error": "missing user id"})
		return
	}

	t, err := handles.GetTaskListByUserId(userId.(string))
	if err != nil {
		// response error json
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// response data list from t
	ctx.JSON(200, gin.H{"data": t})
}
