package typename

import "strings"

// RelativeName examples: "TypeName", "pkg.TypeName".
type RelativeName FullName

func (n RelativeName) WithoutSquares() (str string) {
	str = string(n)
	before, _, ok := strings.Cut(str, "[")
	if !ok {
		return str // no type parameters
	}
	return before
}
