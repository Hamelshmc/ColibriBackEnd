package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

// VerPerfil extrae los valores de perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	// recogemos el parametro  que nos viene en la url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := database.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "No encontro el perfil "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(perfil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
