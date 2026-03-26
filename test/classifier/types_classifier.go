package types

type MyInt int

type MySlice []int

type MyGenericSlice[T any] []T

type MyInterface interface {
	Print()
}

type DoubleDefinedMyInterface MyInterface

type MyStruct struct{}

type DoubleDefinedMyStruct MyStruct

type MyIntPtr *MyInt

type MySlicePtr *MySlice

type MyStructPtr *MyStruct

type MyInterfacePtr *MyInterface
