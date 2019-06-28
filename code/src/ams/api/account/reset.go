package account

import (
	"ams/http"
	"ams/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type ResetRequest struct {
	Password string `form:"password" json:"password" binding:"required"`
}

func HandleReset(_uri string, _group *gin.RouterGroup) {

	_group.POST(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		claims := jwt.ExtractClaims(_context)
		uuid := claims["id"].(string)

		req := ResetRequest{}
		err := _context.ShouldBind(&req)
		http.TryRenderBindError(_context, err)

		dao := model.NewAccountDAO()
		account, err := dao.Find(uuid)
		http.TryRenderDatabaseError(_context, err)

		password := dao.StrengthenPassword(req.Password, account.Username)
		err = dao.UpdatePassword(uuid, password)
		http.TryRenderDatabaseError(_context, err)

		rsp := gin.H{}
		http.RenderOK(_context, rsp)
	})
}
