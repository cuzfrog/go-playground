package main

import "fmt"

type V interface {
	value() int
}

type A int
type D struct {
	A
}

func main() {
	var d D
	fmt.Println(d.value())
	var v1 V = d
	fmt.Println(v1.value())
	var v2 V = new(A)
	fmt.Println(v2.value())
}

func (a A)value() int  {
	return int(a)
}

func (a D)value() int  {
	return a.A.value() + 1
}