package server

import (
	"fmt"
)

type Obj interface {
	ToContent() []byte
}

type idx struct {
	id   [4]byte
	name [8]byte
	size uint16
	path []byte
}

type obj struct {
	size   uint16
	header map[string]string
	body   []byte
}

/* 实现接口方法 */
func (index idx) ToContent() []byte {
	fmt.Println("I am idx")
	return nil
}

func (object obj) ToContent() []byte {
	fmt.Println("I am obj")
	return nil
}
