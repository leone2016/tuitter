package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string)(string, error){
	costo := 8
	//slice tipo byte []byte
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}