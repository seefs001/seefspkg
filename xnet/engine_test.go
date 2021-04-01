package xnet

import (
	"fmt"
	"testing"
)

func TestEngine_Get(t *testing.T) {
	engine := New()
	engine.Get("/api/v1/test/12", func(c *Context) error {
		return nil
	})
	//engine.Get("/api/v1/test/11", func(c *Context) error {
	//	return nil
	//})
	handlerFunc := engine.router.match("/api/v1/test/11")
	if handlerFunc != nil {
		fmt.Println("handlerfun found")
	}
	handlerFunc2 := engine.router.match("/api/v1/test/11")
	if handlerFunc2 != nil {
		fmt.Println("handlerfun2 found")
	}
}
