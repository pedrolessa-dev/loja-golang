package main

import (
	"net/http"

	"github.com/pedrolessa-dev/loja-golang/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
