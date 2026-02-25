package testutil

import (
	"fmt"

	com "github.com/mus-format/common-go"
)

const (
	Impl1DTM com.DTM = iota + 1
	Impl2DTM
)

type MyInterface interface {
	Print()
}

type Impl1 struct{}

func (i Impl1) Print() {
	fmt.Println("Impl1")
}

func (i Impl1) MarshalTypedMUS(bs []byte) (n int) {
	return Impl1DTS.Marshal(i, bs)
}

func (i Impl1) SizeTypedMUS() (n int) {
	return Impl1DTS.Size(i)
}

type Impl2 struct{}

func (i Impl2) Print() {
	fmt.Println("Impl2")
}

func (i Impl2) MarshalTypedMUS(bs []byte) (n int) {
	return Impl2DTS.Marshal(i, bs)
}

func (i Impl2) SizeTypedMUS() (n int) {
	return Impl2DTS.Size(i)
}
