package attrvalues

import "stable-diffusion-sdk/core/httpserver"

func Init() {
	group := httpserver.GetInstance().Group("attr_value")
	group.POST("", addValue)
	group.GET("", getValues)
}
