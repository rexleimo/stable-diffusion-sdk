package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AttrValue struct {
	// ID is the unique identifier of the industry
	ID int64 `json:"id,omitempty" bson:"_id,omitempty"`
	// Name is the name of the industry
	Name     string    `json:"name" binding:"required"`
	EnName   string    `json:"en_name" bson:"en_name" binding:"required"` //由于没钱不调用翻译api
	PID      string    `json:"pid" bson:"pid" binding:"required"`
	Covery   string    `json:"coverage" bson:"cover"`
	CreateAt time.Time `json:"create_at,omitempty" bson:"create_at"`
	UpdateAt time.Time `json:"update_at,omitempty" bson:"update_at"`
}

func (i *AttrValue) TableName() string {
	return "attr_values"
}

func (i *AttrValue) MarshalBSON() ([]byte, error) {
	if i.CreateAt.IsZero() {
		i.CreateAt = time.Now()
	}
	i.UpdateAt = time.Now()

	type my AttrValue
	return bson.Marshal((*my)(i))
}
