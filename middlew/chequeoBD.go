package middlew

import (
	"github.com/leone2016/tuitter/bd"
	"github.com/leone2016/tuitter/bundle"
	"net/http"
)

/*
ChequeoBD: un middleware tiene que recibir un handlerFunc y retornar lo mismo
esta middleware retorna una funcion anonima (no tiene nombre)
		valida que exista estado en la conexión y retorna rel response y requuest
 */

func ChequeoBD( next http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if bd.StatusConeccion()==0 {
			http.Error(w, bundle.ConexionPerdida, 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}