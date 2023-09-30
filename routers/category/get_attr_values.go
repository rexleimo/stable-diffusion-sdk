package category

import (
	"fmt"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getAttrValues(ctx *gin.Context) {
	cid := ctx.Param("id")
	if cid == "" {
		ctx.JSON(400, gin.H{
			"error": "category id is empty",
		})
	}

	var attrTable models.Attr
	var attrValue models.AttrValue

	c, err := mongodb.GetInstance().Collection(attrTable.TableName()).Find(ctx.Request.Context(), bson.M{"pid": cid})
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var queryAttr []models.Attr
	c.All(ctx.Request.Context(), &queryAttr)

	for idx, attr := range queryAttr {
		var queryAttrValue []models.AttrValue
		id := attr.ID.Hex() // 获取ObjectId的字符串
		fmt.Println(id)

		vc, _ := mongodb.GetInstance().Collection(attrValue.TableName()).Find(ctx.Request.Context(), bson.D{
			{Key: "pid", Value: id},
		})
		vc.All(ctx.Request.Context(), &queryAttrValue)
		queryAttr[idx].Values = queryAttrValue
	}

	ctx.JSON(200, gin.H{"data": queryAttr})

}
