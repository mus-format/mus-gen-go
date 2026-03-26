package spec

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/typename"
)

// Type represents the core description (Name + Options) of any MUS type.
type Type struct {
	FullName typename.FullName
	Gops     genopts.Options
}
