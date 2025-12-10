// zonergo/middleware.go
package zonergo

import (
    "log"
    "time"
)

// Logger registra request simples
func Logger() Middleware {
    return func(next HandlerFunc) HandlerFunc {
        return func(c *Context) {
            start := time.Now()
            req := c.Request
            next(c)
            log.Printf("[ZonerGo] %s %s - %v", req.Method, req.URL.Path, time.Since(start))
        }
    }
}

// Recover middleware opcional
func Recover() Middleware {
    return func(next HandlerFunc) HandlerFunc {
        return func(c *Context) {
            defer func() {
                if err := recover(); err != nil {
                    log.Printf("[ZonerGo] panic recovered (middleware): %v", err)
                    c.JSON(500, map[string]string{"error": "internal server error"})
                }
            }()
            next(c)
        }
    }
}
