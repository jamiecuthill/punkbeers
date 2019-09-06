package punkapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const version = "v2"

// Client is an HTTP implementation against the punkapi
type Client struct {
	url  *url.URL
	http *http.Client
}

// NewClient returns a new client that makes calls against the given url
func NewClient(baseURL string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "/"+version)

	return &Client{
		url: u,
		http: &http.Client{
			Timeout: 1 * time.Second,
		},
	}, nil
}

// Beers fetches all beers from the API matching the input criteria
func (c *Client) Beers(input BeersInput) ([]Beer, error) {
	beersURL := *c.url
	beersURL.Path = path.Join(beersURL.Path, "/beers")

	q := beersURL.Query()
	q.Add("page", fmt.Sprintf("%d", input.Page))
	beersURL.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, beersURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %s", resp.Status)
	}

	var beers []Beer
	err = json.NewDecoder(resp.Body).Decode(&beers)

	if len(beers) < 1 {
		return beers, ErrNoMorePages
	}

	return beers, err
}

// AllBeers fetches all pages of Client.Beers
func (c *Client) AllBeers() ([]Beer, error) {
	var err error
	var beers []Beer
	var page uint = 1
	for {
		var b []Beer
		b, err = c.Beers(BeersInput{
			Page: page,
		})
		if err != nil {
			if err == ErrNoMorePages {
				err = nil
			}
			break
		}
		page++
		beers = append(beers, b...)
	}

	return beers, err
}

// BeersInput are parameters for the request for beers
type BeersInput struct {
	Page uint
}
