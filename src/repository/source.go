package repository

import "github.com/boliev/x2tg/src/domain"

type SourceRepository interface {
	getActive() ([]*domain.Source, error)
}
