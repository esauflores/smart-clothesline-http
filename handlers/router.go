package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = nil

func Init() {
	Router = gin.Default()

	Router.Use(RecoveryHandlers())

	// r.ForwardedByClientIP = true
	// r.SetTrustedProxies([]string{os.Getenv("PROXIE")})
	Router.GET("/tendederos", func(c *gin.Context) { GetTendederos(c) })
	Router.PATCH("/tendedero/:modo/:estado", func(c *gin.Context) { PatchTendedero(c) })

	Router.Run(os.Getenv("PORT"))
}
