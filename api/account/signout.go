package account

import (
	"github.com/xtech-cloud/omo-msa-ams/http"

	"github.com/gin-gonic/gin"
)

func HandleSignout(_uri string, _group *gin.RouterGroup) {

	_group.POST(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		rsp := gin.H{}
		http.RenderOK(_context, rsp)
	})
}
