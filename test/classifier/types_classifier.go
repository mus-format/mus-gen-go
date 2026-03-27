package types

type Int int

type Slice []int

type GenericSlice[T any] []T

type Interface interface {
	Print()
}

type DoubleDefinedInterface Interface

type Struct struct{}

type DoubleDefinedStruct Struct

type IntPtr *Int

type SlicePtr *Slice

type StructPtr *Struct

type InterfacePtr *Interface
