package goformclient

import (
	"testing"
	"net/http"
	"time"
)

func TestGetRequestHeader(t *testing.T){

	//Initializtion
	client := httpClient{}
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
              if string(body) != `["one","two"]` {
                       t.Error("Invalid json")
              }


       })

}
func TestTimeOut(t *testing.T){
	client := httpClient{}
	t.Run("ConnectionTimeOut", func(t *testing.T){
		client.generator = &clientGenerator{
		        connectionTimeout: 20,
		}
		timeOut := client.getConnectionTimeout()
		if timeOut != 20{
			t.Error("Invalid Connection Timeout :", timeOut)
		}
		client.generator = &clientGenerator{
			connectionTimeout: 0,
		}
		timeOut = client.getConnectionTimeout()
		if (int(time.Duration(timeOut)) != 1000000000 ){
			t.Error("Invalid Connection Timeout :", timeOut)
		}
		client.generator.disableTimeouts = true
		timeOut = client.getConnectionTimeout()
		if (int(time.Duration(timeOut)) != 0 ){
			t.Error("Invalid Connection Timeout :", timeOut)
		}
       })
       t.Run("ResponseTimeOut", func(t *testing.T){
		client.generator = &clientGenerator{
                        responseTimeout: 20,
	        }
		timeOut := client.getResponseTimeout()
	        if timeOut != 20{
			t.Error("Invalid Response Timeout: ", timeOut)
	        }
	        client.generator = &clientGenerator{
	                responseTimeout: 0,
	        }
	        timeOut = client.getResponseTimeout()
	        if (int(time.Duration(timeOut)) != 5000000000 ){
			t.Error("\n Invalid Response Time : ", timeOut)
	        }
	        client.generator.disableTimeouts = true
	        timeOut = client.getResponseTimeout()
	        if (int(time.Duration(timeOut)) != 0 ){
			t.Error("\n Invalid Response Time : ", timeOut)
	        }

       })

}

func TestGetMaxIdleConnections(t *testing.T){
	client := httpClient{}
        client.generator = &clientGenerator{
                maxIdleConns: 20,
        }
        connections := client.getMaxIdleConnections()
        if connections != 20{
		t.Error("Invalid Connection count: ", connections)
        }
	client.generator = &clientGenerator{
                maxIdleConns: 0,
        }
	connections = client.getMaxIdleConnections()
        if connections != 5{
		t.Error("Invalid connection count: ", connections)
	}

}
func TestCustomClient(t *testing.T){
	customclient := http.Client{}
	client := NewGenerator().
		SetHttpClient(&customclient).
                Generate()
	_, err := client.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", nil)
        if err !=nil{
		t.Error("Failed with Customclient: ", err)
        }

}
