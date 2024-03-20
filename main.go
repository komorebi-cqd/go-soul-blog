package main

import (
	"github/go-soul-blog/internal/routers"
	"net/http"
)

// @title 博客系统
// @version 1.0
// @description 博客系统
func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
