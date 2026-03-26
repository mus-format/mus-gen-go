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

type BuildTimeTypeTestCase struct {
	Name   string
	Setup  BuildTimeTypeSetup
	Params BuildTimeTypeParams
	Want   BuildTimeTypeWant
}

type BuildTimeTypeSetup struct {
	Converter bldr.TypeNameConverter
	Gops      genops.Options
}

type BuildTimeTypeParams struct {
	Type      reflect.Type
	Uops      stopts.UnderlyingTimeOptions
	Validator string
}

type BuildTimeTypeWant struct {
	DefinedType spec.DefinedType
	Err         error
	Mocks       []*mok.Mock
}

func RunBuildTimeTypeTest(t *testing.T, tc BuildTimeTypeTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		builder := bldr.NewTypeBuilder(tc.Setup.Converter, tc.Setup.Gops)
		d, err := builder.BuildTimeType(tc.Params.Type, tc.Params.Uops,
			tc.Params.Validator)
		asserterror.EqualError(t, err, tc.Want.Err)
		asserterror.EqualDeep(t, d, tc.Want.DefinedType)
		asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
	})
}
