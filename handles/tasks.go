package handles

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTaskList(filter primitive.D) ([]models.Task, error) {
	var table models.Task
	c := mongodb.GetInstance().Collection(table.TableName())
	var result []models.Task
	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{Key: "create_at", Value: -1}})
	c2, err := c.Find(context.Background(), filter, &queryOptions)
	if err != nil {
		return nil, err
	}
	c2.All(context.Background(), &result)
	return result, nil
}

func GetTaskOne(filter bson.D) (*models.Task, error) {
	var result models.Task
	c := mongodb.GetInstance().Collection(result.TableName())
	sr := c.FindOne(context.Background(), filter)
	err := sr.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetTaskListByUserId(userId string) ([]models.Task, error) {
	filter := bson.D{{Key: "uid", Value: userId}}
	return GetTaskList(filter)
}

// GetTaskListByTaskId 获取任务列表
func GetTaskListByTaskId(taskId string) (*models.Task, error) {
	filter := bson.D{{Key: "task_id", Value: taskId}}
	return GetTaskOne(filter)
}
