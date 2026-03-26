package builder

import (
	"reflect"
	"testing"

	genops "github.com/mus-format/mus-gen-go/options/gen"
	"github.com/mus-format/mus-gen-go/spec"
	bldr "github.com/mus-format/mus-gen-go/spec/builder"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

type BuildTypedTestCase struct {
	Name   string
	Setup  BuildTypedSetup
	Params BuildTypedParams
	Want   BuildTypedWant
}

type BuildTypedSetup struct {
	Converter bldr.TypeNameConverter
	Gops      genops.Options
}

type BuildTypedParams struct {
	Type reflect.Type
}

type BuildTypedWant struct {
	Typed spec.Type
	Err   error
	Mocks []*mok.Mock
}

func RunBuildTypedTest(t *testing.T, tc BuildTypedTestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		builder := bldr.NewTypeBuilder(tc.Setup.Converter, tc.Setup.Gops)
		tp, err := builder.BuildTyped(tc.Params.Type)
		asserterror.EqualError(t, err, tc.Want.Err)
		asserterror.EqualDeep(t, tp, tc.Want.Typed)
		asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
	})
}
