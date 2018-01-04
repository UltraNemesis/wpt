// locations.go
package wpt

import (
	"errors"
)

const defaultURL = "http://www.webpagetest.org"

const (
	wptQueryLocations   = "getLocations.php"
	wptQueryRunTest     = "runtest.php"
	wptQueryCancelTest  = "cancelTest.php"
	wptQueryTestStatus  = "testStatus.php"
	wptQueryTestResults = "jsonResult.php"
)

var (
	errCreateRequest = errors.New("Error creating request")
	errQueryServer   = errors.New("Error querying WPT server")
	errBadResponse   = errors.New("Bad response code from WPT server")
	errReadBody      = errors.New("Error reading response body")
)

type Location struct {
	Label         string `json:"Label"`
	Location      string `json:"location"`
	Browser       string `json:"Browser"`
	RelayServer   string `json:"relayServer"`
	RelayLocation string `json:"relayLocation"`
	LabelShort    string `json:"labelShort"`
	Default       bool   `json:"default"`
	PendingTests  struct {
		Total        int `json:"Total"`
		HighPriority int `json:"HighPriority"`
		LowPriority  int `json:"LowPriority"`
		Testing      int `json:"Testing"`
		Idle         int `json:"Idle"`
	} `json:"PendingTests"`
}

type LocationsResponse struct {
	StatusCode int                 `json:"statusCode"`
	StatusText string              `json:"statusText"`
	Data       map[string]Location `json:"data"`
}
