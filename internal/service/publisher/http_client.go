package publisher

type HttpClient interface {
	Post(uri string, request interface{}) (int, string, error)
}
