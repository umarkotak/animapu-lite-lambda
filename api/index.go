package api

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", handlers.GetHealth)

	server.Handle(w, r)
}
