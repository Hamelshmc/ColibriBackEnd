package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*YaExisteUsuario recibe un email de parámetro y chequea si ya está en la BD */
func YaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
