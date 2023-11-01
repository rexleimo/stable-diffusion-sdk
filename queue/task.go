package queue

import (
	"context"
	"log"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/sdapi/handle"
	"stable-diffusion-sdk/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

var ch chan models.Task = make(chan models.Task, 10)

func RendererTaskChan() chan models.Task {
	return ch
}

func InstallQueyue(data models.Task) {
	ch <- data
}

func ProcessText2ImgQueue() {
	c := mongodb.GetInstance().Collection("tasks")
	for task := range ch {
		// do something
		s, err := handle.ProcessTask(task)
		if err != nil {
			log.Fatal(err)
			// uploda status to 400
			_, err := c.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: task.ID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 400}}}})
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			_, err := c.UpdateOne(context.Background(),
				bson.D{{Key: "_id", Value: task.ID}},
				bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 200},
					{Key: "images", Value: s}}}},
			)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// 重启服务的时候需要init ququq
func Init() {
	c := mongodb.GetInstance().Collection("tasks")
	c2, _ := c.Find(context.Background(), bson.D{
		{Key: "status", Value: 0},
	})
	var list []models.Task
	err := c2.All(context.Background(), &list)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range list {
		ch <- v
	}
}
