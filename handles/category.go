package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCategoryList() (*[]models.Categories, error) {
	var table models.Categories
	c := mongodb.GetInstance().Collection(table.TableName())
	c2, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []models.Categories
	c2.All(context.Background(), &result)
	return &result, nil
}

func GetCategoryById(id string) (*models.Categories, error) {
	var result models.Categories
	c := mongodb.GetInstance().Collection(result.TableName())
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := c.FindOne(context.Background(), bson.D{
		{Key: "_id", Value: objectId},
	}).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func associateRecurse(all []models.Categories, pid string) []models.Categories {
	var result []models.Categories
	for _, v := range all {
		if v.PID == pid {
			child := associateRecurse(all, v.ID.Hex())
			node := models.Categories{
				ID:       v.ID,
				PID:      v.PID,
				Name:     v.Name,
				EnName:   v.EnName,
				Children: child,
			}
			result = append(result, node)
		}
	}
	return result
}

func GetCategoryListLevel() (*[]models.Categories, error) {
	list, err := GetCategoryList()
	if err != nil {
		return nil, err
	}

	result := associateRecurse(*list, "")

	return &result, nil
}
