package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Estado struct {
	Modo   int `json:"modo" binding:"required"`
	Estado int `json:"estado" binding:"required"`
}

func getEstado(c *gin.Context) {
	c.JSON(200, gin.H{
		"exito": "Llamada exitosa",
	})
}

func patchEstado(c *gin.Context) {
	var estado Estado

	if err := c.ShouldBindBodyWith(&estado, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"error": "Error, llamada inválida por parámetros json",
		})
		return
	}

	cmd := exec.Command("particle", "help")

	out, err := cmd.Output()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error, no se pudo ejecutar: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Llamada exitosa: " + string(out),
	})

}
