package goformclient

import (
        "net/http"
        "time"
)

type clientGenerator struct {

	 headers            http.Header
	 maxIdleConns       int
         connectionTimeout  time.Duration
         responseTimeout    time.Duration
         disableTimeouts    bool
	 client             *http.Client
}

type ClientGenerator  interface {
	 /*Http Header*/
        SetHeaders(headers http.Header) ClientGenerator
        /*Timoute Configuration*/
        DisableTimeouts(disableTimeouts bool) ClientGenerator
        SetConnectionTimeout(timeout time.Duration) ClientGenerator
        SetRequestTimeout(timeout time.Duration) ClientGenerator
        SetMaxIdleConns(i int)ClientGenerator
	SetHttpClient(c *http.Client)ClientGenerator
	Generate() Client

}
//NewGenerator is a  Client generator 
func NewGenerator() ClientGenerator {
        generator := &clientGenerator{}
        return generator
}
//Generate returns final client	
func (c *clientGenerator) Generate() Client {
	client := httpClient {
		generator : c,
	}

	return &client

}
//SetHeaders to set the user configured header
func (c *clientGenerator) SetHeaders(headers http.Header) ClientGenerator {
     c.headers = headers
     return c
}
//SetConnectionTimeout configures the max time for establishing the connection
func (c *clientGenerator) SetConnectionTimeout(timeout time.Duration) ClientGenerator {
     c.connectionTimeout = timeout
     return c
}
//SetRequestTimeout configures time to wait for a response headers from the server
func (c *clientGenerator) SetRequestTimeout(timeout time.Duration) ClientGenerator {
     c.responseTimeout = timeout
     return c
}
// SetMaxIdleConns configures user defined maximum number of idle
func (c *clientGenerator) SetMaxIdleConns(i int) ClientGenerator{
     c.maxIdleConns = i
     return c
}
// DisableTimeouts disable the timeouts
func (c *clientGenerator) DisableTimeouts(disable bool) ClientGenerator {

     c.disableTimeouts = disable
     return c
}

// SetHttpClient  provides functionality to use the customize client rather than creating new one
func (c *clientGenerator) SetHttpClient(client *http.Client)ClientGenerator {

     c.client = client
     return c
}
