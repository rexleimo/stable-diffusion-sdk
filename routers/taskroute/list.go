package taskroute

import (
	"stable-diffusion-sdk/handles"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTaskList(ctx *gin.Context) {
	userId, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(400, gin.H{"error": "missing user id"})
		return
	}

	pageSize := 10
	pageNumber := 1

	// if gin query pageSize value to pageSize
	pageSizeStr := ctx.Query("pageSize")
	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	// if gin query page value to pageNumber
	pageStr := ctx.Query("page")
	if pageStr != "" {
		pageNumber, _ = strconv.Atoi(pageStr)
	}

	t, err := handles.GetTaskListByUserId(userId.(string), int64(pageSize), int64(pageNumber))
	if err != nil {
		// response error json
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// response data list from t
	ctx.JSON(200, gin.H{"data": t})
}

func getTaskListByIds(ctx *gin.Context) {

	var json struct {
		TaskIds []string `json:"task_ids"`
	}

	err2 := ctx.ShouldBindJSON(&json)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": "missing taskIds"})
		return
	}
	t, err := handles.GetTaskListQueyeInTaskId(json.TaskIds)
	if err != nil {
		// response error json
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": t})
}
