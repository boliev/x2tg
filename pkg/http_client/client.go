package http_client

import (
	"io"
	"net/http"
)

type HTTP struct {
}

func (h HTTP) Get(uri string) (string, error) {
	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
