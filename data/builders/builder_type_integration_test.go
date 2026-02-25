package builders

import (
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	prim_testdata "github.com/mus-format/musgen-go/testutil/primitive"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTypeDataBuilderIntegration(t *testing.T) {

	t.Run("BuildDefinedTypeData", func(t *testing.T) {
		var (
			gops = genops.New()
			b    = NewTypeDataBuilder(NewConverter(gops), gops)

			wantTypeData = data.TypeData{
				FullName:       "testutil.MyInt",
				SourceFullName: "int",

				Fields: []data.FieldData{
					{FullName: "int"},
				},
				Gops: gops,
			}
			wantErr error = nil
		)
		d, err := b.BuildDefinedTypeData(reflect.TypeFor[prim_testdata.MyInt](), nil)
		asserterror.EqualError(t, err, wantErr)
		asserterror.EqualDeep(t, d, wantTypeData)
	})

}
