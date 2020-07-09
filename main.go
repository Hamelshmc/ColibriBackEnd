package main

import (
	"log"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/handlers"
)

func main() {
	if database.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la base datos")
		return
	}
	handlers.ManejoRutas()
}
