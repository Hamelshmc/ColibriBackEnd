package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Hamelshmc/ColibriBackEnd/middlewares"
	"github.com/Hamelshmc/ColibriBackEnd/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// ManejoRutas  configuro el puerto  y lo pongo en escucha
func ManejoRutas() {
	router := mux.NewRouter() // mux captura el HTTP

	router.HandleFunc("/registro", middlewares.ChequeoBaseDatos(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBaseDatos(routers.Login)).Methods("POST")
	router.HandleFunc("/tweet", middlewares.ChequeoBaseDatos(middlewares.ValidoJWT(routers.CreoTweet))).Methods("POST")

	router.HandleFunc("/verperfil", middlewares.ChequeoBaseDatos(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/leotweets", middlewares.ChequeoBaseDatos(middlewares.ValidoJWT(routers.LeoTweets))).Methods("GET")

	router.HandleFunc("/modificarperfil", middlewares.ChequeoBaseDatos(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	router.HandleFunc("/eliminarTweet", middlewares.ChequeoBaseDatos(middlewares.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
