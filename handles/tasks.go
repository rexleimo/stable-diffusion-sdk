package handles

import (
	"context"
	"fmt"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTaskList(filter primitive.D, pageSize int64, pageNumber int64) ([]models.Task, error) {
	var table models.Task
	c := mongodb.GetInstance().Collection(table.TableName())
	var result []models.Task
	queryOptions := options.FindOptions{}
	queryOptions.SetSort(bson.D{{Key: "create_at", Value: -1}})
	queryOptions.SetLimit(pageSize)
	queryOptions.SetSkip(pageSize * (pageNumber - 1))
	c2, err := c.Find(context.Background(), filter, &queryOptions)
	if err != nil {
		return nil, err
	}
	c2.All(context.Background(), &result)
	return result, nil
}

func GetTaskListAll(filter primitive.D) ([]models.Task, error) {
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

func GetTaskListByUserId(userId string, pageSize int64, pageNumber int64) ([]models.Task, error) {
	filter := bson.D{{Key: "uid", Value: userId}}
	return GetTaskList(filter, pageSize, pageNumber)
}

// GetTaskListByTaskId 获取任务列表
func GetTaskListByTaskId(taskId string) (*models.Task, error) {
	filter := bson.D{{Key: "task_id", Value: taskId}}
	return GetTaskOne(filter)
}

// GetTaskListQueyeInTaskId
func GetTaskListQueyeInTaskId(taskId []string) ([]models.Task, error) {

	objectIds := make([]primitive.ObjectID, len(taskId))
	for _, v := range taskId {
		oi, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			continue
		}
		objectIds = append(objectIds, oi)
	}

	filter := bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "$in", Value: objectIds},
		}},
	}
	//log filter
	fmt.Println(filter)
	return GetTaskListAll(filter)
}

func DeleteTask(taskId string, userId string) error {
	oid, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return err
	}
	var table models.Task
	c := mongodb.GetInstance().Collection(table.TableName())
	_, err2 := c.DeleteOne(context.Background(), bson.M{"_id": oid, "uid": userId})
	if err2 != nil {
		return err2
	}

	return nil
}
