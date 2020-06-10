package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main1() {
	value := gjson.Get(json, "name")
	println(value.String())
}
