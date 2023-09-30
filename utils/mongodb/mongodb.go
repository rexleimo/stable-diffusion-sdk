package mongodb

import (
	"context"
	"stable-diffusion-sdk/utils/config"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var instance *mongo.Database

func GetInstance() *mongo.Database {

	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		option := options.Client().ApplyURI(config.GetConfig().MongoDB.Host)
		option.SetMaxConnecting(10)
		option.SetMaxPoolSize(10)
		option.SetMinPoolSize(1)

		c, err := mongo.Connect(ctx, option)

		if err != nil {
			panic(err)
		}
		instance = c.Database(config.GetConfig().MongoDB.Dbname)
	})

	return instance

}
