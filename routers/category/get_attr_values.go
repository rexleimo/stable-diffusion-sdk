package category

import (
	"fmt"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAttrValues(ctx *gin.Context) {
	cid := ctx.Param("id")
	if cid == "" {
		ctx.JSON(400, gin.H{
			"error": "category id is empty",
		})
	}

	category, _ := handles.GetCategoryById(cid)
	var attrsIds []primitive.ObjectID
	for _, attr := range category.Attrs {
		oi, _ := primitive.ObjectIDFromHex(attr)
		attrsIds = append(attrsIds, oi)
	}
	var attrTable models.Attr
	var attrValue models.AttrValue

	c, err := mongodb.GetInstance().Collection(attrTable.TableName()).Find(ctx.Request.Context(), bson.D{{
		Key: "_id",
		Value: bson.D{
			{Key: "$in", Value: attrsIds},
		},
	}})
	if err != nil {
		fmt.Println(err)
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

		vc, _ := mongodb.GetInstance().Collection(attrValue.TableName()).Find(ctx.Request.Context(), bson.D{
			{Key: "pid", Value: id},
		})
		vc.All(ctx.Request.Context(), &queryAttrValue)
		queryAttr[idx].Values = queryAttrValue
	}

	ctx.JSON(200, gin.H{"data": queryAttr})

}
