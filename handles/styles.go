package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertStyle(style models.Style) (*mongo.InsertOneResult, error) {
	c := mongodb.GetInstance().Collection(style.TableName())
	return c.InsertOne(context.Background(), style)
}

func UpdateStyleById(id string, style models.Style) (*mongo.UpdateResult, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return mongodb.GetInstance().Collection(style.TableName()).UpdateOne(context.Background(), bson.D{{
		Key:   "_id",
		Value: objectId,
	}}, bson.D{{
		Key:   "$set",
		Value: style,
	}})
}

func GetStyleList(filter primitive.D) ([]models.Style, error) {
	var table models.Style
	c, err := mongodb.GetInstance().Collection(table.TableName()).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var result []models.Style
	err2 := c.All(context.Background(), &result)
	if err2 != nil {
		return nil, err2
	}
	return result, nil
}

func GetStyleOne(filter primitive.D) (*models.Style, error) {
	var table models.Style
	c := mongodb.GetInstance().Collection(table.TableName()).FindOne(context.Background(), filter)
	err := c.Decode(&table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func GetStyleOneById(id string) (*models.Style, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		Key:   "_id",
		Value: objectId,
	}}
	return GetStyleOne(filter)
}
