package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	wd, _ := os.Getwd()
	port := *flag.Int("p", 5555, "Set The Http Port")
	rd := *flag.String("d", wd, "Set The Directory")
	flag.Parse()

	log.Printf("Listen On http://localhost:%d", port)
	log.Printf("Work Directory: %s", wd)
	log.Printf("Root Directory: %s", rd)
	// http.HandleFunc("/upload", uploadHandler)

	handles := getHandles("./ng.json")
	for _, handle := range handles {
		log.Println(handle.Prefix, "->", handle.Forword)
		h := new(ProxyHandle)
		h.Prefix = "/" + handle.Prefix
		h.Forword = handle.Forword
		http.Handle("/"+handle.Prefix+"/", h)
	}

	//static file handler.
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(rd))))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if nil != err {
		log.Fatalln("ERROR:", err.Error())
	}
}
