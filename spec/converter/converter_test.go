package converter_test

import (
	"testing"

	"github.com/mus-format/mus-gen-go/test/converter"
)

func TestTypeNameConverter_ConvertToFullName(t *testing.T) {
	for _, tc := range []converter.ConvertFullTestCase{
		converter.BuildConvertFullPrimTestCase(),
		converter.BuildConvertFullSamePackageTestCase(),
		converter.BuildConvertFullOtherPackageTestCase(),
		converter.BuildConvertFullAliasedPackageTestCase(),
		converter.BuildConvertFullSliceTestCase(),
		converter.BuildConvertFullMapTestCase(),
		converter.BuildConvertFullTypeParamTestCase(),
		converter.BuildConvertFullComplexTestCase(),
	} {
		converter.RunConvertFullTest(t, tc)
	}
}

func TestTypeNameConverter_ConvertToRelativeName(t *testing.T) {
	for _, tc := range []converter.ConvertRelTestCase{
		converter.BuildConvertRelSamePackageTestCase(),
		converter.BuildConvertRelOtherPackageTestCase(),
		converter.BuildConvertRelPrimTestCase(),
		converter.BuildConvertRelComplexTestCase(),
	} {
		converter.RunConvertRelTest(t, tc)
	}
}
