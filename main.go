package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"simulation/simulateTxn/simulation"
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
	outputFile := flag.String("output-file", "", "")
	apiKey := flag.String("api-key", "PS0RMldS1a0GYN-cFxcyhBG_qPZSvAzh", "")
	flag.Parse()

	if *dataRaw == "" {
		fmt.Println("Error: data-raw is required")
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
			"domain": "https://boredapeyartclub.com",
		}
	}

	var callData string
	var toAddr string
	var value string
	var method string
	fromAddr := requestData.AccountAddress
	block := requestData.Block

	simService := simulation.NewBlockaidSimulationService(*apiKey)
	ctx := context.TODO()
	var simulationResult string

	if m, ok := requestData.Data["method"].(string); ok {
		method = m
	}

	if method == "eth_signTypedData_v4" {
		if params, ok := requestData.Data["params"].([]interface{}); ok && len(params) > 0 {
			if account, ok := params[0].(string); ok {
				fromAddr = account
			}
			if msg, ok := params[1].(string); ok {
				callData = msg
			}
		}
		simulationResult, err = simService.SimulateMessage(ctx, callData, fromAddr, block)
	}

	if params, ok := requestData.Data["params"].([]interface{}); ok && len(params) > 0 {
		if txParams, ok := params[0].(map[string]interface{}); ok {
			if data, ok := txParams["data"].(string); ok {
				callData = data
			}
			if to, ok := txParams["to"].(string); ok {
				toAddr = to
			}
			if val, ok := txParams["value"].(string); ok {
				value = val
			}
		}
	}

	if callData == "" {
		fmt.Println("Error: callData is missing in the input")
		return
	}

	fmt.Println("method:", method)
	if method == "eth_sendTransaction" {
		simulationResult, err = simService.SimulateTransaction(ctx, callData, fromAddr, toAddr, value, block)
	} else {
	}
	if err != nil {
		fmt.Println("Error during simulation:", err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(simulationResult), &result)
	if err != nil {
		fmt.Println("Error unmarshaling simulation result:", err)
		return
	}

	responseData := map[string]interface{}{
		"simulation": result,
	}

	output, err := json.MarshalIndent(responseData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		return
	}

	err = ioutil.WriteFile(*outputFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Response saved to %s\n", *outputFile)
}
