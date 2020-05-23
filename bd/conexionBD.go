package bd

import (
	"context"
	"github.com/leone2016/tuitter/bundle"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:BBjkPXa0AXScjpv1VCo7@cluster0-2kquz.mongodb.net/tuitter?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println(bundle.ConexionExitosa)
	return client
}

/*ChequeoConnection es el Ping a la BD */
func StatusConeccion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
