package models

import (
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID               primitive.ObjectID `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
	UID              string             `json:"uid" form:"uid" bson:"uid"`
	CID              string             `json:"cid" form:"cid" bson:"cid"`          // 分类ID
	Status           int                `json:"status" form:"status" bson:"status"` // 操作状态
	SDRendererConfig payload.SDParams   `json:"sd_renderer_config" form:"sd_renderer_config" bson:"sd_renderer_config"`
	Images           []string           `json:"images" form:"images" bson:"images"` // 渲染数据集
	CreateAt         time.Time          `json:"create_at" form:"create_at" bson:"create_at"`
	UpdateAt         time.Time          `json:"update_at" form:"update_at" bson:"update_at"`
	IsDelete         bool               `json:"is_delete" form:"is_delete" bson:"is_delete"`
}

func (*Task) TableName() string {
	return "tasks"
}