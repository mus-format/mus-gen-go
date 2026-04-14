// Package builder provides a builder for MUS type specifications.
package builder

import (
	"reflect"

	"github.com/mus-format/mus-gen-go/classifier"
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	structops "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	veropts "github.com/mus-format/mus-gen-go/options/versioned"
	"github.com/mus-format/mus-gen-go/scanner"
	scnr "github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/spec"
	"github.com/mus-format/mus-gen-go/typename"
)

func NewTypeBuilder(converter TypeNameConverter, gops genopts.Options) TypeBuilder {
	return TypeBuilder{converter, gops}
}

type TypeBuilder struct {
	converter TypeNameConverter
	gops      genopts.Options
}

func (b TypeBuilder) BuildDefinedType(t reflect.Type, tops tpopts.Options) (
	definedTime spec.DefinedType, err error) {
	if !classifier.DefinedBasicType(t) {
		err = b.notDefinedTypeError(t)
		return
	}
	var (
		typeName           = b.parseTypeName(t)
		underlyingTypeName = b.parseUnderlyingTypeName(t)
	)
	definedTime.FullName = typeName
	definedTime.UnderlyingTypeName = underlyingTypeName
	definedTime.Tops = tops
	definedTime.Gops = b.gops
	return
}

func (b TypeBuilder) BuildStructType(t reflect.Type, sops structops.Options) (
	structType spec.StructType, err error,
) {
	if !classifier.DefinedStruct(t) {
		err = b.notStructError(t)
		return
	}
	if err = b.checkFields(t, sops); err != nil {
		return
	}
	var (
		typeName   = b.parseTypeName(t)
		fieldNames = b.parseFieldNames(t)
	)
	structType.FullName = typeName
	structType.Fields = b.makeFields(t, fieldNames)
	structType.Sops = sops
	structType.Gops = b.gops
	return
}

func (b TypeBuilder) BuildTyped(t reflect.Type) (
	tp spec.Type, err error,
) {
	if !(classifier.DefinedBasicType(t) || classifier.DefinedStruct(t) ||
		classifier.DefinedNonEmptyInterface(t)) {
		err = NewUnsupportedTypeError(t)
		return
	}
	tp.FullName = b.parseTypeName(t)
	tp.Gops = b.gops
	return
}

func (b TypeBuilder) BuildInterfaceType(t reflect.Type, iops intropts.Options) (
	interfaceType spec.InterfaceType, err error,
) {
	if !classifier.DefinedInterface(t) {
		err = b.notInterfaceError(t)
		return
	}
	var (
		typeName = b.parseTypeName(t)
		impls    = b.parseImpls(iops)
	)
	interfaceType.FullName = typeName
	interfaceType.Impls = impls
	interfaceType.Iops = iops
	interfaceType.Gops = b.gops
	return
}

func (b TypeBuilder) BuildVersionedType(t reflect.Type, vops veropts.Options) (
	versionedType spec.VersionedType, err error,
) {
	if !classifier.DefinedType(t) {
		err = b.notInterfaceError(t)
		return
	}
	var (
		typeName = b.parseTypeName(t)
		versions = b.parseVersions(vops)
	)
	versionedType.FullName = typeName
	versionedType.Versions = versions
	versionedType.Vops = vops
	versionedType.Gops = b.gops
	return
}

func (b TypeBuilder) BuildTimeType(t reflect.Type, uops stopts.UnderlyingTimeOptions,
	validator string,
) (definedType spec.DefinedType, err error) {
	if !classifier.DefinedStruct(t) {
		err = NewNotStructError(t)
		return
	}
	var tops tpopts.Options
	tpopts.Apply(&tops,
		tpopts.WithTimeUnit(uops.TimeUnit),
		tpopts.WithValidator(validator),
	)

	definedType.FullName = b.parseTypeName(t)
	definedType.UnderlyingTypeName = "time.Time"
	definedType.Tops = tops
	definedType.Gops = b.gops
	return
}

func (b TypeBuilder) BuildAnonType(name typename.FullName, tops tpopts.Options) (
	anonType spec.AnonType, ok bool, err error) {
	op := NewMakeAnonOp(b.converter)

	if err = scnr.Scan(scnr.Config{}, name, op, tops); err != nil {
		if err != ErrStop {
			return
		}
		err = nil
	}
	anonType, ok = op.AnonType()
	return
}

func (b TypeBuilder) CollectAnonTypes(name typename.FullName,
	m map[spec.AnonSerName]spec.AnonType,
	tops tpopts.Options,
) (err error) {
	op := NewCollectAnonOp(m, b.converter)
	return scanner.Scan(scnr.Config{}, name, op, tops)
}

func (b TypeBuilder) checkFields(t reflect.Type, sops structops.Options) (
	err error,
) {
	if len(sops.Fields) == 0 {
		return
	}
	var (
		want   = t.NumField()
		actual = len(sops.Fields)
	)
	if actual != want {
		err = NewWrongFieldsCountError(want)
	}
	return
}

func (b TypeBuilder) makeFields(t reflect.Type, fieldNames []typename.FullName) []spec.FieldType {
	fields := make([]spec.FieldType, len(fieldNames))
	for i := range fieldNames {
		fields[i] = spec.FieldType{
			FullName:  fieldNames[i],
			FieldName: t.Field(i).Name,
			Gops:      b.gops,
		}
	}
	return fields
}

func (b TypeBuilder) parseTypeName(t reflect.Type) (typeName typename.FullName) {
	cname, err := typename.TypeCompleteName(t)
	if err != nil {
		return
	}
	return b.converter.ConvertToFullName(cname)
}

func (b TypeBuilder) parseUnderlyingTypeName(t reflect.Type) (typeName typename.FullName) {
	cname, err := typename.BaseTypeCompleteName(t)
	if err != nil {
		return
	}
	return b.converter.ConvertToFullName(cname)
}

func (b TypeBuilder) parseFieldNames(t reflect.Type) []typename.FullName {
	fieldNames := make([]typename.FullName, t.NumField())
	for i := range t.NumField() {
		fieldNames[i] = b.parseTypeName(t.Field(i).Type)
	}
	return fieldNames
}

func (b TypeBuilder) parseImpls(iops intropts.Options) []typename.FullName {
	impls := make([]typename.FullName, len(iops.Impls))
	for i, impl := range iops.Impls {
		impls[i] = b.parseTypeName(impl)
	}
	return impls
}

func (b TypeBuilder) parseVersions(vops veropts.Options) []typename.FullName {
	versions := make([]typename.FullName, len(vops.Versions))
	for i, version := range vops.Versions {
		versions[i] = b.parseTypeName(version.Type)
	}
	return versions
}

func (b TypeBuilder) notDefinedTypeError(t reflect.Type) error {
	switch {
	case classifier.DefinedStruct(t):
		return NewUnexpectedStructTypeError(t)
	case classifier.DefinedNonEmptyInterface(t):
		return NewUnexpectedInterfaceTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}

func (b TypeBuilder) notStructError(t reflect.Type) error {
	switch {
	case classifier.DefinedBasicType(t):
		return NewUnexpectedDefinedTypeError(t)
	case classifier.DefinedNonEmptyInterface(t):
		return NewUnexpectedInterfaceTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}

func (b TypeBuilder) notInterfaceError(t reflect.Type) error {
	switch {
	case classifier.DefinedBasicType(t):
		return NewUnexpectedDefinedTypeError(t)
	case classifier.DefinedStruct(t):
		return NewUnexpectedStructTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}
