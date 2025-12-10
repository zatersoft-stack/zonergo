# ZonerGo

**ZonerGo** é um micro-framework HTTP minimalista para Golang, focado em APIs rápidas e simples.

## Features

- Roteador HTTP leve (GET, POST, PUT, DELETE)
- Suporte a middlewares globais
- Helpers para respostas `JSON` e `Text`
- Logger de requests
- Recuperação de panic (recover) opcional

## Instalação

```bash
go get github.com/zatersoft-stack/zonergo
```

## Exemplo rápido

```go
r := zonergo.NewRouter()
r.Use(zonergo.Logger(), zonergo.Recover())

r.GET("/ping", func(c *zonergo.Context) {
    c.JSON(200, map[string]string{"message": "pong"})
})

http.ListenAndServe(":8080", r)
```

## Objetivo

Este projeto foi criado por: Melquisedeque Campos como um processo de arquitetura de micro-frameworks em Go
e como portfólio técnico para demonstrar experiência com HTTP, middlewares e design simples.
