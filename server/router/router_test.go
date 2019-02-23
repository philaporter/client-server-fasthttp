package router

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"testing"
)

// Filler
func handler(ctx *fasthttp.RequestCtx) {}

func TestRouter(t *testing.T) {
	t.Log("Test adding individual routes in Router(path string, method []byte, handler func(ctx *fasthttp.RequestCtx)")
	router := Router{}
	router.Router("/sample", []byte("GET"), handler)
	if router.Routes["/sample"].Path != "/sample" {
		t.Fail()
	}
	if !bytes.Equal(router.Routes["/sample"].Method, []byte("GET")) {
		t.Fail()
	}

	t.Log("Sneak some code coverage in for my void, mostly pointless function")
	router.PrintRoutes()
}

func TestRouterFromMap(t *testing.T) {
	t.Log("Test Creating empty map because the argument was nil")
	router := Router{}
	router.RouterFromMap(nil)
	if router.Routes == nil {
		t.Fail()
	}

	t.Log("Test setting the routes that were passed")
	routes := make(map[string]Route)
	routes["/sample"] = Route{
		Path:   "/sample",
		Method: []byte("GET"),
	}
	router.RouterFromMap(routes)
	if "/sample" != router.Routes["/sample"].Path {
		t.Fail()
	}
}

// TODO: Fix this piece of shit
//func TestListenAndServe(t *testing.T) {
//	t.Log("Test fasthttp.ListenAndServe wrapper")
//	go func() {
//		time.Sleep(5000)
//		request := fasthttp.AcquireRequest()
//		response := fasthttp.AcquireResponse()
//		client := fasthttp.Client{}
//
//		// Build the request
//		request.Header.SetHost("localhost:8080")
//		request.SetRequestURI("/sample")
//		request.Header.SetMethodBytes([]byte("GET"))
//		request.Header.SetContentType("application/json")
//
//		err := client.Do(request, response)
//		if err != nil {
//			t.Fail()
//		} else {
//			// How the fuck do I kill this?
//		}
//	}()
//
//	routes := Router{}
//	routes.Routes = make(map[string]Route)
//	routes.Routes["/sample"] = Route{
//		Path:    "/sample",
//		Method:  []byte("GET"),
//		Handler: handler,
//	}
//	err := ListenAndServe("localhost:8080")
//	if err != nil {
//		t.Fail()
//	}
//}
