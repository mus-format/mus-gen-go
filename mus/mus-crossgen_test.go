package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	testutil "github.com/mus-format/musgen-go/testutil/crossgen"
	pkg "github.com/mus-format/musgen-go/testutil/crossgen/pkg"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestCrossGeneration(t *testing.T) {

	t.Run("pkg", func(t *testing.T) {
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/crossgen/pkg"),
		)
		assertfatal.EqualError(t, err, nil)

		// defined type

		err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
		assertfatal.EqualError(t, err, nil)

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testutil/crossgen/pkg/mus-format.gen.go", bs, 0644)
		assertfatal.EqualError(t, err, nil)
	})

	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/crossgen"),
		genops.WithPackage("testutil"),
		genops.WithImport("github.com/mus-format/musgen-go/testutil/crossgen/pkg"),
	)
	assertfatal.EqualError(t, err, nil)

	// defined type

	// err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
	// assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MySlice]())
	assertfatal.EqualError(t, err, nil)

	// struct

	tp := reflect.TypeFor[pkg.MyStruct]()
	err = g.AddStruct(tp)
	assertfatal.EqualError(t, err, nil)

	// defined type

	err = g.AddDefinedType(reflect.TypeFor[testutil.MyMap]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MyArray[testutil.MyMap]]())
	assertfatal.EqualError(t, err, nil)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MyAnotherArray[pkg.MyInt]]())
	assertfatal.EqualError(t, err, nil)

	// struct
	err = g.AddStruct(reflect.TypeFor[testutil.MyStructWithCrossgen]())
	assertfatal.EqualError(t, err, nil)

	// interface

	err = g.AddDTS(tp)
	assertfatal.EqualError(t, err, nil)

	err = g.AddInterface(reflect.TypeFor[pkg.MyInterface](),
		introps.WithImpl(tp),
	)
	assertfatal.EqualError(t, err, nil)

	// generate

	bs, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("../testutil/crossgen/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
