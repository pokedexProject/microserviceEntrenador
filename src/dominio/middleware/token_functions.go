package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	model "github.com/pokedexProject/microserviceEntrenador/dominio"
)

// Clave secreta que no se expone! es una clvve
// del servidor
var jwtKey = []byte("clave_secreta")

// Estructura del token
type Claims struct {
	MasterID string `json:"master_id"`
	jwt.StandardClaims
}

// Crea un token con el ID del usuario
func CreateToken(master *model.Entrenador) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		MasterID: master.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenStr string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "ERROR", err
	}

	if !token.Valid {
		return "ERROR", errors.New("Token inválido")
	}

	return claims.MasterID, nil
}
