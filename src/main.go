package main

import (
	"greeting"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", greeting.Hello)
	http.HandleFunc("/hello-post", greeting.HelloPost)

	log.Println("Sever running on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
