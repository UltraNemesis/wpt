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

func (c *Client) Locations() (*WPTLocations, error) {
	v := url.Values{}
	v.Add("f", "json")

	var response WPTLocations

	resp, err := c.query(wptQueryLocations, v)

	if err == nil {
		parseData(resp, v.Get("f"), &response)
	}

	return &response, err
}

func (c *Client) GetResults(id string) (*WPTResults, error) {
	v := url.Values{}
	v.Add("test", id)
	v.Add("f", "json")

	resp, err := c.query(wptQueryTestResults, v)

	var results WPTResults

	parseData(resp, v.Get("f"), &results)

	return &results, err
}

func (c *Client) GetTestHistory(days string) (*WPTHistory, error) {
	v := url.Values{}

	v.Add("f", "csv")
	v.Add("days", days)
	v.Add("all", "on")

	resp, _ := c.query(wptQueryTestHistory, v)

	var history WPTHistory

	err := parseData(resp, v.Get("f"), &history.Items)

	return &history, err
}

func (c *Client) query(path string, values url.Values) ([]byte, error) {
	url := c.options.URL
	url.Path = path
	url.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)

	log.Println("Making Request : ", url.String())

	if err != nil {
		return nil, errCreateRequest
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, errQueryServer
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errBadResponse
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errReadBody
	}

	//response = parse(body, values.Get("f"), response)

	return body, nil
}
