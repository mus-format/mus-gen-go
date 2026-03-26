package scanner

import (
	"math/big"
	"reflect"
	"testing"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/test/mock"
	"github.com/mus-format/mus-gen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func ScanStructTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[Struct]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					PkgPath: "github.com/mus-format/mus-gen-go/test/scanner",
					Package: "scanner",
					Name:    "Struct",
					Kind:    scanner.KindDefined,
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "struct",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPtrStructTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[*Struct]()
		op    = mock.NewOp[typename.CompleteName]()
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					PkgPath: "github.com/mus-format/mus-gen-go/test/scanner",
					Stars:   "*",
					Package: "scanner",
					Name:    "Struct",
					Kind:    scanner.KindDefined,
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "ptr struct",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanUint8SliceTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		op    = mock.NewOp[typename.CompleteName]()
		tp    = reflect.TypeFor[[]uint8]()
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					PkgPath:  "",
					Stars:    "",
					Package:  "",
					Name:     "[]uint8",
					Kind:     scanner.KindSlice,
					ElemType: "uint8",
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					PkgPath:  "",
					Stars:    "",
					Package:  "",
					Name:     "uint8",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "[]uint8",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanParametrizedStructTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[TripleParamStruct[int, string, uint]]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					PkgPath: "github.com/mus-format/mus-gen-go/test/scanner",
					Package: "scanner",
					Name:    "TripleParamStruct",
					Params:  []typename.CompleteName{"int", "string", "uint"},
					Kind:    scanner.KindDefined,
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessLeftSquare(
		func() {
			//nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionParam,
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessComma(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "string",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionParam,
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessComma(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "uint",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionParam,
				}
				wantTops tpopts.Options = tpopts.Options{}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessRightSquare(
		func() {
			// nothing to do
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "parametrized struct",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanArrayTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[[3]int]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:      "[3]int",
					Kind:      scanner.KindArray,
					ArrLength: "3",
					ElemType:  "int",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "array",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPtrArrayTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[*[3]*int]()
		op    = mock.NewOp[typename.CompleteName]()
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:     "*",
					Name:      "[3]*int",
					Kind:      scanner.KindArray,
					ArrLength: "3",
					ElemType:  "*int",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "ptr array",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanSliceTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[[]int]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "[]int",
					Kind:     scanner.KindSlice,
					ElemType: "int",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "slice",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPtrSliceTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[*[]*int]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "[]*int",
					Kind:     scanner.KindSlice,
					ElemType: "*int",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "ptr slice",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanMapTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[map[int]string]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "map[int]string",
					Kind:     scanner.KindMap,
					KeyType:  "int",
					ElemType: "string",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessLeftSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionKey,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	).RegisterProcessRightSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "string",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "map",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPtrMapTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[*map[*int]*string]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "map[*int]*string",
					Kind:     scanner.KindMap,
					KeyType:  "*int",
					ElemType: "*string",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessLeftSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionKey,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	).RegisterProcessRightSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Stars:    "*",
					Name:     "string",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "ptr map",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanMapInMapTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		tp    = reflect.TypeFor[map[int]map[uint]string]()
		op    = mock.NewOp[typename.CompleteName]()
		tops  = tpopts.Options{}
		mocks = []*mok.Mock{op.Mock}
	)

	op.RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName],
			tops tpopts.Options) (err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "map[int]map[uint]string",
					Kind:     scanner.KindMap,
					KeyType:  "int",
					ElemType: "map[uint]string",
				}
				wantTops = tops
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, wantTops)
			return
		},
	).RegisterProcessLeftSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (
			err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "int",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionKey,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	).RegisterProcessRightSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (
			err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "map[uint]string",
					KeyType:  "uint",
					ElemType: "string",
					Kind:     scanner.KindMap,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	).RegisterProcessLeftSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (
			err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "uint",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionKey,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	).RegisterProcessRightSquare(
		func() {
			// nothing to do
		},
	).RegisterProcessType(
		func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (
			err error) {
			var (
				wantType = scanner.TypeInfo[typename.CompleteName]{
					Name:     "string",
					Kind:     scanner.KindPrim,
					Position: scanner.PositionElem,
				}
			)
			asserterror.EqualDeep(t, tp, wantType)
			asserterror.EqualDeep(t, tops, tpopts.Options{})
			return
		},
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "map in map",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tops,
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPrimitiveTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name: "int",
						Kind: scanner.KindPrim,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		)
		tp    = reflect.TypeFor[int]()
		mocks = []*mok.Mock{op.Mock}
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "primitive",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanPtrPrimitiveTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Stars: "*",
						Name:  "int",
						Kind:  scanner.KindPrim,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		)
		tp    = reflect.TypeFor[*int]()
		mocks = []*mok.Mock{op.Mock}
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "ptr primitive",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanWithoutParamsTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name:     "map[github.com/mus-format/mus-gen-go/test/scanner/scanner.Array[int]]int",
						KeyType:  "github.com/mus-format/mus-gen-go/test/scanner/scanner.Array[int]",
						ElemType: "int",
						Kind:     scanner.KindMap,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessLeftSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						PkgPath:  "github.com/mus-format/mus-gen-go/test/scanner",
						Package:  "scanner",
						Name:     "Array",
						Params:   []typename.CompleteName{"int"},
						Kind:     scanner.KindDefined,
						Position: scanner.PositionKey,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessRightSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Stars:    "",
						Name:     "int",
						Kind:     scanner.KindPrim,
						Position: scanner.PositionElem,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		)
		tp    = reflect.TypeFor[map[Array[int]]int]()
		mocks = []*mok.Mock{op.Mock}
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "WithoutParams",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{WithoutParams: true},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}

func ScanComplexTestCase(t *testing.T) ScanTestCase[typename.CompleteName] {
	var (
		op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName],
				tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						PkgPath: "github.com/mus-format/mus-gen-go/test/scanner",
						Package: "scanner",
						Name:    "DoubleParamStruct",
						Params: []typename.CompleteName{
							"map[github.com/mus-format/mus-gen-go/test/scanner.Int][3]math/big.Int",
							"github.com/mus-format/mus-gen-go/test/scanner.Slice[[]string]",
						},
						Kind: scanner.KindDefined,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessLeftSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name:     "map[github.com/mus-format/mus-gen-go/test/scanner.Int][3]math/big.Int",
						KeyType:  "github.com/mus-format/mus-gen-go/test/scanner.Int",
						ElemType: "[3]math/big.Int",
						Kind:     scanner.KindMap,
						Position: scanner.PositionParam,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessLeftSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						PkgPath:  "github.com/mus-format/mus-gen-go/test",
						Package:  "scanner",
						Name:     "Int",
						Kind:     scanner.KindDefined,
						Position: scanner.PositionKey,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessRightSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name:      "[3]math/big.Int",
						ArrLength: "3",
						ElemType:  "math/big.Int",
						Kind:      scanner.KindArray,
						Position:  scanner.PositionElem,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						PkgPath:  "math",
						Package:  "big",
						Name:     "Int",
						Kind:     scanner.KindDefined,
						Position: scanner.PositionElem,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessComma(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						PkgPath:  "github.com/mus-format/mus-gen-go/test/scanner",
						Package:  "scanner",
						Name:     "Slice",
						Params:   []typename.CompleteName{"[]string"},
						Kind:     scanner.KindDefined,
						Position: scanner.PositionParam,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessLeftSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name:     "[]string",
						ElemType: "string",
						Kind:     scanner.KindSlice,
						Position: scanner.PositionParam,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessType(
			func(tp scanner.TypeInfo[typename.CompleteName], tops tpopts.Options) (err error) {
				var (
					wantType = scanner.TypeInfo[typename.CompleteName]{
						Name:     "string",
						Kind:     scanner.KindPrim,
						Position: scanner.PositionElem,
					}
					wantTops tpopts.Options = tpopts.Options{}
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.EqualDeep(t, tops, wantTops)
				return
			},
		).RegisterProcessRightSquare(
			func() {
				// nothing to do
			},
		).RegisterProcessRightSquare(
			func() {
				// nothing to do
			},
		)
		tp = reflect.TypeFor[DoubleParamStruct[map[Int][3]big.Int,
			Slice[[]string]]]()
		mocks = []*mok.Mock{op.Mock}
	)

	return ScanTestCase[typename.CompleteName]{
		Name: "complex",
		Setup: ScanSetup[typename.CompleteName]{
			Config: scanner.Config{},
			Name:   typename.MustTypeCompleteName(tp),
			Op:     op,
			Tops:   tpopts.Options{},
		},
		Want: ScanWant{
			Mocks: mocks,
		},
	}
}
