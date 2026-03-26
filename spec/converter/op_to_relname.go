package converter

import (
	"fmt"
	"strings"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/typename"
)

func NewToRelativeNameOp(gops genopts.Options) *ToRelativeNameOp {
	return &ToRelativeNameOp{&strings.Builder{}, gops}
}

type ToRelativeNameOp struct {
	b    *strings.Builder
	gops genopts.Options
}

func (o *ToRelativeNameOp) ProcessType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (err error) {
	var pkg typename.Package
	o.b.WriteString(t.Stars)
	switch t.Kind {
	case scanner.KindDefined:
		if pkg, err = o.choosePkg(t); err != nil {
			return
		}
		if pkg != "" {
			o.b.WriteString(string(pkg))
			o.b.WriteString(".")
		}
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

func (o *ToRelativeNameOp) ProcessLeftSquare() {
	o.b.WriteString("[")
}

func (o *ToRelativeNameOp) ProcessComma() {
	o.b.WriteString(",")
}

func (o *ToRelativeNameOp) ProcessRightSquare() {
	o.b.WriteString("]")
}

func (o *ToRelativeNameOp) RelativeName() typename.RelativeName {
	return typename.RelativeName(o.b.String())
}

func (o *ToRelativeNameOp) choosePkg(t scanner.TypeInfo[typename.FullName]) (
	pkg typename.Package, err error) {
	if t.Package == o.gops.Package {
		pkg = ""
	} else {
		pkg = t.Package
	}
	return
}
