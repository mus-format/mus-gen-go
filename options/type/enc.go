package tpopts

import "fmt"

const (
	// NumEncodingUndefined represents an undefined numeric encoding.
	NumEncodingUndefined NumEncoding = iota
	// NumEncodingVarint represents varint numeric encoding.
	NumEncodingVarint
	// NumEncodingVarintPositive represents positive varint numeric encoding.
	NumEncodingVarintPositive
	// NumEncodingRaw represents raw numeric encoding.
	NumEncodingRaw
)

// NumEncoding represents the type of numeric encoding.
type NumEncoding int

// Package returns the package name for the numeric encoding.
func (e NumEncoding) Package() (pkg string) {
	switch e {
	case NumEncodingUndefined, NumEncodingVarint, NumEncodingVarintPositive:
		return "varint"
	case NumEncodingRaw:
		return "raw"
	default:
		panic(fmt.Errorf("undefined %d NumEncoding", e))
	}
}

// LenSer returns the length serializer for the numeric encoding.
func (e NumEncoding) LenSer() string {
	switch e {
	case NumEncodingUndefined, NumEncodingVarintPositive:
		return "varint.PositiveInt"
	case NumEncodingVarint:
		return "varint.Int"
	case NumEncodingRaw:
		return "raw.Int"
	default:
		panic(fmt.Errorf("undefined %d NumEncoding", e))
	}
}
