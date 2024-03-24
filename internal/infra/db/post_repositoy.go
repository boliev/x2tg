package db

import (
	"database/sql"

	"github.com/boliev/x2tg/internal/domain/model"
	"github.com/huandu/go-sqlbuilder"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(DB *sql.DB) *PostRepository {
	return &PostRepository{
		DB: DB,
	}
}

func (pr *PostRepository) MakeSent(post *model.Post, chn *model.Channel) error {
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("sent")
	ib.Cols("source", "channel")
	ib.Values(post.Source, chn.ID)
	sql, args := ib.Build()

	_, err := pr.DB.Query(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) IsSent(post *model.Post, chn *model.Channel) (bool, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("source", "channel")
	sb.From("sent")
	sb.Where(
		sb.And(
			sb.E("source", post.Source),
			sb.E("channel", chn.ID),
		),
	)

	sql, args := sb.Build()
	rows, err := pr.DB.Query(sql, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		return true, nil
	}

	return false, nil

}
