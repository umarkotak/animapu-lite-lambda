package engines

import (
	"context"

	"github.com/umarkotak/animapu-lite-lambda/models"
)

func GetSearchManga(ctx context.Context, queryParams models.QueryParams) ([]models.Manga, models.Meta, error) {
	mangas := []models.Manga{}
	var err error

	mangaScrapper, err := mangaScrapperGenerator(queryParams.Source)
	if err != nil {
		return mangas, models.Meta{}, err
	}

	mangas, err = mangaScrapper.GetSearch(ctx, queryParams)
	if err != nil {
		return mangas, models.Meta{}, err
	}

	return mangas, models.Meta{}, nil
}
