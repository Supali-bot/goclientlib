package goformclient

import (
	"net/http"
)

type httpClient struct{}

type HttpClient interface {
	Get(url string, headers http.Header)(* http.Response, error)
	Post(url string, headers http.Header, body interface{})(* http.Response, error)
	Delete(url string, headers http.Header)(* http.Response, error)

}

func NewClient() HttpClient{
	client := &httpClient{}
	return client
}

func (c *httpClient) Get(url string, headers http.Header)(* http.Response, error){
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{})(* http.Response, error){
        return c.do(http.MethodPost, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header)(* http.Response, error){
	return c.do(http.MethodDelete, url, headers, nil)
}
