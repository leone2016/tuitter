package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
/* Usuario es el modelo de entidad de la base de datos MongoDB*/
type Usuario struct {
	ID 					primitive.ObjectID	`bson:"_id, omitempty" json:"id"`
	Nombre 				string				`bson:"nombre" json:"id,omitempty"`
	Apellidos 			string				`bson:"apellidos" json:"id,omitempty"`
	FechaNacimiento 	time.Time			`bson:"fechaNacimiento" json:"id,omitempty"`
	Email 				string				`bson:"email" json:"id,omitempty"`
	Password 			string				`bson:"password" json:"id,omitempty"`
	Avatar 				string				`bson:"avatar" json:"id,omitempty"`
	Banner 				string				`bson:"banner" json:"id,omitempty"`
	Biografia 			string				`bson:"biografia" json:"id,omitempty"`
	Ubicacion 			string				`bson:"ubicacion" json:"id,omitempty"`
	SitioWeb 			string				`bson:"sitioWeb" json:"id,omitempty"`
}