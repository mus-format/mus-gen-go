package spec

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	veropts "github.com/mus-format/mus-gen-go/options/versioned"
	"github.com/mus-format/mus-gen-go/typename"
)

type VersionedType struct {
	FullName typename.FullName
	Versions []typename.FullName
	Vops     veropts.Options
	Gops     genopts.Options
}
