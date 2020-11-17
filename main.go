package main

import (
	"log"
	"miguelramirez93/go_simple_proxy/data/services"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
)

func main() {

	var AlwaysReject goproxy.FuncHttpsHandler = func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		return goproxy.RejectConnect, host
	}
	proxy := goproxy.NewProxyHttpServer()
	// proxy.Verbose = true
	patterns := services.GetPatterns()

	for _, pattern := range patterns {
		proxy.OnRequest(goproxy.UrlMatches(regexp.MustCompile(".*" + pattern + ".*"))).HandleConnect(AlwaysReject)
	}

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
