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
        fmt.Printf("\n Raw Data : %v \n", raw_data)
        json.Unmarshal([]byte(raw_data), &inputData)
        inputData.DataStruct.ID = uuid
	g_uuid = inputData.DataStruct.ID
	g_version = inputData.DataStruct.Version
        inputData.DataStruct.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
        fmt.Printf("Data ID  ex: %v \n", inputData.DataStruct.ID)
	resp, err := client.Post("http://localhost:8080/v1/organisation/accounts", nil, inputData)
	if err != nil || (resp.StatusCode() != 201) {
		t.Error("nil body was expected")
	}
        fmt.Printf("\nPOST response Code: %d \n", resp.StatusCode())
        fmt.Printf("\nPOST response Code: %s \n", resp.Status())
        fmt.Printf("\nPOST response Code: %s \n", resp.String())
}
func TestPostWithDefaultandXmlHeaders(t *testing.T){
        updatedHeader := make(http.Header)
        updatedHeader.Set("Content-Type", "application/xml")
        updatedHeader.Set("User-Agent", "http-client")
        client := NewGenerator().
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()
        raw_data := ` {
	}`
	inputData :=models.Data{}
        json.Unmarshal([]byte(raw_data), &inputData)
        resp, err := client.Post("http://localhost:8080/v1/organisation/accounts", nil, inputData)
        if err != nil{
                t.Error("nil body was expected")
        }
	client = NewGenerator().
                SetHeaders(updatedHeader).
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()

	resp, err = client.Post("http://localhost:8080/v1/organisation/accounts", nil, inputData)
        if err != nil{
                t.Error("nil body was expected")
        }

        fmt.Printf("\nPOST1 response Code: %d \n", resp.StatusCode())
        fmt.Printf("\nPost response Code: %s \n", resp.Status())
        fmt.Printf("\nPost response Code: %s \n", resp.String())

}
func TestGet(t *testing.T){
	client := NewGenerator().
                DisableTimeouts(true).
                SetMaxIdleConns(20).
                Generate()
	resp, err := client.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", nil)
        if err !=nil{
		t.Error("nil body was expected")
        }
        fmt.Printf("Get Response Code: %d \n", resp.StatusCode())
        bytes := resp.Bytes()
        if  resp.StatusCode() != 200{
		t.Error("status 200  was expected")
        }
        fmt.Printf("\nGet Response Data: %s", string(bytes))
	recvData :=models.Data{}
	err = resp.UnmarshalJson(&recvData)
	if err !=nil {
		t.Error("Error was not expected")
	}
	fmt.Printf("Unmarshaled Data %v", recvData)
        fmt.Printf("\nGET response Code: %d \n", resp.StatusCode())
        fmt.Printf("\nGET response Code: %s \n", resp.Status())
        fmt.Printf("\nGET response Code: %s \n", resp.String())

}
func TestDelete(t *testing.T){
	
	client := NewGenerator().
                SetRequestTimeout(5 * time.Second).
		SetConnectionTimeout(1 * time.Second).
                SetMaxIdleConns(20).
                Generate()
        url := "http://localhost:8080/v1/organisation/accounts/" + g_uuid + "/" + "?" + "version=" + strconv.Itoa(g_version)
        resp, err := client.Delete(url, nil)
        if err !=nil{
		t.Error("no error was expected")
        }
        fmt.Printf("\nDelete response Code: %d \n", resp.StatusCode())
        fmt.Printf("\nDelete response Code: %s \n", resp.Status())
        fmt.Printf("\nDelete response Code: %s \n", resp.String())

}

