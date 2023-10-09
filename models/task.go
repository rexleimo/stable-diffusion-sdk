package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID             primitive.ObjectID `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
	Pormpt         string             `json:"pormpt" form:"pormpt" bson:"pormpt"`
	NegativePrompt string             `json:"negativePrompt" form:"negativePrompt" bson:"negativePrompt"` // 反向提示词
	CID            string             `json:"cid" form:"cid" bson:"cid"`                                  // 分类ID
	Status         int                `json:"status" form:"status" bson:"status"`
}

func (*Task) TableName() string {
	return "tasks"
}
