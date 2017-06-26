package main

import "fmt"
import "testing"
import "log"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type Helper interface {
	Help()
}

type ReadWrite interface {
	Reader
	Writer
}

type R int
type W int
type RW struct {
	H
	*R
	*W
}
type H struct {
	instance interface{}
}

func (r *R) Read() {
	fmt.Println("Read")
}
func (w *W) Write() {
	fmt.Println("Write")
}

func InitHelper(i interface{}) {
	if h, ok := i.(H); ok {
		h.instance = i
		log.Println("ok")
		return
	}
	log.Println("fail")
}
func TestRW(t *testing.T) {
	rw := &RW{}
	InitHelper(rw)
	rw.Read()
}
