package rest

import (
	"encoding/json"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/hdlopez/clean-architecture-golang/apierror"
)

type Client interface {
	Get(url string, h http.Header, v interface{}) (interface{}, error)
}

type client struct {
	readClient *resty.Client
}

func New() Client {
	return &client{
		resty.New(),
	}
}

func (api *client) Get(url string, h http.Header, v interface{}) (interface{}, error) {
	var r *resty.Response
	req := api.readClient.R()
	req.SetError(&apierror.APIError{})

	r, err := req.Get(url)

	if err != nil {
		// returns API error
		return nil, err
	}

	if r.StatusCode() != 200 {
		// returns API error
		return nil, apierror.New(r.StatusCode(), "Status code was not 200")
	}

	if err = json.Unmarshal(r.Body(), v); err != nil {
		// returns API error
		return nil, apierror.New(500, "Unmarshal error")
	}
	return v, nil
}
