package handlers

import (
	"errors"
	"net/http"
	"os"
	"smart-clothesline-http/helpers"
	"smart-clothesline-http/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTendederos() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Variable declaration
		var tendederos []models.Tendedero

		// Query handling
		helpers.OpenDBConnection()
		defer helpers.CloseDBConnection()

		result := helpers.DB.Where("usuario_id = ?", c.Keys["usuario_id"]).Find(&tendederos)
		helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los tendederos"))

		// Response handling
		c.JSON(http.StatusOK, tendederos)
	}
}

func GetTendedero() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Variable declaration
		var tendedero models.Tendedero

		// Param handling
		device_id, err := helpers.GetURLParam(c, "device_id")
		helpers.CheckFatal(err, http.StatusBadRequest, err)

		// Query handling
		helpers.OpenDBConnection()
		defer helpers.CloseDBConnection()

		result := helpers.DB.Where("usuario_id = ? AND id = ?", c.Keys["usuario_id"], device_id).First(&tendedero)
		helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo obtener los tendederos"))

		// Response handling
		c.JSON(http.StatusOK, tendedero)
	}
}

func PatchTendedero() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Param handling
		device_id, err := helpers.GetURLParam(c, "device_id")
		helpers.CheckFatal(err, http.StatusBadRequest, err)

		modo, err := helpers.GetURLParam(c, "modo")
		helpers.CheckFatal(err, http.StatusBadRequest, err)

		estado, err := helpers.GetURLParam(c, "estado")
		helpers.CheckFatal(err, http.StatusBadRequest, err)

		// Validation
		if modo != MODO_AUTO && modo != MODO_MANUAL {
			helpers.Fatal(http.StatusBadRequest, errors.New("modo invalido"))
		}

		if estado != ESTADO_AFUERA && estado != ESTADO_ADENTRO {
			helpers.Fatal(http.StatusBadRequest, errors.New("estado invalido"))
		}

		// File handling
		modo_int, _ := strconv.Atoi(modo)
		estado_int, _ := strconv.Atoi(estado)

		message := strconv.Itoa(modo_int*2 + estado_int)
		helpers.WriteToFile(os.Getenv("DEVICE_PATH"), message)

		// Query handling
		helpers.OpenDBConnection()
		defer helpers.CloseDBConnection()

		var result *gorm.DB

		// Create evento
		result = helpers.DB.Create(&models.Evento{
			Id:          uuid.New().String(),
			Estado:      estado == ESTADO_ADENTRO,
			Modo:        modo == MODO_MANUAL,
			TendederoId: device_id,
		})
		helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo crear el evento"))

		// Update the tendedero
		data := map[string]any{"estado": estado == ESTADO_ADENTRO, "modo": modo == MODO_MANUAL}
		result = helpers.DB.Model(&models.Tendedero{}).Where("usuario_id = ? AND id = ?", c.Keys["usuario_id"], device_id).Updates(data)
		helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo actualizar el tendedero"))

		// Response handling
		c.JSON(http.StatusOK, gin.H{"mensaje": "Actualizado exitoso"})
	}
}
