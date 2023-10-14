package miniapp

import (
	"encoding/json"
	"fmt"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/config"
	"stable-diffusion-sdk/utils/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResJscode2session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func getOpenId(ctx *gin.Context) {

	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(400, gin.H{"error": "code is empty"})
		return
	}
	r := http.GetInstance().R()
	mac := config.GetConfig().MiniAppConfig
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", mac.Appid, mac.Appsecret, code)
	r2, err := r.Get(url)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var resp ResJscode2session
	err2 := json.Unmarshal(r2.Body(), &resp)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}
	open_id := resp.Openid
	u, _ := handles.FindUserByOpenId(open_id)

	if open_id == "" {
		ctx.JSON(400, gin.H{"error": "open_id is empty"})
		return
	}

	saveUser := &models.User{
		OpenId: open_id,
	}

	if u == nil {
		// 不存在的时候添加一条
		ior, err3 := handles.InsertUser(saveUser)
		if err3 != nil {
			ctx.JSON(400, gin.H{"error": err3.Error()})
			return
		}
		saveUser.ID = ior.InsertedID.(primitive.ObjectID)
	} else {
		saveUser = u
	}

	jwt_token := handles.Login(&models.User{
		ID:   saveUser.ID,
		Name: saveUser.Name,
	})

	ctx.JSON(200, gin.H{"data": jwt_token})

}
