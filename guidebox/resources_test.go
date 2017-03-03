package guidebox

import (
	"net/http"
	"testing"
)

func TestNewGuideboxClient(t *testing.T) {

	client := &http.Client{}
	g := NewGuideboxClient(client, "HelloWorld")

	if g == nil {
		t.Error("Expected Guidebox")
	}
}
func TestGuideboxClient_SetParams(t *testing.T) {
	client := &http.Client{}
	g := NewGuideboxClient(client, "v2")
	var params GuideboxParams = GuideboxParams{}
	params.Hello = "world"
	req, err := g.GetShows().ShowId(6569).SetParams(params).Request()
	if err != nil {
		t.Error(err, "There was an error")
	}

	if req.URL.Host != "api-public.guidebox.com" {
		t.Error("Expeted api-public.guidebox.com got ", req.URL.Host)
	}

	if req.URL.Path != "/v2/shows/6569" {
		t.Error("Expected /v2/shows/6569 in path got ", req.URL.Path)
	}

	if req.URL.RawQuery == "" {
		t.Error("Expected a query string and got", req.URL.RawQuery)
	}

}

func TestGuideboxClient_GetShows(t *testing.T) {
	client := &http.Client{}
	g := NewGuideboxClient(client, "v2")

	req, err :=  g.GetShows().Request()

	if err != nil {
		t.Error(err, "There was an error")
	}

	if req.URL.Host != "api-public.guidebox.com" {
		t.Error("Expeted api-public.guidebox.com got ", req.URL.Host)
	}

	if req.URL.Path != "/v2/shows" {
	t.Error("Expected /v2/shows in path got ", req.URL.Path)
	}
}

func TestGuideboxClient_ShowId(t *testing.T) {
	client := &http.Client{}
	g := NewGuideboxClient(client, "Hello world")

	req, err := g.GetShows().ShowId(6569).Request()

	if err != nil {
		t.Error( "There was an error", err)
	}

	if req.URL.Path != "/v2/shows/6569" {
		t.Error("Expected /v2/shows/6569 in path got ", req.URL.Path)
	}
}

