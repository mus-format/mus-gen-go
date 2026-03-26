package mock

import (
	"github.com/mus-format/mus-gen-go/typename"
	"github.com/ymz-ncnk/mok"
)

type ConvertToFullNameFn func(cname typename.CompleteName) typename.FullName
type ConvertToRelativeNameFn func(fname typename.FullName) typename.RelativeName

func NewTypeNameConvertor() TypeNameConvertor {
	return TypeNameConvertor{mok.New("TypeNameConvertor")}
}

type TypeNameConvertor struct {
	*mok.Mock
}

func (c TypeNameConvertor) RegisterConvertToFullName(fn ConvertToFullNameFn) TypeNameConvertor {
	c.Register("ConvertToFullName", fn)
	return c
}

func (c TypeNameConvertor) RegisterConvertToRelativeName(fn ConvertToRelativeNameFn) TypeNameConvertor {
	c.Register("ConvertToRelativeName", fn)
	return c
}

func (c TypeNameConvertor) ConvertToFullName(cname typename.CompleteName) typename.FullName {
	result, err := c.Call("ConvertToFullName", cname)
	if err != nil {
		panic(err)
	}
	return result[0].(typename.FullName)
}

func (c TypeNameConvertor) ConvertToRelativeName(fname typename.FullName) typename.RelativeName {
	result, err := c.Call("ConvertToRelativeName", fname)
	if err != nil {
		panic(err)
	}
	return result[0].(typename.RelativeName)
}
