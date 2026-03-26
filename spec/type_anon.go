package spec

import (
	"fmt"

	"github.com/mus-format/mus-gen-go/typename"
)

const (
	AnonKindUndefined AnonKind = iota
	AnonKindString
	AnonKindArray
	AnonKindByteSlice
	AnonKindSlice
	AnonKindMap
	AnonKindPtr
)

type AnonSerName string

type AnonType struct {
	SerName AnonSerName
	Kind    AnonKind

	ArrType   typename.FullName
	ArrLength string

	LenSer string
	LenVl  string

	KeyType typename.FullName
	KeyVl   string

	ElemType typename.FullName
	ElemVl   string

	// Tops *tpopts.Options
}

type AnonKind int

func (k AnonKind) String() string {
	switch k {
	case AnonKindString:
		return "string"
	case AnonKindArray:
		return "array"
	case AnonKindByteSlice:
		return "byteSlice"
	case AnonKindSlice:
		return "slice"
	case AnonKindMap:
		return "map"
	case AnonKindPtr:
		return "ptr"
	default:
		panic(fmt.Sprintf("unexpected %v AnonSerKind", int(k)))
	}
}
