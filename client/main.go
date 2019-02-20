package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"os"
)

const (
	MethodPost      = "POST"
	ContentTypeJson = "application/json"
)

type JsonRequest struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Count  int    `json:"count"`
}

type ClientConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	RequestCount int    `json:"request_count"`
	Uri          string `json:"uri"`
}

func loadConfig() ClientConfig {
	jsonFile, err := os.Open("client/config.json")
	if err != nil {
		fmt.Println("Failed to open client/config.json ", err)
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonConfig ClientConfig
	err = json.Unmarshal(byteValue, &jsonConfig)
	if err != nil {
		fmt.Println("Failed to unmarshal the JSON defined in client/config.json", err)
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	return jsonConfig
}

func main() {
	clientConfig := loadConfig()
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	client := fasthttp.Client{}

	// Send the configured number of requests
	for i := 0; i < clientConfig.RequestCount; i++ {

		// Build the json request body
		reqA := &JsonRequest{
			Status: "OK",
			Id:     fmt.Sprintf("%d", i+1) + "-EIEIO",
			Count:  i + 1,
		}
		reqB, err := json.Marshal(reqA)
		if err != nil {
			fmt.Println("Error marshaling the json request body: ", err)
			fmt.Println("Exiting...")
			os.Exit(1)
		}

		// Build the request
		request.Header.SetHost(fmt.Sprintf("%s:%d", clientConfig.Host, clientConfig.Port))
		request.SetBody(reqB)
		request.SetRequestURI(clientConfig.Uri)
		request.Header.SetMethodBytes([]byte(MethodPost))
		request.Header.SetContentType(ContentTypeJson)

		err = client.Do(request, response)
		if err != nil {
			fmt.Println("Error sending request ", err)
			fmt.Println("Exiting...")
			os.Exit(1)
		} else {
			fmt.Println("Successfully sent ", string(request.Body()))
		}
	}

}
