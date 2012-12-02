package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and serve error", err.Error())
	}
}
