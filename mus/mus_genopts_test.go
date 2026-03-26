package musgen_test

import (
	"image"
	"os"
	"reflect"
	"testing"

	musgen "github.com/mus-format/mus-gen-go/mus"
	genops "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/test/genopts/cross-package/pkg3"
	custompkg "github.com/mus-format/mus-gen-go/test/genopts/custom_pkg"
	customsername "github.com/mus-format/mus-gen-go/test/genopts/custom_sername"
	importalias "github.com/mus-format/mus-gen-go/test/genopts/import_alias"
	pkg1sub "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg1/sub"
	pkg2sub "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg2/sub"
	"github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg1"
	"github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg2"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenerateGenopts_MultiPackage(t *testing.T) {
	// Foo from pkg1

	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg1"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[pkg1.Foo]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/multi-package/pkg1/foo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

	// Bar from pkg2

	g, err = musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg2"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[pkg2.Bar]())
	assertfatal.EqualError(t, err, nil)
	bs, err = g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/multi-package/pkg2/bar_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateGenopts_CrossPackage(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/cross-package/pkg4"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[pkg3.Foo]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/cross-package/pkg4/foo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateGenopts_CustomPkg(t *testing.T) {
	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/custom_pkg"),
		genops.WithPackage("custompkg"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[custompkg.Foo]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/custom_pkg/foo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateGenopts_ImportAlias(t *testing.T) {

	// Foo from pkg1/sub

	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg1/sub"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[pkg1sub.Foo]())
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/import_alias/pkg1/sub/foo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

	// Bar from pkg2/sub

	g, err = musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg2/sub"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddDefinedType(reflect.TypeFor[pkg2sub.Bar]())
	assertfatal.EqualError(t, err, nil)
	bs, err = g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/import_alias/pkg2/sub/bar_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

	// Zoo from import_alias

	g, err = musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/import_alias"),
		genops.WithImportAlias("github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg1/sub", "pkg1_sub"),
		genops.WithImportAlias("github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg2/sub", "pkg2_sub"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(reflect.TypeFor[importalias.Zoo]())
	assertfatal.EqualError(t, err, nil)
	bs, err = g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/import_alias/zoo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}

func TestGenerateGenopts_CustomSerName(t *testing.T) {
	pointType := reflect.TypeFor[image.Point]()

	g, err := musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/custom_sername/pkg"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(pointType)
	assertfatal.EqualError(t, err, nil)
	bs, err := g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/custom_sername/pkg/image_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)

	fooType := reflect.TypeFor[customsername.Foo]()
	g, err = musgen.NewGenerator(
		genops.WithPkgPath("github.com/mus-format/mus-gen-go/test/genopts/custom_sername"),
		genops.WithSerName(fooType, "CustomSerName"),
		genops.WithImport("github.com/mus-format/mus-gen-go/test/genopts/custom_sername/pkg"),
		genops.WithSerName(pointType, "pkg.Point"),
	)
	assertfatal.EqualError(t, err, nil)
	err = g.AddStruct(fooType)
	assertfatal.EqualError(t, err, nil)
	bs, err = g.Generate()
	assertfatal.EqualError(t, err, nil)
	err = os.WriteFile("../test/genopts/custom_sername/foo_mus.gen.go", bs, 0644)
	assertfatal.EqualError(t, err, nil)
}
