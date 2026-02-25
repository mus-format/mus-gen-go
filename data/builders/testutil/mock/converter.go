package mock

import (
	"github.com/mus-format/musgen-go/typename"
	"github.com/ymz-ncnk/mok"
)

type ConvertToFullNameFn func(cname typename.CompleteName) (fname typename.FullName,
	err error)

type ConvertToRelNameFn func(fname typename.FullName) (rname typename.RelName,
	err error)

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

func (c TypeNameConvertor) RegisterConvertToRelName(fn ConvertToRelNameFn) TypeNameConvertor {
	c.Register("ConvertToRelName", fn)
	return c
}

func (c TypeNameConvertor) ConvertToFullName(cname typename.CompleteName) (
	fname typename.FullName, err error) {
	result, err := c.Call("ConvertToFullName", cname)
	if err != nil {
		panic(err)
	}
	fname = result[0].(typename.FullName)
	err, _ = result[1].(error)
	return
}

func (c TypeNameConvertor) ConvertToRelName(fname typename.FullName) (
	rname typename.RelName, err error) {
	result, err := c.Call("ConvertToRelName", fname)
	if err != nil {
		panic(err)
	}
	rname = result[0].(typename.RelName)
	err, _ = result[1].(error)
	return
}
