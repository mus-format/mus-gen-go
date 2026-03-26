package converter

import (
	"testing"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/spec/converter"
	"github.com/mus-format/mus-gen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
)

type ConvertFullTestCase struct {
	Setup  ConvertFullSetup
	Params ConvertFullParams
	Want   ConvertFullWant
}

type ConvertFullSetup struct {
	Gops genopts.Options
}

type ConvertFullParams struct {
	CName typename.CompleteName
}

type ConvertFullWant struct {
	FName typename.FullName
}

func RunConvertFullTest(t *testing.T, tc ConvertFullTestCase) {
	c := converter.NewTypeNameConverter(tc.Setup.Gops)
	fname := c.ConvertToFullName(tc.Params.CName)
	asserterror.EqualDeep(t, fname, tc.Want.FName)
}
