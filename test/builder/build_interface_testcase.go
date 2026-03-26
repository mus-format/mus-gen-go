package builder

import (
	"reflect"
	"testing"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	intropts "github.com/mus-format/mus-gen-go/options/interface"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

type BuildInterfaceTypeTestCase struct {
	Name   string
	Setup  BuildInterfaceTypeSetup
	Params BuildInterfaceTypeParams
	Want   BuildInterfaceTypeWant
}

type BuildInterfaceTypeSetup struct {
	Converter bldr.TypeNameConverter
	Gops      genops.Options
}

type BuildInterfaceTypeParams struct {
	Type reflect.Type
	Iops intropts.Options
}

type BuildInterfaceTypeWant struct {
	InterfaceType spec.InterfaceType
	Err           error
	Mocks         []*mok.Mock
}

func RunBuildInterfaceTypeTest(t *testing.T, tc BuildInterfaceTypeTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		builder := bldr.NewTypeBuilder(tc.Setup.Converter, tc.Setup.Gops)
		i, err := builder.BuildInterfaceType(tc.Params.Type, tc.Params.Iops)
		asserterror.EqualError(t, err, tc.Want.Err)
		asserterror.EqualDeep(t, i, tc.Want.InterfaceType)
		asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
	})
}
