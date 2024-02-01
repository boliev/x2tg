package db

import (
	"github.com/boliev/x2tg/src/domain"
)

type SourceRepository struct {
}

func NewSourceRepository() *SourceRepository {
	return &SourceRepository{}
}

func (sr *SourceRepository) GetActive() []*domain.Source {
	sources := []*domain.Source{}
	sources = append(sources, &domain.Source{
		Resource: "reddit",
		URL:      "https://www.reddit.com/r/golang/",
	})
	sources = append(sources, &domain.Source{
		Resource: "reddit",
		URL:      "https://www.reddit.com/r/StartledCats/",
	})
	sources = append(sources, &domain.Source{
		Resource: "reddit",
		URL:      "https://www.reddit.com/r/ProgrammerHumor/",
	})

	return sources
}
