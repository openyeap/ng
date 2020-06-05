package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func main() {
	wd, _ := os.Getwd()
	port := flag.Int("p", 5555, "Set The Http Port")
	rd := flag.String("d", wd, "Set The Directory")
	flag.Parse()

	log.Printf("Listen On http://localhost:%d", *port)
	log.Printf("Root Directory: %s", *rd)
	// http.HandleFunc("/upload", uploadHandler)

	viper.SetConfigFile("ng.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if nil != err {
		log.Fatalln("ERROR:", err.Error())
	}

	var host *Host

	err = viper.UnmarshalKey("host", &host)
	if nil != err {
		log.Fatalln("ERROR:", err.Error())
	}
	if host.Static == "" {
		host.Static = *rd
	}
	log.Print(host)
	http.Handle("/", host)

	// for _, item := range servers {
	// 	log.Print(item)
	// 	server := item.(map[interface{}]interface{})

	// 	for k, v := range server {

	// 		log.Print(k)

	// 		log.Print(v)
	// 	}

	// 	log.Println("--------")
	// 	log.Print(server)
	// 	log.Println("--------")
	// 	hostHandle := new(HostHandle)
	// 	http.Handle("/", hostHandle)
	// 	// h.Prefix = "/" + server["Prefix"]
	// 	// h.Forword = server["Forword"]
	// 	// http.Handle("/"+server["Prefix"]+"/", h)
	// }

	// for _, handle := range handles {
	// 	log.Println(handle.Prefix, "->", handle.Forword)
	// 	h := new(ProxyHandle)
	// 	h.Prefix = "/" + handle.Prefix
	// 	h.Forword = handle.Forword
	// 	http.Handle("/"+handle.Prefix+"/", h)
	// }
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if nil != err {
		log.Fatalln("ERROR:", err.Error())
	}
}
