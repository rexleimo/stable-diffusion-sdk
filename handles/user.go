package handles

import (
	"context"
	"errors"
	"fmt"
	"os"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/jwtutils"
	"stable-diffusion-sdk/utils/mongodb"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

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

// 判断用户今天是否签到了
func CheckInUserToday(userId string) error {
	// 判断是否存在 logs文件夹不存在创建
	_, err := os.Stat("./logs")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./logs", 0777)
		}
	}
	dbName := fmt.Sprintf("./logs/bonus-%s.db", time.Now().Format("20060102"))
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		return err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()
	// 迁移表
	db.AutoMigrate(&models.UserCheckInToday{})
	var table models.UserCheckInToday
	db.Find(&table, "user_id = ?", userId)

	// 如果存在 返回已经签到
	if table.ID != 0 {
		return errors.New("already check in today")
	}
	table.UserID = userId
	d := db.Create(&table)

	if d.Error != nil {
		return d.Error
	}

	u, err2 := FindUserById(userId)
	if err2 != nil {
		db.Delete(&table, "user_id = ?", userId)
		return err2
	}
	fmt.Println(u)
	u.Bonus += u.Bonus + 50
	err3 := UpdateUser(u)

	if err3 != nil {
		db.Delete(&table, "user_id = ?", userId)
		return err3
	}

	return nil
}
