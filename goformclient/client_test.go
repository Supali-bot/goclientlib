package goformclient

import (
        "testing"
        "net/http"
        "fmt"
	"github.com/google/uuid"
	"github.com/supalik/fthreeclient/models"
	"strings"
	"encoding/json"
	"strconv"
        "time"
)
var (
	g_uuid string
	g_version int
)
func TestPost(t *testing.T){
	updatedHeader := make(http.Header)
        updatedHeader.Set("Content-Type", "application/json")
        updatedHeader.Set("User-Agent", "http-client")
	client := NewGenerator().
		SetHeaders(updatedHeader).
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()
        uuidWithHyphen := uuid.New()
        uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	url := "http://localhost:8080/v1/organisation/accounts"
	t.Run("POST Valid JSON Data", func(t *testing.T){
	        raw_data := ` {
	          "data": {
	            "type": "accounts",
	            "id": "4538c8264f6640f5bbfe734cf1f5c981",
	            "version": 0,
	            "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	            "attributes": {
	              "country": "GB",
	              "base_currency": "GBP",
	              "account_number": "41426819",
	              "bank_id": "400300",
	              "bank_id_code": "GBDSC",
	              "bic": "NWBKGB22",
	              "iban": "GB11NWBK40030041426819",
	                  "name": [
	                "Samantha Holder1"
	              ],
	              "status": "confirmed"
	            }
	          }
	        }`
	        inputData :=models.Data{}
	        json.Unmarshal([]byte(raw_data), &inputData)
	        inputData.DataStruct.ID = uuid
		g_uuid = inputData.DataStruct.ID
		g_version = inputData.DataStruct.Version
	        inputData.DataStruct.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	        fmt.Printf("Data ID  ex: %v \n", inputData.DataStruct.ID)
		resp, err := client.Post(url, nil, inputData)
		if err != nil || (resp.StatusCode() != http.StatusCreated) {
			t.Error("Error nil was expected recived Code:", resp.StatusCode(), "Expected:", http.StatusCreated)
		}
	})
}
func TestPostWithDefaultandXmlHeaders(t *testing.T){
	/*Header with default data fromat option with invalid data*/
        client := NewGenerator().
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()
        url := "http://localhost:8080/v1/organisation/accounts"
        raw_data := ` {
	}`
	inputData :=models.Data{}
        json.Unmarshal([]byte(raw_data), &inputData)
        resp, err := client.Post(url, nil, inputData)
        if err != nil || resp.StatusCode() != http.StatusBadRequest{
		t.Error("Error recived :", err)
		t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusBadRequest)
        }
	/*Header for XML data format  */
        updatedHeader := make(http.Header)
        updatedHeader.Set("Content-Type", "application/xml")
        updatedHeader.Set("User-Agent", "http-client")
	client = NewGenerator().
                SetHeaders(updatedHeader).
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()

	resp, err = client.Post(url, nil, inputData)
        if err != nil || resp.StatusCode() != http.StatusBadRequest {
		t.Error("Error recived :", err)
		t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusBadRequest)
        }


}
func TestGet(t *testing.T){
	client := NewGenerator().
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()
	t.Run("\nValid GET Request", func(t *testing.T){
		url := "http://localhost:8080/v1/organisation/accounts/" + g_uuid
		resp, err := client.Get(url, nil)
	        if err !=nil || (resp.StatusCode() != http.StatusOK){
			t.Error("Error recived :", err)
			t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusOK)
	        }
	        bytes := resp.Bytes()
	        if string(bytes) == ""{
			t.Error("Received empty body")
	        }
		/*If success unmarshal the received Data*/
		recvData :=models.Data{}
		err = resp.UnmarshalJson(&recvData)
		if err !=nil {
			t.Error("Error was not expected", err)
		}
		fmt.Printf("Unmarshaled Data:  %v", recvData)
	})
	t.Run("Valid GET Request with Invalid ID", func(t *testing.T){
		url := "http://localhost:8080/v1/organisation/accounts/" + "4538c8264f6640f5bbfe734cf1f5c981"
                resp, err := client.Get(url, nil)
                if err !=nil || (resp.StatusCode() != http.StatusNotFound){
                        t.Error("Error recived :", err)
                        t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusNotFound)
                }

	})

}
func TestDelete(t *testing.T){
	client := NewGenerator().
                SetRequestTimeout(5 * time.Second).
		SetConnectionTimeout(1 * time.Second).
                SetMaxIdleConns(20).
                Generate()
		baseUrl := "http://localhost:8080/v1/organisation/accounts/"
	t.Run("Valid DELETE Request", func(t *testing.T){
	        url := baseUrl + g_uuid + "/" + "?" + "version=" + strconv.Itoa(g_version)
	        resp, err := client.Delete(url, nil)
	        if err !=nil || (resp.StatusCode() != http.StatusNoContent){
                        t.Error("Error recived :", err)
                        t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusNoContent)
	        }
	})
	t.Run("Delete with Invalid Vesrion", func(t *testing.T){
	        url := baseUrl + g_uuid + "/" + "?" 
	        resp, err := client.Delete(url, nil)
	        if err !=nil || (resp.StatusCode() != http.StatusBadRequest){
                        t.Error("Error recived :", err)
                        t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusBadRequest)
	        }
	})
	t.Run("Delete with Invalid id", func(t *testing.T){
	        url := baseUrl + "4538c8264f6640f5bbfe734cf1f5c981" + "/" + "?" + "version=" + strconv.Itoa(g_version)
	        resp, err := client.Delete(url, nil)
	        if err !=nil || (resp.StatusCode() != http.StatusNotFound){
                        t.Error("Error recived :", err)
                        t.Error("Recived Status Code:", resp.StatusCode(), "Expected:", http.StatusNotFound)
	        }
	})

}
