package guidebox

// Hacker News API Documentation:
// https://github.com/HackerNews/API

import (
	"net/http"

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
	C      *sling.Sling
}

type GuideboxParams map[string]interface{}

func NewGuideboxClient(client *http.Client, APIkey string) *GuideboxClient {
	return &GuideboxClient{
		sling:  sling.New().Client(client).Base(BaseURL).Path(APIVersion),
		apiKey: APIkey,
	}
}

/*  Data Structs */

/* End Data Structs */

func (client *GuideboxClient) GetShows(args ...interface{}) *sling.Sling {
	client.C = client.sling.New().Path(showsURL)
	return client.C
}

func (client *GuideboxClient) SetParams(params GuideboxParams) *sling.Sling {
	params["api_key"] = client.apiKey
	client.C = client.C.QueryStruct(params)
	return client.C
}

func (client *GuideboxClient) ShowId(id string) *sling.Sling {
	client.C = client.C.Path("/" + id)
	return client.C

}
