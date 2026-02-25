package genops

import (
	"reflect"
	"testing"

	prim_testdata "github.com/mus-format/musgen-go/testutil/primitive"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {

	t.Run("Apply should work", func(t *testing.T) {
		var (
			o             = New()
			wantPkgPath   = "github.com/user/project"
			wantUnsafe    = true
			wantNotUnsafe = true
			wantStream    = true
			wantImports   = []string{"\"import1\""}
			wantSerNames  = map[reflect.Type]string{
				reflect.TypeFor[prim_testdata.MyInt](): "MyAwesomeInt",
			}
		)
		Apply([]SetOption{
			WithPkgPath(wantPkgPath),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o)
		asserterror.Equal(t, o.PkgPath, typename.PkgPath(wantPkgPath))
		asserterror.Equal(t, o.Unsafe, wantUnsafe)
		asserterror.Equal(t, o.NotUnsafe, wantNotUnsafe)
		asserterror.Equal(t, o.Stream, wantStream)
		asserterror.EqualDeep(t, o.Imports, wantImports)
		asserterror.EqualDeep(t, o.SerNames, wantSerNames)
	})

	t.Run("Hash", func(t *testing.T) {
		o1 := New()
		Apply([]SetOption{
			WithPkgPath("github.com/user/project"),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o1)
		o2 := New()
		Apply([]SetOption{
			WithPkgPath("github.com/user/project"),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o2)
		asserterror.Equal(t, o1.Hash(), o2.Hash())
	})

	t.Run("WithImport", func(t *testing.T) {

		t.Run("Should fail if it receives an invalid ImportPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImport(""),
			}, &o)
			asserterror.EqualError(t, err, NewInvalidImportPathError(""))
		})

	})

	t.Run("WithImportAlias", func(t *testing.T) {

		t.Run("Should fail if receives two items with the same alias", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project1", "alias"),
				WithImportAlias("github.com/user/project2", "alias"),
			}, &o)
			asserterror.EqualError(t, err, NewDuplicateImportAlias("alias"))
		})

		t.Run("Should fail if receives two similar pkgPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project", "alias1"),
				WithImportAlias("github.com/user/project", "alias2"),
			}, &o)
			asserterror.EqualError(t, err,
				NewDuplicateImportPath("github.com/user/project"))
		})

		t.Run("Should fail if it receives an invalid ImportPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("", "alias"),
			}, &o)
			asserterror.EqualError(t, err, NewInvalidImportPathError(""))
		})

		t.Run("Should fail if it receives an invalid alias", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project", "123"),
			}, &o)
			asserterror.EqualError(t, err, NewInvalidAliasError("123"))
		})

	})

	t.Run("MarshalSignatureLastParam", func(t *testing.T) {
		var (
			o                             = New()
			wantMarshalSignatureLastParam = msigLastParam
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.MarshalSignatureLastParam(),
			wantMarshalSignatureLastParam)
	})

	t.Run("MarshalSignatureLastParam with Stream option", func(t *testing.T) {
		var (
			o                             = New()
			wantMarshalSignatureLastParam = msigLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.MarshalSignatureLastParam(),
			wantMarshalSignatureLastParam)
	})

	t.Run("MarshalLastParam", func(t *testing.T) {
		var (
			o                         = New()
			wantMarshalLastParam      = mLastParam
			wantMarshalLastParamFirst = mLastParamFirst
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.MarshalLastParam(false), wantMarshalLastParam)
		asserterror.Equal(t, o.MarshalLastParam(true), wantMarshalLastParamFirst)
	})

	t.Run("MarshalLastParam with Stream option", func(t *testing.T) {
		var (
			o                    = New()
			wantMarshalLastParam = mLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.MarshalLastParam(false), wantMarshalLastParam)
		asserterror.Equal(t, o.MarshalLastParam(true), wantMarshalLastParam)
	})

	t.Run("UnmarshalLastParam", func(t *testing.T) {
		var (
			o                           = New()
			wantUnmarshalLastParam      = uLastParam
			wantUnmarshalLastParamFirst = uLastParamFirst
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.UnmarshalLastParam(false), wantUnmarshalLastParam)
		asserterror.Equal(t, o.UnmarshalLastParam(true), wantUnmarshalLastParamFirst)
	})

	t.Run("UnmarshalLastParam with Stream option", func(t *testing.T) {
		var (
			o                      = New()
			wantUnmarshalLastParam = uLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.UnmarshalLastParam(false), wantUnmarshalLastParam)
		asserterror.Equal(t, o.UnmarshalLastParam(true), wantUnmarshalLastParam)
	})

	t.Run("SkipLastParam", func(t *testing.T) {
		var (
			o                 = New()
			wantSkipLastParam = skLastParam
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.SkipLastParam(), wantSkipLastParam)
	})

	t.Run("SkipLastParam with Stream option", func(t *testing.T) {
		var (
			o                 = New()
			wantSkipLastParam = skLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.SkipLastParam(), wantSkipLastParam)
	})

	t.Run("ModImportName", func(t *testing.T) {
		var (
			o                 = New()
			wantModImportName = modImportName
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.ModImportName(), wantModImportName)
	})

	t.Run("ModImportName with Stream option", func(t *testing.T) {
		var (
			o                 = New()
			wantModImportName = modImportNameStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.ModImportName(), wantModImportName)
	})

	t.Run("ExtPackageName", func(t *testing.T) {
		var (
			o                  = New()
			wantExtPackageName = extPackageName
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(t, o.ExtPackageName(), wantExtPackageName)
	})

	t.Run("ExtPackageName with Stream option", func(t *testing.T) {
		var (
			o                  = New()
			wantExtPackageName = extPackageNameStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(t, o.ExtPackageName(), wantExtPackageName)
	})

	t.Run("Package", func(t *testing.T) {
		var (
			o           = New()
			wantPackage = "exts"
		)
		Apply([]SetOption{WithPackage(wantPackage)}, &o)
		asserterror.Equal(t, o.Package, typename.Package(wantPackage))
	})

}
