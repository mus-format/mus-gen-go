package types

type LenString string

type LenSlice []int

type LenArray [3]int

type LenMap map[string]int

type Slice []int

type Array [3]int

type Map map[string]int

type Ptr *int

type ValidString string

type ValidSlice []int

type ValidArray [3]int

type ValidMap map[string]int

type ComplexMap map[string]map[*[]int][][]float32

type PtrStruct struct {
	Ptr *Struct
}

type MapStruct struct {
	Map map[string]Struct
}

type SliceStruct struct {
	Slice []Struct
}

type Struct struct{}
