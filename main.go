package main

import (
	"log"

	"github.com/kingshankha/ZenAgent/middleware"
)

func main() {
	router := middleware.NewRouter()

	s := middleware.NewServer("8080", &router)

	log.Fatal(s.ListenAndServe())
}
