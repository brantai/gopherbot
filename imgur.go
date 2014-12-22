package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ImgurPlugin struct{}

type ImgurResponse struct {
	Data    []ImgurData `json:"data"`
	Status  int         `json:"status"`
	Success bool        `json:"success"`
}

type ImgurData struct {
	Link string `json:"link"`
}

func (p ImgurPlugin) Name() string {
	return "imgur"
}

func (p ImgurPlugin) Execute(command []string) string {
	if len(command) < 1 {
		return topViral()
	} else {
		return search(command)
	}
}

func search(command []string) string {
	var iresp ImgurResponse
	client_id, err := ioutil.ReadFile("imgur_cid")
	if err != nil {
		return "Error reading Imgur client ID file"
	}

	query := strings.Join(command, "%20")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.imgur.com/3/gallery/search?q="+query, nil)
	if err != nil {
		return "Error making a new request"
	}

	req.Header.Add("Authorization", "Client-ID "+string(client_id))

	resp, err := client.Do(req)
	if err != nil {
		return "Error calling Imgur API"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading Imgur response body"
	}

	err = json.Unmarshal(body, &iresp)
	if err != nil {
		return "Error unmarshaling Imgur response"
	}

	if len(iresp.Data) <= 0 {
		return "No image found"
	}

	return iresp.Data[0].Link
}

func topViral() string {
	var iresp ImgurResponse
	client_id, err := ioutil.ReadFile("imgur_cid")
	if err != nil {
		return "Error reading Imgur client ID file"
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.imgur.com/3/gallery/hot/viral/0.json", nil)
	if err != nil {
		return "Error making a new request"
	}

	req.Header.Add("Authorization", "Client-ID "+string(client_id))

	resp, err := client.Do(req)
	if err != nil {
		return "Error calling Imgur API"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading Imgur response body"
	}

	err = json.Unmarshal(body, &iresp)
	if err != nil {
		return "Error unmarshaling Imgur response"
	}

	if len(iresp.Data) <= 0 {
		return "No gif found"
	}

	return iresp.Data[0].Link
}
