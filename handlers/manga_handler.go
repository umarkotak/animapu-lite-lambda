package handlers

import (
	"fmt"
	"strconv"

	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/engines"
	"github.com/umarkotak/animapu-lite-lambda/models"
	"github.com/umarkotak/animapu-lite-lambda/utils"
)

func GetLatestMangas(c *Context) {
	page, _ := strconv.ParseInt(c.Req.URL.Query().Get("page"), 10, 64)
	queryParams := models.QueryParams{
		Source: c.Param("manga_source"),
		Page:   page,
	}

	mangas, meta, err := engines.GetLatestMangas(c.Req.Context(), queryParams)
	if err != nil {
		utils.RenderResponse(c, nil, err, 422)
		return
	}

	c.Writer.Header().Set("Res-From-Cache", fmt.Sprint(meta.FromCache))
	utils.RenderResponse(c, mangas, nil, 200)
}

func GetMangaDetail(c *Context) {
	queryParams := models.QueryParams{
		Source:            c.Param("manga_source"),
		SourceID:          c.Param("manga_id"),
		SecondarySourceID: c.Req.URL.Query().Get("secondary_source_id"),
	}

	manga, meta, err := engines.GetMangaDetail(c.Req.Context(), queryParams)
	if err != nil {
		utils.RenderResponse(c, nil, err, 422)
		return
	}

	c.Writer.Header().Set("Res-From-Cache", fmt.Sprintf("%v", meta.FromCache))
	utils.RenderResponse(c, manga, nil, 200)
}

func GetMangaChapter(c *Context) {
	queryParams := models.QueryParams{
		Source:            c.Param("manga_source"),
		SourceID:          c.Param("manga_id"),
		SecondarySourceID: c.Req.URL.Query().Get("secondary_source_id"),
		ChapterID:         c.Param("chapter_id"),
	}

	chapter, meta, err := engines.GetMangaChapter(c.Req.Context(), queryParams)
	if err != nil {
		utils.RenderResponse(c, nil, err, 422)
		return
	}

	c.Writer.Header().Set("Res-From-Cache", fmt.Sprintf("%v", meta.FromCache))
	utils.RenderResponse(c, chapter, nil, 200)
}
