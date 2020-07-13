package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

// LeoTweetsUsuarios ...
func LeoTweetsUsuarios(w http.ResponseWriter, r *http.Request) {

	// Comprobar antes de que vino algo
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "falta el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "falta el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := database.ObtengoTweetsUsuarios(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
