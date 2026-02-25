package data

import (
	"testing"

	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTypeData(t *testing.T) {

	t.Run("SerializedFiedlds", func(t *testing.T) {

		t.Run("Should work", func(t *testing.T) {
			var (
				wantFirst  = FieldData{FieldName: "first"}
				wantSecond = FieldData{FieldName: "second"}
				wantThird  = FieldData{FieldName: "third"}
				wantLen    = 2
				sops       = structops.New()
			)
			structops.Apply([]structops.SetOption{
				structops.WithField(),
				structops.WithField(typeops.WithIgnore()),
				structops.WithField(),
			}, &sops)

			data := TypeData{
				Fields: []FieldData{wantFirst, wantSecond, wantThird},
				Sops:   sops,
			}

			asserterror.Equal(t, len(data.SerializedFields()), wantLen)
			asserterror.EqualDeep(t, data.SerializedFields(),
				[]FieldData{wantFirst, wantThird})
		})

		t.Run("Should work", func(t *testing.T) {
			var (
				wantFirst  = FieldData{FieldName: "first"}
				wantSecond = FieldData{FieldName: "second"}
				wantThird  = FieldData{FieldName: "third"}
				wantLen    = 3
				sops       = structops.New()
			)
			data := TypeData{
				Fields: []FieldData{wantFirst, wantSecond, wantThird},
				Sops:   sops,
			}
			asserterror.Equal(t, len(data.SerializedFields()), wantLen)
			asserterror.EqualDeep(t, data.SerializedFields(),
				[]FieldData{wantFirst, wantSecond, wantThird})
		})

	})

}
