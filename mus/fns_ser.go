package musgen

import (
	"fmt"
	"strings"
	"unicode"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	cnvtr "github.com/mus-format/mus-gen-go/spec/converter"
	"github.com/mus-format/mus-gen-go/typename"
)

const (
	SuffixMUS      = "MUS"
	SuffixTypedMUS = "TypedMUS"
	SuffixDTM      = "DTM"
)

type SerFns struct {
	converter     cnvtr.TypeNameConverter
	typeBuilder   TypeBuilder
	crossgenTypes map[typename.FullName]struct{}
	serNames      map[typename.FullName]string
	utilFns       TmplFns
	gopts         genopts.Options
}

func NewSerFns(converter cnvtr.TypeNameConverter,
	typeBuilder TypeBuilder,
	crossgenTypes map[typename.FullName]struct{},
	utilFns TmplFns,
	gopts genopts.Options,
) SerFns {

	serNames := map[typename.FullName]string{}
	for t, serName := range gopts.SerNames { // serName could be with pkg
		cname, err := typename.TypeCompleteName(t)
		if err != nil {
			panic(err)
		}
		fullName := converter.ConvertToFullName(cname)
		serNames[fullName] = serName
	}

	return SerFns{
		converter:     converter,
		typeBuilder:   typeBuilder,
		crossgenTypes: crossgenTypes,
		serNames:      serNames,
		utilFns:       utilFns,
		gopts:         gopts,
	}
}

func (f SerFns) Var(name typename.FullName) string {
	return string(f.typeName(name)) + SuffixMUS
}

func (f SerFns) TypedVar(fullName typename.FullName) string {
	return string(f.typeName(fullName)) + SuffixTypedMUS
}

func (f SerFns) DTMVar(fullName typename.FullName) string {
	return string(f.typeName(fullName)) + SuffixDTM
}

func (f SerFns) Type(fullName typename.FullName) string {
	return f.uncapitalize(f.Var(fullName))
}

func (f SerFns) Val(tp any) string {
	return "v"
}

func (f SerFns) Of(fullName typename.FullName, tops tpopts.Options,
	gops genopts.Options,
) string {
	anonType, ok, err := f.typeBuilder.BuildAnonType(fullName, tops)
	if err != nil {
		panic(fmt.Sprintf("can't get serializer for the %v, cause %v", fullName, err))
	}
	if ok {
		return string(anonType.SerName)
	}
	var (
		pkg     string
		serName string
		rules   map[string]string
	)

	switch gops.Mode {
	case genopts.ModeSafe:
		rules = safeModeRules(tops)
	case genopts.ModeNotUnsafe:
		rules = notUnsafeModeRules(tops, gops)
	case genopts.ModeUnsafe:
		rules = unsafeModeRules(tops, gops)
	}
	pkg, pst := rules[string(fullName)]
	if !pst {
		return f.Var(fullName)
	}

	switch fullName {
	case "int", "int64", "int32", "int16", "int8":
		serName = f.capitalize(string(fullName))
		if tops.NumEncoding == tpopts.NumEncodingVarintPositive {
			serName = "Positive" + serName
		}
	case "uint", "uint64", "uint32", "uint16", "uint8", "float64", "float32", "byte":
		serName = f.capitalize(string(fullName))
	case "bool", "string":
		serName = f.capitalize(string(fullName))
	case "[]byte", "[]uint8":
		serName = "ByteSlice"
	case "time.Time":
		serName = rules["time_ser"]
	default:
		panic(fmt.Sprintf("internal bug: unpredicted case %v", fullName))
	}
	return pkg + "." + serName
}

func (f SerFns) Key(anonType spec.AnonType, gops genopts.Options) string {
	return f.Of(anonType.KeyType, tpopts.Options{}, gops)
}

func (f SerFns) Elem(anonType spec.AnonType, gops genopts.Options) string {
	return f.Of(anonType.ElemType, tpopts.Options{}, gops)
}

func (f SerFns) RelName(name typename.FullName) string {
	return string(f.converter.ConvertToRelativeName(name))
}

func (f SerFns) FieldTmplPipe(structType spec.StructType,
	fieldType spec.FieldType, index int, gops genopts.Options) FieldTmplPipe {
	pipe := FieldTmplPipe{
		Val:         f.Val(structType),
		FieldsCount: len(structType.SerializedFields()),
		Field:       fieldType,
		Index:       index,
		Gops:        gops,
	}
	if len(structType.Sops.Fields) > 0 {
		pipe.Tops = structType.Sops.Fields[index].Type
	}
	return pipe
}

func (f SerFns) StringOpts(anonData spec.AnonType) string {
	b := strings.Builder{}
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("stropts.WithLenSer(%v)", anonData.LenSer))
	}
	if anonData.LenVl != "" && anonData.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("stropts.WithLenValidator(%v)", anonData.LenVl))
	}
	return b.String()
}

func (f SerFns) ArrayOpts(anonType spec.AnonType) string {
	var (
		b        = strings.Builder{}
		elemType = f.converter.ConvertToRelativeName(anonType.ElemType)
	)
	if anonType.LenSer != "" && anonType.LenSer != "nil" {
		fmt.Fprintf(&b, "arropts.WithLenSer[%v](%v)", elemType, anonType.LenSer)
	}
	if anonType.ElemVl != "" && anonType.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "arropts.WithElemValidator[%v](%v)", elemType, anonType.ElemVl)
	}
	return b.String()
}

func (f SerFns) ByteSliceOpts(anonType spec.AnonType) string {
	b := strings.Builder{}
	if anonType.LenSer != "" && anonType.LenSer != "nil" {
		fmt.Fprintf(&b, "bslopts.WithLenSer(%v)", anonType.LenSer)
	}
	if anonType.LenVl != "" && anonType.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "bslopts.WithLenValidator(%v)", anonType.LenVl)
	}
	return b.String()
}

func (f SerFns) SliceOpts(anonType spec.AnonType) string {
	var (
		b        = strings.Builder{}
		elemType = f.converter.ConvertToRelativeName(anonType.ElemType)
	)
	if anonType.LenSer != "" && anonType.LenSer != "nil" {
		fmt.Fprintf(&b, "slopts.WithLenSer[%v](%v)", elemType, anonType.LenSer)
	}
	if anonType.LenVl != "" && anonType.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "slopts.WithLenValidator[%v](%v)", elemType, anonType.LenVl)
	}
	if anonType.ElemVl != "" && anonType.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "slopts.WithElemValidator[%v](%v)", elemType, anonType.ElemVl)
	}
	return b.String()
}

func (f SerFns) MapOpts(anonType spec.AnonType) string {
	var (
		b        = strings.Builder{}
		keyType  = f.converter.ConvertToRelativeName(anonType.KeyType)
		elemType = f.converter.ConvertToRelativeName(anonType.ElemType)
	)
	if anonType.LenSer != "" && anonType.LenSer != "nil" {
		fmt.Fprintf(&b, "mapopts.WithLenSer[%v, %v](%v)", keyType, elemType,
			anonType.LenSer)
	}
	if anonType.LenVl != "" && anonType.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "mapopts.WithLenValidator[%v, %v](%v)", keyType, elemType,
			anonType.LenVl)
	}
	if anonType.KeyVl != "" && anonType.KeyVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "mapopts.WithKeyValidator[%v, %v](%v)", keyType, elemType,
			anonType.KeyVl)
	}
	if anonType.ElemVl != "" && anonType.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "mapopts.WithValueValidator[%v, %v](%v)", keyType, elemType, anonType.ElemVl)
	}
	return b.String()
}

func (f SerFns) typeName(name typename.FullName) string {
	if serName, pst := f.serNames[name]; pst {
		return serName
	}
	var rname typename.RelativeName
	if _, pst := f.crossgenTypes[name]; pst {
		rname = typename.RelativeName(name.TypeName())
	} else {
		rname = f.converter.ConvertToRelativeName(name)
	}
	return rname.WithoutSquares()
}

func (f SerFns) capitalize(str string) string {
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (f SerFns) uncapitalize(str string) string {
	r := []rune(str)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
