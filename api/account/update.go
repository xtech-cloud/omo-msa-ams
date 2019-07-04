package account

import (
	"github.com/xtech-cloud/omo-msa-ams/http"
	"github.com/xtech-cloud/omo-msa-ams/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Profile string `form:"profile" json:"profile" binding:"required"`
}

func HandleUpdate(_uri string, _group *gin.RouterGroup) {

	_group.POST(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		claims := jwt.ExtractClaims(_context)
		uuid := claims["id"].(string)

		req := UpdateRequest{}
		err := _context.ShouldBind(&req)
		http.TryRenderBindError(_context, err)

		dao := model.NewAccountDAO()
		err = dao.UpdateProfile(uuid, req.Profile)
		http.TryRenderDatabaseError(_context, err)

		rsp := gin.H{}
		http.RenderOK(_context, rsp)
	})
}
