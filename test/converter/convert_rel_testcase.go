package converter

import (
	"testing"

	genopts "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/spec/converter"
	"github.com/mus-format/mus-gen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
)

type ConvertRelTestCase struct {
	Setup  ConvertRelSetup
	Params ConvertRelParams
	Want   ConvertRelWant
}

type ConvertRelSetup struct {
	Gops genopts.Options
}

type ConvertRelParams struct {
	FName typename.FullName
}

type ConvertRelWant struct {
	RName typename.RelativeName
}

func RunConvertRelTest(t *testing.T, tc ConvertRelTestCase) {
	c := converter.NewTypeNameConverter(tc.Setup.Gops)
	rname := c.ConvertToRelativeName(tc.Params.FName)
	asserterror.EqualDeep(t, rname, tc.Want.RName)
}
