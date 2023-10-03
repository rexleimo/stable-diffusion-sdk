package attrs

import (
	"net/http"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getList(ctx *gin.Context) {
	var table models.Attr
	c := mongodb.GetInstance().Collection(table.TableName())
	c2, err := c.Find(ctx.Request.Context(), bson.D{})
	if err != nil {
		ctx.HTML(400, "error.tmpl", gin.H{"error": err})
		return
	}
	var attrs []models.Attr
	c2.All(ctx.Request.Context(), &attrs)
	ctx.HTML(http.StatusOK, "attrs/list.html", gin.H{
		"title": "属性管理",
		"attrs": attrs,
	})
}
