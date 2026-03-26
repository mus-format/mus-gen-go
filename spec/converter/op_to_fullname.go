package converter

import (
	"fmt"
	"log"
	"strings"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/typename"
)

func NewToFullNameOp(knownPkgs map[typename.Package]typename.PkgPath,
	gops genopts.Options) *ToFullNameOp {
	return &ToFullNameOp{&strings.Builder{}, knownPkgs, gops}
}

type ToFullNameOp struct {
	b         *strings.Builder
	knownPkgs map[typename.Package]typename.PkgPath
	gopts     genopts.Options
}

func (o *ToFullNameOp) ProcessType(t scanner.TypeInfo[typename.CompleteName],
	tops tpopts.Options,
) (err error) {
	var pkg typename.Package
	o.b.WriteString(t.Stars)
	switch t.Kind {
	case scanner.KindDefined:
		if pkg, err = o.choosePkg(t); err != nil {
			return
		}
		o.b.WriteString(string(pkg))
		o.b.WriteString(".")
		o.b.WriteString(string(t.Name))
	case scanner.KindArray:
		o.b.WriteString("[")
		o.b.WriteString(t.ArrLength)
		o.b.WriteString("]")
	case scanner.KindSlice:
		o.b.WriteString("[]")
	case scanner.KindMap:
		o.b.WriteString("map")
	case scanner.KindPrim:
		o.b.WriteString(string(t.Name))
	default:
		return fmt.Errorf("unexpected %v kind", t.Kind)
	}
	return
}

func (o *ToFullNameOp) ProcessLeftSquare() {
	o.b.WriteString("[")
}

func (o *ToFullNameOp) ProcessComma() {
	o.b.WriteString(",")
}

func (o *ToFullNameOp) ProcessRightSquare() {
	o.b.WriteString("]")
}

func (o *ToFullNameOp) FullName() typename.FullName {
	return typename.FullName(o.b.String())
}

func (o *ToFullNameOp) choosePkg(t scanner.TypeInfo[typename.CompleteName]) (
	pkg typename.Package, err error,
) {
	importPath := genopts.ImportPath(t.PkgPath)

	if alias, pst := o.gopts.ImportAliases()[importPath]; pst {
		pkg = typename.Package(alias)
	} else {
		pkg = t.Package
		if t.Position == scanner.PositionParam {
			// We do not know the package alias for the type parameters, see scanner.go
			log.Printf("WARNING: no alias for '%v' in musgen.Generator options\n",
				t.PkgPath)
		}
	}
	if err = o.checkPkg(pkg, t); err != nil {
		return
	}
	o.knownPkgs[pkg] = t.PkgPath
	return
}

func (o *ToFullNameOp) checkPkg(pkg typename.Package,
	t scanner.TypeInfo[typename.CompleteName]) (err error) {
	if pkgPath, pst := o.knownPkgs[pkg]; pst && pkgPath != t.PkgPath {
		err = NewTwoPathsSameAliasError(pkgPath, t.PkgPath, pkg)
	}
	return
}
