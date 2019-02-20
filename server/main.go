package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"os"
)

type JsonConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	RequestCount int    `json:"request_count"`
}

func loadConfig() JsonConfig {
	jsonFile, err := os.Open("server/config.json")
	if err != nil {
		fmt.Println("Failed to open server/config.json ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonConfig JsonConfig
	err = json.Unmarshal(byteValue, &jsonConfig)
	if err != nil {
		fmt.Println("Failed to unmarshal the JSON defined in server/config.json", err)
	}
	return jsonConfig
}

func handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	switch (path) {
	case "/test":
		method := string(ctx.Method())
		if method != "GET" && method != "DELETE" {
			fmt.Println(method)
			fmt.Fprint(ctx, "Body received: ", ctx.Request.Body())
		} else {
			fmt.Fprint(ctx, "My dog is cuter than yours. Runs faster, too.")
		}
	case "/connection/stats":
		fmt.Fprint(ctx, "Total number of requests made on this connection: ", ctx.ConnRequestNum())
	default:
		fmt.Fprint(ctx, "    **   ****     ** \n")
		fmt.Fprint(ctx, "   */*  *///**   */* \n")
		fmt.Fprint(ctx, "  * /* /*  */*  * /* \n")
		fmt.Fprint(ctx, " ******/* * /* ******\n")
		fmt.Fprint(ctx, "/////* /**  /*/////* \n")
		fmt.Fprint(ctx, "    /* /*   /*    /* \n")
		fmt.Fprint(ctx, "    /* / ****     /* \n")
		fmt.Fprint(ctx, "    /   ////      /  \n\n")
		fmt.Fprint(ctx, "Send a GET to /connection/stats for a total number of requests for an active connection.\n")
		fmt.Fprint(ctx, "Send a GET or DELETE to /test for the truth of my dog versus your dog.\n")
		fmt.Fprint(ctx, "Send a PUT or POST to /test for a print out of the request body sent.\n")
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func main() {
	jsonConfig := loadConfig()
	fmt.Println("Starting server, listening on " + fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port))

	err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port), handler)
	if err != nil {
		fmt.Print("The server is tired of serving. Exiting...")
		os.Exit(1)
	}
}
