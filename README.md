# Tuitter" -


## comandos golang
`````Consola
> go run main.go
> go build main.go
...

## subir a heroku
`````Consola
> heroku login
...
...
> heroku git:remote -a tuitter (solo la primera vez)
Antes es necesario habilidar el buildpack
ir a heroku.com > tuitter > settings > boton ADD BUILDPACK > GO
esto sirva para decirle a heroku que los archivos que subo son de golang

> git push heroku master
`````
## INICIA :-)

`````Consola
> go get -t go.mongodb.org/mongo-driver/mongo
> go get go.mongodb.org/mongo-driver/mongo
> go get go.mongodb.org/mongo-driver/mongo/options
> go get go.mongodb.org/mongo-driver/bson
> go get go.mongodb.org/mongo-driver/bson/primitive
> go get golang.org/x/crypto/bcrypt
> go get github.com/gorilla/mux
> go get github.com/rs/cors
> go get github.com/dgrijalva/jwt-go
`````