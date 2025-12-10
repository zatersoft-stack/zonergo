// zonergo/context.go
package zonergo

import (
    "encoding/json"
    "net/http"
)

type Context struct {
    Writer  http.ResponseWriter
    Request *http.Request
}

// JSON helper
func (c *Context) JSON(status int, data interface{}) {
    c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
    c.Writer.WriteHeader(status)
    _ = json.NewEncoder(c.Writer).Encode(data)
}

// Text helper
func (c *Context) Text(status int, text string) {
    c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
    c.Writer.WriteHeader(status)
    _, _ = c.Writer.Write([]byte(text))
}

// Status helper
func (c *Context) Status(status int) {
    c.Writer.WriteHeader(status)
}
