package router

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Route struct {
	Path    string
	Method  []byte
	Handler func(ctx *fasthttp.RequestCtx)
}

type Router struct {
	Routes map[string]Route
}

var MasterChef Router = Router{}

// The Router function will set individual routes
func (r *Router) Router(path string, method []byte, handler func(ctx *fasthttp.RequestCtx)) {
	if r.Routes == nil {
		r.Routes = make(map[string]Route)
	}
	r.Routes[path] = Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
}

// This RouterFromMap function will set the map argument to Router.Routes
func (r *Router) RouterFromMap(routes map[string]Route) {
	if routes == nil {
		log.Println("Creating empty map because the argument was nil")
		r.Routes = make(map[string]Route)
	} else {
		r.Routes = routes
	}
}

// This PrintRoutes function will print out all of the stored routes
func (r *Router) PrintRoutes() {
	for _, v := range r.Routes {
		fmt.Println("Path:", v.Path)
		fmt.Println("Method:", string(v.Method))
	}
}

// This Handler function will be used to satisfy the fasthttp ListenAndServe requirements and will handle all routes
// defined in Router.Routes map
func Handler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	for k, v := range MasterChef.Routes {
		if bytes.Equal([]byte(k), path) && bytes.Equal(v.Method, ctx.Method()) {
			v.Handler(ctx)
		} else {
			continue
		}
	}
}

// Philip ListenAndServe wrapper for fasthttp.ListenAndServer
func ListenAndServe(addr string) error {
	if err := fasthttp.ListenAndServe(addr, Handler); err != nil {
		return err
	}
	return nil
}
