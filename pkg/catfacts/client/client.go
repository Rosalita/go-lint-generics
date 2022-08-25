package catfacts

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type Client struct {
	baseURL string
	client  *retryablehttp.Client
}

func NewClient(baseURL string) (*Client, error) {
	return &Client{
		baseURL: baseURL,
		client:  retryablehttp.NewClient(),
	}, nil
}

// Get makes a http GET request, used retrieve a fact about cats.
func (c *Client) Get(ctx context.Context, path string) (*CatFact, error) {
	reqURL, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	reqURL.Path = path

	req, err := retryablehttp.NewRequestWithContext(ctx, "GET", reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var fact CatFact
		if err := json.Unmarshal(body, &fact); err != nil {
			return nil, err
		}
		return &fact, nil
	default:
		err = errors.New("status code not ok")
		return nil, err
	}
}
