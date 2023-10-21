package stylesroute

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getStyleList(ctx *gin.Context) {
	filter := bson.D{}
	s, err := handles.GetStyleList(filter)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"styles": s})
	}
}
