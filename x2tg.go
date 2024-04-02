package x2tg

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	db "github.com/boliev/x2tg/internal/infra/db"
	parser "github.com/boliev/x2tg/internal/service/parser"
	"github.com/boliev/x2tg/internal/service/publisher"
	"github.com/boliev/x2tg/pkg/http_client"
	"github.com/caarlos0/env/v10"
)

type App struct {
	parsers map[string]parser.Parser
}

type Config struct {
	DbHost     string `env:"DB_HOST,required"`
	DbPort     int    `env:"DB_PORT,required"`
	DbName     string `env:"DB_NAME,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	TgBotToken string `env:"TG_BOT_TOKEN,required"`
}

func (a App) Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("I'm the service! I'm working!")

	httpClient := &http_client.HTTP{}
	a.parsers = make(map[string]parser.Parser)
	a.parsers["reddit"] = parser.NewRedditParser(httpClient)
	config := &Config{}
	if err := env.Parse(config); err != nil {
		slog.Error(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}

	DB, err := db.NewDBConnection(
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
	)

	if err != nil {
		slog.Error(fmt.Sprintf("cannot connect to DB %s", err.Error()))
		os.Exit(1)
	}
	defer DB.Close()

	postRepository := db.NewPostRepository(DB)

	publisher := publisher.NewPublisher(httpClient, postRepository, config.TgBotToken)

	sourceRepository := db.NewSourceRepository(DB)
	sources, err := sourceRepository.GetActive()
	if err != nil {
		slog.Error(fmt.Sprintf("error while retrieveing sources %s", err))
		os.Exit(1)
	}

	for _, source := range sources {
		// fmt.Printf("%#v\n\n", source)
		posts, err := a.parsers["reddit"].Parse(source)
		if err != nil {
			slog.Warn(fmt.Sprintf("error parsing subreddit: %s", err.Error()))
			continue
		}

		err = publisher.Publish(posts, source.Channels)
		if err != nil {
			slog.Warn(fmt.Sprintf("error: %s", err.Error()))
		}

		time.Sleep(5 * time.Second)

		// for _, post := range posts {
		// 	fmt.Printf("Title: %s\nURL: %s\nType: %s\n---\n", post.Title, post.Source, post.Type)
		// }
	}
}
