package gee

import (
	"log"
	"net/http"
)

//handlerFunc defines the request handle used by gee
type handlerFunc func(ctx *Context)

//RouterGroup 路由分组，以使得不同分组可以扩展中间件
type RouterGroup struct {
	prefix      string        //
	middlewares []handlerFunc //support middleware
	parent      *RouterGroup  //support nesting
	engine      *Engine       //all group share a Engine instance (所有分组共享一个路由引擎)
}

//Engine implmment the interface of ServerHTTP
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup //store all groups
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContex(writer, request)
	engine.router.handle(c)
}

//New is the constructor of Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) addRoute(method string, comp string, handler handlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

//GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler handlerFunc) {
	group.addRoute("GET", pattern, handler)
}

//POST defines the method to add POST requesst
func (group *RouterGroup) POST(pattern string, handler handlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//Run defines the method to start an httpserver
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
