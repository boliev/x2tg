package repository

import "github.com/boliev/x2tg/internal/domain/model"

type SourceRepository interface {
	getActive() ([]*model.Source, error)
}
