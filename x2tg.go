package x2tg

import (
	"fmt"

	"github.com/boliev/x2tg/pkg/http_client"
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

	sources := []*domain.Source{}
	sources = append(sources, &domain.Source{
		Resource: "reddit",
		Url:      "https://www.reddit.com/r/golang/",
	})

	for _, source := range sources {
		posts, err := a.parsers[source.Resource].Parse(source)
		if err != nil {
			fmt.Errorf("Cannot parse source %s(%s). Skip", source.Url, source.Resource)
			continue
		}
		for _, post := range posts {
			fmt.Printf("%s\n", post.Title)
			fmt.Printf("%s\n", post.Source)
			fmt.Printf("%s\n", post.Text)
			fmt.Println("-----------")
		}
	}
}
