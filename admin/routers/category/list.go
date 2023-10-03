package category

import (
	"net/http"
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func list(ctx *gin.Context) {
	idTxt := ctx.Query("id")
	list, _ := handles.GetCategoryListLevel()
	info, _ := handles.GetCategoryById(idTxt)
	a, _ := handles.GetAttrs()
	template := ""
	if idTxt != "" {
		template = "category/edit.html"
	} else {
		template = "category/index.html"
	}

	ctx.HTML(http.StatusOK, template, gin.H{
		"title": "分类管理",
		"list":  list,
		"info":  info,
		"attrs": a,
	})
	return
}
