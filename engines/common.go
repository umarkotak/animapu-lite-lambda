package engines

import (
	"fmt"

	"github.com/umarkotak/animapu-lite-lambda/scrappers"
)

func mangaScrapperGenerator(mangaSource string) (scrappers.MangaScrapper, error) {
	var mangaScrapper scrappers.MangaScrapper

	switch mangaSource {
	case "mangabat":
		mangaScrapper := scrappers.NewMangabat()
		return &mangaScrapper, nil
	case "mangasee":
		mangaScrapper := scrappers.NewMangasee()
		return &mangaScrapper, nil
	}

	return mangaScrapper, fmt.Errorf("manga scrapper not found")
}
