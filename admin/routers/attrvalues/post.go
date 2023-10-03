package attrvalues

import (
	"fmt"
	"net/http"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"

	"github.com/gin-gonic/gin"
)

func posts(ctx *gin.Context) {
	a, err := handles.GetAttrs()
	fmt.Println(a)
	if err != nil {
		ctx.HTML(500, "public/error.html", gin.H{"error": err})
	}
	ctx.HTML(http.StatusOK, "attrvalues/posts.html", gin.H{
		"attrs": a,
	})
}

func create(ctx *gin.Context) {
	var json models.AttrValue
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err2 := handles.CreateAttrValue(json)

	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": true})
}
