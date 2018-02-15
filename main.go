package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	SetUpAPICalls()

	svr := http.Server{
		Addr:           ":8080",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 8175, // if it's good enough for Apache, it's good enough for me
	}
	fmt.Println("Serving...")
	svr.ListenAndServe()
}
