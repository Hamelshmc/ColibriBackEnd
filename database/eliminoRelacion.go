package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// EliminoRelacion borra la relacion en la base de datos
func EliminoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("relacion")

	_, err := collection.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
