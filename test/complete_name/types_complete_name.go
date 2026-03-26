package types

// primitive
type Int int

// container
type Array [3]int
type Slice []int
type Map map[int]string

// pointer
type IntPtr *int
type ArrayPtr *[3]int
type SlicePtr *[]int
type MapPtr *map[int]string
type DoubleIntPtr **int

// struct
type Struct struct {
	Int int
}

type ParametrizedStruct[T any] struct {
	T T
}
