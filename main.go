package main

import (
	"goblog/bootstrap"
	"goblog/tests/app/http/middlewares"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
