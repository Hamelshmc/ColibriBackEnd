package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Hamelshmc/ColibriBackEnd/middlewares"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// ManejoRutas  configuro el puerto  y lo pongo en escucha
func ManejoRutas() {
	router := mux.NewRouter() // mux captura el HTTP

	router.HandleFunc("/registro", middlewares.ChequeoBaseDatos(routers.Resgistro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
