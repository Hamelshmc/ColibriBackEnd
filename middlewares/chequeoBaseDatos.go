package middlewares

import (
	"net/http"

	"github.com/Hamelshmc/ColibriBackEnd/database"
)

// ChequeoBaseDatos , comprobar si la base de datos esta activa para que continue la ejecución
func ChequeoBaseDatos(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ChequeoConexion() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
