package snet

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/seefs001/seefslib/sync/errgroup"
)

type Engine struct {
	mutex sync.Mutex

	// router
	router *Route

	pool   sync.Pool
	addr   []string
	config Config
}

type Config struct {
	Mode         int
	Address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func New(config ...Config) *Engine {
	engine := &Engine{
		pool: sync.Pool{
			New: func() interface{} {
				return new(Context)
			},
		},
		addr: make([]string, 0),
		router: &Route{
			router: make(map[string]HandlerFunc),
			notFoundHandler: func(c *Context) error {
				return c.JSON(404, &Map{
					"msg": "not found router",
				})
			},
			errorHandler: func(c *Context) error {
				return c.JSON(500, &Map{
					"msg": "server err",
				})
			},
		},
		config: Config{},
	}

	engine.router.Engine = engine

	if len(config) > 0 {
		engine.config = config[0]
	}

	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := &Context{
		Method:  req.Method,
		URLPath: req.URL.Path,
		Writer:  w,
		Request: req,
		Engine:  e,
	}
	e.router.match(ctx)
}

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.router.Get(pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandlerFunc) {
	e.router.Post(pattern, handler)
}

func (e *Engine) Put(pattern string, handler HandlerFunc) {
	e.router.Put(pattern, handler)
}

func (e *Engine) Delete(pattern string, handler HandlerFunc) {
	e.router.Delete(pattern, handler)
}

func (e *Engine) Patch(pattern string, handler HandlerFunc) {
	e.router.Patch(pattern, handler)
}

func (e *Engine) Head(pattern string, handler HandlerFunc) {
	e.router.Head(pattern, handler)
}

func (e *Engine) Options(pattern string, handler HandlerFunc) {
	e.router.Options(pattern, handler)
}

func (e *Engine) AddAddr(addr string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.addr = append(e.addr, addr)
}

// Run 启动 HTTP 服务
func (e *Engine) Run(addr string) (err error) {
	g := errgroup.Group{}

	e.addr = append(e.addr, addr)

	fmt.Println(e.addr)
	for i := 0; i < len(e.addr); i++ {
		g.Go(func(context.Context) error {
			return http.ListenAndServe(e.addr[i], e)
		})
	}

	if err = g.Wait(); err != nil {
		return err
	}
	return nil
}
