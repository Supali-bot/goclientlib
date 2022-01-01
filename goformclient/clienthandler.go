package goformclient

import (
        "fmt"
        "errors"
	"net"
        "net/http"
	"encoding/json"
        "encoding/xml"
        "strings"
	"bytes"
	"time"
)
const (
        defaultMaxIdleConnections = 5
	defaultResponseTimeout = 5 * time.Second
	defaultConnectionTimeout = 2 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{})([]byte, error){
	if body == nil{
		return nil, nil
	}
        fmt.Println("Content-type in getRequestBody %v", contentType)
	switch strings.ToLower(contentType){
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
        default:
		return json.Marshal(body)
	}

}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{})(*http.Response, error) {

        fmt.Println("Inside Client handler")

	reqHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(reqHeaders.Get("Content-type"), body)

        if err != nil{
                return nil,  err
        }

	//client := http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

        if err != nil{
                return nil,  errors.New("unable to handle request")
        }

	request.Header = reqHeaders

	client := c.getHttpClient()
        return client.Do(request)

}
func (c *httpClient) getMaxIdleConnections() int {
        if c.generator.maxIdleConns >0 {
		return c.generator.maxIdleConns
	}

	return defaultMaxIdleConnections

}
func (c *httpClient) getResponseTimeout() time.Duration {
        if c.generator.responseTimeout >0 {
                return c.generator.responseTimeout
        }

	if c.generator.disableTimeouts{
		return 0
	}

        return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
        if c.generator.connectionTimeout >0 {
                return c.generator.connectionTimeout
        }

	if c.generator.disableTimeouts{
		return 0
	}
        return defaultConnectionTimeout
}


func (c *httpClient) getHttpClient() *http.Client{
	if c.client != nil{
		return c.client
	}
	c.client = &http.Client{
		Timeout:  c.getConnectionTimeout() + c.getResponseTimeout(),
                Transport: &http.Transport{
                        MaxIdleConns:          c.getMaxIdleConnections(),
                        ResponseHeaderTimeout: c.getResponseTimeout(),
                        DialContext: (&net.Dialer{
                                Timeout: c.getConnectionTimeout(),
                        }).DialContext,
                },
        }

	return c.client

}
func (c *httpClient) getRequestHeaders(requestsHeaders http.Header) http.Header {

	responseHeader := make(http.Header)
	//Common Request header
        for header, value := range c.generator.headers{
                if len(value) >0{
                        responseHeader.Set(header, value[0])
                }
        }
        //Custome Request header
        for header, value := range requestsHeaders{
                if len(value) >0{
                        responseHeader.Set(header, value[0])
                }
        }
        return responseHeader

}
