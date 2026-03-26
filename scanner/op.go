package scanner

import tpopts "github.com/mus-format/mus-gen-go/options/type"

type Op[T QualifiedName] interface {
	ProcessType(t TypeInfo[T], tops tpopts.Options) error
	ProcessLeftSquare()
	ProcessComma()
	ProcessRightSquare()
}

type ignoreOp[T QualifiedName] struct{}

func (o ignoreOp[T]) ProcessType(t TypeInfo[T], tops tpopts.Options) (err error) {
	return
}
func (o ignoreOp[T]) ProcessLeftSquare()  {}
func (o ignoreOp[T]) ProcessComma()       {}
func (o ignoreOp[T]) ProcessRightSquare() {}
