package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

/*CreoTweet permite grabar el tweet en la base de datos */
func CreoTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	creoTweet := models.CreoTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := database.CrearTweet(creoTweet)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, inténtenlo de nuevo"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
