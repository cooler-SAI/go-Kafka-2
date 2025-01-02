package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Server is listening on port 8080\n")
	fmt.Println(fprintf)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
