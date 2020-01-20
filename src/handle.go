package main

import (
	// "context"
	"fmt"
	"strings"
	// "github.com/bogdanovich/dns_resolver"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyHandle struct {
	Prefix  string
	Forword string
}

func (this *ProxyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())

	url, err := url.Parse(this.Forword)
	if err != nil {
		log.Println(err)
		customHttpResponse(w, err.Error(), 500)
		return
	}

	log.Printf("path:%s", r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, this.Prefix)
	log.Printf("path:%s", path)
	r.Host = url.Host
	r.URL.Path = path
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

// var fileHandle = http.FileServer(http.Dir(DEFAULTROOT))
func customHttpResponse(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}
