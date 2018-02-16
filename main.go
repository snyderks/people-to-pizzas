package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	SetUpAPICalls()
	Conf, _ = ConfigRead("config.json")
	svr := http.Server{
		Addr:           Conf.HTTPPort,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 8175, // if it's good enough for Apache, it's good enough for me
	}
	fmt.Println("Serving on", Conf.HTTPPort+"...")
	svr.ListenAndServe()
}
