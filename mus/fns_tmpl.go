package musgen

import (
	"bytes"
	"regexp"
	"text/template"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	"github.com/mus-format/mus-gen-go/typename"
)

const (
	PtrPattern   = `(^\*)(.+$)`
	ArrayPattern = `[\d+](.+$)`
)

var (
	ptrRe   = regexp.MustCompile(PtrPattern)
	arrayRe = regexp.MustCompile(ArrayPattern)
)

type TmplFns struct {
	tmpl *template.Template
}

func (f TmplFns) PtrType(name typename.FullName) (ok bool) {
	return ptrRe.MatchString(string(name))
}

func (f TmplFns) ArrayType(name typename.FullName) (ok bool) {
	return arrayRe.MatchString(string(name))
}

func (f TmplFns) WithComma(str string) string {
	if str == "" {
		return str
	}
	return ", " + str
}

func (f TmplFns) Minus(a int, b int) int {
	return a - b
}

func (f TmplFns) ByteSliceStream(name typename.FullName,
	gopts genopts.Options) bool {
	return gopts.Stream && (name == "[]byte" || name == "[]uint8")
}

func (f TmplFns) TimeSer(topts tpopts.Options) string {
	return topts.TimeUnit.Ser()
}

func (f TmplFns) Include(name string, pipeline any) (str string, err error) {
	var buf bytes.Buffer
	if err = f.tmpl.ExecuteTemplate(&buf, name, pipeline); err != nil {
		return
	}
	str = buf.String()
	return
}

func (f TmplFns) MarshalSignatureLastParam(gops genopts.Options) string {
	if gops.Stream {
		return "w mus.Writer"
	}
	return "bs []byte"
}

func (f TmplFns) MarshalLastParam(first bool, gops genopts.Options) string {
	if gops.Stream {
		return "w"
	}
	if first {
		return "bs"
	}
	return "bs[n:]"
}

func (f TmplFns) UnmarshalSignatureLastParam(gops genopts.Options) string {
	if gops.Stream {
		return "r mus.Reader"
	}
	return "bs []byte"
}

func (f TmplFns) UnmarshalLastParam(first bool, gops genopts.Options) string {
	if gops.Stream {
		return "r"
	}
	if first {
		return "bs"
	}
	return "bs[n:]"
}

func (f TmplFns) SkipSignatureLastParam(gops genopts.Options) string {
	if gops.Stream {
		return "r mus.Reader"
	}
	return "bs []byte"
}

func (f TmplFns) SkipLastParam(gops genopts.Options) string {
	if gops.Stream {
		return "r"
	}
	return "bs[n:]"
}
func (f TmplFns) ModImportName(gops genopts.Options) string {
	if gops.Stream {
		return "mus"
	}
	return "mus"
}

type FieldTmplPipe struct {
	Val         string
	FieldsCount int
	Field       spec.FieldType
	Index       int
	Tops        tpopts.Options
	Gops        genopts.Options
}
