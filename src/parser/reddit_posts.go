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
}
