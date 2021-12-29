package  main
//package  fthreeclient

import (
	"fmt"
	"io/ioutil"
	"github.com/supalik/fthreeclient/goformclient"
)

func main(){

	client := goformclient.NewClient()
	resp, err := client.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", nil)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("Response Code: %d \n", resp.StatusCode)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Response Data: %s", string(bytes))
}
