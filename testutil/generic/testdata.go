package testutil

import (
	com "github.com/mus-format/common-go"
)

const ImplDTM com.DTM = iota + 1

type MyInt int

type MyArray[T any] [3]T

type MySlice[T any] []T

type MyMap[T comparable, V any] map[T]V

type MyStruct[T any] struct {
	T   T
	Int int
}

type MyDoubleParamStruct[T, V any] struct {
	T T
	V V
}

type MyTripleParamStruct[T, V, N any] struct {
	T T
	V V
	N N
}

type MyInterface[T any] interface {
	Print(t T)
}

type Impl[T any] struct{}

func (i Impl[T]) Print(t T) {}
