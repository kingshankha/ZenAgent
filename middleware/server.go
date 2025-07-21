package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewServer(port string, router *http.Handler) *http.Server {

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        *router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Add coloring to log output using ANSI escape codes
	log.Printf("\033[1;32mServer is running on port %s\033[0m", s.Addr)
	log.Fatal(s.ListenAndServe())
	return s
}
