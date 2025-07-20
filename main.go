package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kingshankha/ZenAgent/middleware"
)

func main() {
	router := middleware.NewRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
