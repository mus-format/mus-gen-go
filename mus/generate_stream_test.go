package musgen

import (
	"testing"

	"github.com/mus-format/mus-gen-go/test/types/mode"
	safe "github.com/mus-format/mus-gen-go/test/types/mode/stream_safe"
	"github.com/mus-format/mus-stream-go/test"
)

func TestGeneratedStream_Versioned(t *testing.T) {
	test.TestVersioned(t, safe.VersionedMUS,
		test.Version(mode.FooV1(mode.FooV1{Num: 11}), safe.FooV1TypedMUS, mode.Versioned("11")),
		test.Version(mode.FooV2("str"), safe.FooV2TypedMUS, mode.Versioned("str")),
	)
	test.TestVersionedSkip(t, safe.VersionedMUS,
		test.VersionSkip[mode.FooV1, mode.Versioned](mode.FooV1(mode.FooV1{Num: 11}), safe.FooV1TypedMUS),
		test.VersionSkip[mode.FooV2, mode.Versioned](mode.FooV2("str"), safe.FooV2TypedMUS),
	)
}
