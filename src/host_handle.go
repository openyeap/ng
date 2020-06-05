package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Host struct {
	Root   string
	Port   string
	Routes []Route
}
type Route struct {
	Name    string
	Uri     string
	Asserts []string
	Filters []string
}

func assert(asserts []string, r *http.Request) bool {
	// # time in from end
	// # host in key...
	// # mothod in key...
	// # path in key...
	// # ip in key...
	// # query has key[=v]...
	// # cookie has key[=v]...
	// # header has key[=v]...
	for _, assert := range asserts {
		m := strings.Split(assert, " ")

		switch m[0] {
		case "time":
			if false {
				return false
			}
		case "host":
			if !IsExits(r.Host, m[1:]) {
				return false
			}
		case "method":
			if !IsExits(r.Method, m[1:]) {
				return false
			}
		case "path":
			if !IsExits(r.URL.Path, m[1:]) {
				return false
			}

		case "ip":
			if !IsExits(r.RemoteAddr, m[1:]) {
				return false
			}
		case "query":
			if strings.Index(r.URL.RawQuery, m[1]) < 0 {
				return false
			}
		case "cookie":
			if strings.Index(r.Header.Get("cookie"), m[1]) < 0 {
				return false
			}
		case "header":
			if false {
				return false
			}
		}

	}
	return true
}
func IsExits(input string, keys []string) bool {
	for _, item := range keys {
		if strings.Index(input, item) == 0 {
			return true
		}
	}
	return false
}

func Filter(filters []string, r *http.Request) {
	// path insert append remove
	// header k v
	// cookie k v
	for _, filter := range filters {
		m := strings.Split(filter, " ")
		switch m[0] {
		case "path":
			switch m[1] {
			case "insert":
				r.URL.Path = strings.Join(append(m[2:], r.URL.Path), "/")
			case "append":
				r.URL.Path = strings.Join(append([]string{r.URL.Path}, m[2:]...), "/")
			case "remove":
				path := strings.Split(r.URL.Path, "/")
				for _, s := range m[2:] {

					i, err := strconv.Atoi(s)
					if err != nil {
						log.Println(err)
						continue
					}
					if i == 0 {
						path = path[1:]
					} else {
						if i == -1 {
							path = path[0 : len(path)-1]
						} else {
							path = append(path[:i], path[i+1:]...)
						}
					}

				}

				r.URL.Path = strings.Join(path, "/")
				log.Println(r.URL.Path)
			}

		case "header":
			if len(m) == 3 {
				r.Header.Set(m[1], m[2])
			} else {
				r.Header.Del(m[1])
			}
		case "cookie":
			c := &http.Cookie{
				Name:     m[1],
				Value:    m[2],
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
			}
			if len(m) == 3 {
				c.Expires = time.Now().Add(1)
			}
			r.AddCookie(c)
		}

	}

}

func (this *Host) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, router := range this.Routes {
		url, err := url.Parse(router.Uri)
		if err != nil {
			log.Println(err)
			continue
		}

		if assert(router.Asserts, r) {
			Filter(router.Filters, r)
			log.Println(router)

			r.Host = url.Host

			proxy := httputil.NewSingleHostReverseProxy(url)

			proxy.ServeHTTP(w, r)
			return
		}
	}
	http.FileServer(http.Dir(this.Root)).ServeHTTP(w, r)
}
