package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("My HTTP server learning")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This my WebSITE: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
