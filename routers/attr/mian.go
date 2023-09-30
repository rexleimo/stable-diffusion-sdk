package attrs

import "stable-diffusion-sdk/core/httpserver"

func Init() {
	group := httpserver.GetInstance().Group("attr")
	group.POST("", addAttr)
	group.GET("", getAttrs)

}
