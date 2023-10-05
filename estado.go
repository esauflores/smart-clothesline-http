package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type EstadoJSON struct {
	Modo   int `json:"modo" binding:"required"`
	Estado int `json:"estado" binding:"required"`
}

func getEstado(c *gin.Context, db *gorm.DB) {
	c.JSON(200, gin.H{
		"exito": "Llamada exitosa",
	})
}

func patchEstado(c *gin.Context, db *gorm.DB) {
	var estado EstadoJSON

	if err := c.ShouldBindBodyWith(&estado, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"error": "Error, llamada inválida por parámetros json",
		})
		return
	}

	const devicePath = "/dev/ttyACM0"
	text := strconv.Itoa(estado.Modo*2 + estado.Estado)

	file, err := os.OpenFile(devicePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error, falló la apertura del puerto serial",
		})
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error, falló la escritura en el puerto serial",
		})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Llamada exitosa: " + string(out),
	})

}
