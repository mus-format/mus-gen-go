package testutil

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	com "github.com/mus-format/common-go"
)

const MyIntDTM com.DTM = iota + 1

type MyStruct struct {
	Int int
	Str string
}

type ComplexStruct struct {
	Bool bool
	Byte byte

	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Float32 float32
	Float64 float64

	String string

	ByteSlice   []byte
	StructSlice []MyStruct

	Array [3]int

	PtrString *string
	PtrStruct *MyStruct
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[MyInt][]MyStruct

	Time time.Time

	Alias MySlice

	Interface MyInterface
}

type MyInt int

func (i MyInt) Print() {
	fmt.Println("MyInt")
}

type MySlice []string

type MyInterface interface{ Print() }

type DoubleDefinedMyStruct MyStruct

func MakeComplexStruct() ComplexStruct {
	var (
		str      = "hello world"
		myStruct = MyStruct{Int: 1}
		arr      = [3]int{1, 2, 3}
		m        = map[float32]map[MyInt][]MyStruct{
			32.0: {
				MyInt(1): []MyStruct{{Int: 1}},
			},
		}
	)
	return ComplexStruct{
		Bool: true,
		Byte: 3,

		Int8:  8,
		Int16: 16,
		Int32: 32,
		Int64: 64,

		Uint8:  8,
		Uint16: 16,
		Uint32: 32,
		Uint64: 64,

		Float32: 32.0,
		Float64: 64.0,

		String:      "hello world",
		ByteSlice:   []byte{1, 2, 3},
		StructSlice: []MyStruct{{Int: 1}, {Int: 2}, {Int: 3}},

		Array: [3]int{1, 2, 3},

		PtrString: &str,
		PtrStruct: &myStruct,
		NilPtr:    nil,
		PtrArray:  &arr,

		Map: m,

		Time: time.Unix(int64(time.Now().Second()), 0),

		Alias: MySlice{},

		Interface: MyInt(3),
	}
}

func EqualComplexStruct(s, v ComplexStruct) bool {
	opt := cmp.FilterValues(func(x1, x2 any) bool {
		var (
			t1 = reflect.TypeOf(x1)
			t2 = reflect.TypeOf(x2)
		)
		if t1.Kind() == reflect.Pointer && t2.Kind() == reflect.Pointer {
			if t1.Elem().Kind() == reflect.String && t2.Elem().Kind() == reflect.String {
				return true
			}
			if t1.Elem().Kind() == reflect.Struct && t2.Elem().Kind() == reflect.Struct {
				return true
			}
			if t1.Elem().Kind() == reflect.Array && t2.Elem().Kind() == reflect.Array {
				return true
			}
		}
		return false
	}, cmp.Comparer(func(i1, i2 any) bool {
		str1, ok1 := i1.(*string)
		str2, ok2 := i2.(*string)
		if ok1 && ok2 {
			if str1 != nil && str2 != nil {
				return *str1 == *str2
			}
			if str1 == nil && str2 == nil {
				return true
			}
			return false
		}
		sct1, ok1 := i1.(*MyStruct)
		sct2, ok2 := i1.(*MyStruct)
		if ok1 && ok2 {
			if sct1 != nil && sct2 != nil {
				return reflect.DeepEqual(*sct1, *sct2)
			}
			if sct1 == nil && sct2 == nil {
				return true
			}
			return false
		}
		arr1, ok1 := i1.(*[3]int)
		arr2, ok2 := i1.(*[3]int)
		if ok1 && ok2 {
			if arr1 != nil && arr2 != nil {
				return reflect.DeepEqual(arr1, arr2)
			}
			if arr1 == nil && arr2 == nil {
				return true
			}
			return false
		}
		return true
	}))
	// fmt.Println(cmp.Diff(s, v, opt))
	return cmp.Equal(s, v, opt)
}
