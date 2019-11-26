package main

import (
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	tick   *time.Ticker
}

type apiTransport struct {
	http.RoundTripper
	key string
}

type Result struct {
	AddressComponents []struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		//more fields
	} `json:"geometry"`
	PlaceID string `json:"place_id"`
	//more fields
}

func (a apiTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	q := r.URL.Query()
	q.Set("key", a.key)
	r.URL.RawQuery = q.Encode()
	return a.RoundTripper.RoundTrip(r)
}

func NewClient(tick time.Duration, key string) *Client {
	return &Client{
		client: &http.Client{
			Transport: apiTransport{http.DefaultTransport, key},
		},
		tick: time.NewTicker(tick),
	}
}
