package musgen

import (
	"image"
	"reflect"
	"testing"
	"time"

	"github.com/mus-format/mus-gen-go/test/genopts/cross-package/pkg3"
	"github.com/mus-format/mus-gen-go/test/genopts/cross-package/pkg4"
	custompkg "github.com/mus-format/mus-gen-go/test/genopts/custom_pkg"
	customsername "github.com/mus-format/mus-gen-go/test/genopts/custom_sername"
	importalias "github.com/mus-format/mus-gen-go/test/genopts/import_alias"
	pkg1sub "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg1/sub"
	pkg2sub "github.com/mus-format/mus-gen-go/test/genopts/import_alias/pkg2/sub"
	"github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg1"
	"github.com/mus-format/mus-gen-go/test/genopts/multi-package/pkg2"
	"github.com/mus-format/mus-gen-go/test/types"
	"github.com/mus-format/mus-gen-go/test/types/mode"
	"github.com/mus-format/mus-gen-go/test/types/mode/notunsafe"
	"github.com/mus-format/mus-gen-go/test/types/mode/safe"
	"github.com/mus-format/mus-gen-go/test/types/mode/stream_notunsafe"
	"github.com/mus-format/mus-gen-go/test/types/mode/stream_safe"
	"github.com/mus-format/mus-gen-go/test/types/mode/stream_unsafe"
	"github.com/mus-format/mus-gen-go/test/types/mode/unsafe"
	"github.com/mus-format/mus-go/test"
	teststream "github.com/mus-format/mus-stream-go/test"
)

func TestGenerated_Defined(t *testing.T) {
	test.Test([]types.Int{1}, types.IntMUS, t)
	test.Test([]types.RawInt{1}, types.RawIntMUS, t)
	test.Test([]types.ValidInt{1}, types.ValidIntMUS, t)
}

func TestGenerated_Struct(t *testing.T) {
	test.Test([]types.SimpleStruct{{Num: 1, Str: "str"}}, types.SimpleStructMUS, t)
	test.Test([]types.TimeStruct{types.TimeStruct(time.Now().Truncate(time.Second).UTC())}, types.TimeStructMUS, t)
	test.Test([]types.InnerStruct{{Str: "str"}}, types.InnerStructMUS, t)
	test.Test([]types.EmbeddingStruct{{Num: 1, InnerStruct: types.InnerStruct{Str: "str"}}}, types.EmbeddingStructMUS, t)
	test.Test([]types.IgnoreStruct{{Num: 1}}, types.IgnoreStructMUS, t)
	test.Test([]types.ValidStruct{{Num: 1, Str: "str"}}, types.ValidStructMUS, t)
	test.Test([]types.ParametricStruct[int]{{Field: 1}}, types.ParametricStructMUS, t)
}

func TestGenerated_Anon(t *testing.T) {
	test.Test([]types.LenString{"str"}, types.LenStringMUS, t)
	test.Test([]types.LenSlice{[]int{1}}, types.LenSliceMUS, t)
	test.Test([]types.LenArray{[3]int{1, 2, 3}}, types.LenArrayMUS, t)
	test.Test([]types.LenMap{{"str": 1}}, types.LenMapMUS, t)
	test.Test([]types.Slice{[]int{1}}, types.SliceMUS, t)
	test.Test([]types.Array{[3]int{1, 2, 3}}, types.ArrayMUS, t)
	test.Test([]types.Map{{"str": 1}}, types.MapMUS, t)
	i := 1
	test.Test([]types.Ptr{&i}, types.PtrMUS, t)
	test.Test([]types.ValidString{"str"}, types.ValidStringMUS, t)
	test.Test([]types.ValidSlice{[]int{1}}, types.ValidSliceMUS, t)
	test.Test([]types.ValidArray{[3]int{1, 2, 3}}, types.ValidArrayMUS, t)
	test.Test([]types.ValidMap{{"str": 1}}, types.ValidMapMUS, t)
}

func TestGenerated_ComplexAnon(t *testing.T) {
	bs := make([]byte, types.ComplexMapMUS.Size(complexMap))
	types.ComplexMapMUS.Marshal(complexMap, bs)
	v, n, err := types.ComplexMapMUS.Unmarshal(bs)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(bs) {
		t.Errorf("unexpected n, want %v actual %v", len(bs), n)
	}
	if !EqualComplexMap(complexMap, v) {
		t.Errorf("unexpected v, want %v actual %v", complexMap, v)
	}
}

func TestGenerated_Interface(t *testing.T) {
	test.Test([]types.Interface{types.Impl1{Num: 1}, types.Impl2("str")}, types.InterfaceMUS, t)
	test.Test([]types.MarshallerInterface{types.Impl1{Num: 1}, types.Impl2("str")}, types.MarshallerInterfaceMUS, t)
}

func TestGenerated_InterfaceRegister(t *testing.T) {
	test.Test([]types.InterfaceRegister{types.Impl3{Num: 1}, types.Impl4("str")}, types.InterfaceRegisterMUS, t)
	test.Test([]types.MarshallerInterfaceRegister{types.Impl5{Num: 1}, types.Impl6("str")}, types.MarshallerInterfaceRegisterMUS, t)
}

func TestGenerated_Versioned(t *testing.T) {
	test.TestVersioned(t, types.VersionedMUS,
		test.Version(types.Ver1(11), types.Ver1TypedMUS, types.Versioned("11")),
		test.Version(types.Ver2("str"), types.Ver2TypedMUS, types.Versioned("str")),
	)
	test.TestVersionedSkip(t, types.VersionedMUS,
		test.VersionSkip[types.Ver1, types.Versioned](types.Ver1(11), types.Ver1TypedMUS),
		test.VersionSkip[types.Ver2, types.Versioned](types.Ver2("str"), types.Ver2TypedMUS),
	)
}

func TestGenerated_VersionedRegister(t *testing.T) {
	test.TestVersioned(t, types.VersionedRegisterMUS,
		test.Version(types.Ver3(11), types.Ver3TypedMUS, types.VersionedRegister("11")),
		test.Version(types.Ver4("str"), types.Ver4TypedMUS, types.VersionedRegister("str")),
	)
	test.TestVersionedSkip(t, types.VersionedRegisterMUS,
		test.VersionSkip[types.Ver3, types.VersionedRegister](types.Ver3(11), types.Ver3TypedMUS),
		test.VersionSkip[types.Ver4, types.VersionedRegister](types.Ver4("str"), types.Ver4TypedMUS),
	)
}

func TestGenerated_Typed(t *testing.T) {
	test.Test([]types.TypedInt{1}, types.TypedIntMUS, t)
}

func TestGenerated_Genopts(t *testing.T) {
	// MultiPackage
	test.Test([]pkg1.Foo{1}, pkg1.FooMUS, t)
	test.Test([]pkg2.Bar{{Foo: 1}}, pkg2.BarMUS, t)
	// CrossPackage
	test.Test([]pkg3.Foo{1}, pkg4.FooMUS, t)
	// CustomPkg
	test.Test([]custompkg.Foo{1}, custompkg.FooMUS, t)
	// ImportAlias
	test.Test([]pkg1sub.Foo{1}, pkg1sub.FooMUS, t)
	test.Test([]pkg2sub.Bar{"str"}, pkg2sub.BarMUS, t)
	test.Test([]importalias.Zoo{{Foo: 1, Bar: "str"}}, importalias.ZooMUS, t)
	// CustomSerName
	test.Test([]customsername.Foo{{Point: image.Point{X: 1, Y: 2}}},
		customsername.CustomSerNameMUS, t)
}

func TestGenerated_Validation(t *testing.T) {
	// ValidInt
	test.TestValidation(types.ValidInt(-1), types.ValidIntMUS, types.ErrNegativeNum, t)

	// ValidStruct
	test.TestValidation(types.ValidStruct{Num: -1, Str: "str"}, types.ValidStructMUS,
		types.ErrNegativeNum, t)
	test.TestValidation(types.ValidStruct{Num: 1, Str: ""}, types.ValidStructMUS,
		types.ErrEmptyStr, t)
	test.TestValidation(types.ValidStruct{Num: 10, Str: "hello"}, types.ValidStructMUS,
		types.ErrNotValidStruct, t)

	// ValidString
	test.TestValidation(types.ValidString(""), types.ValidStringMUS,
		types.ErrZeroLen, t)

	// ValidSlice
	test.TestValidation(types.ValidSlice([]int{}), types.ValidSliceMUS,
		types.ErrZeroLen, t)
	test.TestValidation(types.ValidSlice([]int{-1}), types.ValidSliceMUS,
		types.ErrNegativeNum, t)

	// ValidArray
	test.TestValidation(types.ValidArray([3]int{1, -1, 3}), types.ValidArrayMUS,
		types.ErrNegativeNum, t)

	// ValidMap
	test.TestValidation(types.ValidMap(map[string]int{}), types.ValidMapMUS,
		types.ErrZeroLen, t)
	test.TestValidation(types.ValidMap(map[string]int{"": 1}), types.ValidMapMUS,
		types.ErrEmptyStr, t)
	test.TestValidation(types.ValidMap(map[string]int{"str": -1}), types.ValidMapMUS,
		types.ErrNegativeNum, t)
}

func TestGenerated_Mode(t *testing.T) {
	test.Test([]mode.FullStruct{fullStruct}, safe.FullStructMUS, t)
	test.Test([]mode.FullStruct{fullStruct}, unsafe.FullStructMUS, t)
	test.Test([]mode.FullStruct{fullStruct}, notunsafe.FullStructMUS, t)
}

func TestGenerated_ModeStream(t *testing.T) {
	teststream.Test([]mode.FullStruct{fullStruct}, stream_safe.FullStructMUS, t)
	teststream.Test([]mode.FullStruct{fullStruct}, stream_unsafe.FullStructMUS, t)
	teststream.Test([]mode.FullStruct{fullStruct}, stream_notunsafe.FullStructMUS, t)
}

// -----------------------------------------------------------------------------

var (
	i          = 1
	fullStruct = mode.FullStruct{
		Int:        1,
		Int64:      2,
		Int32:      3,
		Int16:      4,
		Int8:       5,
		Uint:       6,
		Uint64:     7,
		Uint32:     8,
		Uint16:     9,
		Uint8:      10,
		Float64:    11.1,
		Float32:    12.2,
		Byte:       13,
		Bool:       true,
		Str:        "str",
		ByteSlice:  []byte{1, 2, 3},
		Uint8Slice: []uint8{4, 5, 6},
		Time:       time.Now().Truncate(time.Second).UTC(),
		PtrInt:     &i,
		ArrayInt:   [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		SliceInt:   []int{11, 12, 13},
		MapStrInt:  map[string]int{"str": 14},
		Defined:    mode.FullDefined(15),
		Interface:  mode.FullInterfaceImpl("str"),
		Versioned:  mode.Versioned("current version"),
	}
	sl         = []int{1, 2}
	complexMap = types.ComplexMap{
		"str": {
			&sl: {{1.1, 1.2}, {1.3}},
		},
	}
)

func EqualComplexMap(cm1, cm2 types.ComplexMap) bool {
	if len(cm1) != len(cm2) {
		return false
	}
	for k1, v1 := range cm1 {
		v2, ok := cm2[k1]
		if !ok {
			return false
		}
		if len(v1) != len(v2) {
			return false
		}
		for pk1, pv1 := range v1 {
			found := false
			for pk2, pv2 := range v2 {
				if reflect.DeepEqual(*pk1, *pk2) && reflect.DeepEqual(pv1, pv2) {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
