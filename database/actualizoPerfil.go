package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ActualizoPerfil permite modificar el perfil del usuario */
func ActualizoPerfil(usuario models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("usuarios")

	// Va tener una clave string  y los valores van ser  una interface
	registro := make(map[string]interface{})
	// Comprobamos los campos que recibimos
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}
	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}
	registro["fechaNacimiento"] = usuario.FechaNacimiento
	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}
	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}
	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}
	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}
	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}

	updatePerfil := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	// Indicando a que usuario voy actualizar sus datos
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filtro, updatePerfil)
	if err != nil {
		return false, err
	}

	return true, nil
}
