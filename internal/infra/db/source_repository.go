package db

import (
	"database/sql"

	domain "github.com/boliev/x2tg/internal/domain/model"
	"github.com/huandu/go-sqlbuilder"
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
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("s.id", "s.resource", "s.url", "s.is_active, c.id, c.tg_id, c.name")
	sb.From(sb.As("sources", "s"))
	sb.JoinWithOption(sqlbuilder.InnerJoin, sb.As("sources_channels", "sc"), "sc.source_id = s.id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, sb.As("channels", "c"), "c.id = sc.channel_id")
	sb.Where(sb.E("s.is_active", true))

	sql, args := sb.Build()
	rows, err := sr.DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sourcesMap := make(map[int]*domain.Source)
	for rows.Next() {
		source := &domain.Source{}
		channel := &domain.Channel{}
		if err := rows.Scan(
			&source.ID,
			&source.Resource,
			&source.URL,
			&source.IsActive,
			&channel.ID,
			&channel.TgIg,
			&channel.Name,
		); err != nil {
			return nil, err
		}

		if _, ok := sourcesMap[source.ID]; ok {
			sourcesMap[source.ID].Channels = append(sourcesMap[source.ID].Channels, channel)
		} else {
			source.Channels = append(source.Channels, channel)
			sourcesMap[source.ID] = source
		}
	}

	sources := []*domain.Source{}
	for _, source := range sourcesMap {
		sources = append(sources, source)
	}

	return sources, nil
}
