package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fprintf)
}

func main() {
	fmt.Printf("Server is listening on port 8080\n")

	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		fmt.Println("Server is listening on port 8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, err) {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
