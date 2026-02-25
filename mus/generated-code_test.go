package musgen

import (
	"testing"
	"time"

	mus_testdata "github.com/mus-format/mus-go/testutil"
	muss_testdata "github.com/mus-format/mus-stream-go/testutil"
	"github.com/mus-format/musgen-go/testutil"
	contr_testdata "github.com/mus-format/musgen-go/testutil/container"
	crossgen_testdata "github.com/mus-format/musgen-go/testutil/crossgen"
	crossgen_pkg "github.com/mus-format/musgen-go/testutil/crossgen/pkg"
	generic_testdata "github.com/mus-format/musgen-go/testutil/generic"
	intr_testdata "github.com/mus-format/musgen-go/testutil/interface"
	intrm_testdata "github.com/mus-format/musgen-go/testutil/interface_marshaller"
	notunsafe_testdata "github.com/mus-format/musgen-go/testutil/notunsafe"
	ptr_testdata "github.com/mus-format/musgen-go/testutil/pointer"
	prim_testdata "github.com/mus-format/musgen-go/testutil/primitive"
	intrr_testdata "github.com/mus-format/musgen-go/testutil/register_interface"
	ser_testdata "github.com/mus-format/musgen-go/testutil/ser"
	ser_pkg "github.com/mus-format/musgen-go/testutil/ser/pkg"
	stream_testdata "github.com/mus-format/musgen-go/testutil/stream"
	stream_notunsafe_testdata "github.com/mus-format/musgen-go/testutil/stream_notunsafe"
	stream_unsafe_testdata "github.com/mus-format/musgen-go/testutil/stream_unsafe"
	struct_testdata "github.com/mus-format/musgen-go/testutil/struct"
	sttime_testdata "github.com/mus-format/musgen-go/testutil/struct_time"
	twopkg_testdata "github.com/mus-format/musgen-go/testutil/two_pkg"
	pkg_testdata "github.com/mus-format/musgen-go/testutil/two_pkg/pkg"
	unsafe_testdata "github.com/mus-format/musgen-go/testutil/unsafe"
)

func TestGeneratedCode(t *testing.T) {
	t.Run("primitive", func(t *testing.T) {
		t.Run("bool", func(t *testing.T) {
			cases := []prim_testdata.MyBool{true}
			mus_testdata.Test(cases, prim_testdata.MyBoolMUS, t)
		})

		t.Run("byte", func(t *testing.T) {
			cases := []prim_testdata.MyByte{3}
			mus_testdata.Test(cases, prim_testdata.MyByteMUS, t)
		})

		t.Run("float32", func(t *testing.T) {
			cases := []prim_testdata.MyFloat32{3.2}
			mus_testdata.Test(cases, prim_testdata.MyFloat32MUS, t)
		})

		t.Run("float64", func(t *testing.T) {
			cases := []prim_testdata.MyFloat64{6.4}
			mus_testdata.Test(cases, prim_testdata.MyFloat64MUS, t)
		})

		t.Run("int", func(t *testing.T) {
			cases := []prim_testdata.AllMyInt{1}
			mus_testdata.Test(cases, prim_testdata.AllMyIntMUS, t)

			testutil.TestUnmarshalError(0, testutil.ErrZeroValue,
				prim_testdata.AllMyIntMUS, t)
		})

		t.Run("string", func(t *testing.T) {
			cases := []prim_testdata.AllMyString{"abc"}
			mus_testdata.Test(cases, prim_testdata.AllMyStringMUS, t)

			testutil.TestUnmarshalError("", testutil.ErrZeroValue,
				prim_testdata.AllMyStringMUS, t)
		})
	})

	t.Run("pointer", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			n := 3
			cases := []ptr_testdata.MyIntPtr{&n}
			mus_testdata.Test(cases, ptr_testdata.MyIntPtrMUS, t)
		})
	})

	t.Run("container", func(t *testing.T) {
		t.Run("array", func(t *testing.T) {
			cases := []contr_testdata.MyArray{{1, 2, 3}}
			mus_testdata.Test(cases, contr_testdata.MyArrayMUS, t)

			testutil.TestUnmarshalError(contr_testdata.AllMyArray{},
				testutil.ErrZeroValue,
				contr_testdata.AllMyArrayMUS, t)
		})

		t.Run("byte_slice", func(t *testing.T) {
			cases := []contr_testdata.MyByteSlice{{1, 2, 3}}
			mus_testdata.Test(cases, contr_testdata.MyByteSliceMUS, t)

			testutil.TestUnmarshalError(contr_testdata.AllMyByteSlice{1},
				testutil.ErrTooLong,
				contr_testdata.AllMyByteSliceMUS, t)
		})

		t.Run("slice", func(t *testing.T) {
			cases := []contr_testdata.MySlice{{1, 2, 3}}
			mus_testdata.Test(cases, contr_testdata.MySliceMUS, t)

			testutil.TestUnmarshalError(contr_testdata.AllMySlice{1},
				testutil.ErrTooLong,
				contr_testdata.AllMySliceMUS, t)
		})

		t.Run("map", func(t *testing.T) {
			cases := []contr_testdata.MyMap{{1: "hello world"}}
			mus_testdata.Test(cases, contr_testdata.MyMapMUS, t)

			testutil.TestUnmarshalError(contr_testdata.AllMyMap{1: "hello world"},
				testutil.ErrTooLong,
				contr_testdata.AllMyMapMUS, t)
		})
	})

	t.Run("struct", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructSer(v, struct_testdata.ComplexStructMUS, t)
	})

	t.Run("interface", func(t *testing.T) {
		cases := []intr_testdata.MyInterface{
			intr_testdata.Impl1{Str: "abs"},
			intr_testdata.Impl2(3),
		}
		mus_testdata.Test(cases, intr_testdata.MyInterfaceMUS, t)
	})

	t.Run("any interface", func(t *testing.T) {
		cases := []intr_testdata.MyAnyInterface{
			intr_testdata.Impl1{Str: "abs"},
			intr_testdata.Impl2(3),
		}
		mus_testdata.Test(cases, intr_testdata.MyAnyInterfaceMUS, t)
	})

	t.Run("interface_marshaller", func(t *testing.T) {
		cases := []intrm_testdata.MyInterface{
			intrm_testdata.Impl1{},
			intrm_testdata.Impl2{},
		}
		mus_testdata.Test(cases, intrm_testdata.MyInterfaceMUS, t)
	})

	t.Run("register interface", func(t *testing.T) {
		cases := []intrr_testdata.MyInterface{
			intrr_testdata.Impl1{Str: "abs"},
			intrr_testdata.Impl2(3),
		}
		mus_testdata.Test(cases, intrr_testdata.MyInterfaceMUS, t)
	})

	t.Run("unsafe", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructSer(v, unsafe_testdata.ComplexStructMUS, t)
	})

	t.Run("notunsafe", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructSer(v, notunsafe_testdata.ComplexStructMUS,
			t)
	})

	t.Run("stream", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructStreamSer(v,
			stream_testdata.ComplexStructMUS, t)

		t.Run("interface_marshaller", func(t *testing.T) {
			cases := []stream_testdata.MyInterface{
				stream_testdata.Impl1{Str: "abc"},
				stream_testdata.Impl2(3),
			}
			muss_testdata.Test(cases, stream_testdata.MyInterfaceMUS, t)
		})
	})

	t.Run("stream_unsafe", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructStreamSer(v,
			stream_unsafe_testdata.ComplexStructMUS, t)
	})

	t.Run("stream_notunsafe", func(t *testing.T) {
		v := struct_testdata.MakeComplexStruct()
		struct_testdata.TestComplexStructStreamSer(v, stream_notunsafe_testdata.ComplexStructMUS, t)
	})

	t.Run("two_pkg", func(t *testing.T) {
		cases := []twopkg_testdata.MySlice{{pkg_testdata.MyInt(1)}}
		mus_testdata.Test(cases, twopkg_testdata.MySliceMUS, t)
	})

	t.Run("generic", func(t *testing.T) {
		// MyInt
		cases1 := []generic_testdata.MySlice[generic_testdata.MyInt]{
			{generic_testdata.MyInt(1)},
		}
		mus_testdata.Test(cases1, generic_testdata.MySliceMUS, t)

		// MyStruct
		cases2 := []generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]{
			{
				T:   generic_testdata.MySlice[generic_testdata.MyInt]{1, 2, 3},
				Int: 1,
			},
		}
		mus_testdata.Test(cases2, generic_testdata.MyStructMUS, t)

		// MyInterface
		cases3 := []generic_testdata.MyInterface[generic_testdata.MyInt]{
			generic_testdata.Impl[generic_testdata.MyInt]{},
		}
		mus_testdata.Test(cases3, generic_testdata.MyInterfaceMUS, t)
	})

	t.Run("ser", func(t *testing.T) {
		t.Run("struct", func(t *testing.T) {
			cases := []ser_testdata.MyStruct{
				{
					MyInt: ser_pkg.MyInt(1),
				},
			}
			mus_testdata.Test(cases, ser_testdata.MyAwesomeStructMUS, t)
		})

		t.Run("interface", func(t *testing.T) {
			cases := []ser_testdata.MyInterface{
				ser_pkg.MyInt(1),
			}
			mus_testdata.Test(cases, ser_testdata.MyAwesomeInterfaceMUS, t)
		})
	})

	t.Run("crossgen", func(t *testing.T) {
		cases := []crossgen_testdata.MyMap{
			{
				1: crossgen_pkg.MyStruct{
					Int:     2,
					MyInt:   crossgen_pkg.MyInt(3),
					MySlice: crossgen_pkg.MySlice{crossgen_pkg.MyInt(4)},
				},
			},
		}
		mus_testdata.Test(cases, crossgen_testdata.MyMapMUS, t)
	})

	t.Run("struct_time", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			cases := []sttime_testdata.MyDefaultTime{
				sttime_testdata.MyDefaultTime(
					time.Unix(int64(time.Now().Second()), 0),
				),
			}
			mus_testdata.Test(cases, sttime_testdata.MyDefaultTimeMUS, t)
		})

		t.Run("micro", func(t *testing.T) {
			cases := []sttime_testdata.MyMicroTime{
				sttime_testdata.MyMicroTime(
					time.UnixMicro(time.Now().UnixMicro()),
				),
			}
			mus_testdata.Test(cases, sttime_testdata.MyMicroTimeMUS, t)
		})
	})
}
