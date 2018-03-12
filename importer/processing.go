package main

import (
	"net/http/pprof"

	"github.com/julienschmidt/httprouter"
)

func initProfiling() *httprouter.Router {
	router := httprouter.New()
	router.Handler("GET", "/heap", pprof.Handler("heap"))
	router.Handler("GET", "/goroutine", pprof.Handler("goroutine"))
	return router
}
