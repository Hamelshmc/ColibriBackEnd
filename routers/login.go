package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/auth"
	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// Login , loguearse en la app  y devuelve un token
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Usuario y Contraseña invalida"+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "Email es obligatorio", 400)
		return
	}

	documento, existe := database.ComprobarLogin(usuario.Email, usuario.Password)
	if !existe {
		http.Error(w, "Usuario y Contraseña invalida", 400)
		return
	}
	jwtKey, err := auth.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Hubo un error al generar al token ", 400)
		return
	}

	respuesta := models.RepuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	eexpirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: eexpirationTime,
	})

}
