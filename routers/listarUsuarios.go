package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

// ListarUsuarios ...
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {

	tipo := r.URL.Query().Get("tipo")
	pagina := r.URL.Query().Get("pagina")
	busqueda := r.URL.Query().Get("busqueda")

	// Comprobar antes de que vino algo para luego convertirlo a numero
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "falta el parámetro página", http.StatusBadRequest)
		return
	}
	//  lo convertimos a un numero
	paginaTemporal, err := strconv.Atoi(pagina)
	if err != nil {
		http.Error(w, "falta el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(paginaTemporal)

	resultado, estado := database.ObtenerTodosUsuarios(IDUsuario, page, busqueda, tipo)
	if !estado {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resultado)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
