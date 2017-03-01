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


func TestGuideboxClient_GetShows(t *testing.T) {
	client := &http.Client{}
	g := NewGuideboxClient(client, "hello world")

	req, err :=  g.GetShows().Request()

	if err != nil {
		t.Error(err, "There was an error")
	}

	if req.URL.Host != "api-public.guidebox.com" {
		t.Error("Expeted api-public.guidebox.com got ", req.URL.Host)
	}

	if req.URL.Path != "/shows" {
	t.Error("Expected /shows in path got ", req.URL.Path)
	}
}