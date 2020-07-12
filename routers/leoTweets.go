package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

/*LeoTweets Leo los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	// recogemos el parametro  que nos viene en la url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "falta el parámetro id", http.StatusBadRequest)
		return
	}

	// Comprobar antes de que vino algo para luego convertirlo a numero
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "falta el parámetro página", http.StatusBadRequest)
		return
	}
	//  lo convertimos a un numero
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "falta el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(pagina)

	respuesta, correcto := database.BuscoTweets(ID, page)
	if !correcto {
		http.Error(w, "no se pudo leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(respuesta)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
