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
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return 0, "", err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	res, err := client.Do(req)

	// res, err := http.Get(uri)
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
