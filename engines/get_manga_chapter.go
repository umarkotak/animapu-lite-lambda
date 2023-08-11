package engines

import (
	"context"

	"github.com/umarkotak/animapu-lite-lambda/models"
)

func GetMangaChapter(ctx context.Context, queryParams models.QueryParams) (models.Chapter, models.Meta, error) {
	var err error
	chapter := models.Chapter{}

	mangaScrapper, err := mangaScrapperGenerator(queryParams.Source)
	if err != nil {
		return chapter, models.Meta{}, err
	}

	chapter, err = mangaScrapper.GetChapter(ctx, queryParams)
	if err != nil {
		return chapter, models.Meta{}, err
	}

	return chapter, models.Meta{}, nil
}
