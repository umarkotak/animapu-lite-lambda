package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/handlers"
	"github.com/umarkotak/animapu-lite-lambda/utils"
)

// https://animapu-lite-lambda.vercel.app/
func Handler(w http.ResponseWriter, r *http.Request) {
	logrus.SetFormatter(&utils.Formatter{})

	if r.Method == "OPTIONS" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Animapu-User-Uid, X-Visitor-Id, X-From-Path, Server, X-Vercel-Cache, X-Vercel-Id",
		)
		json.NewEncoder(w).Encode(map[string]interface{}{})
		return
	}

	server := New()

	server.GET("/", handlers.GetHealth)

	server.GET("/mangas/:manga_source/latest", handlers.GetLatestMangas)
	server.GET("/mangas/:manga_source/detail/:manga_id", handlers.GetMangaDetail)
	server.GET("/mangas/:manga_source/read/:manga_id/:chapter_id", handlers.GetMangaChapter)
	server.GET("/mangas/:manga_source/search", handlers.GetSearchManga)

	server.Handle(w, r)
}
