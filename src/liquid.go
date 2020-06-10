package main

import (
	"fmt"
	"log"

	"gopkg.in/osteele/liquid.v1"
)

func main3() {

	engine := liquid.NewEngine()
	template := `<h1>{{ page.title }}</h1>`
	bindings := map[string]interface{}{
		"page": map[string]string{
			"title": "Introduction",
		},
	}
	out, err := engine.ParseAndRenderString(template, bindings)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
}
