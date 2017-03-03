package guidebox

// Hacker News API Documentation:
// https://github.com/HackerNews/API

import (
	"net/http"

	"fmt"

	"github.com/dghubble/sling"
)

const (
	BaseURL    = "http://api-public.guidebox.com"
	APIVersion = "v2"
	showsURL   = "/shows"

	userURL    = "/user/"
	topURL     = "/topstories.json"
	newURL     = "/newstories.json"
	bestURL    = "/beststories.json"
	askURL     = "/askstories.json"
	showURL    = "/showstories.json"
	jobsURL    = "/jobstories.json"
	updatesURL = "/updates.json"
	maxItemURL = "/maxitem.json"
)

type GuideboxClient struct {
	sling  *sling.Sling
	apiKey string
	region string
	Path    string
}

type GuideboxParams map[string]interface{}

func NewGuideboxClient(client *http.Client, APIkey string) *GuideboxClient {
	return &GuideboxClient{
		sling:  sling.New().Base(BaseURL),
		apiKey: APIkey,
	}
}

/*  Data Structs */

/* End Data Structs */

func (client *GuideboxClient) Request() (*http.Request, error) {
	return client.sling.Path(client.Path).Request()
}

func (client *GuideboxClient) GetShows(args ...interface{}) *GuideboxClient {
	client.Path =  fmt.Sprintf("%v%v", APIVersion, showsURL)
	return client
}

func (client *GuideboxClient) SetParams(params GuideboxParams) *sling.Sling {
	params["api_key"] = client.apiKey
	return client.sling.Path(client.Path).QueryStruct(params)
}

func (client *GuideboxClient) ShowId(id interface{}) *GuideboxClient {
	client.Path = fmt.Sprintf("%v/%v",client.Path, id)
	return client

}
