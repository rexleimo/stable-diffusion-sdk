package attrs

import (
	"context"
	"fmt"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
)

func post(ctx *gin.Context) {
	idTxt := ctx.Query("id")

	if idTxt == "" {
		ctx.HTML(200, "attrs/post.html", gin.H{
			"title": "属性编辑",
			"info": &models.Attr{
				Name:   "",
				EnName: "",
			},
		})
		return
	}

	info, err := handles.GetAttrsById(idTxt)
	if err != nil {
		fmt.Println(err)
		ctx.HTML(500, "public/error.html", gin.H{"error": err})
		return
	}

	ctx.HTML(200, "attrs/post.html", gin.H{
		"title": "属性编辑",
		"info":  info,
	})

}

func create(ctx *gin.Context) {
	var json models.Attr
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c := mongodb.GetInstance().Collection(json.TableName())

	result, err2 := c.InsertOne(context.Background(), json)

	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": true, "data": result.InsertedID})
}
