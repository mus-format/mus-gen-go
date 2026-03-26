package scanner

import "github.com/mus-format/mus-gen-go/typename"

const (
	KindUndefined Kind = iota
	KindDefined
	KindArray
	KindSlice
	KindMap
	KindPrim
)

const (
	PositionUndefined Position = iota
	PositionKey
	PositionElem
	PositionParam
)

type TypeInfo[T QualifiedName] struct {
	PkgPath typename.PkgPath // always in the format "PkgPath/Pkg.Type", thanks
	// to the scanner
	Stars     string
	Package   typename.Package
	Name      typename.TypeName
	Params    []T
	ArrLength string

	KeyType  T
	ElemType T

	Kind     Kind
	Position Position
}

type Kind int

type Position int
