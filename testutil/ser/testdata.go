package testutil

import (
	com "github.com/mus-format/common-go"
	another "github.com/mus-format/musgen-go/testutil/ser/pkg"
)

const MyAwesomeStructDTM com.DTM = iota + 2

type MySlice []another.MyInt

type MyStruct struct {
	MyInt another.MyInt
}

type MyInterface interface {
	Print()
}
