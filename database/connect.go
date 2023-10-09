package database

import (
	"fmt"
	"log"
	"os"
	"smart-clothesline-http/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB = nil

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/El_Salvador",
		os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"), os.Getenv("DBPORT"))
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("No se pudo conectar con la base de datos")
		return
	}

	Connection = conn
	conn.AutoMigrate(&models.Tendedero{})
	log.Println("Conexión exitosa con la base de datos")

}
