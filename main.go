package main

import (
	"github.com/leone2016/tuitter/bd"
	"github.com/leone2016/tuitter/handlers"
	"log"
)

func main() {
	if bd.StatusConeccion() == 0 {
		log.Fatal("Sin conexiòn en la BD")
		return
	}
	handlers.Config()
}
