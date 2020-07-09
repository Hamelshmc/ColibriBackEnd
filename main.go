package main

import (
	"log"

	"github.com/Hamelshmc/ColibriBackEnd/bd"
	"github.com/Hamelshmc/ColibriBackEnd/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la base datos")
		return
	}
	handlers.ManejoRutas()
}
