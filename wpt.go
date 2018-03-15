// wpt.go
package wpt

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	DefualtStatusPollingInterval = 5
	DefaultTestTimeout           = 30
)

var (
	StatusPollingInterval time.Duration = DefualtStatusPollingInterval
	TestTimeout           time.Duration = DefaultTestTimeout
)

type ClientOptions struct {
	URL    *url.URL
	APIKey string
}

type Client struct {
	options    *ClientOptions
	httpClient *http.Client
}

type Test struct {
	client   *Client
	options  *TestOptions
	Response *TestResponse
}

func NewClient(options *ClientOptions) (*Client, error) {

	if options.URL == nil {
		options.URL, _ = url.Parse(defaultURL)
	}

	client := &Client{
		options:    options,
		httpClient: &http.Client{},
	}

	return client, nil
}

func (c *Client) NewTest(options *TestOptions) (*Test, error) {
	return &Test{
		client:  c,
		options: options,
	}, nil
}

func (t *Test) Run() (*TestResponse, error) {
	v, _ := query.Values(t.options)

	if len(t.client.options.APIKey) > 0 {
		v.Add("k", t.client.options.APIKey)
	}

	v.Add("f", "json")

	resp, err := t.client.query(wptQueryRunTest, v)

	var response TestResponse

	if err != nil {
		log.Println(err)

	} else {
		parseData(resp, v.Get("f"), &response)
	}

	t.Response = &response

	return &response, err
}

func (t *Test) RunSync() {
	t.Run()
	t.monitor()
}

func (t *Test) monitor() {
	for {
		select {

		default:
			time.Sleep(StatusPollingInterval * time.Second)

			status, _ := t.client.GetStatus(t.Response.Data.TestId)

			log.Println(status.StatusCode)

			switch status.StatusCode {
			case wptStatusTestSuccess, wptStatusTestNotFound, wptStatusTestCancelled:
				log.Println("Exiting")
				return
			}

		}
	}
}

func (c *Client) GetLocations() (*WPTLocations, error) {
	v := url.Values{}
	v.Add("f", "json")

	var response WPTLocations

	resp, err := c.query(wptQueryLocations, v)

	if err == nil {
		parseData(resp, v.Get("f"), &response)
	}

	return &response, err
}

func (c *Client) GetStatus(testId string) (*TestStatus, error) {
	v := url.Values{}
	v.Add("test", testId)
	v.Add("f", "json")

	resp, err := c.query(wptQueryTestStatus, v)

	var testStatus TestStatus

	parseData(resp, v.Get("f"), &testStatus)

	return &testStatus, err
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

func (c *Client) GetTestHistory(days int, from string, filter string) (*WPTHistory, error) {
	v := url.Values{}

	v.Add("f", "csv")
	v.Add("days", strconv.Itoa(days))
	v.Add("all", "on")

	if len(from) > 0 {
		v.Add("from", from)
	}

	if len(filter) > 0 {
		v.Add("filter", filter)
	}

	resp, _ := c.query(wptQueryTestHistory, v)

	var history WPTHistory

	err := parseData(resp, v.Get("f"), &history.Items)

	return &history, err
}

func (c *Client) CancelTest(testId string) error {
	v := url.Values{}

	v.Add("test", testId)
	v.Add("k", c.options.APIKey)

	_, err := c.query(wptQueryCancelTest, v)

	return err
}

func (c *Client) query(path string, values url.Values) (string, error) {
	url := c.options.URL
	url.Path = path
	url.RawQuery = values.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)

	log.Println("Making Request : ", url.String())

	if err != nil {
		return "", errCreateRequest
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return "", errQueryServer
	}

	if resp.StatusCode != http.StatusOK {
		return "", errBadResponse
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errReadBody
	}

	//response = parse(body, values.Get("f"), response)

	return string(body), nil
}
