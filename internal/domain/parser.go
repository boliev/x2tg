package domain

type Parser interface {
	Parse(source *Source) ([]*Post, error)
}
