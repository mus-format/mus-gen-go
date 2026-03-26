package spec

import (
	genops "github.com/mus-format/mus-gen-go/options/gen"
	structops "github.com/mus-format/mus-gen-go/options/struct"
	"github.com/mus-format/mus-gen-go/typename"
)

type StructType struct {
	FullName typename.FullName
	Fields   []FieldType
	Sops     structops.Options
	Gops     genops.Options
}

func (d StructType) SerializedFields() (sl []FieldType) {
	if len(d.Sops.Fields) == 0 {
		return d.Fields
	}
	sl = make([]FieldType, 0, len(d.Fields))
	for i, f := range d.Fields {
		if !d.Sops.Fields[i].Ignore {
			sl = append(sl, f)
		}
	}
	return
}
