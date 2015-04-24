package gochallenge3

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

// simple wrapper for Instagram REST API
type ImageSource interface {
	Search(s string) []string
}


type InstagramImageSource struct {
	clientID string
}

func NewInstagramImageSource(clientID string) *InstagramImageSource {
	return &InstagramImageSource{clientID: clientID}
}

func (i *InstagramImageSource) Search(s string) ([]string, error) {
	u, err := i.instagramAPIUrl(s)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if body == nil {
		panic("oops")
	}

	var result []string
	return result, nil
}

func (i *InstagramImageSource) instagramAPIUrl(searchTag string) (string, error) {
	searchTag, err := UrlEncode(searchTag)
	if err != nil {
		return "", err
	}

	u, err := url.Parse("https://api.instagram.com/v1/tags")
	if err != nil {
		return "", err
	}

	u.Path += "/" + searchTag
	parameters := url.Values{}
	parameters.Add("client_id", i.clientID)
	u.RawQuery = parameters.Encode()

	return u.String(), nil
}