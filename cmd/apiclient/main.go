package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go server is up & running!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	serverAddr := ":8080"

	log.Printf("Server is running at http://localhost%s\n", serverAddr)

	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
