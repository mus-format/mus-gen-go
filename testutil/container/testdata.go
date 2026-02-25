package testutil

import (
	generic_testdata "github.com/mus-format/musgen-go/testutil/generic"
)

type MyArray [3]int
type LenEncodingMyArray [3]int
type ElemEncodingMyArray [3]int
type ElemValidMyArray [3]int
type ValidMyArray [3]int
type AllMyArray [3]int

type MyByteSlice []byte
type LenEncodingMyByteSlice []byte
type LenValidMyByteSlice []byte
type ValidMyByteSlice []byte
type AllMyByteSlice []byte

type MySlice []int
type LenEncodingMySlice []int
type LenValidMySlice []int
type ElemEncodingMySlice []int
type ElemValidMySlice []int
type ValidMySlice []int
type AllMySlice []int

type MyMap map[int]string
type LenEncodingMyMap map[int]string
type LenValidMyMap map[int]string
type KeyEncodingMyMap map[int]string
type KeyValidMyMap map[int]string
type ElemEncodingMyMap map[int]string
type ElemValidMyMap map[int]string
type ValidMyMap map[int]string
type AllMyMap map[int]string

type MyArrayPtr *[3]int

type ComplexMap map[[3]int]map[*MyByteSlice][]generic_testdata.MySlice[generic_testdata.MyInt]
