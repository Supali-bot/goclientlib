package goformclient

import (
        "testing"
//        "net/http"
        "fmt"
	"github.com/google/uuid"
	"github.com/supalik/fthreeclient/models"
	"strings"
	"encoding/json"
//        "time"
)

func TestPost(t *testing.T){
	client := NewGenerator().
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
        inputData.DataStruct.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
        fmt.Printf("Data ID  ex: %v \n", inputData.DataStruct.ID)
	//Post(url string, headers http.Header, body interface{})(*Response, error){
	_, err := client.Post("http://localhost:8080/v1/organisation/accounts", nil, inputData)
	if err != nil{
		t.Error("nil body was expected")
	}
}

