package account

import (
	"ams/http"
	"ams/model"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func HandleSignup(_uri string, _group *gin.RouterGroup) {

	_group.POST(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		req := SignupRequest{}
		err := _context.ShouldBind(&req)
		http.TryRenderBindError(_context, err)

		dao := model.NewAccountDAO()
		account := model.Account{
			UUID:     model.NewUUID(),
			Username: req.Username,
			Password: dao.StrengthenPassword(req.Password, req.Username),
			Profile:  "",
		}
		err = dao.Insert(account)
		http.TryRenderDatabaseError(_context, err)

		rsp := gin.H{
			"uuid": account.UUID,
		}
		http.RenderOK(_context, rsp)
	})
}
