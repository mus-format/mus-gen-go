package scanner

// generic
type TripleParamStruct[T, V, N any] struct{}
type DoubleParamStruct[T, V any] struct{}
type Array[T any] [3]T
type Slice[T any] []T
type Int int

// struct
type Struct struct {
	Int int
}
