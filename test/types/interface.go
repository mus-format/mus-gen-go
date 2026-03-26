package types

import com "github.com/mus-format/common-go"

const (
	Impl1DTM com.DTM = iota + 1
	Impl2DTM
)

type Impl1 struct {
	Num int
}

func (i Impl1) MarshalTypedMUS(bs []byte) (n int) {
	return Impl1TypedMUS.Marshal(i, bs)
}

func (i Impl1) SizeTypedMUS() (size int) {
	return Impl1TypedMUS.Size(i)
}

type Impl2 string

func (i Impl2) MarshalTypedMUS(bs []byte) (n int) {
	return Impl2TypedMUS.Marshal(i, bs)
}

func (i Impl2) SizeTypedMUS() (size int) {
	return Impl2TypedMUS.Size(i)
}

type Interface any

type MarshallerInterface any
