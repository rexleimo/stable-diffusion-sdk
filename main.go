package main

import "stable-diffusion-sdk/routers"

func main() {
	// s, _ := handle.Text2ImgApi()

	// for _, v := range s {
	// 	data, _ := base64.StdEncoding.DecodeString(v)
	// 	ioutil.WriteFile("./"+"test"+".png", data, 0644)
	// }

	routers.Init()

}
