package models

import (
	"fmt"
	"strings"
)

func (m *Manga) GenerateLatestChapter() {
	if len(m.Chapters) > 0 {
		m.LatestChapterID = m.Chapters[0].ID
		m.LatestChapterNumber = m.Chapters[0].Number
		m.LatestChapterTitle = m.Chapters[0].Title
	}
}

func (m *Manga) GetUniqueKey() string {
	return fmt.Sprintf(
		"%v:%v:%v:%v", m.Source, m.SourceID, m.SecondarySource, m.SecondarySourceID,
	)
}

func (m *Manga) GetFbUniqueKey() string {
	return strings.ReplaceAll(fmt.Sprintf(
		"%v:%v:%v:%v", m.Source, m.SourceID, m.SecondarySource, m.SecondarySourceID,
	), ".", "dot")
}
