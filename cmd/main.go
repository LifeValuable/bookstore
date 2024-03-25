package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/LifeValuable/bookstore/pkg/routers"
)

func main() {
	router := mux.NewRouter()
	routers.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
}
