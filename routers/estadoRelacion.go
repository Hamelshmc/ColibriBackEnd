package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// EstadoRelacion comprobar si hay relacion entre los usuarios
func EstadoRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	var respuesta models.RespuestaEstadoRelacion

	status, err := database.ObtengoRelacion(relacion)
	if err != nil || !status {
		respuesta.Status = false
	} else {
		respuesta.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
