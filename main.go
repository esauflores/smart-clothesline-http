package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/El_Salvador",
		os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), os.Getenv("DBPORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("falló la conexión a la base de datos")
	}

	r := gin.Default()
	// r.ForwardedByClientIP = true
	// r.SetTrustedProxies([]string{os.Getenv("PROXIE")})
	r.GET("/estado", func(c *gin.Context) { getEstado(c, db) })
	r.PATCH("/estado", func(c *gin.Context) { patchEstado(c, db) })

	r.Run(os.Getenv("PORT"))
}
