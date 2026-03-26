// Package converter provides functionality for converting between different Go
// type name formats (e.g., complete, full, relative).
package converter

import (
	"fmt"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	scnr "github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/typename"
)

func NewTypeNameConverter(gops genopts.Options) TypeNameConverter {
	return TypeNameConverter{
		knownPkgs: map[typename.Package]typename.PkgPath{gops.Package: gops.PkgPath},
		gopts:     gops,
	}
}

type TypeNameConverter struct {
	knownPkgs map[typename.Package]typename.PkgPath
	gopts     genopts.Options
}

// ConvertToFullName converts a complete type name to a full type name.
//
// During this process, it modifies the type packages according to the specified
// aliases in the gen options.
func (c TypeNameConverter) ConvertToFullName(cname typename.CompleteName) (
	fname typename.FullName) {
	op := NewToFullNameOp(c.knownPkgs, c.gopts)
	if err := scnr.Scan(scnr.Config{}, cname, op, tpopts.Options{}); err != nil {
		panic(fmt.Sprintf("can't convert %v to FullName, cause: %v", cname, err))
	}
	return op.FullName()
}

func (c TypeNameConverter) ConvertToRelativeName(fname typename.FullName) (
	rname typename.RelativeName) {
	op := NewToRelativeNameOp(c.gopts)
	if err := scnr.Scan(scnr.Config{}, fname, op, tpopts.Options{}); err != nil {
		panic(fmt.Sprintf("can't convert %v to RelativeName, cause: %v", fname, err))
	}
	return op.RelativeName()
}
