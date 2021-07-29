package main

import (
	"fmt"
	"net/http"

	"github.com/seefs001/seefspkg/xnet"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	if err != nil {
		panic(err)
	}
}

func main() {
	engine := xnet.New()
	// http.HandleFunc("/",indexHandler)
	engine.Get("/", func(c *xnet.Context) error {
		return c.JSON(200, &xnet.Map{
			"/":    "1",
			"path": c.Path(),
		})
	})
	engine.Get("/test", func(c *xnet.Context) error {
		return c.JSON(200, &xnet.Map{
			"/":    "1",
			"path": c.Path(),
		})
	})
	engine.AddAddr(":9998")
	engine.Run(":8000")
}
