package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"irfanabduhu.io/go-chatroom/v1/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndPoint))
	return mux
}
