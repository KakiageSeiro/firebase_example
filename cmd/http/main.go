package main

import (
	helloworld "gitlab.com/mc_go/firebase_example"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld.HelloHTTP)
	http.ListenAndServe(":8080", nil)
}