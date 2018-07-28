package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	log.Println("Sever running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
