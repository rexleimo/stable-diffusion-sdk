package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAttrValuesList() (*[]models.AttrValue, error) {
	var table models.AttrValue
	c, err := mongodb.GetInstance().Collection(table.TableName()).Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []models.AttrValue
	err2 := c.All(context.Background(), &result)
	if err2 != nil {
		return nil, err2
	}
	return &result, nil
}

func CreateAttrValue(json models.AttrValue) (*mongo.InsertOneResult, error) {
	return mongodb.GetInstance().Collection(json.TableName()).InsertOne(context.TODO(), json)
}
