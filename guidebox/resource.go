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

func (client *GuideboxClient) GetItem(itemID int) (Item, error) {
	item := Item{}
	request_url := itemURL + strconv.Itoa(itemID) + ".json"
	_, err := client.sling.New().Get(request_url).ReceiveSuccess(&item)

	if err != nil {
		return item, err
	}

	if item.ID == 0 {
		return item, errors.New("Invalid item ID, no such item exists.")
	}

	return item, err
}

func (client *GuideboxClient) GetUser(userID string) (User, error) {
	user := User{}
	request_url := userURL + userID + ".json"
	_, err := client.sling.New().Get(request_url).ReceiveSuccess(&user)

	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("Invalid user ID, user does not exist or has no activity.")
	}

	return user, err
}

func (client *GuideboxClient) GetTopStories() ([]int, error) {
	var topStories []int

	_, err := client.sling.New().Get(topURL).ReceiveSuccess(&topStories)

	if err != nil {
		return topStories, err
	}

	return topStories, nil
}

func (client *GuideboxClient) GetNewStories() ([]int, error) {
	var newStories []int

	_, err := client.sling.New().Get(newURL).ReceiveSuccess(&newStories)

	if err != nil {
		return newStories, err
	}

	return newStories, nil
}

func (client *GuideboxClient) GetMaxItem() (int, error) {
	var max int

	_, err := client.sling.New().Get(maxItemURL).ReceiveSuccess(&max)

	if err != nil {
		return max, err
	}

	return max, nil
}

func (client *GuideboxClient) GetBestStories() ([]int, error) {
	var bestStories []int

	_, err := client.sling.New().Get(bestURL).ReceiveSuccess(&bestStories)

	if err != nil {
		return bestStories, err
	}

	return bestStories, nil
}

func (client *GuideboxClient) GetShowGuideboxStories() ([]int, error) {
	var showGuideboxStories []int

	_, err := client.sling.New().Get(showURL).ReceiveSuccess(&showGuideboxStories)

	if err != nil {
		return showGuideboxStories, err
	}

	return showGuideboxStories, nil
}

func (client *GuideboxClient) GetAskGuideboxStories() ([]int, error) {
	var askGuideboxStories []int

	_, err := client.sling.New().Get(askURL).ReceiveSuccess(&askGuideboxStories)

	if err != nil {
		return askGuideboxStories, err
	}

	return askGuideboxStories, nil
}

func (client *GuideboxClient) GetJobStories() ([]int, error) {
	var jobStories []int

	_, err := client.sling.New().Get(jobsURL).ReceiveSuccess(&jobStories)

	if err != nil {
		return jobStories, err
	}

	return jobStories, nil
}

func (client *GuideboxClient) GetUpdate() (Update, error) {
	update := Update{}

	_, err := client.sling.New().Get(updatesURL).ReceiveSuccess(&update)

	if err != nil {
		return update, err
	}

	return update, nil
}
