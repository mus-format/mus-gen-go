package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	typeops "github.com/mus-format/musgen-go/options/type"
	generic_testdata "github.com/mus-format/musgen-go/testutil/generic"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenericTypeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/generic"),
		genops.WithPackage("testutil"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testutil",
			"common_testdata"),
	)
	assertfatal.EqualError(t, err, nil)

	// defined type

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyInt]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MySlice[generic_testdata.MyInt]]())
	assertfatal.EqualError(t, err, nil)

	//
	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyArray[int]]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyMap[generic_testdata.MyArray[int], generic_testdata.MyInt]](),
		typeops.WithKey(
			typeops.WithValidator("common_testdata.ValidateZeroValue[MyArray[int]]"),
		),
		typeops.WithElem(
			typeops.WithValidator("common_testdata.ValidateZeroValue[MyInt]"),
		),
	)
	assertfatal.EqualError(t, err, nil)
	//

	// struct

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyDoubleParamStruct[int, generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyTripleParamStruct[generic_testdata.MySlice[generic_testdata.MyInt], generic_testdata.MyInterface[generic_testdata.MyInt], generic_testdata.MyDoubleParamStruct[int, generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]]]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.Impl[generic_testdata.MyInt]]())
	assertfatal.EqualError(t, err, nil)

	// DTS

	tp := reflect.TypeFor[generic_testdata.Impl[generic_testdata.MyInt]]()
	err = g.AddDTS(tp)
	assertfatal.EqualError(t, err, nil)

	// interface

	err = g.AddInterface(reflect.TypeFor[generic_testdata.MyInterface[generic_testdata.MyInt]](),
		introps.WithImpl(tp),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("../testutil/generic/mus-format.gen.go", bs, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
