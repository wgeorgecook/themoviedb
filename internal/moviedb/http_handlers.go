package moviedb

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var httpClient = http.Client{
	Timeout: 30 * time.Second,
}

func generateHttpRequest(method, endpoint, token string, payload []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("bearer %s", token))

	return req, nil
}

func makeHttpRequest(req *http.Request) ([]byte, error) {
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
