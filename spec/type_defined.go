package spec

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/typename"
)

type DefinedType struct {
	FullName           typename.FullName
	UnderlyingTypeName typename.FullName // e.g., 'int' for 'type MyInt int'

	Tops tpopts.Options
	Gops genopts.Options
}
