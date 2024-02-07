package db

import (
	"database/sql"

	domain "github.com/boliev/x2tg/internal/domain/model"
)

type SourceRepository struct {
	DB *sql.DB
}

func NewSourceRepository(DB *sql.DB) *SourceRepository {
	return &SourceRepository{
		DB: DB,
	}
}

func (sr *SourceRepository) GetActive() ([]*domain.Source, error) {
	rows, err := sr.DB.Query("SELECT id, resource, url, is_active FROM sources where is_active = $1", true)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sources := []*domain.Source{}
	for rows.Next() {
		source := &domain.Source{}
		if err := rows.Scan(&source.ID, &source.Resource, &source.URL, &source.IsActive); err != nil {
			return sources, err
		}

		sources = append(sources, source)
	}

	return sources, nil
}
