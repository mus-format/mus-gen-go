package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	testutil "github.com/mus-format/musgen-go/testutil/two_pkg"
	pkg "github.com/mus-format/musgen-go/testutil/two_pkg/pkg"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestTwoPkgGeneration(t *testing.T) {

	t.Run("First pkg", func(t *testing.T) {
		tp := reflect.TypeFor[pkg.MyIntSerName]()
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/two_pkg/pkg"),
			genops.WithSerName(tp, "MyAwesomeInt"),
		)
		assertfatal.EqualError(t, err, nil)

		err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
		assertfatal.EqualError(t, err, nil)

		err = g.AddDefinedType(reflect.TypeFor[pkg.MySlice]())
		assertfatal.EqualError(t, err, nil)

		err = g.AddDefinedType(tp)
		assertfatal.EqualError(t, err, nil)

		// generate

		bs, err := g.Generate()
		assertfatal.EqualError(t, err, nil)
		err = os.WriteFile("../testutil/two_pkg/pkg/mus-format.gen.go", bs, 0644)
		assertfatal.EqualError(t, err, nil)
	})

	t.Run("Second pkg", func(t *testing.T) {
		tp := reflect.TypeFor[pkg.MyIntSerName]()
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testutil/two_pkg"),
			genops.WithPackage("testutil"),
			genops.WithImport("github.com/mus-format/musgen-go/testutil/two_pkg/pkg"),
			genops.WithSerName(tp, "pkg.MyAwesomeInt"),
		)
		assertfatal.EqualError(t, err, nil)

		err = g.AddDefinedType(reflect.TypeFor[testutil.MySlice]())
		assertfatal.EqualError(t, err, nil)

		err = g.AddDefinedType(reflect.TypeFor[testutil.MySliceSerName]())
		assertfatal.EqualError(t, err, nil)

		// generate

		bs, err := g.Generate()
		assertfatal.EqualError(t, err, nil)
		err = os.WriteFile("../testutil/two_pkg/mus-format.gen.go", bs, 0644)
		assertfatal.EqualError(t, err, nil)
	})

}
