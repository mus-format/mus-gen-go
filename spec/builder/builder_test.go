package builder_test

import (
	"testing"

	"github.com/mus-format/mus-gen-go/test/builder"
)

func TestTypeBuilder_BuildDefinedType(t *testing.T) {
	for _, tc := range []builder.BuildDefinedTypeTestCase{
		builder.BuildDefinedPrimitiveTestCase(),
		builder.BuildDefinedSliceTestCase(),
		builder.BuildDefinedNotTestCase(),
		builder.BuildDefinedForStructTestCase(),
		builder.BuildDefinedForInterfaceTestCase(),
		builder.BuildDefinedForUnsupportedTypeTestCase(),
	} {
		builder.RunBuildDefinedTypeTest(t, tc)
	}
}
func TestTypeBuilder_BuildStructType(t *testing.T) {
	for _, tc := range []builder.BuildStructTypeTestCase{
		builder.BuildWrongFieldsCountTestCase(),
		builder.BuildEmptyStructTestCase(),
		builder.BuildStructWithFieldsTestCase(),
		builder.BuildStructForDefinedTypeTestCase(),
		builder.BuildStructForInterfaceTypeTestCase(),
		builder.BuildStructForUnsupportedTypeTestCase(),
	} {
		builder.RunBuildStructTypeTest(t, tc)
	}
}
func TestTypeBuilder_BuildTyped(t *testing.T) {
	for _, tc := range []builder.BuildTypedTestCase{
		builder.BuildTypedDefinedTestCase(),
		builder.BuildTypedStructTestCase(),
		builder.BuildTypedInterfaceTestCase(),
		builder.BuildTypedNotSupportedTestCase(),
	} {
		builder.RunBuildTypedTest(t, tc)
	}
}

func TestTypeBuilder_BuildInterfaceType(t *testing.T) {
	for _, tc := range []builder.BuildInterfaceTypeTestCase{
		builder.BuildInterfaceTestCase(),
		builder.BuildInterfaceForDefinedTypeTestCase(),
		builder.BuildInterfaceForStructTestCase(),
		builder.BuildInterfaceForUnsupportedTypeTestCase(),
	} {
		builder.RunBuildInterfaceTypeTest(t, tc)
	}
}

func TestTypeBuilder_BuildTimeType(t *testing.T) {
	for _, tc := range []builder.BuildTimeTypeTestCase{
		builder.BuildTimeTestCase(),
		builder.BuildTimeForNotStructTestCase(),
	} {
		builder.RunBuildTimeTypeTest(t, tc)
	}
}
