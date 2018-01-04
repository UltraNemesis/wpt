// util.go
package wpt

import (
	"encoding/json"
)

func parse(data []byte, format string, response interface{}) error {
	var err error = nil

	switch format {
	case "json":
		err = json.Unmarshal(data, &response)
	case "csv":

	}

	return err
}
