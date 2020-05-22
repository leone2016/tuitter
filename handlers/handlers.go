package handlers

import(
	"github.com/rs/cors@v1.7.0"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)
/**
*  Cuando se use la api, va entrar en esta funciòn
*  mux: captura el http y le va dar un manejo al response y al request que viene en el llamado
 de la api y va procesar, para saber si en el body/header del llamado hay informaciòn y
 va mandar la respuesta al navegador, por ejemplo decirle a al frontEnd que el registro se creo (201) o
 si hubo un error en el servicio (404)
*/
func Config(){

	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "1991"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler ))
}