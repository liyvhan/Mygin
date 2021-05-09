package gee

import (
	"net/http"
)

//handlerFunc defines the request handle used by gee
type handlerFunc func(ctx *Context)

//Engine implmment the interface of ServerHTTP
type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c:=newContex(writer,request)
	engine.router.handle(c)
}

//New is the constructor of Engine
func New() *Engine {
	return &Engine {router: newRouter()}
}

func (engine *Engine) addRoute (method string, pattern string, handler handlerFunc)  {
	engine.router.addRoute(method,pattern,handler)
}

//GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler handlerFunc) {
	engine.addRoute("GET",pattern,handler)
}

//POST defines the method to add POST requesst
func (engine *Engine) POST(pattern string, handler handlerFunc) {
	engine.addRoute("POST",pattern,handler)
}

//Run defines the method to start an httpserver
func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr,engine)
}


