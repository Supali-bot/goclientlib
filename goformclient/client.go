package goformclient

import (
//	"net"
	"net/http"
	"fmt"
//	"time"
)

type httpClient struct{

	generator *clientGenerator
	client *http.Client

}

type Client interface {
	/*CRD operations*/
	Get(url string, headers http.Header)(* http.Response, error)
	Post(url string, headers http.Header, body interface{})(* http.Response, error)
	Delete(url string, headers http.Header)(* http.Response, error)

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
