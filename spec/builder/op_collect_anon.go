package builder

import (
	"crypto/md5"
	"strings"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/spec"
	"github.com/mus-format/mus-gen-go/typename"
)

func NewCollectAnonOp(m map[spec.AnonSerName]spec.AnonType,
	converter TypeNameConverter) *CollectAnonOp {
	return &CollectAnonOp{m: m, converter: converter}
}

// CollectAnonOp
type CollectAnonOp struct {
	m             map[spec.AnonSerName]spec.AnonType
	firstAnonType *spec.AnonType
	converter     TypeNameConverter
}

func (o *CollectAnonOp) ProcessType(typeInfo scanner.TypeInfo[typename.FullName],
	tops tpopts.Options) (err error) {
	var (
		anonType spec.AnonType
		ok       bool
	)
	if typeInfo.Stars != "" {
		return o.processPtrType(typeInfo, tops)
	}
	switch typeInfo.Kind {
	case scanner.KindPrim:
		anonType, ok = o.makePrimType(typeInfo, tops)
	case scanner.KindArray:
		anonType, ok = o.makeArrayType(typeInfo, tops)
	case scanner.KindSlice:
		if typeInfo.ElemType == "byte" || typeInfo.ElemType == "uint8" {
			anonType, ok = o.makeByteSliceType(typeInfo, tops)
		} else {
			anonType, ok = o.makeSliceType(typeInfo, tops)
		}
	case scanner.KindMap:
		anonType, ok = o.makeMapType(typeInfo, tops)
	}
	if ok {
		o.putToMap(anonType)
	}
	return
}

func (o *CollectAnonOp) ProcessLeftSquare() {}

func (o *CollectAnonOp) ProcessComma() {}

func (o *CollectAnonOp) ProcessRightSquare() {}

func (o *CollectAnonOp) FirstAnonType() *spec.AnonType {
	return o.firstAnonType
}

func (o *CollectAnonOp) processPtrType(typeInfo scanner.TypeInfo[typename.FullName],
	tops tpopts.Options) (err error) {
	d, _ := o.makePtrType(typeInfo, tops)
	o.putToMap(d)
	typeInfo.Stars = trimOneStar(typeInfo.Stars)
	return o.ProcessType(typeInfo, tpopts.Options{})
}

func (o *CollectAnonOp) makePtrType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (d spec.AnonType, ok bool) {
	str := t.Stars + string(typename.MakeFullName(t.Package, t.Name))
	return spec.AnonType{
		SerName:  anonSerName(spec.AnonKindPtr, t, tops),
		ElemType: typename.FullName(trimOneStar(str)),
		Kind:     spec.AnonKindPtr,
	}, true
}

func (o *CollectAnonOp) makePrimType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (d spec.AnonType, ok bool) {
	if string(t.Name) == "string" {
		lenSer := "nil"
		if tops.LenEnc != tpopts.NumEncodingUndefined &&
			tops.LenEnc != tpopts.NumEncodingVarintPositive {
			lenSer = tops.LenEnc.LenSer()
			ok = true
		}
		lenVl := "nil"
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
			ok = true
		}
		if ok {
			return spec.AnonType{
				SerName: anonSerName(spec.AnonKindString, t, tops),
				Kind:    spec.AnonKindString,
				LenSer:  lenSer,
				LenVl:   lenVl,
			}, ok
		}
	}
	return
}

func (o *CollectAnonOp) makeArrayType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (d spec.AnonType, ok bool) {
	var (
		lenSer = "nil"
		elemVl = "nil"
	)
	if tops.LenEnc != tpopts.NumEncodingUndefined {
		lenSer = tops.LenEnc.LenSer()
	}
	if tops.ElemValidator != "" {
		elemVl = o.validatorStr(t.ElemType, tops.ElemValidator)
	}
	return spec.AnonType{
		SerName:   anonSerName(spec.AnonKindArray, t, tops),
		Kind:      spec.AnonKindArray,
		ArrType:   typename.MakeFullName(t.Package, t.Name),
		ArrLength: t.ArrLength,
		LenSer:    lenSer,
		ElemType:  t.ElemType,
		ElemVl:    elemVl,
	}, true
}

func (o *CollectAnonOp) makeByteSliceType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (d spec.AnonType, ok bool) {
	if tops.LenEnc != tpopts.NumEncodingUndefined ||
		tops.LenValidator != "" {
		var (
			lenSer = "nil"
			lenVl  = "nil"
		)
		if tops.LenEnc != tpopts.NumEncodingUndefined {
			lenSer = tops.LenEnc.LenSer()
		}
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
		}
		return spec.AnonType{
			SerName: anonSerName(spec.AnonKindByteSlice, t, tops),
			Kind:    spec.AnonKindByteSlice,
			LenSer:  lenSer,
			LenVl:   lenVl,
		}, true
	}
	return
}

func (o *CollectAnonOp) makeSliceType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options) (d spec.AnonType, ok bool) {
	var (
		lenSer = "nil"
		lenVl  = "nil"
		elemVl = "nil"
	)
	if tops.LenEnc != tpopts.NumEncodingUndefined {
		lenSer = tops.LenEnc.LenSer()
	}
	if tops.LenValidator != "" {
		lenVl = o.validatorStr("int", tops.LenValidator)
	}
	if tops.ElemValidator != "" {
		elemVl = o.validatorStr(t.ElemType, tops.ElemValidator)
	}
	return spec.AnonType{
		SerName:  anonSerName(spec.AnonKindSlice, t, tops),
		Kind:     spec.AnonKindSlice,
		LenSer:   lenSer,
		LenVl:    lenVl,
		ElemType: t.ElemType,
		ElemVl:   elemVl,
	}, true
}

func (o *CollectAnonOp) makeMapType(t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options,
) (d spec.AnonType, ok bool) {
	var (
		lenSer = "nil"
		lenVl  = "nil"
		keyVl  = "nil"
		elemVl = "nil"
	)
	if tops.LenEnc != tpopts.NumEncodingUndefined {
		lenSer = tops.LenEnc.LenSer()
	}
	if tops.LenValidator != "" {
		lenVl = o.validatorStr("int", tops.LenValidator)
	}
	if tops.KeyValidator != "" {
		keyVl = o.validatorStr(t.KeyType, tops.KeyValidator)
	}
	if tops.ElemValidator != "" {
		elemVl = o.validatorStr(t.ElemType, tops.ElemValidator)
	}
	return spec.AnonType{
		SerName:  anonSerName(spec.AnonKindMap, t, tops),
		Kind:     spec.AnonKindMap,
		LenSer:   lenSer,
		LenVl:    lenVl,
		KeyType:  t.KeyType,
		KeyVl:    keyVl,
		ElemType: t.ElemType,
		ElemVl:   elemVl,
	}, true
}

func (o *CollectAnonOp) putToMap(anonType spec.AnonType) {
	if o.firstAnonType == nil {
		o.firstAnonType = &anonType
	}
	o.m[anonType.SerName] = anonType
}

func (o *CollectAnonOp) validatorStr(name typename.FullName, vl string) string {
	relName := o.converter.ConvertToRelativeName(name)
	return "com.ValidatorFn[" + string(relName) + "](" + vl + ")"
}

func anonSerName(kind spec.AnonKind, t scanner.TypeInfo[typename.FullName],
	tops tpopts.Options) spec.AnonSerName {
	bs := []byte(kind.String())
	switch kind {
	case spec.AnonKindString:
		bs = append(bs, []byte(tops.LenEnc.Package())...)
		bs = append(bs, []byte(tops.LenValidator)...)
	case spec.AnonKindArray:
		bs = append(bs, []byte(t.ElemType)...)
		bs = append(bs, []byte(tops.LenEnc.Package())...)
		bs = append(bs, []byte(tops.ElemValidator)...)
	case spec.AnonKindByteSlice:
		bs = append(bs, []byte(tops.LenEnc.Package())...)
		bs = append(bs, []byte(tops.LenValidator)...)
	case spec.AnonKindSlice:
		bs = append(bs, []byte(t.ElemType)...)
		bs = append(bs, []byte(tops.LenEnc.Package())...)
		bs = append(bs, []byte(tops.LenValidator)...)
		bs = append(bs, []byte(tops.ElemValidator)...)
	case spec.AnonKindMap:
		bs = append(bs, []byte(t.KeyType)...)
		bs = append(bs, []byte(t.ElemType)...)
		bs = append(bs, []byte(tops.LenEnc.Package())...)
		bs = append(bs, []byte(tops.LenValidator)...)
		bs = append(bs, []byte(tops.KeyValidator)...)
		bs = append(bs, []byte(tops.ElemValidator)...)
	case spec.AnonKindPtr:
		bs = append(bs, []byte(t.ElemType)...)
	}
	h := md5.Sum(bs)
	return spec.AnonSerName(kind.String() + Base64KeywordEncoding.EncodeToString(h[:]))
}

func trimOneStar(s string) string {
	if strings.HasPrefix(s, "*") {
		return s[1:]
	}
	return s
}
