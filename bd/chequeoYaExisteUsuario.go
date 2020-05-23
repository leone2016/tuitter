package bd

import (
	"context"
	"github.com/leone2016/tuitter/bundle"
	"github.com/leone2016/tuitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ChequeoYaExisteUsuario(email string)(models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(bundle.NAMEDATABASE)
	col := db.Collection(bundle.ENTITY_USUARIO)

	condicion := bson.M{"email":email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
