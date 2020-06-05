package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	conf := flag.String("c", "ng.yaml", "config file")
	flag.Parse()
	viper.SetConfigFile(*conf)
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

	http.Handle("/", host)
	log.Printf("Listen On http://localhost:%s", host.Port)
	log.Printf("Root Directory: %s", host.Root)

	err = http.ListenAndServe(":"+host.Port, nil)
	if nil != err {
		log.Fatalln("ERROR:", err.Error())
	}
}
