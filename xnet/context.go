package xnet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Context struct {
	sync.RWMutex

	Writer  http.ResponseWriter
	Request *http.Request

	Method  string
	URLPath string

	Engine *Engine

	middleware []HandlerFunc
	StatusCode int
	header     map[string][]string
}

type Map map[string]interface{}

// Status 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Path() string {
	return c.URLPath
}

// SetHeader 设置请求头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) JSON(code int, obj interface{}) error {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		return err
	}
	return nil
}

// 返回 format 字符串
func (c *Context) String(code int, format string, values ...interface{}) error {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		return err
	}
	return nil
}

// Data 返回 Data
func (c *Context) Data(code int, data []byte) error {
	c.Status(code)
	_, err := c.Writer.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// HTML 返回 HTML
func (c *Context) HTML(code int, html string) error {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		return err
	}
	return nil
}
