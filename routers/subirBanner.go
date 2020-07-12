package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Hamelshmc/ColibriBackEnd/database"
	"github.com/Hamelshmc/ColibriBackEnd/models"
)

// SubirBanner  al servidor
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	// Lo procesamos como un formulario
	imagenBanner, handler, _ := r.FormFile("banner")

	var extension = strings.Split(handler.Filename, ".")[1]

	//Nombre del archivo que suban poner el ID de usuario para identificar su avatar
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	guardarArchivo, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		http.Error(w, "Error al subir el  banner"+err.Error(), http.StatusBadRequest)
		return
	}

	// guardo el  Archivo
	_, err = io.Copy(guardarArchivo, imagenBanner)
	if err != nil {
		http.Error(w, "Error al copiar  la banner en la ruta"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario

	usuario.Banner = IDUsuario + "." + extension

	status, err := database.ActualizoPerfil(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al guardar   la imagen en la base datos"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
