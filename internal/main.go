package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey yo, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("8080 portu dinleniyor..")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
