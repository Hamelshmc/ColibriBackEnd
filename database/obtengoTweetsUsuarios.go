package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ObtengoTweetsUsuarios ..
func ObtengoTweetsUsuarios(ID string, pagina int) ([]models.DevuelvoTweetsUsuarios, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("relacion")

	skip := (pagina - 1) * 20 // Paginado

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}}) // Filtro por usuario

	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid", // usuariorelacionid de la colección
			"foreignField": "userid",            // userID de la colección  tweet
			"as":           "tweet",             // Alias de la colección
		}}) // Join de collection

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"}) // aplanamos  los  datos de arrays

	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}}) // ordenamos por fecha

	condiciones = append(condiciones, bson.M{"$skip": skip}) // Cuantos elementos debe saltarse
	condiciones = append(condiciones, bson.M{"$limit": 20})  // la cantidad de tweets que voy a traer

	resultado, _ := collection.Aggregate(ctx, condiciones)
	var almacenTweets []models.DevuelvoTweetsUsuarios
	err := resultado.All(ctx, &almacenTweets)
	if err != nil {
		return almacenTweets, false
	}
	return almacenTweets, true

}
