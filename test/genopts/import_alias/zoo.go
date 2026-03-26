package import_alias

import (
	foo "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg1/sub"
	bar "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg2/sub"
)

type Zoo struct {
	Foo foo.Foo
	Bar bar.Bar
}
