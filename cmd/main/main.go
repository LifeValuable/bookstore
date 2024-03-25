package main

import (
	"log"
	"net/http"

	"github.com/LifeValuable/bookstore/pkg/routers"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	router := mux.NewRouter()
	routers.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
