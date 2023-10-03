package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAttrsById(id string) (*models.Attr, error) {
	var result models.Attr
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := mongodb.GetInstance().Collection(result.TableName()).FindOne(nil, bson.M{"_id": objectId}).Decode(&result)
	return &result, err
}

func GetAttrs() (*[]models.Attr, error) {
	var table models.Attr
	c, err := mongodb.GetInstance().Collection(table.TableName()).Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}
	var result []models.Attr
	err2 := c.All(context.Background(), &result)
	if err2 != nil {
		return nil, err2
	}
	return &result, nil
}
