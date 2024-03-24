package repository

import "github.com/boliev/x2tg/internal/domain/model"

type PostRepository interface {
	MakeSent(post *model.Post, chn *model.Channel) error
	IsSent(post *model.Post, chn *model.Channel) (bool, error)
}
