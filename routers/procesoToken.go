package routers

import (
	"errors"
	"strings"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
	"github.com/dgrijalva/jwt-go"
)

// Email valor de Email usado en todas las rutas
var Email string

// IDUsuario es el devuelto del modelo, que usará en todas las rutas
var IDUsuario string

// ProcesoToken , proceso token para extraer su contenido
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("HolaHamiltonMercadoCuellar")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, ID := database.YaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, ID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}
