package middlewares

import (
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

// ChequeoBaseDatos i
func ChequeoBaseDatos(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ChequeoConexion() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
