package main

import (
	"fmt"
	"unsafe"
)

// go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
// fieldalignment -fix ./...

type Foo struct {
	a bool
	b int64
	c bool
	d bool
}

func main() {
	x := &Foo{}
	y := Foo{}

	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))
}
