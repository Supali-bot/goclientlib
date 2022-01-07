package goformclient

import (
        "net/http"
        //"fmt"
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

func NewGenerator() ClientGenerator {
        generator := &clientGenerator{}
        return generator
}

func (c *clientGenerator) Generate() Client {
	client := httpClient {
		generator : c,
	}

	return &client

}
func (c *clientGenerator) SetHeaders(headers http.Header) ClientGenerator {
     c.headers = headers
     return c
}

func (c *clientGenerator) SetConnectionTimeout(timeout time.Duration) ClientGenerator {
     c.connectionTimeout = timeout
     return c
}

func (c *clientGenerator) SetRequestTimeout(timeout time.Duration) ClientGenerator {
     c.responseTimeout = timeout
     return c
}

func (c *clientGenerator) SetMaxIdleConns(i int) ClientGenerator{
     c.maxIdleConns = i
     return c
}
func (c *clientGenerator) DisableTimeouts(disable bool) ClientGenerator {

     c.disableTimeouts = disable
     return c
}
func (c *clientGenerator) SetHttpClient(client *http.Client)ClientGenerator {

     c.client = client
     return c
}
