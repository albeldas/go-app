package main

import (
	"fmt"
	"log"
	"net/http"
)
// Sample app
func handler(w http.ResponseWriter, r *http.Request) {
	title := "Allcloud DOE Demo app canary develop v1.4.6"
	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
