package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func getHandles(path string) []ProxyHandle {

	handles := []ProxyHandle{}
	file, err := os.Open(path)
	if err != nil {
		return handles
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return handles
	}
	json.Unmarshal([]byte(data), &handles)
	return handles
}
