// main.go
package main

import (
	"github.com/UltraNemesis/wpt"
)

func main() {
	options := &wpt.Options{}
	client, _ := wpt.NewClient(options)
	client.Locations()
}
