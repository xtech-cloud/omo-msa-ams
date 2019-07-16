package account

import (
	"github.com/xtech-cloud/omo-msa-ams/http"
	"github.com/xtech-cloud/omo-msa-ams/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HandleCurrent(_uri string, _group *gin.RouterGroup) {

	_group.GET(_uri, func(_context *gin.Context) {
		defer http.CatchRenderError()

		claims := jwt.ExtractClaims(_context)
		uuid := claims["id"].(string)

		dao := model.NewAccountDAO()
		account, err := dao.Find(uuid)

		http.TryRenderDatabaseError(_context, err)

		rsp := gin.H{
			"uuid":    account.UUID,
			"profile": account.Profile,
		}
		http.RenderOK(_context, rsp)
	})
}
