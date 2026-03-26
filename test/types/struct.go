package types

import (
	"time"
)

type SimpleStruct struct {
	Num int
	Str string
}

type IgnoreStruct struct {
	Num int
	Str string
}

type TimeStruct time.Time

type EmbeddingStruct struct {
	Num int
	InnerStruct
}

type InnerStruct struct {
	Str string
}

type ValidStruct struct {
	Num int
	Str string
}

type ParametricStruct[T any] struct {
	Field T
}
