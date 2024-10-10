package main

import (
	"fmt"
	"log"
	"net/http"
)

func urlHandler(w http.ResponseWriter, r *http.Request, s string) {
	output := r.URL
	log.Printf("%v\n", output)
	fmt.Fprintf(w, s)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))

	// Handle requests to the "/" path
	http.Handle("/", fs)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		urlHandler(w, r, "This section is all about me")
	})

	log.Println("Listening on Port :8008****")

	log.Fatal(http.ListenAndServe(":8008", nil))
}
