package bd

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
/**
variables que empiezan en Mayúsculas se exportan
varibles que empiezan en minúsculas son de uso local
 */
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:BBjkPXa0AXScjpv1VCo7@cluster0-2kquz.mongodb.net/tuitter?authSource=tuitter")

/*
* conectarBD() es la función que me permite conectar con MONGODB
* retorna un objeto de tipo mongo.cliente
* en Golang no esta bien visto, que exista variables globales, archivos de funciones,
dentro del llamado de una api go ha creado los CONTEXTOS, que son un espacio en memoria,
con lo cual se va compartir cosas como setear un contexto de ejecución.

* se puede ejecutar una instrucción mongo y dentro del contexto decirle que no puede  superar los 30 seg

* los context sirven para comunicar información entre ejecución y ejecución, ademas nos permite setear una serie
de valores, como por ejemplo un timeOut, con esto se evita que cuando exita un error por el lado de la base la aplicación se cuelgue
context.TODO(): se le dice a go que se conecte a mongo sin ningun tipo de timeOut
*/
func conectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) // aqui no se usa  := ya que solo se lo utiliza cuando se crea la variable
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con MongoDB")
	return client
}
/**
StatusConeccion() es el ping a la base de datos
 */
func StatusConeccion() int{
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1 // si no hubo error
}
