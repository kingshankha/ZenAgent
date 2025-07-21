package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kingshankha/ZenAgent/api/chat"
)

// creates and returns a router with the routes
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page!")
	})

	log.Printf("\033[1;34m Home : http://localhost:%s/\033[0m", "8080")

	mux.HandleFunc("/chat", chat.ChatPostHandler) // Register the chat handler

	log.Printf("\033[1;34m Chat : http://localhost:%s/chat\033[0m", "8080")

	return mux
}
