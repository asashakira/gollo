package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang Server!")
}

func main() {
	http.HandleFunc("/", handler)
	port := "8081"
	fmt.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
