package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequest = errors.New("request return an non success code")
)

type Client[T any] struct{}

func (c *Client[T]) Get(url string) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, ErrInvalidRequest
	}

	var data T
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
