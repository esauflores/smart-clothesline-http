package handlers

import (
	"errors"
	"net/http"
	"smart-clothesline-http/helpers"
	"smart-clothesline-http/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get body
		var body struct {
			Nombres    string `json:"nombres"`
			Apellidos  string `json:"apellidos"`
			Email      string `json:"email"`
			Contrasena string `json:"contrasena"`
		}
		err := c.ShouldBindJSON(&body)
		helpers.CheckFatal(err, http.StatusBadRequest, errors.New("no se pudo obtener los datos del usuario"))

		// Hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Contrasena), bcrypt.DefaultCost)
		helpers.CheckFatal(err, http.StatusInternalServerError, errors.New("no se pudo encriptar la contraseña"))

		// Create user
		helpers.OpenDBConnection()
		defer helpers.CloseDBConnection()

		result := helpers.DB.Create(
			&models.Usuario{
				Id:         uuid.New().String(),
				Nombres:    body.Nombres,
				Apellidos:  body.Apellidos,
				Email:      body.Email,
				Contrasena: string(hash),
			},
		)
		helpers.CheckFatal(result.Error, http.StatusInternalServerError, errors.New("no se pudo crear el usuario"))

		// Response handling
		c.JSON(http.StatusOK, gin.H{"message": "usuario creado"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get body
		var body struct {
			Email      string `json:"email"`
			Contrasena string `json:"contrasena"`
		}
		err := c.ShouldBindJSON(&body)
		helpers.CheckFatal(err, http.StatusBadRequest, errors.New("no se pudo obtener los datos del usuario"))

		// Query handling
		helpers.OpenDBConnection()
		defer helpers.CloseDBConnection()

		var usuario models.Usuario
		result := helpers.DB.Where("email = ?", body.Email).First(&usuario)
		helpers.CheckFatal(result.Error, http.StatusUnauthorized, errors.New("usuario o contraseña incorrecta"))

		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(body.Contrasena))
		helpers.CheckFatal(err, http.StatusUnauthorized, errors.New("usuario o contraseña incorrecta"))

		// Create JWT
		tokenString, err := helpers.CreateJWT(map[string]any{"id": usuario.Id})
		helpers.CheckFatal(err, http.StatusInternalServerError, errors.New("no se pudo crear el token"))

		// Response handling
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
