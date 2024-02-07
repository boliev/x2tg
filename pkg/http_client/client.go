package http_client

import (
	"io"
	"net/http"
)

type HTTP struct {
}

func (h HTTP) Get(uri string) (int, string, error) {
	res, err := http.Get(uri)
	if err != nil {
		return 0, "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(body), nil
}
