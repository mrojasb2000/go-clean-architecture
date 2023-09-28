package main

import (
	"Customers/pkg/infrastructure/entrypoints"
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe("localhost:3000", entrypoints.Routes()))
}
