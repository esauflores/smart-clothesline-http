package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

func main() {
	r := gin.Default()

	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{os.Getenv("PROXIE")})

	r.GET("/estado", func(c *gin.Context) { getEstado(c) })
	r.PATCH("/estado", func(c *gin.Context) { patchEstado(c) })

	r.Run(os.Getenv("PORT"))
}
