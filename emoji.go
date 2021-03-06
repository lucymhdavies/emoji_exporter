package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

//
// REST API
//
// https://github.com/emojitracker/emojitrack-rest-api

type EmojiRankingsResponse []struct {
	Char  string `json:"char"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func EmojiRankingsRequest() ([]byte, *http.Response, error) {
	// Modified from code generated by curl-to-Go: https://mholt.github.io/curl-to-go

	url := "https://api.emojitracker.com/v1/rankings"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("%s", err)
		return []byte{}, nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("%s", err)
		return []byte{}, resp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"url": url}).Errorf("%s", err)
		return []byte{}, resp, err
	}

	return body, resp, nil
}

func Rankings() (EmojiRankingsResponse, error) {

	// init an empty response
	response := EmojiRankingsResponse{}

	// body, resp, err
	body, resp, err := EmojiRankingsRequest()
	if err != nil {
		log.Errorf("%s", err)
		return response, err
	}
	if resp.StatusCode != 200 {
		log.Errorf("Error code %s, Error: %s", resp.StatusCode, err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Errorf("%s", err)
		return response, err
	}

	return response, nil
}

// TODO
// https://github.com/emojitracker/emojitrack-streamer-spec
