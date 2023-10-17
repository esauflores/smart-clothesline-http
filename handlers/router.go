package handlers

import (
	"log"
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
	router.Use(HTTPRecoveryHandler())

	// login routes
	router.POST("/login", func(c *gin.Context) { Login(c) })
	router.POST("/signup", func(c *gin.Context) { Signup(c) })

	// tendedero routes
	router.GET("/tendederos", func(c *gin.Context) { GetTendederos(c) })
	router.GET("/tendedero/:device_id", func(c *gin.Context) { GetTendedero(c) })
	router.PATCH("/tendedero/:device_id/:modo/:estado", func(c *gin.Context) { PatchTendedero(c) })

	// usuario routes
	router.GET("/usuarios", func(c *gin.Context) { GetUsuarios(c) })

	// random id route
	router.GET("/random_id", func(c *gin.Context) { uuid := uuid.New(); c.JSON(http.StatusOK, gin.H{"uuid": uuid.String()}) })

	// Run the server
	router.Run(os.Getenv("SERVER_PORT"))
}

func HTTPRecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if httpError, ok := err.(helpers.HTTPError); ok {
					// Log the HTTP error message
					log.Println("HTTP Error:", httpError)

					// Return the HTTP error as the response
					c.JSON(httpError.StatusCode, gin.H{"error": httpError.Message})
				} else {
					// Log the panic message
					log.Println("Panic:", err)

					// For other panics, return a generic error response
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				}

				// Abort the current request
				c.Abort()
			}
		}()

		// Continue processing the request
		c.Next()
	}
}
