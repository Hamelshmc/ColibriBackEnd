package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// Registro ,Registra el usuario en la base datos
func Registro(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(usuario.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := database.YaExisteUsuario(usuario.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	// Doble comprobación
	_, status, err := database.RegistroUsuario(usuario)
	if err != nil {
		http.Error(w, "No se registro correctamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se registro correctamente "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
