// cmd/example/main.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/zatersoft-stack/zonergo/zonergo"
)

func main() {
	r := zonergo.NewRouter()

	// Middlewares globais
	r.Use(zonergo.Logger(), zonergo.Recover())

	// Rotas
	r.GET("/", func(c *zonergo.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Bem-vindo ao ZonerGo!",
			"docs":    "/docs",
		})
	})

	r.GET("/ping", func(c *zonergo.Context) {
		c.Text(http.StatusOK, "pong")
	})

	r.POST("/echo", func(c *zonergo.Context) {
		var body map[string]interface{}
		_ = json.NewDecoder(c.Request.Body).Decode(&body)
		c.JSON(http.StatusOK, map[string]interface{}{
			"received": body,
		})
	})

	// Pequena “pseudo-doc”
	r.GET("/docs", func(c *zonergo.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"routes": []map[string]string{
				{"method": "GET", "path": "/", "desc": "Mensagem de boas-vindas"},
				{"method": "GET", "path": "/ping", "desc": "Healthcheck simples"},
				{"method": "POST", "path": "/echo", "desc": "Retorna o JSON enviado no corpo"},
			},
		})
	})

	http.ListenAndServe(":8080", r)
}
