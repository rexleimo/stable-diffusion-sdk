package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	OpenId   string             `json:"open_id" bson:"open_id"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt time.Time          `json:"update_at" bson:"update_at"`
}

type BaseClaims struct {
	jwt.RegisteredClaims
}

func (*User) TableName() string {
	return "users"
}
