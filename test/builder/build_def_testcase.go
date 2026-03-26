package builder

import (
	"reflect"
	"testing"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

type BuildDefinedTypeTestCase struct {
	Name   string
	Setup  BuildDefinedTypeSetup
	Params BuildDefinedTypeParams
	Want   BuildDefinedTypeWant
}

type BuildDefinedTypeSetup struct {
	Converter bldr.TypeNameConverter
	Gops      genops.Options
}

type BuildDefinedTypeParams struct {
	Type reflect.Type
	Tops tpopts.Options
}

type BuildDefinedTypeWant struct {
	DefinedType spec.DefinedType
	Err         error
	Mocks       []*mok.Mock
}

func RunBuildDefinedTypeTest(t *testing.T, tc BuildDefinedTypeTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		builder := bldr.NewTypeBuilder(tc.Setup.Converter, tc.Setup.Gops)
		d, err := builder.BuildDefinedType(tc.Params.Type, tc.Params.Tops)
		asserterror.EqualError(t, err, tc.Want.Err)
		asserterror.EqualDeep(t, d, tc.Want.DefinedType)
		asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
	})
}
