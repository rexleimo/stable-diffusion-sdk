package attrvalues

import (
	"fmt"
	"net/http"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"

	"github.com/gin-gonic/gin"
)

func posts(ctx *gin.Context) {

	idQuery := ctx.Query("id")

	a, err := handles.GetAttrs()
	if err != nil {
		ctx.HTML(500, "public/error.html", gin.H{"error": err})
		return
	}

	var info *models.AttrValue

	if idQuery != "" {
		info, err = handles.GetAttrValueInfoById(idQuery)
		if err != nil {
			fmt.Println(err)
			ctx.HTML(500, "public/error.html", gin.H{"error": err})
			return
		}
		ctx.HTML(http.StatusOK, "attrvalues/edit.html", gin.H{
			"attrs": a,
			"info":  info,
		})
	} else {
		ctx.HTML(http.StatusOK, "attrvalues/posts.html", gin.H{
			"attrs": a,
		})
	}

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

func update(ctx *gin.Context) {
	idQuery := ctx.Query("id")
	if idQuery == "" {
		ctx.JSON(500, gin.H{"error": "id is required"})
		return
	}
	var json models.AttrValue
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err2 := handles.UpdateAttrValue(idQuery, json)

	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": true})
}
