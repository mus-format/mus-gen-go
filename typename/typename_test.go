package typename

import (
	"reflect"
	"testing"

	types "github.com/mus-format/mus-gen-go/test/complete_name"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTypeCompleteName(t *testing.T) {
	testCases := []struct {
		t                reflect.Type
		wantCompleteName CompleteName
		wantErr          error
	}{
		{
			t:                reflect.TypeFor[int](),
			wantCompleteName: "int",
		},
		{
			t:                reflect.TypeFor[*int](),
			wantCompleteName: "*int",
		},
		{
			t:                reflect.TypeFor[[3]int](),
			wantCompleteName: "[3]int",
		},
		{
			t:                reflect.TypeFor[*[3]int](),
			wantCompleteName: "*[3]int",
		},
		{
			t:                reflect.TypeFor[[3]*int](),
			wantCompleteName: "[3]*int",
		},
		{
			t:                reflect.TypeFor[[]int](),
			wantCompleteName: "[]int",
		},
		{
			t:                reflect.TypeFor[*[]int](),
			wantCompleteName: "*[]int",
		},
		{
			t:                reflect.TypeFor[[]*int](),
			wantCompleteName: "[]*int",
		},
		{
			t:                reflect.TypeFor[map[int]string](),
			wantCompleteName: "map[int]string",
		},
		{
			t:                reflect.TypeFor[*map[int]string](),
			wantCompleteName: "*map[int]string",
		},
		{
			t:                reflect.TypeFor[map[*int]string](),
			wantCompleteName: "map[*int]string",
		},
		{
			t:                reflect.TypeFor[map[int]*string](),
			wantCompleteName: "map[int]*string",
		},
		{
			t:                reflect.TypeFor[types.Struct](),
			wantCompleteName: "github.com/mus-format/mus-gen-go/test/complete_name/types.Struct",
		},
		{
			t:                reflect.TypeFor[*types.Struct](),
			wantCompleteName: "*github.com/mus-format/mus-gen-go/test/complete_name/types.Struct",
		},
		{
			t:       reflect.TypeFor[**int](),
			wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
		},
		{
			t:       reflect.TypeFor[[3]**int](),
			wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
		},
		{
			t:       reflect.TypeFor[[]**int](),
			wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
		},
		{
			t:       reflect.TypeFor[map[**int]string](),
			wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
		},
		{
			t:       reflect.TypeFor[map[int]**string](),
			wantErr: NewMultiPointerError(reflect.TypeFor[**string]()),
		},
		{
			t:       reflect.TypeFor[struct{}](),
			wantErr: NewUnsupportedTypeError(reflect.TypeFor[struct{}]()),
		},
		{
			t:                reflect.TypeFor[types.ParametrizedStruct[types.Struct]](),
			wantCompleteName: "github.com/mus-format/mus-gen-go/test/complete_name/types.ParametrizedStruct[github.com/mus-format/mus-gen-go/test/complete_name.Struct]",
		},
	}
	for _, c := range testCases {
		cname, err := TypeCompleteName(c.t)
		asserterror.EqualError(t, err, c.wantErr)
		asserterror.Equal(t, cname, c.wantCompleteName)
	}
}

func TestBaseTypeCompletename(t *testing.T) {
	testCases := []struct {
		t                reflect.Type
		wantCompleteName CompleteName
		wantErr          error
	}{
		{
			t:                reflect.TypeFor[types.IntPtr](),
			wantCompleteName: "*int",
		},
		{
			t:                reflect.TypeFor[types.Int](),
			wantCompleteName: "int",
		},
		{
			t:                reflect.TypeFor[types.Array](),
			wantCompleteName: "[3]int",
		},
		{
			t:                reflect.TypeFor[types.ArrayPtr](),
			wantCompleteName: "*[3]int",
		},
		{
			t:                reflect.TypeFor[types.Slice](),
			wantCompleteName: "[]int",
		},
		{
			t:                reflect.TypeFor[types.SlicePtr](),
			wantCompleteName: "*[]int",
		},
		{
			t:                reflect.TypeFor[types.Map](),
			wantCompleteName: "map[int]string",
		},
		{
			t:                reflect.TypeFor[types.MapPtr](),
			wantCompleteName: "*map[int]string",
		},
		{
			t:       reflect.TypeFor[types.DoubleIntPtr](),
			wantErr: NewMultiPointerError(reflect.TypeFor[types.DoubleIntPtr]()),
		},
		{
			t:       reflect.TypeFor[int](),
			wantErr: ErrTypeMismatch,
		},
		{
			t:       reflect.TypeFor[*int](),
			wantErr: ErrTypeMismatch,
		},
		{
			t:       reflect.TypeFor[types.Struct](),
			wantErr: NewUnsupportedTypeError(reflect.TypeFor[types.Struct]()),
		},
	}
	for _, c := range testCases {
		cname, err := BaseTypeCompleteName(c.t)
		asserterror.EqualError(t, err, c.wantErr)
		asserterror.Equal(t, cname, c.wantCompleteName)
	}
}

func TestFullName_Package(t *testing.T) {
	testCases := []struct {
		name    FullName
		wantPkg Package
	}{
		{name: "pkg.Type", wantPkg: "pkg"},
		{name: "Type", wantPkg: ""},
		{name: "pkg123.Type", wantPkg: "pkg123"},
		{name: "pkg.Type[pkg.Type]", wantPkg: "pkg"},
		{name: "", wantPkg: ""},
	}
	for _, c := range testCases {
		asserterror.Equal(t, c.name.Package(), c.wantPkg)
	}
}

func TestFullName_RelName(t *testing.T) {
	testCases := []struct {
		name         FullName
		wantTypeName TypeName
	}{
		{name: "pkg.Type", wantTypeName: "Type"},
		{name: "Type", wantTypeName: "Type"},
		{name: "pkg123.Type", wantTypeName: "Type"},
		{name: "pkg.Type[pkg.Type]", wantTypeName: "Type[pkg.Type]"},
		{name: "", wantTypeName: ""},
	}
	for _, c := range testCases {
		asserterror.Equal(t, c.name.TypeName(), c.wantTypeName)
	}
}

func TestRelName_WithoutSqaures(t *testing.T) {
	testCases := []struct {
		name    RelativeName
		wantStr string
	}{
		{
			name:    RelativeName("pkg.Type"),
			wantStr: "pkg.Type",
		},
		{
			name:    RelativeName("pkg.Type[Param1, Param2]"),
			wantStr: "pkg.Type",
		},
	}
	for _, c := range testCases {
		str := c.name.WithoutSquares()
		asserterror.Equal(t, str, c.wantStr)
	}
}

// package

func TestStrToPkg(t *testing.T) {
	testCases := []struct {
		str     string
		wantPkg Package
		wantErr error
	}{
		{
			str:     "pkg",
			wantPkg: Package("pkg"),
		},
		{
			str:     "+++",
			wantErr: NewInvalidPackageError("+++"),
		},
	}
	for _, c := range testCases {
		pkg, err := StrToPackage(c.str)
		asserterror.EqualError(t, err, c.wantErr)
		asserterror.Equal(t, pkg, c.wantPkg)
	}
}

// pkg_path

func TestStrToPkgPath(t *testing.T) {
	testCases := []struct {
		str         string
		wantPkgPath PkgPath
		wantErr     error
	}{
		{
			str:         "github.com/user/project",
			wantPkgPath: PkgPath("github.com/user/project"),
		},
		{
			str:     "+++",
			wantErr: NewInvalidPkgPathError("+++"),
		},
	}
	for _, c := range testCases {
		pkgPath, err := StrToPkgPath(c.str)
		asserterror.EqualError(t, err, c.wantErr)
		asserterror.Equal(t, pkgPath, c.wantPkgPath)
	}
}

func TestPkgPath_Base(t *testing.T) {
	testCases := []struct {
		pkgPath PkgPath
		wantPkg Package
	}{
		{
			pkgPath: PkgPath("github.com/user/project"),
			wantPkg: Package("project"),
		},
		{
			pkgPath: "+++",
			wantPkg: Package("+++"),
		},
	}
	for _, c := range testCases {
		pkg := Package(c.pkgPath.Base())
		asserterror.Equal(t, pkg, c.wantPkg)
	}
}
