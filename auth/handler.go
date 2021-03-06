package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/xtech-cloud/omo-msa-ams/core"
	"github.com/xtech-cloud/omo-msa-ams/model"

	"github.com/gin-gonic/gin"
)

type signin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

var ISS = "ams"
var Secret = "ams-secret"

type account struct {
	ID string
}

func BindAuthHandler(_router *gin.Engine, _uri string, _group string) *gin.RouterGroup {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "ams",
		Key:         []byte(Secret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			core.Logger.Debugf("handle payload: %v", data)
			if v, ok := data.(*account); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"iss": ISS,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			core.Logger.Debugf("handle identity: %v", claims)
			return &account{
				ID: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			core.Logger.Debug("Authenticator")

			var signinVals signin
			if err := c.ShouldBind(&signinVals); err != nil {
				core.Logger.Error(err)
				return "", jwt.ErrMissingLoginValues
			}

			core.Logger.Info(signinVals)

			dao := model.NewAccountDAO()
			username := signinVals.Username
			password := dao.StrengthenPassword(signinVals.Password, signinVals.Username)

			accountVal, err := dao.WhereUsername(username)
			if nil != err {
				core.Logger.Error(err)
				return "", jwt.ErrFailedAuthentication
			}

			if accountVal.UUID == "" {
				core.Logger.Info("account not found")
				return "", jwt.ErrFailedAuthentication
			}

			if accountVal.Password != password {
				core.Logger.Info("password not matched")
				return "", jwt.ErrFailedAuthentication
			}

			return &account{
				ID: accountVal.UUID,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			core.Logger.Debug("Authorizator")
			if v, ok := data.(*account); ok {
				dao := model.NewAccountDAO()
				accountVal, err := dao.Find(v.ID)
				if nil != err {
					return false
				}
				if accountVal.UUID == "" {
					return false
				}
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			core.Logger.Info("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		//TokenLookup:   "cookie:token",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		core.Logger.Fatal("JWT Error:" + err.Error())
	}
	_router.POST(_uri, authMiddleware.LoginHandler)
	/*
		_router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
		})
	*/
	group := _router.Group(_group)
	group.GET("/api/refresh_token", authMiddleware.RefreshHandler)
	group.Use(authMiddleware.MiddlewareFunc())
	return group
}
