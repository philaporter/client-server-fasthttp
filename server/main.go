package main

import (
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

type Route struct {
	Path    string
	Method  []byte
	Handler func()
}

type Router struct {
	Routes map[string]Route
}

// temp var for testing
var MasterChef Router = Router{}

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

func (r *Router) Router(path string, method []byte, handler func()) {
	if r.Routes == nil {
		r.Routes = make(map[string]Route)
	}
	r.Routes[path] = Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

func (r *Router) RouterFromMap(routes map[string]Route) {
	if routes == nil {
		r.Routes = make(map[string]Route)
	} else {
		r.Routes = routes
	}
}

func (r *Router) PrintRoutes() {
	for k, v := range r.Routes {
		fmt.Printf("%s:%v", k, v)
	}
}

// temp for testing
func setupRouter() {
	routes := make(map[string]Route)
	routes["/temp"] = Route{
		Path:    "/temp",
		Method:  []byte("GET"),
		Handler: GetHandler,
	}
	MasterChef.RouterFromMap(routes)
	MasterChef.PrintRoutes()
}

func GetHandler() {
	fmt.Println("hey there, mr get")
}

func handler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	switch (path) {
	case "/temp":
		MasterChef.Routes["/temp"].Handler()
	case "/test":
		method := string(ctx.Method())
		if method != "GET" && method != "DELETE" {
			fmt.Println("Body received: ", string(ctx.Request.Body()))
			fmt.Fprint(ctx, "Body received: ", ctx.Request.Body())
			//fmt.Println("Total number of requests made on this connection: ", ctx.ConnRequestNum())
		} else {
			fmt.Println(ctx, "My dog is cuter than yours. Runs faster, too.")
			fmt.Fprint(ctx, "My dog is cuter than yours. Runs faster, too.")
		}
	case "/connection/stats":
		fmt.Fprint(ctx, "Total number of requests made on this connection: ", ctx.ConnRequestNum())
	default:
		fmt.Println("    **   ****     ** \n")
		fmt.Println("   */*  *///**   */* \n")
		fmt.Println("  * /* /*  */*  * /* \n")
		fmt.Println(" ******/* * /* ******\n")
		fmt.Println("/////* /**  /*/////* \n")
		fmt.Println("    /* /*   /*    /* \n")
		fmt.Println("    /* / ****     /* \n")
		fmt.Println("    /   ////      /  \n\n")
		fmt.Println("Send a GET to /connection/stats for a total number of requests for an active connection.\n")
		fmt.Println("Send a GET or DELETE to /test for the truth of my dog versus your dog.\n")
		fmt.Println("Send a PUT or POST to /test for a print out of the request body sent.\n")
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
	setupRouter()

	// TODO: Make the fasthttp handler iterate through the Routes to see what handler to call instead of using a switch

	jsonConfig := loadConfig()
	fmt.Println("Starting server, listening on " + fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port))
	err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", jsonConfig.Host, jsonConfig.Port), handler)

	if err != nil {
		fmt.Print("The server is tired of serving. Exiting...")
		os.Exit(1)
	}
}
