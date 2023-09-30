package category

import (
	"stable-diffusion-sdk/core/httpserver"
)

func Init() {
	group := httpserver.GetInstance().Group("category")
	group.POST("", create)
	group.GET("", query)
	group.GET("/attr/:id", getAttrValues)
}
