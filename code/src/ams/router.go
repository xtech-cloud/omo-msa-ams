package main

import (
	"ams/api/account"

	"github.com/gin-gonic/gin"
)

func route(_pubGroup *gin.RouterGroup, _authGroup *gin.RouterGroup) {

	account.HandleSignup("/signup", _pubGroup)
	account.HandleSignout("/signout", _authGroup)
	// /account
	{
		account.HandleCurrent("/current", _authGroup)
	}
}
