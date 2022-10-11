package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// read env vars, no inference.
var secret string = os.Getenv("bitbucket_secret")
var appname string = os.Getenv("app_name")
var token string = os.Getenv("bitbucket_token")
var apiurl string = os.Getenv("apiurl")
var webhookurl string = os.Getenv("webhookurl")

type Payload struct {
	Events        []string      `json:"events"`
	Active        bool          `json:"active"`
	Statistics    Statistics    `json:"statistics"`
	Configuration Configuration `json:"configuration"`
	URL           string        `json:"url"`
	Name          string        `json:"name"`
}
type Statistics struct {
}
type Configuration struct {
	CreatedBy string `json:"createdBy"`
	Secret    string `json:"secret"`
}

func main() {

	data := Payload{
		Events:        []string{"repo:refs_changed"}, // the events that BBS should monitor, when this occurs the trigger is fired and a webhook call is made
		Active:        true,
		Statistics:    Statistics{},
		Configuration: Configuration{"bitbucket", secret},
		URL:           webhookurl,    // this is the URL that BBS should send a event payload to when a commit is made.
		Name:          "Argo Events", // the name given to the webhook that gets created in BBS
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	// The request is sent to the api url with the repo-slug name appended
	req, err := http.NewRequest("POST", apiurl+appname+"/webhooks", body)
	if err != nil {
		// handle err
	}
	// set headers, grab token from env variable
	req.Header.Set("Authorization", "Bearer "+token+"")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

}
