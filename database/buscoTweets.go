package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BuscoTweets lee los tweets de un perfil de manera paginada
func BuscoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("tweet")

	var almacenTweets []*models.DevuelvoTweets

	// Para buscar en la base datos
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)                               // Cuantos tweets me va cargar
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // Ordenarlos por fecha de forma descendente
	opciones.SetSkip((pagina - 1) * 20)                 // aquí esta la magia del paginado

	contenido, err := collection.Find(ctx, condicion, opciones)
	if err != nil {
		return almacenTweets, false
	}
	//
	for contenido.Next(context.TODO()) {

		var fila models.DevuelvoTweets
		err := contenido.Decode(&fila)
		if err != nil {
			return almacenTweets, false
		}
		// Agremos
		almacenTweets = append(almacenTweets, &fila)
	}
	return almacenTweets, true
}
