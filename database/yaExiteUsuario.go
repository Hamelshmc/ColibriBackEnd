package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

//YaExisteUsuario , comprobar si existe el usuario en la base datos
func YaExisteUsuario(email string) (models.Usuario, bool, string) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := collection.FindOne(context, condicion).Decode(resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
