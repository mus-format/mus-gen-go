// Package musgen provides a code generator for the MUS binary serialization
// format.
package musgen

import (
	"bytes"
	"embed"
	"go/parser"
	"go/token"
	"io/fs"
	"reflect"
	"text/template"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	cnvtr "github.com/mus-format/mus-gen-go/spec/converter"
	"github.com/mus-format/mus-gen-go/typename"
	"golang.org/x/tools/imports"
)

// Generator is responsible for generating MUS serialization code.
type Generator struct {
	baseTmpl      *template.Template
	typeBuilder   bldr.TypeBuilder
	crossgenTypes map[typename.FullName]struct{}
	anonMap       map[spec.AnonSerName]spec.AnonType
	genSl         []fileData
	dtmTypes      [][]string
	bs            []byte
	gops          genops.Options
}

// NewGenerator creates a new Generator.
func NewGenerator(opts ...genops.SetOption) (g *Generator, err error) {
	gops := genops.New()
	if err = genops.Apply(opts, &gops); err != nil {
		return
	}
	var (
		crossgenTypes = map[typename.FullName]struct{}{}
		converter     = cnvtr.NewTypeNameConverter(gops)
		typeBuilder   = bldr.NewTypeBuilder(converter, gops)
	)
	baseTmpl := template.New("base")
	registerFns(typeBuilder, converter, crossgenTypes, baseTmpl, gops)
	err = registerTemplates(baseTmpl)
	if err != nil {
		panic(err)
	}
	g = &Generator{
		baseTmpl:      baseTmpl,
		typeBuilder:   typeBuilder,
		crossgenTypes: crossgenTypes,
		anonMap:       map[spec.AnonSerName]spec.AnonType{},
		genSl:         []fileData{},
		bs:            []byte{},
		gops:          gops,
	}
	return
}

// AddDefinedType adds the specified type to the Generator to produce a
// serializer for it. This method supports types defined with the following
// source types: number, string, array, slice, map, pointer.
func (g *Generator) AddDefinedType(t reflect.Type, opts ...tpopts.SetOption) (
	err error,
) {
	tops := tpopts.Options{}
	tpopts.Apply(&tops, opts...)
	definedType, err := g.typeBuilder.BuildDefinedType(t, tops)
	if err != nil {
		return
	}
	g.fillCrossgen(t, definedType.FullName)
	g.typeBuilder.CollectAnonTypes(definedType.UnderlyingTypeName, g.anonMap, tops)
	g.genSl = append(g.genSl, fileData{definedTypeSerTmpl, definedType})
	return
}

// AddStruct adds the specified type to the Generator to produce a
// serializer for it. This method supports types definined with the struct
// source type.
func (g *Generator) AddStruct(t reflect.Type, opts ...stopts.SetOption) (
	err error,
) {
	sops := stopts.Options{}
	stopts.Apply(opts, &sops)
	if sops.UnderlyingTime != nil {
		return g.addStructUnderlyingTime(t, sops)
	}
	structType, err := g.typeBuilder.BuildStructType(t, sops)
	if err != nil {
		return
	}
	for i, fd := range structType.SerializedFields() {
		var tops tpopts.Options
		if len(sops.Fields) > 0 {
			tops = sops.Fields[i].Type
		}
		g.typeBuilder.CollectAnonTypes(fd.FullName, g.anonMap, tops)
	}
	g.fillCrossgen(t, structType.FullName)
	g.genSl = append(g.genSl, fileData{structSerTmpl, structType})
	return
}

// AddTyped adds the specified type to the Generator to produce a typed
// definition for it. This method supports all types acceptable by the
// AddDefinedType, AddStruct, and AddInterface methods.
func (g *Generator) AddTyped(t reflect.Type, opts ...tpopts.SetOption) (
	err error,
) {
	tops := tpopts.Options{}
	tpopts.Apply(&tops, opts...)
	tp, err := g.typeBuilder.BuildTyped(t)
	if err != nil {
		return
	}
	g.genSl = append(g.genSl, fileData{typedSerTmpl, tp})
	return
}

// AddInterface adds the specified type to the Generator to produce a
// serializer for it. This method supports types definined with the interface
// source type.
func (g *Generator) AddInterface(t reflect.Type, opts ...intropts.SetOption) (
	err error,
) {
	iops := intropts.Options{}
	intropts.Apply(&iops, opts...)
	interfaceType, err := g.typeBuilder.BuildInterfaceType(t, iops)
	if err != nil {
		return
	}
	g.fillCrossgen(t, interfaceType.FullName)
	g.genSl = append(g.genSl, fileData{interfaceSerTmpl, interfaceType})
	return
}

// RegisterInterface registers an interface type and all of its implementations
// with the code generator.
//
// DTM values are generated automatically, so there is no need to assign them
// manually.
//
// This helper method is equivalent to calling, in order:
//
//	AddStruct/AddDefinedType → AddTyped → AddInterface
func (g *Generator) RegisterInterface(t reflect.Type,
	opts ...intropts.SetRegisterOption,
) (err error) {
	rops := intropts.RegisterOptions{}
	if opts != nil {
		intropts.ApplyRegister(&rops, opts...)
	}
	var (
		l     = len(rops.StructImpls) + len(rops.DefinedTypeImpls)
		iops  = make([]intropts.SetOption, 0, l)
		types = make([]string, 0, l)
	)
	for _, impl := range rops.StructImpls {
		err = g.AddStruct(impl.Type, impl.Opts...)
		if err != nil {
			return
		}
		err = g.AddTyped(impl.Type)
		if err != nil {
			return
		}
		iops = append(iops, intropts.WithImpl(impl.Type))
		types = append(types, impl.Type.Name())
	}
	for _, impl := range rops.DefinedTypeImpls {
		err = g.AddDefinedType(impl.Type, impl.Opts...)
		if err != nil {
			return
		}
		err = g.AddTyped(impl.Type)
		if err != nil {
			return
		}
		iops = append(iops, intropts.WithImpl(impl.Type))
		types = append(types, impl.Type.Name())
	}
	g.addDTM(types)
	return g.AddInterface(t, iops...)
}

func (g *Generator) addDTM(types []string) {
	g.dtmTypes = append(g.dtmTypes, types)
}

// Generate produces the serialization code.
//
// The output is intended to be saved to a file.
func (g *Generator) Generate() (bs []byte, err error) {
	tmp := g.generatePackage()
	tmp = append(tmp, g.generateImports()...)
	tmp = append(tmp, g.generateDTMs()...)
	tmp = append(tmp, g.generateAnonymosDefinitions()...)
	tmp = append(tmp, g.generateSerializers()...)
	bs, err = imports.Process("", tmp, nil)
	if err != nil {
		err = ErrCodeGenFailed
		return tmp, err
	}
	err = g.checkSyntax(bs)
	if err != nil {
		err = ErrCodeGenFailed
		return tmp, err
	}
	return
}

func (g *Generator) addStructUnderlyingTime(t reflect.Type, sops stopts.Options) (
	err error,
) {
	uops := stopts.UnderlyingTimeOptions{}
	stopts.UnderlyingTimeApply(sops.UnderlyingTime, &uops)
	definedType, err := g.typeBuilder.BuildTimeType(t, uops, sops.Validator)
	if err != nil {
		return
	}
	g.fillCrossgen(t, definedType.FullName)
	g.genSl = append(g.genSl, fileData{definedTypeSerTmpl, definedType})
	return
}

func (g *Generator) generatePackage() (bs []byte) {
	return g.generatePart(packageTmpl, g.gops)
}

func (g *Generator) generateImports() (bs []byte) {
	return g.generatePart(importsTmpl, g.gops)
}

func (g *Generator) generateAnonymosDefinitions() (bs []byte) {
	return g.generatePart(anonDefinitionsTmpl, struct {
		Map map[spec.AnonSerName]spec.AnonType
		Ops genops.Options
	}{
		Map: g.anonMap,
		Ops: g.gops,
	})
}

func (g *Generator) generateSerializers() (bs []byte) {
	for _, item := range g.genSl {
		bs = append(bs, g.generatePart(item.fileName, item.typeData)...)
	}
	return
}

func (g *Generator) generatePart(tmplName string, a any) (bs []byte) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := g.baseTmpl.ExecuteTemplate(buf, tmplName, a)
	if err != nil {
		panic(err)
	}
	bs = buf.Bytes()
	return
}

func (g *Generator) generateDTMs() (bs []byte) {
	buf := bytes.Buffer{}
	for _, types := range g.dtmTypes {
		err := g.baseTmpl.ExecuteTemplate(&buf, dtmsDefinitionTmpl, types)
		if err != nil {
			panic(err)
		}
	}
	bs = buf.Bytes()
	return
}

func (g *Generator) fillCrossgen(t reflect.Type, fullName typename.FullName) {
	if g.crossGeneration(t) {
		g.crossgenTypes[fullName] = struct{}{}
	}
}

func (g *Generator) crossGeneration(t reflect.Type) bool {
	return t.PkgPath() != string(g.gops.PkgPath)
}

func (g *Generator) checkSyntax(bs []byte) (err error) {
	var (
		fs  = token.NewFileSet()
		src = string(bs)
	)
	_, err = parser.ParseFile(fs, "", src, parser.AllErrors)
	return
}

type fileData struct {
	fileName string
	typeData any
}

//go:embed templates/*.tmpl
var templatesFS embed.FS

func registerFns(typeBuilder bldr.TypeBuilder, converter cnvtr.TypeNameConverter,
	crossgenTypes map[typename.FullName]struct{},
	tmpl *template.Template,
	gopts genops.Options,
) {
	var (
		tmplFns = TmplFns{tmpl: tmpl}
		serFns  = NewSerFns(converter, typeBuilder, crossgenTypes, tmplFns, gopts)
	)
	m := map[string]any{}
	m["Tmpl"] = func() TmplFns {
		return tmplFns
	}
	m["Ser"] = func() SerFns {
		return serFns
	}
	tmpl.Funcs(m)
}

func registerTemplates(tmpl *template.Template) (err error) {
	subFS, _ := fs.Sub(templatesFS, "templates")
	_, err = tmpl.ParseFS(subFS, "*.tmpl")
	return
}
