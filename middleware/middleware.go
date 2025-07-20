package middleware

import (
	"fmt"
	"net/http"

	"github.com/kingshankha/ZenAgent/handlers"
)

// creates and returns a router with the routes
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About page")
	})
	mux.HandleFunc("/chat", handlers.ChatPostHandler) // Register the chat handler
	return mux
}
