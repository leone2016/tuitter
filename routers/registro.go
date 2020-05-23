package routers

import (
	"encoding/json"
	"github.com/leone2016/tuitter/bd"
	"github.com/leone2016/tuitter/bundle"
	"github.com/leone2016/tuitter/models"
	"net/http"
)

/*

	json.NewDecoder : 	decodifica lo que llega en el Body (EL BODY DE UN HTTP ES U STRING, una vez que se lee se destruye
						en memoria)
 */

func Registro( w http.ResponseWriter, r *http.Request ){

	var t models.Usuario
	err := json.NewDecoder( r.Body ).Decode( &t )
	if err != nil {
		http.Error( w, bundle.ErrorDatos+ err.Error(), http.StatusBadRequest)
		return
	}

	if len( t.Email ) == 0 {
		http.Error( w, bundle.EmailRequerido, http.StatusBadRequest)
		return
	}

	if len( t.Password ) < 6 {
		http.Error( w, bundle.MinContrasenia, http.StatusBadRequest)
		return
	}

	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error( w, bundle.ErrorExisteUsuario, http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error( w, bundle.ErrorInsertarRegistroAntes+err.Error(), http.StatusBadRequest)
	}

	if !status{
		http.Error( w, bundle.ErrorInsertarRegistroDespues, http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)

}