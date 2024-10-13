package main

import (
	"log"
	"net/http"
	"os"
)

const (
	port     = ":8002"
	certFile = "./tls/cert.pem"
	keyFile  = "./tls/key.pem"
)

func main() {

	gateway := NewAPIGateway("http://inventory-service:8001", "http://order-service:8003")
	//log.Fatal(http.ListenAndServe(":8002", gateway.Router))

	// Check cert.pem, fallback to http if not found
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Printf("Certificate file '%s' not found, running on HTTP", certFile)
		log.Fatal(http.ListenAndServe(port, gateway.Router))
		return
	}

	// Check key.pem, fallback to http if not found
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Printf("Key file '%s' not found, running on HTTP", keyFile)
		log.Fatal(http.ListenAndServe(port, gateway.Router))
		return
	}

	log.Println("Certificate and key files found, running on HTTPS")
	log.Fatal(http.ListenAndServeTLS(port, certFile, keyFile, gateway.Router))

}
