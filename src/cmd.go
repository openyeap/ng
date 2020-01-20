package main

import (
	"fmt"
	"os/exec"
)

func run() {

	f, err := exec.Command("python", "/").Output()
	if err != nil {
		fmt.Println("发生错误：")
		fmt.Println(err.Error())
	}
	fmt.Println("执行结果：")
	fmt.Println(string(f))
}
