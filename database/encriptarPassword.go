package database

import "golang.org/x/crypto/bcrypt"

//EncriptarPassword , encriptamos la password
func EncriptarPassword(password string) (string, error) {
	cost := 8 // Algoritmo , como el nivel de seguridad de la password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
