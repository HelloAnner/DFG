package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {

	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello From Docker")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "28080"
	}

	fmt.Printf("Server listening on port %s\n", port)

	http.ListenAndServe(":"+port, nil)
}
