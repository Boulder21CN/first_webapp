package main

import (
	"fmt"
	"net/http"
)

var num int

func handler(writer http.ResponseWriter, request *http.Request) {
	num++
	fmt.Println(request.URL.Path[0:], num)
	fmt.Fprintf(writer, "Hello world,%s,%v!", request.URL.Path[1:], num)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
