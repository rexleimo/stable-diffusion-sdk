package httpserver

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once
var instance *gin.Engine

func GetInstance() *gin.Engine {
	once.Do(func() {
		instance = gin.Default()
	})
	return instance
}
