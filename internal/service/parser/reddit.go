package parser

import (
	"encoding/json"
	"fmt"
	"path"

	domain "github.com/boliev/x2tg/internal/domain/model"
)

type Reddit struct {
	client HttpClient
}

func NewRedditParser(client HttpClient) *Reddit {
	return &Reddit{
		client: client,
	}
}

func (r *Reddit) Parse(source *domain.Source) ([]*domain.Post, error) {
	fmt.Printf("Source: %s (%s)!\n", source.Resource, source.URL)
	sub := r.getSubredditFromUrl(source.URL)
	posts, err := r.getTopPosts(sub)
	if err != nil {
		return nil, err
	}

	return r.toDomain(posts), nil
}

func (r *Reddit) getSubredditFromUrl(subreddit string) string {
	return path.Base(subreddit)
}

func (r *Reddit) getTopPosts(sub string) ([]redditPost, error) {
	topUrl := fmt.Sprintf("https://www.reddit.com/r/%s/top/.json?t=day", sub)
	code, jsn, err := r.client.Get(topUrl)
	if err != nil {
		return nil, err
	}

	if code >= 300 {
		return nil, r.processError(jsn)
	}

	data := &reddisPostList{}
	if err := json.Unmarshal([]byte(jsn), &data); err != nil {
		return nil, err
	}

	return data.Data.Children, nil
}

func (r *Reddit) toDomain(posts []redditPost) []*domain.Post {
	domainPosts := []*domain.Post{}
	for _, post := range posts {
		contentType := domain.TYPE_TEXT
		content := post.Data.Selftext
		source := post.Data.Url
		media := []string{}
		if post.Data.IsVideo {
			contentType = domain.TYPE_VIDEO
			content = post.Data.Media.RedditVideo.FallbackUrl
		} else if post.Data.PostHint == "image" {
			contentType = domain.TYPE_PIC
			content = post.Data.Url
		} else if len(post.Data.GalleryData.Items) > 0 {
			contentType = domain.TYPE_GALLERY
			source = post.Data.Permalink
			for _, pic := range post.Data.GalleryData.Items {
				media = append(media, fmt.Sprintf("https://i.redd.it/%s.jpg", pic.MediaId))
			}
		}
		domainPosts = append(domainPosts, &domain.Post{
			Title:   post.Data.Title,
			Source:  source,
			Content: content,
			Type:    contentType,
			Media:   media,
		})
	}

	return domainPosts
}

func (r *Reddit) processError(jsn string) error {
	data := &RedditError{}
	if err := json.Unmarshal([]byte(jsn), &data); err != nil {
		return err
	}

	return fmt.Errorf("%d: %s", data.Error, data.Message)
}
