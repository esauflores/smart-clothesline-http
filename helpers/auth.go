package helpers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateJWT(m map[string]any) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()

	for key, value := range m {
		claims[key] = value
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("no se pudo validar el token")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("no se pudo validar el token")
	}

	return token, nil
}

func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no se pudo obtener los claims")
	}

	return claims, nil
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get bearer token
		authHeader := c.GetHeader("Authorization")

		// Check if not contains Bearer
		if len(authHeader) < len("Bearer ") {
			Fatal(http.StatusUnauthorized, errors.New("no se ha encontrado el token"))
		}
		tokenString := authHeader[len("Bearer "):]

		// Validate token
		token, err := ValidateJWT(tokenString)
		CheckFatal(err, http.StatusUnauthorized, err)

		// Get claims
		data := token.Claims.(jwt.MapClaims)

		// Set user id
		c.Set("usuario_id", data["id"])

		// Continue
		c.Next()
	}
}
