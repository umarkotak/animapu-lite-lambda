package models

type (
	QueryParams struct {
		Source            string `json:"source"`
		SourceID          string `json:"source_id"`
		SecondarySourceID string `json:"secondary_source_id"`
		Page              int64  `json:"page"`
		ChapterID         string `json:"chapter_id"`
		Title             string `json:"title"`
	}

	Manga struct {
		ID                  string       `json:"id"`
		SourceID            string       `json:"source_id"`
		SecondarySourceID   string       `json:"secondary_source_id"`
		Source              string       `json:"source"`
		SecondarySource     string       `json:"secondary_source"`
		Title               string       `json:"title"`
		Description         string       `json:"description"`
		Genres              []string     `json:"genres"`
		Status              string       `json:"status"`
		Rating              string       `json:"rating"`
		LatestChapterID     string       `json:"latest_chapter_id"`
		LatestChapterNumber float64      `json:"latest_chapter_number"`
		LatestChapterTitle  string       `json:"latest_chapter_title"`
		ChapterPaginated    bool         `json:"chapter_paginated"`
		Chapters            []Chapter    `json:"chapters"`
		CoverImages         []CoverImage `json:"cover_image"`
		PopularityPoint     int64        `json:"popularity_point"`
		ReadCount           int64        `json:"read_count"`
		Star                bool         `json:"star"`
		LastChapterRead     float64      `json:"last_chapter_read"`
		LastLink            string       `json:"last_link"`
		Weight              int64        `json:"weight"`
		FollowCount         int64        `json:"follow_count"`
	}

	CoverImage struct {
		Index     int64    `json:"index"`
		ImageUrls []string `json:"image_urls"`
	}

	Chapter struct {
		ID                string         `json:"id"`
		SourceID          string         `json:"source_id"`
		Source            string         `json:"source"`
		SourceLink        string         `json:"source_link"`
		SecondarySourceID string         `json:"secondary_source_id"`
		SecondarySource   string         `json:"secondary_source"`
		Title             string         `json:"title"`
		Index             int64          `json:"index"`
		Number            float64        `json:"number"`
		ChapterImages     []ChapterImage `json:"chapter_images"`
	}

	ChapterImage struct {
		Index     int64    `json:"index"`
		ImageUrls []string `json:"image_urls"`
	}

	Meta struct {
		FromCache bool `json:"from_cache"`
	}
)
