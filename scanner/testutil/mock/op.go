package mock

import (
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/ymz-ncnk/mok"
)

type TypeFn[T scanner.QualifiedName] func(t scanner.Type[T], tops *typeops.Options) error
type LeftSquareFn func()
type CommaFn func()
type RightSquareFn func()

func NewOp[T scanner.QualifiedName]() Op[T] {
	return Op[T]{mok.New("Op")}
}

type Op[T scanner.QualifiedName] struct {
	*mok.Mock
}

func (o Op[T]) RegisterProcessType(fn TypeFn[T]) Op[T] {
	o.Register("ProcessType", fn)
	return o
}

func (o Op[T]) RegisterProcessLeftSquare(fn LeftSquareFn) Op[T] {
	o.Register("ProcessLeftSquare", fn)
	return o
}

func (o Op[T]) RegisterProcessComma(fn CommaFn) Op[T] {
	o.Register("ProcessComma", fn)
	return o
}

func (o Op[T]) RegisterProcessRightSquare(fn RightSquareFn) Op[T] {
	o.Register("ProcessRightSquare", fn)
	return o
}

func (o Op[T]) ProcessType(t scanner.Type[T], tops *typeops.Options) (err error) {
	result, err := o.Call("ProcessType", t, mok.SafeVal[*typeops.Options](tops))
	if err != nil {
		panic(err)
	}
	err, _ = result[0].(error)
	return
}

func (o Op[T]) ProcessLeftSquare() {
	_, err := o.Call("ProcessLeftSquare")
	if err != nil {
		panic(err)
	}
}

func (o Op[T]) ProcessComma() {
	_, err := o.Call("ProcessComma")
	if err != nil {
		panic(err)
	}
}

func (o Op[T]) ProcessRightSquare() {
	_, err := o.Call("ProcessRightSquare")
	if err != nil {
		panic(err)
	}
}
