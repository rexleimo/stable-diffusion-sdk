package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	ID             primitive.ObjectID `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name" form:"name" bson:"name" binding:"required"`
	EnName         string             `json:"en_name" form:"en_name" bson:"en_name" binding:"required"` //由于没钱不调用翻译api
	PID            string             `json:"pid" form:"pid" bson:"pid"`
	Cover          string             `json:"cover" form:"cover" bson:"cover"`
	Sort           int                `json:"sort" form:"sort" bson:"sort"`
	IsDeleted      int                `json:"is_deleted" form:"is_deleted" bson:"is_deleted"`
	Attrs          []string           `json:"attrs" form:"attrs[]" bson:"attrs"`
	Checkpoint     string             `json:"checkpoint" form:"checkpoint"`
	Pormpt         string             `json:"pormpt" form:"pormpt" bson:"prompt"`
	NegativePrompt string             `json:"negative_prompt" form:"negative_prompt" bson:"negative_prompt"`
	SamplerIndex   string             `json:"sampler_index" form:"sampler_index" bson:"sampler_index"`
	CfgScale       int32              `json:"cfg_scale" form:"cfg_scale" bson:"cfg_scale"`
	Steps          int32              `json:"steps" form:"steps" bson:"steps"`
	IsSize         bool               `json:"is_size" form:"is_size" bson:"is_size"`
	Imgw           int32              `json:"img_w" form:"img_w" bson:"img_w" binding:"required"`
	Imgh           int32              `json:"img_h" form:"img_h" bson:"img_h" binding:"required"`

	CreateAt  time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty" bson:"update_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at" bson:"deleted_at"`
	Children  []Categories
}

func (*Categories) TableName() string {
	return "categories"
}

func (i *Categories) MarshalBSON() ([]byte, error) {
	if i.CreateAt.IsZero() {
		i.CreateAt = time.Now()
	}
	i.UpdateAt = time.Now()

	type my Categories
	return bson.Marshal((*my)(i))
}
