package builder

// Defined

type Int int

type Slice []int

type Struct struct{}

type Interface interface {
	Foo()
}

// Structs

type StructWithFields struct {
	Int int
}

type StructWithTwoFields struct {
	Int1 int
	Int2 int
}
