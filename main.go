package main

import (
	"bytes"
	"net/http"
)

func main() {
	url := "place webhook url here"

	var json = []byte(`{
			"content": "<@865734393570132000>",
			"embeds": [
    			{
      				"title": "A new Episode of the Kirby TV Channel is available!",
      				"color": 16716799
    			}
  			]
		}`)

	_, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
}
