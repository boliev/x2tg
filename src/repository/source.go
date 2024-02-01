package repository

type SourceRepository interface {
	getActive() []*Source
}
