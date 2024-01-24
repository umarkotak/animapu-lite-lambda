package scrappers

import (
	"context"

	"github.com/umarkotak/animapu-lite-lambda/models"
)

type (
	MangaScrapper interface {
		GetHome(ctx context.Context, queryParams models.QueryParams) ([]models.Manga, error)
		GetDetail(ctx context.Context, queryParams models.QueryParams) (models.Manga, error)
		GetSearch(ctx context.Context, queryParams models.QueryParams) ([]models.Manga, error)
		GetChapter(ctx context.Context, queryParams models.QueryParams) (models.Chapter, error)
	}
)
