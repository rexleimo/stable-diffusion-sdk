package attrvalues

import (
	"net/http"
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func list(ctx *gin.Context) {
	av, _ := handles.GetAttrValuesList()
	ctx.HTML(http.StatusOK, "attrvalues/list.html", gin.H{"list": av})
}
