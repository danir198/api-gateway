package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	gateway := NewAPIGateway("http://inventory-service:8001")
	//log.Fatal(http.ListenAndServe(":8002", gateway.Router))

	certFile := "./tls/cert.pem"
	keyFile := "./tls/key.pem"

	// Check if certificate and key files exist
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Println("Certificate file not found, running on HTTP")
		log.Fatal(http.ListenAndServe(":8002", gateway.Router))
		return
	}

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Println("Key file not found, running on HTTP")
		log.Fatal(http.ListenAndServe(":8002", gateway.Router))
		return
	}

	log.Println("Certificate and key files found, running on HTTPS")
	log.Fatal(http.ListenAndServeTLS(":8002", certFile, keyFile, gateway.Router))

	log.Fatal(http.ListenAndServeTLS(":8002", certFile, keyFile, gateway.Router))

}
