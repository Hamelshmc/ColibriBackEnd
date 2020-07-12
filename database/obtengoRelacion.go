package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ObtengoRelacion ...
func ObtengoRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         relacion.UsuarioID,
		"usuariorelacionid": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
