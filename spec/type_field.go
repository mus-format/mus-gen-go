package spec

import (
	genops "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/typename"
)

type FieldType struct {
	FullName  typename.FullName
	FieldName string
	Gops      genops.Options
}
