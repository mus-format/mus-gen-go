package testutil

import (
	"fmt"
)

// const (
// 	Impl1DTM com.DTM = iota
// 	Impl2DTM
// )

type MyInterface interface {
	Print()
}

type Impl1 struct {
	Str string
}

func (i Impl1) Print() {
	fmt.Println("impl1")
}

type Impl2 int

func (i Impl2) Print() {
	fmt.Println("impl2")
}
