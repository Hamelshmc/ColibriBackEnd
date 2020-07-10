package auth

import (
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"github.com/dgrijalva/jwt-go"
)

// GeneroJWT ...
func GeneroJWT(usuario models.Usuario) (string, error) {

	miClave := []byte("HolaHamiltonMercadoCuellar")

	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellidos":        usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia":        usuario.Biografia,
		"ubicacion":        usuario.Ubicacion,
		"sitioweb":         usuario.SitioWeb,
		"_id":              usuario.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(miClave)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil

}
