package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Style struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Cover          string             `json:"cover" form:"cover" bson:"cover"`
	Checkpoint     string             `json:"checkpoint" bson:"checkpoint"`
	Pormpt         string             `json:"pormpt" form:"pormpt" bson:"prompt"`
	NegativePrompt string             `json:"negative_prompt" form:"negative_prompt" bson:"negative_prompt"`
	SamplerIndex   string             `json:"sampler_index" form:"sampler_index" bson:"sampler_index"`
	CfgScale       float32            `json:"cfg_scale,string" form:"cfg_scale" bson:"cfg_scale"`
	Steps          float32            `json:"steps,string" form:"steps" bson:"steps"`
	Sort           int8               `json:"sort" bson:"sort"`
	CreateAt       time.Time          `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdateAt       time.Time          `json:"update_at,omitempty" bson:"update_at,omitempty"`
	DeletedAt      time.Time          `json:"deleted_at" form:"deleted_at" bson:"deleted_at"`
}

func (*Style) TableName() string {
	return "styles"
}
