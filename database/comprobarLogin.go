package database

import (
	"github.com/Hamelshmc/ColibriBackEnd/models"
	"golang.org/x/crypto/bcrypt"
)

// ComprobarLogin ...
func ComprobarLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := YaExisteUsuario(email)
	if !encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBaseDatos := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBaseDatos, passwordBytes)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
