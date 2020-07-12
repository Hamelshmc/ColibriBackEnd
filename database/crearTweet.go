package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*CrearTweet graba el Tweet en la base de datos*/
func CrearTweet(tweet models.CreoTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("tweet")

	insertoTweet := bson.M{
		"userid":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}
	result, err := collection.InsertOne(ctx, insertoTweet)
	if err != nil {
		return "", false, err
	}

	// Extrae la clave del ultimo campo insertado
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
