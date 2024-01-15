package parser

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/boliev/x2tg/src/domain"
)

type Reddit struct {
	client HttpClient
}

func NewRedditParser(client HttpClient) *Reddit {
	return &Reddit{
		client: client,
	}
}

func (r Reddit) Parse(source *domain.Source) ([]*domain.Post, error) {
	fmt.Printf("Source: %s (%s)!\n", source.Resource, source.Url)
	sub := r.getSubredditFromUrl(source.Url)
	posts, err := r.getTopPosts(sub)
	if err != nil {
		return nil, err
	}

	return r.toDomain(posts), nil
}

func (r Reddit) getSubredditFromUrl(subreddit string) string {
	return path.Base(subreddit)
}

func (r Reddit) getTopPosts(sub string) ([]redditPost, error) {
	topUrl := fmt.Sprintf("https://www.reddit.com/r/%s/top/.json?t=day", sub)
	jsn, err := r.client.Get(topUrl)
	if err != nil {
		return nil, err
	}

	data := &reddisPostList{}
	if err := json.Unmarshal([]byte(jsn), &data); err != nil {
		return nil, err
	}

	return data.Data.Children, nil
}

func (r Reddit) toDomain(posts []redditPost) []*domain.Post {
	domainPosts := []*domain.Post{}
	for _, post := range posts {
		domainPosts = append(domainPosts, &domain.Post{
			Title:  post.Data.Title,
			Source: post.Data.Url,
			Text:   post.Data.Selftext,
		})
	}

	return domainPosts
}
