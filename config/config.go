package config

import (
	"fmt"
	"net/http"
	"net/url"
)

// FoxbitApi represents a Foxbit API Client connection
type FoxbitApi = FoxbitAPI

var publicMethods = []string{
	"OMSId",
}

// FoxbitApi represents a Foxbit API Client connection
type FoxbitAPI struct {
	key    string
	secret string
	client *http.Client
}

// New creates a new connection
func New(key, secret string) (*FoxbitAPI, error) {
	foxbitAPI := FoxbitAPI{
		key:    key,
		secret: secret,
		client: http.DefaultClient,
	}
	return &foxbitAPI, nil
}

func (api *FoxbitAPI) Query(endPoint string, data map[string]string) (interface{}, error) {
	return queryPublic(endPoint, values, nil)
}

func (api *FoxbitAPI) queryPublic(method string, values url.Values, typ interface{}) (interface{}, error) {
	url := fmt.Sprintf("%s/%s/public/%s", APIURL, APIVersion, method)
	resp, err := api.makeWs(url, values)

	return resp, err
}
