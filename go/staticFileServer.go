package main

import (
	"log"
	"net/http"
	h "goserver/handlers"
)

func main() {
	// Static assets for client
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.HandleFunc("/api/pdf", h.MakePdf)
	http.HandleFunc("/api/pdf/upload", h.UploadText)

	// TLS
	port := ":443"
	log.Println("Listening on port", port)
	go log.Fatal(http.ListenAndServeTLS(port, "certificate.pem", "key.pem", nil))
}