package utils

import (
	"stable-diffusion-sdk/core/httpserver"
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once
var instance *gin.RouterGroup

func GetAdminRouter() *gin.RouterGroup {
	once.Do(func() {
		group := httpserver.GetInstance().Group("admin")
		instance = group
	})
	return instance
}
