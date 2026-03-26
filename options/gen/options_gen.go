// Package genopts provides options for MUS code generation.
package genopts

import (
	"go/token"
	"reflect"

	"github.com/mus-format/mus-gen-go/typename"
	"golang.org/x/mod/module"
)

// ImportPath represents a standard Go import path.
type ImportPath string

// Alias represents a Go package alias used in imports.
type Alias string

// Options specifies configuration options for code generation.
type Options struct {
	PkgPath typename.PkgPath
	Package typename.Package

	Imports  []string
	SerNames map[reflect.Type]string

	Mode   Mode
	Stream bool

	importAliases map[ImportPath]Alias
}

// New returns a new Options struct with initialized internal maps.
func New() Options {
	return Options{
		SerNames:      make(map[reflect.Type]string),
		importAliases: make(map[ImportPath]Alias),
	}
}

// ImportAliases returns a map of import paths to their respective aliases.
func (o Options) ImportAliases() map[ImportPath]Alias {
	return o.importAliases
}

// SetOption defines a function type that applies a configuration option to an
// Options struct.
type SetOption func(o *Options) error

// WithPkgPath returns a SetOption that configures the package path for the
// generated file. The path must match the standard Go package path format
// (e.g., github.com/user/project/pkg) and could be obtained like:
//
//	pkgPath := reflect.TypeFor[YourType]().PkgPath()
func WithPkgPath(str string) SetOption {
	return func(o *Options) (err error) {
		o.PkgPath, err = typename.StrToPkgPath(str)
		if err != nil {
			return
		}
		if o.Package == "" {
			o.Package, err = typename.StrToPackage(o.PkgPath.Base())
		}
		return
	}
}

// WithPackage returns a SetOption that sets the package name for the generated
// file.
func WithPackage(str string) SetOption {
	return func(o *Options) (err error) {
		o.Package, err = typename.StrToPackage(str)
		return
	}
}

// WithImport returns a SetOption that adds the given import path to the list
// of imports.
func WithImport(importPath string) SetOption {
	return func(o *Options) (err error) {
		if err = module.CheckImportPath(importPath); err != nil {
			err = NewInvalidImportPathError(importPath)
			return
		}
		o.Imports = append(o.Imports, `"`+importPath+`"`)
		return
	}
}

// WithImportAlias returns a SetOption that adds the given import path and alias
// to the list of imports.
func WithImportAlias(importPath, alias string) SetOption {
	return func(o *Options) (err error) {
		if err = module.CheckImportPath(importPath); err != nil {
			err = NewInvalidImportPathError(importPath)
			return
		}
		if !token.IsIdentifier(alias) {
			err = NewInvalidAliasError(alias)
			return
		}
		o.Imports = append(o.Imports, alias+" "+`"`+importPath+`"`)
		o.importAliases[ImportPath(importPath)] = Alias(alias)
		return
	}
}

// WithSerName returns a SetOption that registers a custom serializer name for a
// specific type.
func WithSerName(t reflect.Type, serName string) SetOption {
	return func(o *Options) (err error) {
		o.SerNames[t] = serName
		return
	}
}

// WithUnsafe returns a SetOption that enables unsafe code generation.
func WithUnsafe() SetOption {
	return func(o *Options) (err error) {
		o.Mode = ModeUnsafe
		return
	}
}

// WithNotUnsafe returns a SetOption that enables unsafe code generation without
// side effects. When applied, the generator will avoid unsafe operations for
// the string type.
func WithNotUnsafe() SetOption {
	return func(o *Options) (err error) {
		o.Mode = ModeNotUnsafe
		return
	}
}

// WithStream returns a SetOption that enables streaming code generation.
// When applied, the generator will produce code that can process data
// incrementally rather than requiring complete in-memory representations.
func WithStream() SetOption {
	return func(o *Options) (err error) {
		o.Stream = true
		return
	}
}

// Validate checks if the options are valid and returns an error if not.
func (o Options) Validate() (err error) {
	if o.PkgPath == "" {
		return ErrEmptyPkgPath
	}
	var aliases = make(map[Alias]struct{})
	for _, alias := range o.importAliases {
		if _, ok := aliases[alias]; ok {
			return NewDuplicateImportAlias(alias)
		}
		aliases[alias] = struct{}{}
	}
	return
}

// Apply applies a set of options to the provided Options struct and validates
// the result.
func Apply(ops []SetOption, o *Options) (err error) {
	for i := range ops {
		if ops[i] != nil {
			err = ops[i](o)
			if err != nil {
				return
			}
		}
	}
	addImportAlias(o)
	return o.Validate()
}

func addImportAlias(o *Options) {
	o.importAliases[ImportPath(o.PkgPath)] = Alias(o.Package)
}
