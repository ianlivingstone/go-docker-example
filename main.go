package main

import (
        "os"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
	if port == "" {
            log.Fatal("PORT environment variable was not set")
        }

        log.Println("Attempting to listen on port: ", port)
        err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
