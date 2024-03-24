package publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boliev/x2tg/internal/domain/model"
	"github.com/boliev/x2tg/internal/domain/repository"
)

type Publisher struct {
	client         HttpClient
	postRepository repository.PostRepository
}

func NewPublisher(client HttpClient, postRepository repository.PostRepository) *Publisher {
	return &Publisher{
		client:         client,
		postRepository: postRepository,
	}
}

func (p *Publisher) Publish(posts []*model.Post, channels []*model.Channel) error {
	for _, cnh := range channels {
		for _, post := range posts {
			isSent, err := p.postRepository.IsSent(post, cnh)

			if err != nil {
				return err
			}

			if isSent {
				continue
			}

			uri := p.getPostMessageUri(post)

			request, err := p.buildRequest(cnh, post)
			if err != nil {
				fmt.Printf("cannot build request %s %s\n", post.Source, err)
			}
			code, message, err := p.client.Post(uri, request)
			if err != nil {
				fmt.Printf("cannot publish post %s %s\n", post.Source, err)
			}
			if code >= 300 {
				fmt.Printf("cannot publish post %s: code %d, %s\n", post.Source, code, message)
			}

			p.makeSent(post, cnh)
			time.Sleep(5 * time.Second)
		}
	}

	return nil
}

func (p *Publisher) getPostMessageUri(post *model.Post) string {
	if post.Type == model.TYPE_GALLERY {
		return "https://api.telegram.org//sendMediaGroup"
	}
	if post.Type == model.TYPE_PIC {
		return "https://api.telegram.org//sendPhoto"
	}
	if post.Type == model.TYPE_VIDEO {
		return "https://api.telegram.org//sendVideo"
	}

	return "https://api.telegram.org//sendMessage"
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
	} else if post.Type == model.TYPE_PIC {
		request, err = p.buildPhotoRequest(cnh, post)
		if err != nil {
			return nil, err
		}
	} else if post.Type == model.TYPE_VIDEO {
		request, err = p.buildVideoRequest(cnh, post)
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

func (p *Publisher) buildPhotoRequest(cnh *model.Channel, post *model.Post) (*postRequest, error) {
	request := &postRequest{
		ChatId:                cnh.TgIg,
		ParseMode:             "html",
		DisableWebPagePreview: false,
		Photo:                 post.Media[0],
		Caption:               fmt.Sprintf("%s\nhttp://reddit.com%s", post.Title, post.Source),
	}

	return request, nil
}

func (p *Publisher) buildVideoRequest(cnh *model.Channel, post *model.Post) (*postRequest, error) {
	request := &postRequest{
		ChatId:                cnh.TgIg,
		ParseMode:             "html",
		DisableWebPagePreview: false,
		Video:                 post.Media[0],
		Caption:               fmt.Sprintf("%s\nhttp://reddit.com%s", post.Title, post.Source),
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

func (p *Publisher) makeSent(post *model.Post, cnh *model.Channel) error {
	err := p.postRepository.MakeSent(post, cnh)
	if err != nil {
		return err
	}

	return nil
}
