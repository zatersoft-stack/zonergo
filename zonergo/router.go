// zonergo/router.go
package zonergo

import (
    "log"
    "net/http"
)

type HandlerFunc func(*Context)
type Middleware func(HandlerFunc) HandlerFunc

type Router struct {
    routes      map[string]map[string]HandlerFunc // method -> path -> handler
    middlewares []Middleware
}

func NewRouter() *Router {
    return &Router{
        routes: make(map[string]map[string]HandlerFunc),
    }
}

func (r *Router) Use(mw ...Middleware) {
    r.middlewares = append(r.middlewares, mw...)
}

func (r *Router) addRoute(method, path string, handler HandlerFunc) {
    if _, ok := r.routes[method]; !ok {
        r.routes[method] = make(map[string]HandlerFunc)
    }
    r.routes[method][path] = handler
}

func (r *Router) GET(path string, handler HandlerFunc) {
    r.addRoute(http.MethodGet, path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
    r.addRoute(http.MethodPost, path, handler)
}

func (r *Router) PUT(path string, handler HandlerFunc) {
    r.addRoute(http.MethodPut, path, handler)
}

func (r *Router) DELETE(path string, handler HandlerFunc) {
    r.addRoute(http.MethodDelete, path, handler)
}

// Implementa http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    pathMap, ok := r.routes[req.Method]
    if !ok {
        http.NotFound(w, req)
        return
    }

    handler, ok := pathMap[req.URL.Path]
    if !ok {
        http.NotFound(w, req)
        return
    }

    ctx := &Context{
        Writer:  w,
        Request: req,
    }

    finalHandler := handler
    // Aplica middlewares em ordem reversa (clássico pattern)
    for i := len(r.middlewares) - 1; i >= 0; i-- {
        finalHandler = r.middlewares[i](finalHandler)
    }

    defer func() {
        if err := recover(); err != nil {
            log.Printf("[ZonerGo] panic recovered: %v", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
        }
    }()

    finalHandler(ctx)
}
