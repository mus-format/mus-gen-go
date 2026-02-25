package builders

import (
	"testing"

	"github.com/mus-format/musgen-go/data"
	"github.com/mus-format/musgen-go/data/builders/testutil/mock"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func TestAnonDataBuilder(t *testing.T) {

	t.Run("Make", func(t *testing.T) {
		var (
			tops     *typeops.Options = nil
			gops                      = genops.New()
			wantData                  = data.AnonData{
				AnonSerName: anonSerName(data.AnonMap, "map[[3]int]map[string]pkg.Type", tops, gops),
				Kind:        data.AnonMap,

				LenSer: "nil",
				LenVl:  "nil",

				KeyType: "[3]int",
				KeyVl:   "nil",

				ElemType: "map[string]pkg.Type",
				ElemVl:   "nil",
			}
			wantOk        = true
			wantErr error = nil
		)

		gops = genops.New()
		b := NewAnonDataBuilder(NewTypeDataBuilder(nil, gops), gops)
		d, ok, err := b.Build("map[[3]int]map[string]pkg.Type", tops)

		asserterror.EqualError(t, err, wantErr)
		asserterror.Equal(t, ok, wantOk)
		asserterror.EqualDeep(t, d, wantData)
	})

	t.Run("Collect", func(t *testing.T) {

		t.Run("string", func(t *testing.T) {
			var (
				tops = typeops.Options{}
				gops = genops.New()
			)
			typeops.Apply([]typeops.SetOption{
				typeops.WithLenEncoding(typeops.Raw),
				typeops.WithLenValidator("ValidateLength"),
			}, &tops)
			var (
				anonName = anonSerName(data.AnonString, "string", &tops, gops)
				wantM    = map[data.AnonSerName]data.AnonData{
					anonName: {
						AnonSerName: anonName,
						Kind:        data.AnonString,

						LenSer: "raw.Int",
						LenVl:  "com.ValidatorFn[int](ValidateLength)",

						Tops: &tops,
					},
				}
				m    = map[data.AnonSerName]data.AnonData{}
				conv = mock.NewTypeNameConvertor().RegisterConvertToRelName(
					func(fname typename.FullName) (rname typename.RelName, err error) {
						return "int", nil
					},
				)
				mocks = []*mok.Mock{conv.Mock}
			)

			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(conv, gops), gops)
			b.Collect("string", m, &tops)
			asserterror.EqualDeep(t, m, wantM)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("array", func(t *testing.T) {
			var (
				tops = typeops.Options{}
				gops = genops.New()
			)
			typeops.Apply([]typeops.SetOption{
				typeops.WithLenEncoding(typeops.Raw),
				typeops.WithElem(
					typeops.WithValidator("ValidateInt"),
				),
			}, &tops)
			var (
				anonName = anonSerName(data.AnonArray, "[3]int", &tops, gops)
				wantM    = map[data.AnonSerName]data.AnonData{
					anonName: {
						AnonSerName: anonName,
						Kind:        data.AnonArray,

						ArrType:   "[3]int",
						ArrLength: "3",

						LenSer: "raw.Int",

						ElemType: "int",
						ElemVl:   "com.ValidatorFn[int](ValidateInt)",

						Tops: &tops,
					},
				}
				m    = map[data.AnonSerName]data.AnonData{}
				conv = mock.NewTypeNameConvertor().RegisterConvertToRelName(
					func(fname typename.FullName) (rname typename.RelName, err error) {
						return "int", nil
					},
				)
				mocks = []*mok.Mock{conv.Mock}
			)
			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(conv, gops), gops)
			b.Collect("[3]int", m, &tops)
			asserterror.EqualDeep(t, m, wantM)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("slice", func(t *testing.T) {
			var (
				tops = typeops.Options{}
				gops = genops.New()
			)
			typeops.Apply([]typeops.SetOption{
				typeops.WithLenEncoding(typeops.Raw),
				typeops.WithLenValidator("ValidateLength"),
				typeops.WithElem(
					typeops.WithValidator("ValidateInt"),
				),
			}, &tops)
			var (
				anonName = anonSerName(data.AnonSlice, "[]int", &tops, gops)
				wantM    = map[data.AnonSerName]data.AnonData{
					anonName: {
						AnonSerName: anonName,
						Kind:        data.AnonSlice,

						LenSer: "raw.Int",
						LenVl:  "com.ValidatorFn[int](ValidateLength)",

						ElemType: "int",
						ElemVl:   "com.ValidatorFn[int](ValidateInt)",

						Tops: &tops,
					},
				}
				m    = map[data.AnonSerName]data.AnonData{}
				conv = mock.NewTypeNameConvertor().RegisterConvertToRelName(
					func(fname typename.FullName) (rname typename.RelName, err error) {
						return "int", nil
					},
				).RegisterConvertToRelName(
					func(fname typename.FullName) (rname typename.RelName, err error) {
						return "int", nil
					},
				)
				mocks = []*mok.Mock{conv.Mock}
			)
			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(conv, gops), gops)
			b.Collect("[]int", m, &tops)
			asserterror.EqualDeep(t, m, wantM)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("map", func(t *testing.T) {
			var (
				tops      *typeops.Options = nil
				gops                       = genops.New()
				anonName1                  = anonSerName(data.AnonMap, "map[[3]int]map[string]pkg.Type", tops, gops)
				anonName2                  = anonSerName(data.AnonArray, "[3]int", tops, gops)
				anonName3                  = anonSerName(data.AnonMap, "map[string]pkg.Type", tops, gops)
				wantM                      = map[data.AnonSerName]data.AnonData{
					anonName1: {
						AnonSerName: anonName1,
						Kind:        data.AnonMap,

						LenSer: "nil",
						LenVl:  "nil",

						KeyType: "[3]int",
						KeyVl:   "nil",

						ElemType: "map[string]pkg.Type",
						ElemVl:   "nil",
					},
					anonName2: {
						AnonSerName: anonName2,
						Kind:        data.AnonArray,

						ArrType:   "[3]int",
						ArrLength: "3",

						LenSer: "nil",

						ElemType: "int",
						ElemVl:   "nil",
					},
					anonName3: {
						AnonSerName: anonName3,
						Kind:        data.AnonMap,

						LenSer: "nil",
						LenVl:  "nil",

						KeyType: "string",
						KeyVl:   "nil",

						ElemType: "pkg.Type",
						ElemVl:   "nil",
					},
				}
				m = map[data.AnonSerName]data.AnonData{}
			)
			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(nil, gops), gops)
			b.Collect("map[[3]int]map[string]pkg.Type", m, tops)
			asserterror.EqualDeep(t, m, wantM)
		})

		t.Run("pointer", func(t *testing.T) {
			var (
				tops     = typeops.Options{}
				gops     = genops.New()
				anonName = anonSerName(data.AnonPtr, "*int", &tops, gops)
				wantM    = map[data.AnonSerName]data.AnonData{
					anonName: {
						AnonSerName: anonName,
						Kind:        data.AnonPtr,

						ElemType: "int",

						Tops: &tops,
					},
				}
				m = map[data.AnonSerName]data.AnonData{}
			)
			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(nil, gops), gops)
			b.Collect("*int", m, &tops)
			asserterror.EqualDeep(t, m, wantM)
		})

		t.Run("pointer of defined type", func(t *testing.T) {
			var (
				tops     = typeops.Options{}
				gops     = genops.New()
				anonName = anonSerName(data.AnonPtr, "*pkg.Type", &tops, gops)
				wantM    = map[data.AnonSerName]data.AnonData{
					anonName: {
						AnonSerName: anonName,
						Kind:        data.AnonPtr,

						ElemType: "pkg.Type",

						Tops: &tops,
					},
				}
				m = map[data.AnonSerName]data.AnonData{}
			)
			gops = genops.New()
			b := NewAnonDataBuilder(NewTypeDataBuilder(nil, gops), gops)
			b.Collect("*pkg.Type", m, &tops)
			asserterror.EqualDeep(t, m, wantM)
		})

		t.Run("double pointer", func(t *testing.T) {
			tops := typeops.Options{}
			typeops.Apply([]typeops.SetOption{
				typeops.WithElem(nil),
			}, &tops)
			var (
				gops      = genops.New()
				anonName1 = anonSerName(data.AnonPtr, "**int", &tops, gops)
				anonName2 = anonSerName(data.AnonPtr, "*int", tops.Elem, gops)
				wantM     = map[data.AnonSerName]data.AnonData{
					anonName1: {
						AnonSerName: anonName1,
						Kind:        data.AnonPtr,

						ElemType: "*int",

						Tops: &tops,
					},
					anonName2: {
						AnonSerName: anonName2,
						Kind:        data.AnonPtr,

						ElemType: "int",

						Tops: tops.Elem,
					},
				}
				m = map[data.AnonSerName]data.AnonData{}
			)
			b := NewAnonDataBuilder(NewTypeDataBuilder(nil, gops), gops)
			b.Collect("**int", m, &tops)
			asserterror.EqualDeep(t, m, wantM)
		})

	})

}
