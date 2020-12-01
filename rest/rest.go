package rest

import (
	"encoding/json"
	"net/http"

	resty "github.com/go-resty/resty/v2"
	"github.com/hdlopez/greeting-api-golang/apierror"
)

type Client interface {
	Get(url string, h http.Header, v interface{}) error
}

type client struct {
	readClient *resty.Client
}

func New() Client {
	return &client{
		resty.New(),
	}
}

func (api *client) Get(url string, h http.Header, v interface{}) error {
	var r *resty.Response
	req := api.readClient.R()
	req.SetError(&apierror.APIError{})

	if h != nil {
		for k := range h {
			req.SetHeader(k, h.Get(k))
		}
	}

	r, err := req.Get(url)

	if err != nil {
		// returns API error
		return err
	}

	if r.StatusCode() != 200 {
		// returns API error
		return apierror.New(r.StatusCode(), "Status code was not 200")
	}

	if err = json.Unmarshal(r.Body(), v); err != nil {
		// returns API error
		return apierror.New(500, "Unmarshal error")
	}
	return nil
}
