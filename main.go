package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luismoroco/gorilla-mux-api/db"
	"github.com/luismoroco/gorilla-mux-api/routes"
)

func main() {
	db.SetupConnectionWithPostgreSQL()
	rootRouter := mux.NewRouter()

	rootRouter.HandleFunc("/", routes.IndexRootHandler)

	http.ListenAndServe(":3000", rootRouter)
}
