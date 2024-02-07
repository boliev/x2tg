package parser

import "github.com/boliev/x2tg/internal/domain/model"

type Parser interface {
	Parse(source *model.Source) ([]*model.Post, error)
}
