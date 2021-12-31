package goformclient

import (
	"testing"
	"net/http"
	"fmt"
)

func TestGetRequestHeader(t *testing.T){

	//Initializtion
	client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "http-client")
	client.Headers =commonHeaders

	//Execution
	reqHeaders := make(http.Header)
	reqHeaders.Set("X-Request-ID", "ABC-a23")

	finalHeader := client.getRequestHeaders(reqHeaders)
	//validation

	if len(finalHeader) !=3 {
		t.Error("Expected 3 headers")
	}

	if finalHeader.Get("X-Request-ID") != "ABC-a23"{
		t.Error("Invalid request Id")
	}
	if finalHeader.Get("User-Agent") != "http-client"{
		t.Error("Invalid request Id")
	}
	if finalHeader.Get("Content-Type") != "application/json"{
		t.Error("Invalid request Id")
	}
}

func TestGetRequestBody(t *testing.T){
	//Initializtion
        client := httpClient{}
        t.Run("nobodyNilResponse", func(t *testing.T){
		body, err := client.getRequestBody("",nil)
		if err !=nil{
			t.Error("No error expected")
		}
		if body != nil{
			 t.Error("nil body was expected")
		}
       })
       t.Run("JsonBody", func(t *testing.T){
	       requestBody := []string{"one", "two"}
	       body, err := client.getRequestBody("application/json", requestBody)
               if err !=nil{
                        t.Error("No error expected")
               }
	      fmt.Println(err)
	      fmt.Println(string(body))
              if string(body) != `["one","two"]` {
                       t.Error("Invalid json")
              }


       })

       t.Run("xmlBody", func(t *testing.T){
       })
       t.Run("defaultBody", func(t *testing.T){
       })
}
