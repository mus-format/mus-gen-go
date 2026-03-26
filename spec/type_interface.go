package spec

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	"github.com/mus-format/mus-gen-go/typename"
)

type InterfaceType struct {
	FullName typename.FullName
	Impls    []typename.FullName
	Iops     intropts.Options
	Gops     genopts.Options
}
