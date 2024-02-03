package x2tg

import (
	"fmt"

	"github.com/boliev/x2tg/pkg/http_client"
	db "github.com/boliev/x2tg/src/db"
	"github.com/boliev/x2tg/src/domain"
	parser "github.com/boliev/x2tg/src/parser"
)

type App struct {
	parsers map[string]domain.Parser
}

func (a App) Run() {
	fmt.Println("I'm the service! I'm working!")

	httpClient := &http_client.HTTP{}
	a.parsers = make(map[string]domain.Parser)
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
		fmt.Printf("URL: %s\n", source.URL)
	}
}
