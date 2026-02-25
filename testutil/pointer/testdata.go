package testutil

import "fmt"

type MyIntPtr *int
type MyDoubleIntPtr **int
type MyArrayPtr *[3]int
type MySlicePtr *[]int
type MyMapPtr *map[int]string
type MyStructPtr *MyStruct
type MyInterfacePtr *MyInterface

type MyStruct struct{}

func (s MyStruct) Print() {
	fmt.Println("MyStruct")
}

type MyInterface interface{ Print() }
