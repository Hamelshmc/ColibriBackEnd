package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroUsuario(usuario models.Usuario) (string, bool, error) {

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := collection.InsertOne(context, usuario)

	if err != nil {
		return "", false, err
	}

	ObjiD, _ := result.InsertedID.(primitive.ObjectID)
	return ObjiD.String(), true, nil
}
