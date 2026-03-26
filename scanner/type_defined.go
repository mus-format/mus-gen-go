package scanner

import (
	"regexp"
	"strings"

	"github.com/mus-format/mus-gen-go/typename"
)

const (
	// Works for strings like "*github.com/user-name/project/some/some.Type" and
	// "foo.Type" (when module name is just foo).
	FullNamePattern = `^(\**)(?:([.\w/-]+)/)*(?:([\w-]+)\.)([\w]+)(?:·\d+)*(?:\[(.+)\](?:·\d+)*)?$`
)

func ParseDefinedType[T QualifiedName](name T) (t TypeInfo[T], ok bool) {
	return DefinedType[T](name).Parse()
}

type DefinedType[T QualifiedName] string

func (s DefinedType[T]) Parse() (t TypeInfo[T], ok bool) {
	re := regexp.MustCompile(FullNamePattern)
	match := re.FindStringSubmatch(string(s))
	if len(match) != 6 {
		return
	}
	return s.makeDefinedType(match), true
}

func (s DefinedType[T]) makeDefinedType(match []string) (t TypeInfo[T]) {
	t.PkgPath = typename.PkgPath(match[2])
	t.Stars = match[1]
	t.Package = typename.Package(match[3])
	t.Name = typename.TypeName(match[4])
	t.Params = s.splitTypeParams(match[5])
	return
}

// splitTypeParams properly splits generic parameters while keeping nested
// brackets intact.
func (s DefinedType[T]) splitTypeParams(params string) (result []T) {
	var buf strings.Builder
	depth := 0
	for _, char := range params {
		switch char {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				// Found a top-level comma, finalize the current param
				result = append(result, T(strings.TrimSpace(buf.String())))
				buf.Reset()
				continue
			}
		}
		buf.WriteRune(char)
	}
	// Add the last parameter if there’s any remaining content
	if buf.Len() > 0 {
		result = append(result, T(strings.TrimSpace(buf.String())))
	}
	return
}
