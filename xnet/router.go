package xnet

import (
	"net/http"
	"sync"
)

type Router interface {
	Get(path string, handlers HandlerFunc) Router
	Post(path string, handlers HandlerFunc) Router
	Put(path string, handlers HandlerFunc) Router
	Delete(path string, handlers HandlerFunc) Router
	Options(path string, handlers HandlerFunc) Router

	Add(method, path string, handlers HandlerFunc) Router
	All(path string, handlers HandlerFunc) Router
	Group(prefix string, handlers HandlerFunc) Router
	Mount(prefix string, engine *Engine) Router
}

type HandlerFunc func(c *Context) error

type Route struct {
	sync.RWMutex

	Engine     *Engine
	middleware []HandlerFunc
	//trees      *Tree
	// route数量
	routesCount int
	// handler数量
	//handlerCount int
	router          map[string]HandlerFunc
	notFoundHandler HandlerFunc
	errorHandler    HandlerFunc
}

type Tree struct {
	root *Node
}

type Node struct {
	key string
	// path records a request path
	path   string
	method string
	handle HandlerFunc
	// depth records Node's depth
	depth int
	// children records Node's children node
	children map[string]*Node
	// isPattern flag
	isPattern bool
	// middleware records middleware stack
	middleware []HandlerFunc
}

func NewNode(key string, depth int, handler HandlerFunc) *Node {
	return &Node{
		key:      key,
		depth:    depth,
		children: make(map[string]*Node),
	}
}

func NewTree() *Tree {
	return &Tree{
		root: NewNode("/", 1, func(c *Context) error {
			return nil
		}),
	}
}

func (r *Route) addRoute(method, path string, handler HandlerFunc) {
	//pathSeg := strings.Split(path, "/")
	//tree := r.trees.root
	//for i, s := range pathSeg {
	//	if i == 0 {
	//		continue
	//	}
	//	tree.children[s] = NewNode(s, i,nil)
	//	tree.children[s].method = method
	//	tree.children[s].handle = handler
	//	tree = tree.children[s]
	//}
	//return *r
	key := method + "-" + path
	r.router[key] = handler

	r.Lock()
	defer r.Unlock()
	r.routesCount++
}

func (r *Route) Get(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodGet, pattern, handler)
}

func (r *Route) Post(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodPost, pattern, handler)
}

func (r *Route) Put(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodPut, pattern, handler)
}

func (r *Route) Delete(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodDelete, pattern, handler)
}

func (r *Route) Patch(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodPatch, pattern, handler)
}

func (r *Route) Head(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodHead, pattern, handler)
}

func (r *Route) Options(pattern string, handler HandlerFunc) {
	r.addRoute(http.MethodOptions, pattern, handler)
}

func (r *Route) match(ctx *Context) {
	key := ctx.Method + "-" + ctx.URLPath
	if handler, ok := r.router[key]; ok {
		if err := handler(ctx); err != nil {
			_ = r.errorHandler(ctx)
		}
	} else {
		_ = r.notFoundHandler(ctx)
	}

	//pathSeg := strings.Split(path, "/")
	//tree := r.trees.root
	//
	//for i, s := range pathSeg {
	//	if i == 0 {
	//		continue
	//	}
	//	_, ok := tree.children[s]
	//	if ok {
	//		tree = tree.children[s]
	//	}else {
	//		return nil
	//	}
	//}
	//return tree.handle
}
