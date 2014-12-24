package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RandomgifPlugin struct{}

type RandomGif struct {
	Token  string `json:"token"`
	GifUrl string `json:"gif_url"`
}

func init() {
	Plugins = append(Plugins, RandomgifPlugin{})
}

func (p RandomgifPlugin) Help() string {
	return "Returns a random gif from gifs.com. Usage: !randomgif"
}

func (p RandomgifPlugin) Name() string {
	return "randomgif"
}

func (p RandomgifPlugin) Execute(command []string) string {
	_ = command // This plugin doesn't need command bits, but the interface defines ti
	resp, err := http.Get("http://gifs.com/r.json")
	if err != nil {
		return "Couldn't get random gif"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Couldn't read response body for random gif"
	}

	var rg RandomGif
	err = json.Unmarshal(body, &rg)
	if err != nil {
		return "Couldn't unmarshal data for random gif"
	}

	return rg.GifUrl
}
