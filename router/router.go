package router

import (
	"net/http"
	"small_account/handler"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Index())
	mux.HandleFunc("PUT /register", handler.Register())
	mux.HandleFunc("POST /login", handler.Login())
	mux.HandleFunc("GEt /me", handler.GETMe())
	mux.HandleFunc("POST /me", handler.POSTMe())
	return mux
}
