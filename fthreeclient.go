package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Hello, world.")
	//httpMethod: = "GET"
	url := "http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	client := http.Client{}

	res, err := client.Get(url)

	if err != nil{
		panic(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(bytes))
}
