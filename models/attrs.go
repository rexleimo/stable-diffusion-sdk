package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attr struct {
	// ID is the unique identifier of the industry
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	// Name is the name of the industry
	Name   string `json:"name" binding:"required"`
	EnName string `json:"en_name" bson:"en_name" binding:"required"` //由于没钱不调用翻译api
	PID    string `json:"pid" bson:"pid" binding:"required"`
	Covery string `json:"coverage" bson:"cover"`

	Values []AttrValue `json:"values" bson:"-"`

	CreateAt time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty" bson:"update_at,omitempty"`
}

func (i *Attr) TableName() string {
	return "attrs"
}

func (i *Attr) MarshalBSON() ([]byte, error) {
	if i.CreateAt.IsZero() {
		i.CreateAt = time.Now()
	}
	i.UpdateAt = time.Now()

	type my Attr
	return bson.Marshal((*my)(i))
}
