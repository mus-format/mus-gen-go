package scanner

import (
	"regexp"

	"github.com/mus-format/mus-gen-go/typename"
)

const PrimitivePattern = `^(\**)(uint8|uint16|uint32|uint64|int8|int16|int32|` +
	`int64|uint|int|byte|bool|string|float32|float64)`

var PrimitiveRe = regexp.MustCompile(PrimitivePattern)

func ParsePrimitiveType[T QualifiedName](name T) (t TypeInfo[T], ok bool) {
	return PrimitiveType[T](name).Parse()
}

type PrimitiveType[T QualifiedName] string

func (s PrimitiveType[T]) Parse() (t TypeInfo[T], ok bool) {
	match := PrimitiveRe.FindStringSubmatch(string(s))
	if len(match) != 3 {
		return
	}
	t = TypeInfo[T]{
		Stars: match[1],
		Name:  typename.TypeName(match[2]),
	}
	ok = true
	return
}
