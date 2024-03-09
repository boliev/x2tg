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
	Title       string
	Permalink   string
	Url         string
	Selftext    string
	IsVideo     bool        `json:"is_video"`
	PostHint    string      `json:"post_hint"`
	GalleryData GalleryData `json:"gallery_data"`
	Media       RedditMedia
}

type RedditMedia struct {
	RedditVideo RedditVideo `json:"reddit_video"`
}

type RedditVideo struct {
	FallbackUrl string `json:"fallback_url"`
}

type GalleryData struct {
	Items []GalleyItem
}

type GalleyItem struct {
	MediaId string `json:"media_id"`
}
