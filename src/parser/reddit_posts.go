package parser

type reddisPostList struct {
	Data reddisPostListData
}

type reddisPostListData struct {
	Children []redditPost
}

type redditPost struct {
	Kind string
	Data redditPostData
}

type redditPostData struct {
	Title    string
	Url      string
	Selftext string
	IsVideo  bool   `json:"is_video"`
	PostHint string `json:"post_hint"`
	Media    RedditMedia
}

type RedditMedia struct {
	RedditVideo RedditVideo `json:"reddit_video"`
}

type RedditVideo struct {
	FallbackUrl string `json:"fallback_url"`
}
