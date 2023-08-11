package api

import (
	"encoding/json"
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/handlers"
)

// https://animapu-lite-lambda.vercel.app/
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "options" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	server := New()

	server.GET("/", handlers.GetHealth)

	server.GET("/mangas/:manga_source/latest", handlers.GetLatestMangas)
	server.GET("/mangas/:manga_source/detail/:manga_id", handlers.GetHealth)
	server.GET("/mangas/:manga_source/read/:manga_id/:chapter_id", handlers.GetHealth)
	server.GET("/mangas/:manga_source/search", handlers.GetHealth)

	server.Handle(w, r)
}
