// main.go
package main

import (
	"encoding/json"
	"log"

	"io/ioutil"
	"os"

	"github.com/UltraNemesis/wpt"
)

func main() {
	f, _ := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	log.SetOutput(f)

	options := &wpt.Options{}
	client, _ := wpt.NewClient(options)
	resp, _ := client.Locations()
	log.Println(resp)

	results, _ := client.GetResults("180105_S9_6c347b11f8fc33eb77f3e4f69a11710c")
	out, _ := json.Marshal(results)
	ioutil.WriteFile("results.json", out, 777)

	history, _ := client.GetTestHistory("1")
	log.Println(history.Items)
}
