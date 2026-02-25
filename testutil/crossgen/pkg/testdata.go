package pkg

import "fmt"

type MyInt int

type MySlice []MyInt

type MyArray[T any] [3]T

type MyAnotherArray[T any] [3]T

type MyStruct struct {
	Int     int
	MyInt   MyInt
	MySlice MySlice
}

func (s MyStruct) Print() {
	fmt.Println("MyStruct")
}

type MyInterface interface {
	Print()
}
