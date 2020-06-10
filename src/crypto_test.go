package main

import "fmt"

func Test() {
	fmt.Println(Md5("hello"))
	fmt.Println(Sha1("hello"))
	fmt.Println(Sha2("hello"))
}
