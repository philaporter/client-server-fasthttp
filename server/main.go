package main

import (
	"client-server-fasthttp/server/router"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"os"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Because reasons?
func loadConfig() ServerConfig {
	jsonFile, err := os.Open("server/config.json")
	if err != nil {
		fmt.Println("Failed to open server/config.json ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var serverConfig ServerConfig
	err = json.Unmarshal(byteValue, &serverConfig)
	if err != nil {
		fmt.Println("Failed to unmarshal the JSON defined in server/config.json", err)
	}
	return serverConfig
}

// Just for testing
func MethodHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("method:::", string(ctx.Method()))
	fmt.Println("path:::", string(ctx.Path()))
}

// temp for testing
func setupRouter() {
	routes := make(map[string]router.Route)
	routes["/temp"] = router.Route{
		Path:    "/temp",
		Method:  []byte("GET"),
		Handler: MethodHandler,
	}
	router.MasterChef.RouterFromMap(routes)
	router.MasterChef.PrintRoutes()
}

func main() {
	// All temp testing nonsense
	routes := make(map[string]router.Route)
	routes["/master"] = router.Route{
		Path:    "/master",
		Method:  []byte("GET"),
		Handler: MethodHandler,
	}

	routes["/chef"] = router.Route{
		Path:    "/chef",
		Method:  []byte("POST"),
		Handler: MethodHandler,
	}
	router.MasterChef.RouterFromMap(routes)
	router.MasterChef.PrintRoutes()

	// Just random unnecessary bullcrap
	jsonConfig := loadConfig()
	fmt.Println("Starting server, listening on " + fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port))

	// Awesomeness. Did you look at my router/router.go?
	err := router.ListenAndServe(fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port))

	if err != nil {
		fmt.Print("The server is tired of serving. Exiting...")
		os.Exit(1)
	}
}
