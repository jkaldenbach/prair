package main

import (
	"log"
	"net/http"

	m "kaldenbach.design/prayr/models"
	"kaldenbach.design/prayr/routes"
)

func main() {
	db := m.InitDb()
	defer db.Close()

	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":7770", router))
	log.Printf("Server running on port 7770")
}
