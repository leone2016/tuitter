package bd

import (
	"context"
	"github.com/leone2016/tuitter/bundle"
	"github.com/leone2016/tuitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


/*
se utiliza defer cancel, para dar de baja a ese espacio de memorio, similar a unSubscribe
 */
func InsertarRegistro( u models.Usuario )( string, bool, error ) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(bundle.NAMEDATABASE)
	col := db.Collection(bundle.ENTITY_USUARIO)

	u.Password, _= EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)


	if err != nil {
		// de acuerdo a los valores que se van a retornar en la funcion
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}