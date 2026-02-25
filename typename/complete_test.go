package typename

import (
	"reflect"
	"testing"

	container_testdata "github.com/mus-format/musgen-go/testutil/container"
	ptr_testdata "github.com/mus-format/musgen-go/testutil/pointer"
	prim_testdata "github.com/mus-format/musgen-go/testutil/primitive"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestTypeCompleteName(t *testing.T) {

	t.Run("primitive", func(t *testing.T) {
		testCases := []struct {
			t                reflect.Type
			wantCompleteName CompleteName
			wantErr          error
		}{
			{
				t:                reflect.TypeFor[int](),
				wantCompleteName: "int",
			},
			{
				t:                reflect.TypeFor[*int](),
				wantCompleteName: "*int",
			},
			{
				t:                reflect.TypeFor[[3]int](),
				wantCompleteName: "[3]int",
			},
			{
				t:                reflect.TypeFor[*[3]int](),
				wantCompleteName: "*[3]int",
			},
			{
				t:                reflect.TypeFor[[3]*int](),
				wantCompleteName: "[3]*int",
			},
			{
				t:                reflect.TypeFor[[]int](),
				wantCompleteName: "[]int",
			},
			{
				t:                reflect.TypeFor[*[]int](),
				wantCompleteName: "*[]int",
			},
			{
				t:                reflect.TypeFor[[]*int](),
				wantCompleteName: "[]*int",
			},
			{
				t:                reflect.TypeFor[map[int]string](),
				wantCompleteName: "map[int]string",
			},
			{
				t:                reflect.TypeFor[*map[int]string](),
				wantCompleteName: "*map[int]string",
			},
			{
				t:                reflect.TypeFor[map[*int]string](),
				wantCompleteName: "map[*int]string",
			},
			{
				t:                reflect.TypeFor[map[int]*string](),
				wantCompleteName: "map[int]*string",
			},
			{
				t:                reflect.TypeFor[struct_testdata.MyStruct](),
				wantCompleteName: "github.com/mus-format/musgen-go/testutil/struct/testutil.MyStruct",
			},
			{
				t:                reflect.TypeFor[*struct_testdata.MyStruct](),
				wantCompleteName: "*github.com/mus-format/musgen-go/testutil/struct/testutil.MyStruct",
			},
			{
				t:       reflect.TypeFor[**int](),
				wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
			},
			{
				t:       reflect.TypeFor[[3]**int](),
				wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
			},
			{
				t:       reflect.TypeFor[[]**int](),
				wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
			},
			{
				t:       reflect.TypeFor[map[**int]string](),
				wantErr: NewMultiPointerError(reflect.TypeFor[**int]()),
			},
			{
				t:       reflect.TypeFor[map[int]**string](),
				wantErr: NewMultiPointerError(reflect.TypeFor[**string]()),
			},
			{
				t:       reflect.TypeFor[struct{}](),
				wantErr: NewUnsupportedTypeError(reflect.TypeFor[struct{}]()),
			},
		}
		for _, c := range testCases {
			cname, err := TypeCompleteName(c.t)
			asserterror.EqualError(t, err, c.wantErr)
			asserterror.Equal(t, cname, c.wantCompleteName)
		}
	})

	t.Run("SourceTypeCompleteName", func(t *testing.T) {
		testCases := []struct {
			t                reflect.Type
			wantCompleteName CompleteName
			wantErr          error
		}{
			{
				t:                reflect.TypeFor[ptr_testdata.MyIntPtr](),
				wantCompleteName: "*int",
			},
			{
				t:                reflect.TypeFor[prim_testdata.MyInt](),
				wantCompleteName: "int",
			},
			{
				t:                reflect.TypeFor[container_testdata.MyArray](),
				wantCompleteName: "[3]int",
			},
			{
				t:                reflect.TypeFor[ptr_testdata.MyArrayPtr](),
				wantCompleteName: "*[3]int",
			},
			{
				t:                reflect.TypeFor[container_testdata.MySlice](),
				wantCompleteName: "[]int",
			},
			{
				t:                reflect.TypeFor[ptr_testdata.MySlicePtr](),
				wantCompleteName: "*[]int",
			},
			{
				t:                reflect.TypeFor[container_testdata.MyMap](),
				wantCompleteName: "map[int]string",
			},
			{
				t:                reflect.TypeFor[ptr_testdata.MyMapPtr](),
				wantCompleteName: "*map[int]string",
			},
			{
				t:       reflect.TypeFor[ptr_testdata.MyDoubleIntPtr](),
				wantErr: NewMultiPointerError(reflect.TypeFor[ptr_testdata.MyDoubleIntPtr]()),
			},
			{
				t:       reflect.TypeFor[int](),
				wantErr: ErrTypeMismatch,
			},
			{
				t:       reflect.TypeFor[*int](),
				wantErr: ErrTypeMismatch,
			},
			{
				t:       reflect.TypeFor[struct_testdata.MyStruct](),
				wantErr: NewUnsupportedTypeError(reflect.TypeFor[struct_testdata.MyStruct]()),
			},
		}
		for _, c := range testCases {
			cname, err := SourceTypeCompleteName(c.t)
			asserterror.EqualError(t, err, c.wantErr)
			asserterror.Equal(t, cname, c.wantCompleteName)
		}
	})

}
