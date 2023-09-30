package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	EnName    string             `json:"en_name" bson:"en_name"` //由于没钱不调用翻译api
	PID       string             `json:"pid" bson:"pid"`
	Cover     string             `json:"cover" bson:"cover"`
	Sort      int                `json:"sort" bson:"sort"`
	DeletedAt string             `json:"deleted_at" bson:"deleted_at"`
	IsDeleted int                `json:"is_deleted" bson:"is_deleted"`
	CreateAt  time.Time          `json:"create_at,omitempty" bson:"create_at"`
	UpdateAt  time.Time          `json:"update_at,omitempty" bson:"update_at"`
	Children  []*Categories
}

func (*Categories) TableName() string {
	return "categories"
}

func (i *Categories) MarshalBSON() ([]byte, error) {
	fmt.Println("xx")
	if i.CreateAt.IsZero() {
		i.CreateAt = time.Now()
	}
	i.UpdateAt = time.Now()

	type my Categories
	return bson.Marshal((*my)(i))
}
