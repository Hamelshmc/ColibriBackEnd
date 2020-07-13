package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConexion variable global  que es una función
var MongoConexion = ConectarBaseDatos()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://admin:1234@colibri.dvvo4.mongodb.net/Colibri?retryWrites=true&w=majority")

// ConectarBaseDatos  nos conecta con la base datos de MongoDB
func ConectarBaseDatos() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions) // Context-->  para comunicar información entre ejecución en ejecución.
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	// Para saber si la base datos esta disponible
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la base datos!!")
	return client
}

// ChequeoConexion me devuelve 0 si hubo un error y un 1 si fue exitosa
func ChequeoConexion() int {
	err := MongoConexion.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
