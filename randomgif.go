package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RandomGif struct {
	Token  string `json:"token"`
	GifUrl string `json:"gif_url"`
}

func randomgif() string {
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
