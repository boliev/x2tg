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

	sourceRepository := db.NewSourceRepository()
	sources := sourceRepository.GetActive()

	for _, source := range sources {
		fmt.Printf("URL: %s\n", source.URL)
		// posts, err := a.parsers[source.Resource].Parse(source)
		// if err != nil {
		// 	fmt.Errorf("Cannot parse source %s(%s). Skip", source.Url, source.Resource)
		// 	continue
		// }
		// for _, post := range posts {
		// 	fmt.Printf("Title: %s\n", post.Title)
		// 	fmt.Printf("Source: %s\n", post.Source)
		// 	fmt.Printf("Type: %s\n", post.Type)
		// 	fmt.Printf("Content: %s\n", post.Content)
		// 	fmt.Println("-----------")
		// }
	}
}
