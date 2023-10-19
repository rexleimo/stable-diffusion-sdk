package styles

import "github.com/gin-gonic/gin"

func renderList(ctx *gin.Context) {
	ctx.HTML(200, "styles/index.html", gin.H{})
}
