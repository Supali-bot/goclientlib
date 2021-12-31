package goformclient

import (
//	"net"
	"net/http"
	"fmt"
	"time"
)

type httpClient struct{
	client *http.Client
	maxIdleConns int
	connectionTimeout  time.Duration
	responseTimeout  time.Duration
	Headers http.Header
}

type HttpClient interface {
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetRequestTimeout(timeout time.Duration)
	SetMaxIdleConns(i int)
	Get(url string, headers http.Header)(* http.Response, error)
	Post(url string, headers http.Header, body interface{})(* http.Response, error)
	Delete(url string, headers http.Header)(* http.Response, error)

}

func NewClient() HttpClient{
	httpClient := &httpClient{}
	return httpClient
}

func (c *httpClient) SetHeaders(headers http.Header){
     c.Headers = headers
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration){
     c.connectionTimeout = timeout
}

func (c *httpClient) SetRequestTimeout(timeout time.Duration){
     c.responseTimeout = timeout
}

func (c *httpClient) SetMaxIdleConns(i int){
     c.maxIdleConns = i
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
