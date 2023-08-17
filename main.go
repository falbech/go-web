package main

import (
	"go-web/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe("localhost:8000", nil)
}
