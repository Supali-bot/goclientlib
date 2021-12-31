package goformclient

import (
        "fmt"
        "errors"
        "net/http"
	"encoding/json"
        "encoding/xml"
        "strings"
	"bytes"
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

	client := http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

        if err != nil{
                return nil,  errors.New("unable to handle request")
        }

	request.Header = reqHeaders
        return client.Do(request)

}
func (c *httpClient) getRequestHeaders(requestsHeaders http.Header) http.Header {

	responseHeader := make(http.Header)
	//Common Request header
        for header, value := range c.Headers{
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
