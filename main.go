package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/elazarl/goproxy"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Println("Listen :8080")
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
