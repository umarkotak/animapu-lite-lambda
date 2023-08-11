package engines

import (
	"context"

	"github.com/umarkotak/animapu-lite-lambda/models"
)

func GetMangaDetail(ctx context.Context, queryParams models.QueryParams) (models.Manga, models.Meta, error) {
	manga := models.Manga{}
	var err error

	mangaScrapper, err := mangaScrapperGenerator(queryParams.Source)
	if err != nil {
		return manga, models.Meta{}, err
	}

	manga, err = mangaScrapper.GetDetail(ctx, queryParams)
	if err != nil {
		return manga, models.Meta{}, err
	}

	return manga, models.Meta{}, nil
}
