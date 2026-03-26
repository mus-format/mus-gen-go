package typename

import (
	"go/token"
)

type Package string

func StrToPackage(str string) (pkg Package, err error) {
	if !token.IsIdentifier(str) {
		err = NewInvalidPackageError(str)
		return
	}
	pkg = Package(str)
	return
}
