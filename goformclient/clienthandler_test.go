package goformclient

import (
	"testing"
	"net/http"
	"fmt"
	"time"
)

func TestGetRequestHeader(t *testing.T){

	//Initializtion
	client := httpClient{}
	//client := httpClient{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "http-client")
	client.generator = &clientGenerator{
		headers: commonHeaders,
	}

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
func TestConnectionTimeot(t *testing.T){
	client := httpClient{}
	client.generator = &clientGenerator{
		connectionTimeout: 20,
	}
	timeOut := client.getConnectionTimeout()
	if timeOut != 20{
                        t.Error("Received Timeout ")
        }
	fmt.Printf("received timeout 20 %d",timeOut)
	client.generator = &clientGenerator{
                connectionTimeout: 0,
        }
	timeOut1 := client.getConnectionTimeout()
	fmt.Printf("\n received timeout 0 %d",timeOut1)
        if (int(time.Duration(timeOut1)) != 1000000000 ){
                t.Error("\n Received Timeout 0")
        }
	client.generator.disableTimeouts = true

	timeOut1 = client.getConnectionTimeout()
        if (int(time.Duration(timeOut1)) != 0 ){
                t.Error("\n Received Timeout 0")
        }
}
func TestResponseTimeout(t *testing.T){
        client := httpClient{}
        client.generator = &clientGenerator{
                responseTimeout: 20,
        }
        timeOut := client.getResponseTimeout()
        if timeOut != 20{
                        t.Error("Received Timeout ")
        }
        fmt.Printf("received timeout 20 %d",timeOut)
        client.generator = &clientGenerator{
                responseTimeout: 0,
        }
        timeOut1 := client.getResponseTimeout()
        fmt.Printf("\n received timeout 0 %d",timeOut1)
        if (int(time.Duration(timeOut1)) != 5000000000 ){
                t.Error("\n Received Timeout 0")
        }
        client.generator.disableTimeouts = true

        timeOut1 = client.getResponseTimeout()
        if (int(time.Duration(timeOut1)) != 0 ){
                t.Error("\n Received Timeout 0")
        }
}

func TestGetMaxIdleConnections(t *testing.T){
	client := httpClient{}
        client.generator = &clientGenerator{
                maxIdleConns: 20,
        }
        timeOut := client.getMaxIdleConnections()
        if timeOut != 20{
                        t.Error("Received Timeout ")
        }
	client.generator = &clientGenerator{
                maxIdleConns: 0,
        }
	timeOut = client.getMaxIdleConnections()
        if timeOut != 5{
                        t.Error("Received Timeout ")
	}

}
