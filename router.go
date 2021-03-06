package main

import (
	"github.com/xtech-cloud/omo-msa-ams/api/account"

	"github.com/gin-gonic/gin"
)

func route(_pubGroup *gin.RouterGroup, _authGroup *gin.RouterGroup) {

	account.HandleSignup("/signup", _pubGroup)
	account.HandleSignout("/signout", _authGroup)
	// /account
	{
		account.HandleCurrent("/current", _authGroup)
		account.HandleUpdate("/update", _authGroup)
		account.HandleReset("/reset", _authGroup)
	}
}
