package main

import (
	"log"
	"net/http"
)

func main() {
	gateway := NewAPIGateway("http://inventory-service:8001")
	log.Fatal(http.ListenAndServe(":8002", gateway.Router))
}
