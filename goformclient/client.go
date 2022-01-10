package goformclient

import (
	"net/http"
	"sync"
)

type httpClient struct{

	generator *clientGenerator
	client *http.Client
	clientOnce sync.Once

}

type Client interface {
	/*CRD operations*/
	Get(url string, headers http.Header)(*Response, error)
	Post(url string, headers http.Header, body interface{})(*Response, error)
	Delete(url string, headers http.Header)(*Response, error)

}

//Get to get the data from the defined server
func (c *httpClient) Get(url string, headers http.Header)(*Response, error){
	return c.do(http.MethodGet, url, headers, nil)
}
//Post Create a new entry to the defined server
func (c *httpClient) Post(url string, headers http.Header, body interface{})(*Response, error){
        return c.do(http.MethodPost, url, headers, body)
}
//Delete removes the entry from the server
func (c *httpClient) Delete(url string, headers http.Header)(*Response, error){
	return c.do(http.MethodDelete, url, headers, nil)
}
