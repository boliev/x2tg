package repository

import "github.com/boliev/x2tg/internal/domain"

type SourceRepository interface {
	getActive() ([]*domain.Source, error)
}
