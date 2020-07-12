package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// CreoRelacion  añadirla a la base de datos
func CreoRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("relacion")

	_, err := collection.InsertOne(ctx, relacion)
	if err != nil {
		return false, err
	}

	return true, nil

}
