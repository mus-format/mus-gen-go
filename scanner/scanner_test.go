package scanner_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/scanner/testutil/mock"
	generic_testdata "github.com/mus-format/musgen-go/testutil/generic"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

// TODO Add tests for any.

func TestScan(t *testing.T) {

	t.Run("struct", func(t *testing.T) {
		var (
			stOps = &typeops.Options{}

			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath: "github.com/mus-format/musgen-go/testutil/struct",
							Package: "testutil",
							Name:    "MyStruct",
							Kind:    scanner.Defined,
						}
						wantTops = stOps
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})

			tp  = reflect.TypeFor[struct_testdata.MyStruct]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), op, stOps)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("ptr struct", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath: "github.com/mus-format/musgen-go/testutil/struct",
							Stars:   "*",
							Package: "testutil",
							Name:    "MyStruct",
							Kind:    scanner.Defined,
						}
						wantTops *typeops.Options = nil
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})

			tp  = reflect.TypeFor[*struct_testdata.MyStruct]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("[]uint8", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath:  "",
							Stars:    "",
							Package:  "",
							Name:     "[]uint8",
							Kind:     scanner.Slice,
							ElemType: "uint8",
						}
						wantTops *typeops.Options = nil
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath:  "",
							Stars:    "",
							Package:  "",
							Name:     "uint8",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops *typeops.Options = nil
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})

			tp  = reflect.TypeFor[[]uint8]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("parametrized struct", func(t *testing.T) {
		// type MyStruct[T, V, N any] struct{}

		var (
			stOps = &typeops.Options{
				Key:  &typeops.Options{},
				Elem: &typeops.Options{},
			}

			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath: "github.com/mus-format/musgen-go/testutil/generic",
							Package: "testutil",
							Name:    "MyTripleParamStruct",
							Params:  []typename.CompleteName{"int", "string", "uint"},
							Kind:    scanner.Defined,
						}
						wantTops = stOps
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				//nothing to do
			}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Name:     "int",
						Kind:     scanner.Prim,
						Position: scanner.Param,
					}
					wantTops *typeops.Options = nil
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			}).RegisterProcessComma(func() {
				// nothing to do
			}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Name:     "string",
						Kind:     scanner.Prim,
						Position: scanner.Param,
					}
					wantTops *typeops.Options = nil
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			}).RegisterProcessComma(func() {
				// nothing to do
			}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Name:     "uint",
						Kind:     scanner.Prim,
						Position: scanner.Param,
					}
					wantTops *typeops.Options = nil
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			}).RegisterProcessRightSquare(func() {
				// nothing to do
			})
			tp  = reflect.TypeFor[generic_testdata.MyTripleParamStruct[int, string, uint]]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, stOps)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("array", func(t *testing.T) {
		var (
			elemTops = &typeops.Options{}
			arrTops  = &typeops.Options{
				Elem: elemTops,
			}
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:      "[3]int",
							Kind:      scanner.Array,
							ArrLength: "3",
							ElemType:  "int",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Name:     "int",
						Kind:     scanner.Prim,
						Position: scanner.Elem,
					}
					wantTops = tops
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			})
			tp  = reflect.TypeFor[[3]int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, arrTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("ptr array", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:     "*",
							Name:      "[3]*int",
							Kind:      scanner.Array,
							ArrLength: "3",
							ElemType:  "*int",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Stars:    "*",
						Name:     "int",
						Kind:     scanner.Prim,
						Position: scanner.Elem,
					}
					wantTops = tops
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			})

			tp  = reflect.TypeFor[*[3]*int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("slice", func(t *testing.T) {
		var (
			elemTops = &typeops.Options{}
			slTops   = &typeops.Options{
				Elem: elemTops,
			}
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "[]int",
							Kind:     scanner.Slice,
							ElemType: "int",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "int",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops = elemTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})
			tp  = reflect.TypeFor[[]int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, slTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("ptr slice", func(t *testing.T) {
		var (
			elemTops = &typeops.Options{}
			slTops   = &typeops.Options{
				Elem: elemTops,
			}
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "*",
							Name:     "[]*int",
							Kind:     scanner.Slice,
							ElemType: "*int",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "*",
							Name:     "int",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops = elemTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})
			tp  = reflect.TypeFor[*[]*int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, slTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("map", func(t *testing.T) {
		var (
			keyTops   = &typeops.Options{}
			valueTops = &typeops.Options{}
			mapTops   = &typeops.Options{
				Key:  keyTops,
				Elem: valueTops,
			}
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "map[int]string",
							Kind:     scanner.Map,
							KeyType:  "int",
							ElemType: "string",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "int",
							Kind:     scanner.Prim,
							Position: scanner.Key,
						}
						wantTops = keyTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "string",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops = valueTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})

			tp  = reflect.TypeFor[map[int]string]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, mapTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("ptr map", func(t *testing.T) {
		var (
			keyTops   = &typeops.Options{}
			valueTops = &typeops.Options{}
			mapTops   = &typeops.Options{
				Key:  keyTops,
				Elem: valueTops,
			}
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "*",
							Name:     "map[*int]*string",
							Kind:     scanner.Map,
							KeyType:  "*int",
							ElemType: "*string",
						}
						wantTops = tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "*",
							Name:     "int",
							Kind:     scanner.Prim,
							Position: scanner.Key,
						}
						wantTops = keyTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "*",
							Name:     "string",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops = valueTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})

			tp  = reflect.TypeFor[*map[*int]*string]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, mapTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("map in map", func(t *testing.T) {
		var (
			key2Tops   = &typeops.Options{}
			value2Tops = &typeops.Options{}

			keyTops   = &typeops.Options{}
			valueTops = &typeops.Options{
				Key:  key2Tops,
				Elem: value2Tops,
			}
			mapTops = &typeops.Options{
				Key:  keyTops,
				Elem: valueTops,
			}

			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "",
							Name:     "map[int]map[uint]string",
							Kind:     scanner.Map,
							KeyType:  "int",
							ElemType: "map[uint]string",
						}
						wantTops = mapTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (
					err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "",
							Name:     "int",
							Kind:     scanner.Prim,
							Position: scanner.Key,
						}
						wantTops = keyTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (
					err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "",
							Name:     "map[uint]string",
							KeyType:  "uint",
							ElemType: "string",
							Kind:     scanner.Map,
							Position: scanner.Elem,
						}
						wantTops = valueTops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (
					err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "",
							Name:     "uint",
							Kind:     scanner.Prim,
							Position: scanner.Key,
						}
						wantTops = key2Tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (
					err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars:    "",
							Name:     "string",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops = value2Tops
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})
			tp  = reflect.TypeFor[map[int]map[uint]string]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, mapTops)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}

	})

	t.Run("primitive", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name: "int",
							Kind: scanner.Prim,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})
			tp  = reflect.TypeFor[int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("ptr primitive", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Stars: "*",
							Name:  "int",
							Kind:  scanner.Prim,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				})
			tp  = reflect.TypeFor[*int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("WithoutParams", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "map[github.com/mus-format/musgen-go/testutil/generic/testutil.MyArray[int]]int",
							KeyType:  "github.com/mus-format/musgen-go/testutil/generic/testutil.MyArray[int]",
							ElemType: "int",
							Kind:     scanner.Map,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						PkgPath:  "github.com/mus-format/musgen-go/testutil/generic",
						Package:  "testutil",
						Name:     "MyArray",
						Params:   []typename.CompleteName{"int"},
						Kind:     scanner.Defined,
						Position: scanner.Key,
					}
					wantTops *typeops.Options
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(func(tp scanner.Type[typename.CompleteName],
				tops *typeops.Options) (err error) {
				var (
					wantType = scanner.Type[typename.CompleteName]{
						Stars:    "",
						Name:     "int",
						Kind:     scanner.Prim,
						Position: scanner.Elem,
					}
					wantTops *typeops.Options
				)
				asserterror.EqualDeep(t, tp, wantType)
				asserterror.Equal(t, tops, wantTops)
				return
			})
			tp  = reflect.TypeFor[map[generic_testdata.MyArray[int]]int]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil,
				scanner.WithoutParams())

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("complex", func(t *testing.T) {
		var (
			op = mock.NewOp[typename.CompleteName]().RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName],
					tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath: "github.com/mus-format/musgen-go/testutil/generic",
							Package: "testutil",
							Name:    "MyDoubleParamStruct",
							Params: []typename.CompleteName{
								"map[github.com/mus-format/musgen-go/testutil/generic.MyInt][3]math/big.Int",
								"github.com/mus-format/musgen-go/testutil/generic.MySlice[[]string]",
							},
							Kind: scanner.Defined,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "map[github.com/mus-format/musgen-go/testutil/generic.MyInt][3]math/big.Int",
							KeyType:  "github.com/mus-format/musgen-go/testutil/generic.MyInt",
							ElemType: "[3]math/big.Int",
							Kind:     scanner.Map,
							Position: scanner.Param,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath:  "github.com/mus-format/musgen-go/testutil",
							Package:  "generic",
							Name:     "MyInt",
							Kind:     scanner.Defined,
							Position: scanner.Key,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:      "[3]math/big.Int",
							ArrLength: "3",
							ElemType:  "math/big.Int",
							Kind:      scanner.Array,
							Position:  scanner.Elem,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath:  "math",
							Package:  "big",
							Name:     "Int",
							Kind:     scanner.Defined,
							Position: scanner.Elem,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessComma(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							PkgPath:  "github.com/mus-format/musgen-go/testutil/generic",
							Package:  "generic",
							Name:     "MySlice",
							Params:   []typename.CompleteName{"[]string"},
							Kind:     scanner.Defined,
							Position: scanner.Param,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessLeftSquare(func() {
				// nothing to do
			}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "[]string",
							ElemType: "string",
							Kind:     scanner.Slice,
							Position: scanner.Param,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessType(
				func(tp scanner.Type[typename.CompleteName], tops *typeops.Options) (err error) {
					var (
						wantType = scanner.Type[typename.CompleteName]{
							Name:     "string",
							Kind:     scanner.Prim,
							Position: scanner.Elem,
						}
						wantTops *typeops.Options
					)
					asserterror.EqualDeep(t, tp, wantType)
					asserterror.Equal(t, tops, wantTops)
					return
				}).RegisterProcessRightSquare(func() {
				// nothing to do
			}).RegisterProcessRightSquare(func() {
				// nothing to do
			})
			tp  = reflect.TypeFor[generic_testdata.MyDoubleParamStruct[map[generic_testdata.MyInt][3]big.Int, generic_testdata.MySlice[[]string]]]()
			err = scanner.Scan(typename.MustTypeCompleteName(tp), &op, nil)

			mocks = []*mok.Mock{op.Mock}
		)
		asserterror.EqualError(t, err, nil)
		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

}

type MyOp[T scanner.QualifiedName] struct {
	sl []scanner.Type[T]
}

func (o *MyOp[T]) Type(t scanner.Type[T], tops *typeops.Options) (err error) {
	o.sl = append(o.sl, t)
	fmt.Println(t)
	return
}

func (o *MyOp[T]) LeftSquare()  {}
func (o *MyOp[T]) Comma()       {}
func (o *MyOp[T]) RightSquare() {}
