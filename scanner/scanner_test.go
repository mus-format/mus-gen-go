package scanner_test

import (
	"testing"

	testscanner "github.com/mus-format/mus-gen-go/test/scanner"
	"github.com/mus-format/mus-gen-go/typename"
)

func TestScan(t *testing.T) {
	for _, tc := range []testscanner.ScanTestCase[typename.CompleteName]{
		testscanner.ScanStructTestCase(t),
		testscanner.ScanPtrStructTestCase(t),
		testscanner.ScanUint8SliceTestCase(t),
		testscanner.ScanParametrizedStructTestCase(t),
		testscanner.ScanArrayTestCase(t),
		testscanner.ScanPtrArrayTestCase(t),
		testscanner.ScanSliceTestCase(t),
		testscanner.ScanPtrSliceTestCase(t),
		testscanner.ScanMapTestCase(t),
		testscanner.ScanPtrMapTestCase(t),
		testscanner.ScanMapInMapTestCase(t),
		testscanner.ScanPrimitiveTestCase(t),
		testscanner.ScanPtrPrimitiveTestCase(t),
		testscanner.ScanWithoutParamsTestCase(t),
		testscanner.ScanComplexTestCase(t),
	} {
		testscanner.RunScanTestCase(t, tc)
	}
}
