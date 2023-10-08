package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateAttrValue(id string, json models.AttrValue) (*mongo.UpdateResult, error) {
	oi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return mongodb.GetInstance().Collection(json.TableName()).UpdateOne(context.Background(), bson.M{"_id": oi}, bson.M{"$set": json})
}

func GetAttrValueInfo(bson primitive.D) (*models.AttrValue, error) {

	var result models.AttrValue
	c := mongodb.GetInstance().Collection(result.TableName())
	err2 := c.FindOne(context.Background(), bson).Decode(&result)
	if err2 != nil {
		return nil, err2
	}
	return &result, nil
}

func GetAttrValueInfoById(id string) (*models.AttrValue, error) {
	oi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return GetAttrValueInfo(bson.D{
		{Key: "_id", Value: oi},
	})
}
