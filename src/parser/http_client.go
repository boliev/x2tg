package parser

type HttpClient interface {
	Get(string) (string, error)
}
