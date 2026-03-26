package typename

import (
	"regexp"
	"strings"
)

const pkgTypeNamePattern = `^(\w+)\.(.+)$`

var pkgTypeNameRe = regexp.MustCompile(pkgTypeNamePattern)

// FullName examples: "*pkg.TypeName", "[]pkg.TypeName",
// "pkg.TypeName[pkg.Typename]".
type FullName string

func MakeFullName(pkg Package, name TypeName) FullName {
	if pkg == "" {
		return FullName(name)
	}
	return FullName(strings.Join([]string{string(pkg), string(name)}, "."))
}

func (n FullName) Package() Package {
	match := pkgTypeNameRe.FindStringSubmatch(string(n))
	if len(match) != 3 {
		return ""
	}
	return Package(match[1])
}

func (n FullName) TypeName() TypeName {
	match := pkgTypeNameRe.FindStringSubmatch(string(n))
	if len(match) != 3 {
		return TypeName(n)
	}
	return TypeName(match[2])
}
