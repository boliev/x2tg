package x2tg

import (
	"fmt"

	db "github.com/boliev/x2tg/internal/infra/db"
	parser "github.com/boliev/x2tg/internal/service/parser"
	"github.com/boliev/x2tg/pkg/http_client"
)

type App struct {
	parsers map[string]parser.Parser
}

func (a App) Run() {
	fmt.Println("I'm the service! I'm working!")

	httpClient := &http_client.HTTP{}
	a.parsers = make(map[string]parser.Parser)
	a.parsers["reddit"] = parser.NewRedditParser(httpClient)

	DB, err := db.NewDBConnection("localhost", 5432, "x2tg", "123456", "x2tg")
	if err != nil {
		panic(fmt.Sprintf("cannot connect to DB %s", err.Error()))
	}
	defer DB.Close()

	sourceRepository := db.NewSourceRepository(DB)
	sources, err := sourceRepository.GetActive()
	if err != nil {
		panic(fmt.Sprintf("error while retrieveing sources %s", err))
	}

	for _, source := range sources {
		posts, err := a.parsers["reddit"].Parse(source)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			continue
		}

		for _, post := range posts {
			fmt.Printf("Title: %s\nURL: %s\nType: %s\n---\n", post.Title, post.Source, post.Type)
		}
	}
}
