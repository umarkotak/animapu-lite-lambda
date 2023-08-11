package api

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/handlers"
	"github.com/umarkotak/animapu-lite-lambda/utils"
)

// https://animapu-lite-lambda.vercel.app/
func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(CORSMiddleware)

	server.GET("/", handlers.GetHealth)

	server.GET("/mangas/:manga_source/latest", handlers.GetLatestMangas)
	server.GET("/mangas/:manga_source/detail/:manga_id", handlers.GetHealth)
	server.GET("/mangas/:manga_source/read/:manga_id/:chapter_id", handlers.GetHealth)
	server.GET("/mangas/:manga_source/search", handlers.GetHealth)

	server.Handle(w, r)
}

func CORSMiddleware(c *Context) {
	if c.Req.Method == "OPTIONS" {
		utils.RenderResponse(c, map[string]interface{}{}, nil, 200)
		return
	}
	c.Next()
}
