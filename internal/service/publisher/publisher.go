package publisher

import (
	"encoding/json"
	"fmt"

	"github.com/boliev/x2tg/internal/domain/model"
)

type Publisher struct {
	client HttpClient
}

func NewPublisher(client HttpClient) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) Publish(posts []*model.Post, channels []*model.Channel) error {
	for _, cnh := range channels {
		for _, post := range posts {
			uri := "https://api.telegram.org//sendMessage"
			if post.Type == model.TYPE_GALLERY {
				uri = "https://api.telegram.org//sendMediaGroup"
			}

			request, err := p.buildRequest(cnh, post)
			if err != nil {
				fmt.Printf("error in %s %s\n", post.Source, err)
			}
			code, message, err := p.client.Post(uri, request)
			if err != nil {
				fmt.Printf("error in %s %s\n", post.Source, err)
			}
			if code >= 300 {
				fmt.Printf("error in %s: code %d, %s\n", post.Source, code, message)
			}
		}
	}

	return nil
}

func (p *Publisher) buildRequest(cnh *model.Channel, post *model.Post) (*postRequest, error) {
	// photo request
	// video request
	var request *postRequest
	var err error
	if post.Type == model.TYPE_GALLERY {
		request, err = p.buildMediaGroupRequest(cnh, post)
		if err != nil {
			return nil, err
		}
	} else {
		request = &postRequest{
			ChatId:                cnh.TgIg,
			ParseMode:             "html",
			DisableWebPagePreview: false,
			Text:                  p.buildPostContent(post),
		}
	}

	return request, nil
}

func (p *Publisher) buildMediaGroupRequest(cnh *model.Channel, post *model.Post) (*postRequest, error) {
	media := []*PostMediaField{}
	for index, m := range post.Media {
		caption := ""
		if index == 0 {
			caption = fmt.Sprintf("%s\nhttp://reddit.com%s", post.Title, post.Source)
		}
		media = append(media, &PostMediaField{
			Type:     "photo",
			ParseMod: "html",
			Media:    m,
			Caption:  caption,
		})
	}
	jsonMedia, err := json.Marshal(media)
	if err != nil {
		return nil, fmt.Errorf("cannot build json from gallery %s", post.Source)
	}
	request := &postRequest{
		ChatId:                cnh.TgIg,
		DisableWebPagePreview: false,
		Media:                 string(jsonMedia),
	}

	return request, nil
}

func (p *Publisher) buildPostContent(post *model.Post) string {
	content := post.Content
	if len(content) > 3000 {
		content = content[:3000]
	}

	template := fmt.Sprintf("<b>%s</b>\n%s", post.Title, content)
	if post.Type == model.TYPE_TEXT {
		template = template + fmt.Sprintf("\n%s", post.Source)
	}

	return template
}
