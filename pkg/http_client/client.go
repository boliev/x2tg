package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HTTP struct {
}

func (h *HTTP) Get(uri string) (int, string, error) {
	res, err := http.Get(uri)
	if err != nil {
		return 0, "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	h.logToFile(uri, body)

	return res.StatusCode, string(body), nil
}

func (h *HTTP) Post(uri string, request interface{}) (int, string, error) {
	requestJson, err := json.Marshal(request)
	if err != nil {
		return 0, "", err
	}

	h.logToFile(uri, requestJson)

	res, err := http.Post(uri, "application/json", bytes.NewReader(requestJson))
	if err != nil {
		return 0, "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, "", err
	}

	return res.StatusCode, string(body), nil
}

func (h *HTTP) logToFile(uri string, json []byte) {
	f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	if _, err := f.Write([]byte(fmt.Sprintf("\n\n%s\n", uri))); err != nil {
		fmt.Println(err.Error())
	}
	if _, err := f.Write(json); err != nil {
		fmt.Println(err.Error())
	}
	if err := f.Close(); err != nil {
		fmt.Println(err.Error())
	}
}
