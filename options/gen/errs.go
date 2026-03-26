package genopts

import (
	"errors"
	"fmt"
)

// ErrEmptyPkgPath is returned when the package path option is not set.
var ErrEmptyPkgPath = errors.New("option PkgPath is not set")

// NewInvalidImportPathError returns an error indicating that an import path is
// invalid.
func NewInvalidImportPathError(val string) error {
	return fmt.Errorf("invalid '%v' import path", val)
}

// NewInvalidAliasError returns an error indicating that a package alias is
// invalid.
func NewInvalidAliasError(val string) error {
	return fmt.Errorf("invalid '%v' package alias", val)
}

// NewDuplicateImportPath returns an error indicating that an import path is
// duplicated.
func NewDuplicateImportPath(path ImportPath) error {
	return fmt.Errorf("duplicate '%v' import path in musgen.Generator options", path)
}

// NewDuplicateImportAlias returns an error indicating that a package alias is
// duplicated.
func NewDuplicateImportAlias(alias Alias) error {
	return fmt.Errorf("duplicate '%v' package alias in musgen.Generator options",
		alias)
}
