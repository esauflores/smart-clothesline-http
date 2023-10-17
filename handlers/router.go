package handlers

import (
	"net/http"
	"os"
	"smart-clothesline-http/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Constants
const MODO_AUTO = "0"
const MODO_MANUAL = "1"

const ESTADO_AFUERA = "0"
const ESTADO_ADENTRO = "1"

func InitRouter() {
	router := gin.Default()
	router.Use(helpers.HTTPRecoveryHandler())

	// login routes
	router.POST("/login", Login())
	router.POST("/signup", Signup())

	// tendedero routes
	router.GET("/tendederos", helpers.AuthCheck(), GetTendederos())
	router.GET("/tendedero/:device_id", helpers.AuthCheck(), GetTendedero())
	router.PATCH("/tendedero/:device_id/:modo/:estado", helpers.AuthCheck(), PatchTendedero())

	// usuario routes
	router.GET("/usuarios", helpers.AuthCheck(), GetUsuarios)

	// random id route
	router.GET("/random_id", func(c *gin.Context) { uuid := uuid.New(); c.JSON(http.StatusOK, gin.H{"uuid": uuid.String()}) })

	// Run the server
	router.Run(os.Getenv("SERVER_PORT"))
}
