package genopts

import (
	"reflect"
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestNew(t *testing.T) {
	o := New()
	asserterror.Equal(t, o.SerNames != nil, true)
	asserterror.Equal(t, o.importAliases != nil, true)
}

func TestImportAliases(t *testing.T) {
	o := New()
	o.importAliases["path"] = "alias"
	asserterror.Equal(t, len(o.ImportAliases()), 1)
	asserterror.Equal(t, o.ImportAliases()["path"], Alias("alias"))
}

func TestWithPkgPath(t *testing.T) {
	pkgPath := "github.com/user/project"
	o := New()

	asserterror.EqualError(t, WithPkgPath(pkgPath)(&o), nil)
	asserterror.Equal(t, string(o.PkgPath), pkgPath)
	asserterror.Equal(t, string(o.Package), "project")

	// Test invalid path
	asserterror.Equal(t, WithPkgPath("a b")(&o) != nil, true)
}

func TestWithPackage(t *testing.T) {
	pkg := "project"
	o := New()

	asserterror.EqualError(t, WithPackage(pkg)(&o), nil)
	asserterror.Equal(t, string(o.Package), pkg)

	// Test invalid package
	asserterror.Equal(t, WithPackage("a b")(&o) != nil, true)
}

func TestWithImport(t *testing.T) {
	path := "fmt"
	o := New()

	asserterror.EqualError(t, WithImport(path)(&o), nil)
	asserterror.Equal(t, len(o.Imports), 1)
	asserterror.Equal(t, o.Imports[0], `"`+path+`"`)

	// Test invalid path
	asserterror.Equal(t, WithImport("a b")(&o) != nil, true)
}

func TestWithImportAlias(t *testing.T) {
	path := "fmt"
	alias := "f"
	o := New()

	asserterror.EqualError(t, WithImportAlias(path, alias)(&o), nil)
	asserterror.Equal(t, len(o.Imports), 1)
	asserterror.Equal(t, o.Imports[0], alias+" "+`"`+path+`"`)
	asserterror.Equal(t, string(o.importAliases[ImportPath(path)]), alias)

	// Test invalid path
	asserterror.Equal(t, WithImportAlias("a b", alias)(&o) != nil, true)

	// Test invalid alias
	asserterror.Equal(t, WithImportAlias(path, "a b")(&o) != nil, true)
}

func TestWithSerName(t *testing.T) {
	typ := reflect.TypeFor[int]()
	serName := "Int"
	o := New()

	asserterror.EqualError(t, WithSerName(typ, serName)(&o), nil)
	asserterror.Equal(t, o.SerNames[typ], serName)
}

func TestWithUnsafe(t *testing.T) {
	o := New()
	asserterror.EqualError(t, WithUnsafe()(&o), nil)
	asserterror.Equal(t, o.Mode, ModeUnsafe)
}

func TestWithNotUnsafe(t *testing.T) {
	o := New()
	asserterror.EqualError(t, WithNotUnsafe()(&o), nil)
	asserterror.Equal(t, o.Mode, ModeNotUnsafe)
}

func TestWithStream(t *testing.T) {
	o := New()
	asserterror.EqualError(t, WithStream()(&o), nil)
	asserterror.Equal(t, o.Stream, true)
}

func TestValidate(t *testing.T) {
	o := New()

	// Empty PkgPath
	asserterror.EqualError(t, o.Validate(), ErrEmptyPkgPath)

	o.PkgPath = "path"
	o.importAliases["path1"] = "alias"
	o.importAliases["path2"] = "alias"

	// Duplicate alias
	asserterror.EqualError(t, o.Validate(), NewDuplicateImportAlias("alias"))

	// Valid
	delete(o.importAliases, "path2")
	asserterror.EqualError(t, o.Validate(), nil)
}

func TestApply(t *testing.T) {
	o := New()
	ops := []SetOption{
		WithPkgPath("github.com/user/project"),
		WithStream(),
		nil,
	}

	asserterror.EqualError(t, Apply(ops, &o), nil)
	asserterror.Equal(t, string(o.PkgPath), "github.com/user/project")
	asserterror.Equal(t, o.Stream, true)
	asserterror.Equal(t, string(o.importAliases[ImportPath(o.PkgPath)]), string(o.Package))

	// Test error in Apply
	ops = append(ops, WithPkgPath("a b"))
	asserterror.Equal(t, Apply(ops, &o) != nil, true)
}

func TestErrors(t *testing.T) {
	path := ImportPath("path")
	asserterror.Equal(t, NewDuplicateImportPath(path).Error(), NewDuplicateImportPath(path).Error())
	asserterror.Equal(t, NewInvalidImportPathError("path").Error(), NewInvalidImportPathError("path").Error())
	asserterror.Equal(t, NewInvalidAliasError("alias").Error(), NewInvalidAliasError("alias").Error())
}
