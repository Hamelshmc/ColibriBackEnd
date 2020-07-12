package routers

import (
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// BajaRelacion ...
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	status, err := database.EliminoRelacion(relacion)

	if err != nil {
		http.Error(w, "Ocurrió un error "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
