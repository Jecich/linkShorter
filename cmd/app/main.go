package main

import (
	"net/http"
	"shorterUrl/pkg/api"
	"shorterUrl/pkg/db"
)

type server struct {
	api *api.API
}

func main() {
	db.InitDB()
	defer db.DB.Close()
	srv := new(server)

	srv.api = api.New()

	_ = http.ListenAndServe(":8080", srv.api.Router())

}
