package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	testutil "github.com/mus-format/musgen-go/testutil/ser"
	another "github.com/mus-format/musgen-go/testutil/ser/pkg"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithSerGeneration(t *testing.T) {

	var (
		myIntType       = reflect.TypeFor[another.MyInt]()
		myStructType    = reflect.TypeFor[testutil.MyStruct]()
		myInterfaceType = reflect.TypeFor[testutil.MyInterface]()
	)

	t.Run("Another pkg", func(t *testing.T) {
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/ser/pkg"),
			genops.WithPackage("another"),
			genops.WithSerName(myIntType, "MyAwesomeInt"),
		)
		assertfatal.EqualError(t, err, nil)

		// defined type

		err = g.AddDefinedType(reflect.TypeFor[another.MyInt]())
		if err != nil {
			t.Fatal(err)
		}

		// dts

		err = g.AddDTS(myIntType)
		if err != nil {
			t.Fatal(err)
		}

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testutil/ser/pkg/mus-format.gen.go", bs, 0644)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testdata pkg", func(t *testing.T) {
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/ser"),
			genops.WithPackage("testutil"),
			genops.WithImportAlias("github.com/mus-format/musgen-go/testutil/ser/pkg",
				"another"),
			genops.WithSerName(myIntType, "another.MyAwesomeInt"),
			genops.WithSerName(myStructType, "MyAwesomeStruct"),
			genops.WithSerName(myInterfaceType, "MyAwesomeInterface"),
		)
		assertfatal.EqualError(t, err, nil)

		// defined type
		err = g.AddDefinedType(reflect.TypeFor[testutil.MySlice]())
		if err != nil {
			t.Fatal(err)
		}

		// struct type

		err = g.AddStruct(myStructType)
		if err != nil {
			t.Fatal(err)
		}

		// dts

		err = g.AddDTS(myStructType)
		if err != nil {
			t.Fatal(err)
		}

		// interface

		err = g.AddInterface(myInterfaceType, introps.WithImpl(myIntType))
		if err != nil {
			t.Fatal(err)
		}

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testutil/ser/mus-format.gen.go", bs, 0644)
		if err != nil {
			t.Fatal(err)
		}

	})

}
