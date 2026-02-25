package testutil

import (
	"fmt"

	com "github.com/mus-format/common-go"
	muss "github.com/mus-format/mus-stream-go"
)

const (
	Impl1DTM com.DTM = iota + 1
	Impl2DTM
	MyIntDTM
)

type MyInterface interface {
	Print()
}

type Impl1 struct {
	Str string
}

func (i Impl1) Print() {
	fmt.Println("Impl1")
}

func (i Impl1) MarshalTypedMUS(w muss.Writer) (n int, err error) {
	return Impl1DTS.Marshal(i, w)
}

func (i Impl1) SizeTypedMUS() (n int) {
	return Impl1DTS.Size(i)
}

type Impl2 int

func (i Impl2) Print() {
	fmt.Println("Impl2")
}

func (i Impl2) MarshalTypedMUS(w muss.Writer) (n int, err error) {
	return Impl2DTS.Marshal(i, w)
}

func (i Impl2) SizeTypedMUS() (n int) {
	return Impl2DTS.Size(i)
}
