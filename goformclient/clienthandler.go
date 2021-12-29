package goformclient

import (
        "fmt"
        "errors"
        "net/http"
)

func (c * httpClient) do(method string, url string, headrer http.Header, body interface{})(*http.Response, error) {

        fmt.Println("Client handler")

        client := http.Client{}
	request, err := http.NewRequest(method, url, nil)

        if err != nil{
                return nil,  errors.New("unable to handle request")
        }
	return client.Do(request)

}

