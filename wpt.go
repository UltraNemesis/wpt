// wpt.go
package wpt

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Options struct {
	URL    *url.URL
	APIKey string
}

type Client struct {
	options    *Options
	httpClient *http.Client
}

func NewClient(options *Options) (*Client, error) {

	if options.URL == nil {
		options.URL, _ = url.Parse(defaultURL)
	}

	client := &Client{
		options:    options,
		httpClient: &http.Client{},
	}

	return client, nil
}

func (c *Client) Locations() {
	v := url.Values{}
	v.Add("f", "json")

	var response LocationsResponse

	c.query(wptQueryLocations, v, &response)

	log.Println(response)
}

func (c *Client) query(path string, values url.Values, response interface{}) error {
	url := c.options.URL
	url.Path = path
	url.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return errCreateRequest
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return errQueryServer
	}

	if resp.StatusCode != http.StatusOK {
		return errBadResponse
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return errReadBody
	}

	response = parse(body, values.Get("f"), response)

	return nil
}
