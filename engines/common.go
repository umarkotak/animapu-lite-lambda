package engines

import (
	"fmt"
)

func mangaScrapperGenerator(mangaSource string) (MangaScrapper, error) {
	var mangaScrapper MangaScrapper

	switch mangaSource {
	case "mangabat":
		mangaScrapper := NewMangabat()
		return &mangaScrapper, nil
	}

	return mangaScrapper, fmt.Errorf("manga scrapper not found")
}
