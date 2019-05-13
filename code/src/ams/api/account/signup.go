package account

import (
	"ams/http"
	"ams/model"
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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
		http.TryRenderDatabaseError(_context, err)

		dao := model.NewAccountDAO()
		account := model.Account{
			UUID:     newGUID(),
			Username: req.Username,
			Password: req.Password,
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

func newGUID() string {
	uid, _ := uuid.NewV4()
	hex := md5.Sum(uid.Bytes())
	return fmt.Sprintf("%x", hex)
}