package handles

import (
	"context"
	"fmt"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/jwtutils"
	"stable-diffusion-sdk/utils/mongodb"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(user *models.User) (*mongo.InsertOneResult, error) {
	c := mongodb.GetInstance().Collection(user.TableName())
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	user.Name = fmt.Sprintf("RexAi用户%d", time.Now().UnixMilli())
	return c.InsertOne(context.Background(), user)
}

func UpdateUser(user *models.User) error {
	c := mongodb.GetInstance().Collection(user.TableName())
	user.UpdateAt = time.Now()
	_, err := c.UpdateOne(context.Background(), bson.D{{
		Key:   "_id",
		Value: user.ID,
	}}, bson.D{{
		Key:   "$set",
		Value: user,
	}})

	if err != nil {
		return err
	}

	return nil
}

func FindUserById(id string) (*models.User, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return FindUser(bson.D{{Key: "_id", Value: objectId}})
}

func FindUserByOpenId(open_id string) (*models.User, error) {
	return FindUser(bson.D{{Key: "open_id", Value: open_id}})
}

func FindUser(f primitive.D) (*models.User, error) {
	var result *models.User
	c := mongodb.GetInstance().Collection(result.TableName())
	sr := c.FindOne(context.Background(), f)
	err := sr.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Login(user *models.User) string {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(2 * time.Hour).Unix(),
		"aud":  "rexai.top",
		"iss":  "rexai",
	}
	s, err := jwtutils.SignedString(claims)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return s
}
