package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// ManejoRutas  configuro el puerto  y lo pongo en escucha
func ManejoRutas() {
	router := mux.NewRouter() // mux captura el HTTP

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
