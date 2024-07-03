package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestData struct {
	Block          string                 `json:"block"`
	Chain          string                 `json:"chain"`
	AccountAddress string                 `json:"account_address"`
	Data           map[string]interface{} `json:"data"`
	Options        []string               `json:"options"`
	Metadata       map[string]string      `json:"metadata"`
}

func main() {
	dataRaw := flag.String("data-raw", "", "")
	jwtToken := flag.String("jwt", "", "")
	flag.Parse()

	if *dataRaw == "" {
		fmt.Println("Error: data-raw is required")
		return
	}

	if *jwtToken == "" {
		fmt.Println("Error: jwt token is required")
		return
	}

	var requestData RequestData
	err := json.Unmarshal([]byte(*dataRaw), &requestData)
	if err != nil {
		fmt.Println("Error parsing data-raw:", err)
		return
	}

	requestData.Options = []string{"simulation"}

	if requestData.Metadata == nil {
		requestData.Metadata = map[string]string{
			"domain": "https://shibicoin.org",
		}
	}

	updatedDataRaw, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling updated data-raw:", err)
		return
	}

	url := "https://demo.blockaid.io/api/tx/jsonRpc"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(updatedDataRaw))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", *jwtToken))
	req.Header.Set("content-type", "application/json")

	startTime := time.Now()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(startTime)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
	fmt.Println("Request duration:", duration)
}
