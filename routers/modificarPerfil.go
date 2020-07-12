package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

/*ModificarPerfil actualiza  el perfil de usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = database.ActualizoPerfil(usuario, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar  actualizar  inténtelo de  nuevo "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado actualizar ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
