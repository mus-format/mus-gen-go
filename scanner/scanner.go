// Package scanner provides a type scanner that decomposes Go type names into
// their individual components and processes them.
package scanner

import (
	"path/filepath"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/typename"
)

type QualifiedName interface {
	typename.CompleteName | typename.FullName
}

func Scan[T QualifiedName](cfg Config, name T, op Op[T], tops tpopts.Options) (
	err error) {
	return New(cfg, op, tops).Scan(name)
}

// Scanner is used to scan type names and perform operations on them.
//
// For example, it will call op.ProcessType for each TypeN type in:
// - Type1[Type2]
// - map[Type3]Type4
//
// The Go reflect produce for the type params "PkgPath.Type", not the
// "PkgPath/Pkg.Type". So the scanner uses a PkgPath base value for the
// TypeInfo.Package. For example:
//
// "example.com/pkg.Type" -> "example.com/pkg/pkg.Type"
type Scanner[T QualifiedName] struct {
	cfg  Config
	op   Op[T]
	tops tpopts.Options
}

func New[T QualifiedName](cfg Config, op Op[T], opts tpopts.Options) Scanner[T] {
	return Scanner[T]{cfg, op, opts}
}

func (s Scanner[T]) Scan(name T) error {
	return s.scan(name, 0)
}

func (s Scanner[T]) scan(name T, position Position) (err error) {
	if t, ok := ParseDefinedType(name); ok {
		return s.scanDefined(t, position)
	}
	if t, keyType, elemType, kind, ok := ParseContainerType(name); ok {
		return s.scanContainer(t, keyType, elemType, kind, position)
	}
	if t, ok := ParsePrimitiveType(name); ok {
		return s.scanPrimitive(t, position)
	}
	return NewUnsupportedQualifiedNameError(name)
}

func (s Scanner[T]) scanDefined(t TypeInfo[T], position Position) (err error) {
	t.Kind = KindDefined
	t.Position = position

	// For generic parameters, we might not have its full PkgPath because
	// reflect.Type.String() only includes the package name for type arguments.
	if t.Position == PositionParam {
		t = fixParamPkgPath(t)
	}

	if err = s.op.ProcessType(t, s.tops); err != nil {
		return
	}
	pS := s
	if s.cfg.WithoutParams {
		pS.op = ignoreOp[T]{}
	}
	pS.tops = tpopts.Options{}
	for i := range t.Params {
		if i == 0 {
			pS.op.ProcessLeftSquare()
		}
		if err = pS.scan(t.Params[i], PositionParam); err != nil {
			return
		}
		if i < len(t.Params)-1 {
			pS.op.ProcessComma()
		}
		if i == len(t.Params)-1 {
			pS.op.ProcessRightSquare()
		}
	}
	return
}

func (s Scanner[T]) scanContainer(t TypeInfo[T], keyType, elemType T,
	kind Kind, position Position,
) (err error) {
	t.Kind = kind
	t.Position = position
	if err = s.op.ProcessType(t, s.tops); err != nil {
		return
	}
	if keyType != "" {
		s.op.ProcessLeftSquare()
		keyScanner := s
		keyScanner.tops = tpopts.Options{}
		if err = keyScanner.scan(keyType, PositionKey); err != nil {
			return
		}
		s.op.ProcessRightSquare()
	}
	elemScanner := s
	elemScanner.tops = tpopts.Options{}
	return elemScanner.scan(elemType, PositionElem)
}

func (s Scanner[T]) scanPrimitive(t TypeInfo[T], position Position) (err error) {
	t.Kind = KindPrim
	t.Position = position
	return s.op.ProcessType(t, s.tops)
}

// fixParamPkgPath is a workaround for the fact that Go's reflect.Type.String()
// for type parameters doesn't include the Pkg, so the format is "PkgPath.Type"
// instead of "PkgPath/Pkg.Type".
func fixParamPkgPath[T QualifiedName](t TypeInfo[T]) TypeInfo[T] {
	t.PkgPath = typename.PkgPath(filepath.Join(string(t.PkgPath), string(t.Package)))
	return t
}
