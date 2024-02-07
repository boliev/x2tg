package parser

type HttpClient interface {
	Get(string) (int, string, error)
}
