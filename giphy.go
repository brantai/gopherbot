package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type GiphyPlugin struct{}

type GiphyResponse struct {
	Data []GiphyData `json:"data"`
}

type GiphyData struct {
	Url string `json:"url"`
}

func (p GiphyPlugin) Help() string {
	return "Returns the first search result from giphy.com. Usage: !giphy <search term>"
}

func (p GiphyPlugin) Name() string {
	return "giphy"
}

func (p GiphyPlugin) Execute(command []string) string {
	if len(command) < 1 {
		return "herp"
	}
	var gresp GiphyResponse
	apikey := configMap["giphy_key"]
	if apikey == "" {
		return "No giphy_key in config file"
	}
	resp, err := http.Get("http://api.giphy.com/v1/gifs/search?q=" + strings.Join(command, "+") + "&api_key=" + string(apikey) + "&limit=1")
	if err != nil {
		return "Error calling Giphy API"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading Giphy response body"
	}

	err = json.Unmarshal(body, &gresp)
	if err != nil {
		return "Error unmarshaling Giphy response"
	}

	if len(gresp.Data) <= 0 {
		return "No gif found"
	}

	return gresp.Data[0].Url

}
