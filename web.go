package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "<!doctype html><html><body><h1>My Go Web App</h1><h2>Hello from New York!</h2></body></html>")
}
