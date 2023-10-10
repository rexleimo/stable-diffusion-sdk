package main

import (
	adminrouters "stable-diffusion-sdk/admin/routers"
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/queue"
	"stable-diffusion-sdk/routers"
)

func main() {
	httpserver.GetInstance().LoadHTMLGlob("templates/**/*")
	routers.Init()
	adminrouters.Init()
	go queue.ProcessText2ImgQueue()
	httpserver.GetInstance().Run(":7100")
}
