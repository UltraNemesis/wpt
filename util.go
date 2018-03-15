// util.go
package wpt

import (
	"encoding/json"
	"log"

	"github.com/gocarina/gocsv"
)

func parseData(data string, format string, response interface{}) error {
	var err error = nil

	switch format {
	case "json":
		err = json.Unmarshal([]byte(data), response)
	case "csv":
		err = gocsv.UnmarshalString(data, response)
	default:
		log.Println("Unknown format : ", format)
	}

	if err != nil {
		log.Println(err)
	}

	return err
}
