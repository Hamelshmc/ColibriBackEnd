package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

/*ObtenerBanner enviá el Avatar a */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Falta el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := database.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Banner no encontrado", http.StatusBadRequest)
		return
	}

	// aquí es donde copio la imagen en el HTTP
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el banner", http.StatusBadRequest)
	}
}
