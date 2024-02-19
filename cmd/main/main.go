package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harshavallabhaneni56/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhost:8000", r))

}
