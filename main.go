package main

import (
	"log"
	"net/http"

	"github.com/piazin/store-go/routes"
)

func main() {
	routes.LoadRoutes()
	log.Print("Listening on 8000")
	http.ListenAndServe(":8000", nil)
}

