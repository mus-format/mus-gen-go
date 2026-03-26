package builder

import "github.com/mus-format/mus-gen-go/typename"

type TypeNameConverter interface {
	ConvertToFullName(cname typename.CompleteName) (fname typename.FullName)
	ConvertToRelativeName(fname typename.FullName) (rname typename.RelativeName)
}
