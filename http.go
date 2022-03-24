package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequest = errors.New("request return an non success code")
)

type Client[T any] struct {
	HttpClient *http.Client
}

// NewClient initialize with custom HTTP Client
func NewClient[T any](httpClient *http.Client) Client[T] {
	return Client[T]{
		HttpClient: httpClient,
	}
}

// Get calls the given url and decode the response
// to the given type of the Client
func (c *Client[T]) Get(url string) (*T, error) {
	httpClient := http.DefaultClient
	if c.HttpClient != nil {
		httpClient = c.HttpClient
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
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
