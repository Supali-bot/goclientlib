package goformclient

import (
	"net/http"
	"fmt"
)

type httpClient struct{
	Headers http.Header
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header)(* http.Response, error)
	Post(url string, headers http.Header, body interface{})(* http.Response, error)
	Delete(url string, headers http.Header)(* http.Response, error)

}

func NewClient() HttpClient{
	client := &httpClient{}
	return client
}

func (c *httpClient) SetHeaders(headers http.Header){
     c.Headers = headers
}
func (c *httpClient) Get(url string, headers http.Header)(* http.Response, error){
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{})(* http.Response, error){
	fmt.Println("Body %v", body)
        return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header)(* http.Response, error){
	return c.do(http.MethodDelete, url, headers, nil)
}
