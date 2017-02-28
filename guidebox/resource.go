package guidebox

// Hacker News API Documentation:
// https://github.com/HackerNews/API

import (
	"github.com/dghubble/sling"
	"net/http"
	"strconv"
	"errors"
	"strings"
	"fmt"
	"path"
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
}

func NewGuideboxClient(client *http.Client, APIkey string, args ...interface{}) *GuideboxClient {
	region := "US"
	if args["region"] != nil {
		region = args["region"].(string)
	}
	return &GuideboxClient{
		sling:  sling.New().Client(client).Base(BaseURL).Path(APIVersion),
		apiKey: APIkey,
		region: region,
	}
}

/*  Data Structs */

/* End Data Structs */

func (client *GuideboxClient) GetShows(args ...interface{}) (result map[string]interface{}, err error) {
	var offset string
	var limit string
	var sources string
	var platform string
	var tags []string

	params := map[string]interface{}{}

	params["api_key"] = client.apiKey
	switch {
	case offset:
		params["offset"] = offset
	case limit:
		params["limit"] = limit
	case sources:
		params["sources"] = sources
	case platform:
		params["platform"] = platform
	case tags:
		params["tags"] = strings.Join(tags, ",")
	}

	_, err = client.sling.New().Get(showsURL).QueryStruct(params).ReceiveSuccess(&result)

	if err != nil{
		return nil, err
	}

	return result, err
}



