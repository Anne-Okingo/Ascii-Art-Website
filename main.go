package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	sava "ascii-art-web/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	port := ":8080"
	http.HandleFunc("/", sava.Handler)
	http.HandleFunc("/ascii-art", sava.HandleasciiArt)
	log.Printf("Server is running on http://localhost%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
