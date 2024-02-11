package model

type Source struct {
	ID       int
	Resource string
	URL      string
	Channels []*Channel
	IsActive bool
}
