package builder

import (
	"reflect"
	"testing"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	stopts "github.com/mus-format/mus-gen-go/options/struct"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

type BuildStructTypeTestCase struct {
	Name   string
	Setup  BuildStructTypeSetup
	Params BuildStructTypeParams
	Want   BuildStructTypeWant
}

type BuildStructTypeSetup struct {
	Converter bldr.TypeNameConverter
	Gops      genops.Options
}

type BuildStructTypeParams struct {
	Type reflect.Type
	Sops stopts.Options
}

type BuildStructTypeWant struct {
	StructType spec.StructType
	Err        error
	Mocks      []*mok.Mock
}

func RunBuildStructTypeTest(t *testing.T, tc BuildStructTypeTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		builder := bldr.NewTypeBuilder(tc.Setup.Converter, tc.Setup.Gops)
		s, err := builder.BuildStructType(tc.Params.Type, tc.Params.Sops)
		asserterror.EqualError(t, err, tc.Want.Err)
		asserterror.EqualDeep(t, s, tc.Want.StructType)
		asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
	})
}
