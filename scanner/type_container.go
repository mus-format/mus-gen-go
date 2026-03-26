package scanner

import (
	"regexp"
	"strings"

	"github.com/mus-format/mus-gen-go/typename"
)

const (
	ArrayPattern = `^(\**)(\[(\d+)\](.+$))`
	SlicePattern = `^(\**)(\[\](.+$))`
)

var (
	ArrayRe = regexp.MustCompile(ArrayPattern)
	SliceRe = regexp.MustCompile(SlicePattern)
)

func ParseContainerType[T QualifiedName](name T) (t TypeInfo[T], keyType, elemStr T,
	kind Kind, ok bool) {
	if t, elemStr, ok = ContainerType[T](name).ParseArray(); ok {
		kind = KindArray
		return
	}
	if t, elemStr, ok = ContainerType[T](name).ParseSlice(); ok {
		kind = KindSlice
		return
	}
	if t, keyType, elemStr, ok = ContainerType[T](name).ParseMap(); ok {
		kind = KindMap
		return
	}
	return
}

type ContainerType[T QualifiedName] string

func (s ContainerType[T]) ParseArray() (t TypeInfo[T], elemType T, ok bool) {
	match := ArrayRe.FindStringSubmatch(string(s))
	if len(match) != 5 {
		return
	}
	t = TypeInfo[T]{
		Stars:     match[1],
		Name:      typename.TypeName(match[2]),
		ArrLength: match[3],
		ElemType:  T(match[4]),
	}
	elemType = t.ElemType
	ok = true
	return
}

func (s ContainerType[T]) ParseSlice() (t TypeInfo[T], elemType T, ok bool) {
	match := SliceRe.FindStringSubmatch(string(s))
	if len(match) != 4 {
		return
	}
	t = TypeInfo[T]{
		Stars:    match[1],
		Name:     typename.TypeName(match[2]),
		ElemType: T(match[3]),
	}
	elemType = t.ElemType
	ok = true
	return
}

func (s ContainerType[T]) ParseMap() (t TypeInfo[T], keyType, elemType T, ok bool) {
	str := string(s)
	// Extract leading asterisks (stars)
	i := 0
	for i < len(str) && str[i] == '*' {
		i++
	}
	starsTmp := str[:i]
	str = str[i:]

	const prefix = "map["
	if !strings.HasPrefix(str, prefix) {
		ok = false
		return
	}
	stars := starsTmp

	// Skip "map[" prefix
	start := len(prefix)
	depth := 1
	j := start

	for j < len(str) && depth > 0 {
		switch str[j] {
		case '[':
			depth++
		case ']':
			depth--
		}
		j++
	}

	if depth != 0 {
		ok = false
		return // unbalanced brackets
	}

	t = TypeInfo[T]{
		Stars: stars,
	}
	t.KeyType = T(strings.TrimSpace(str[start : j-1]))
	t.ElemType = T(strings.TrimSpace(str[j:]))
	t.Name = typename.TypeName("map[" + string(t.KeyType) + "]" + string(t.ElemType))

	keyType = t.KeyType
	elemType = t.ElemType
	ok = true
	return
}
