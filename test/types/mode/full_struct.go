package mode

import (
	"strconv"
	"time"
)

type FullStruct struct {
	Int        int
	Int64      int64
	Int32      int32
	Int16      int16
	Int8       int8
	Uint       uint
	Uint64     uint64
	Uint32     uint32
	Uint16     uint16
	Uint8      uint8
	Float64    float64
	Float32    float32
	Byte       byte
	Bool       bool
	Str        string
	ByteSlice  []byte
	Uint8Slice []uint8
	Time       time.Time

	// Pointers
	PtrInt *int

	// Arrays
	ArrayInt [10]int

	// Slices (non-byte)
	SliceInt []int

	// Maps
	MapStrInt map[string]int

	// Defined
	Defined FullDefined

	// Interface
	Interface FullInterface

	// Versioned
	Versioned Versioned
}

type FullDefined int

type FullInterface any

type FullInterfaceImpl string

type Versioned FooV2

type FooV2 string

type FooV1 struct {
	Num int
}

func MigrateFooV1(v FooV1) Versioned {
	return Versioned(strconv.Itoa(v.Num))
}
