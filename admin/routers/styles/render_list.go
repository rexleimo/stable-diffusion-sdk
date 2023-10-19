package styles

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func renderList(ctx *gin.Context) {

	list, _ := handles.GetStyleList(bson.D{})

	ctx.HTML(200, "styles/index.html", gin.H{
		"list": list,
	})
}
