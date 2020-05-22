package main
import(
	"github.com/leone2016/tuitter/handlers"
	"github.com/leone2016/tuitter/bd"
	"log"
)
func main(){
	if bd.StatusConeccion() == 0 {
		log.Fatal("Sin conexiòn en la BD")
		return
	}
	handlers.Config()
}
