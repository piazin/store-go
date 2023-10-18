package main

import (
	"log"
	"net/http"

	"github.com/piazin/store-go/routes"
)

func main() {
	routes.LoadRoutes()
	log.Print("Listening on 3000")
	http.ListenAndServe(":8000", nil)
}

