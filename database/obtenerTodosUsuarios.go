package database

import (
	"context"
	"time"

	"github.com/Hamelshmc/ColibriBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ObtenerTodosUsuarios , es buscador con paginado y  filtro
func ObtenerTodosUsuarios(ID string, pagina int64, textoBusquedad string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("Colibri")
	collection := database.Collection("usuarios")

	var almacenUsuarios []*models.Usuario

	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((pagina - 1) * 20)

	// utilizo un expresión regular para le de igual si es mayúscula o minúscula
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + textoBusquedad},
	}

	resultado, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return almacenUsuarios, false
	}

	var encontrado, incluir bool

	for resultado.Next(ctx) {
		var usuario models.Usuario
		err := resultado.Decode(&usuario)
		if err != nil {
			return almacenUsuarios, false
		}

		// comprabamos si existe una relación creada con la ID que recibimos
		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = usuario.ID.Hex()

		incluir = false

		encontrado, _ = ObtengoRelacion(relacion)

		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}

		//  resteo los valores que no vamos a utilizar
		if incluir {
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Banner = ""
			usuario.Email = ""

			almacenUsuarios = append(almacenUsuarios, &usuario) // guardamos
		}
	}
	err = resultado.Err()
	if err != nil {
		return almacenUsuarios, false
	}
	err = resultado.Close(ctx)
	if err != nil {
		return almacenUsuarios, false
	}
	return almacenUsuarios, true

}
