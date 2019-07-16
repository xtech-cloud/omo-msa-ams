package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/xtech-cloud/omo-msa-ams/auth"
	"github.com/xtech-cloud/omo-msa-ams/core"
	"github.com/xtech-cloud/omo-msa-ams/model"
)

func main() {
	core.SetupLogger()

	core.Logger.Info("initialize model")

	model.SetupEnv()
	model.AutoMigrateDatabase()

	httpAddrArg := os.Getenv("AMS_HTTP_ADDR")
	if "" == httpAddrArg {
		httpAddrArg = ":80"
	}
	auth.ISS = os.Getenv("AMS_ISS")
	if "" == auth.ISS {
		auth.ISS = "ams"
	}
	auth.Secret = os.Getenv("AMS_Secret")
	if "" == auth.Secret{
		auth.Secret = "ams-secret"
	}

	router := gin.Default()

	// 跨域调用
	corsConf := cors.Config{
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "access-control-allow-headers", "origin", "authorization", "access-control-allow-origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	corsConf.AllowAllOrigins = true
	router.Use(cors.New(corsConf))

	// 不需要认证
	pubGroup := router.Group("/ams")
	// 需要认证
	authGroup := auth.BindAuthHandler(router, "/ams/signin", "/ams/auth")

	route(pubGroup, authGroup)

	core.Logger.Infof("serve at %v", httpAddrArg)
	router.Run(httpAddrArg)
}
