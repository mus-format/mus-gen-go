package scanner

import (
	"testing"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

type ScanTestCase[T scanner.QualifiedName] struct {
	Name   string
	Setup  ScanSetup[T]
	Params ScanParams[T]
	Want   ScanWant
}

type ScanSetup[T scanner.QualifiedName] struct {
	Config scanner.Config
	Name   T
	Op     scanner.Op[T]
	Tops   tpopts.Options
}

type ScanParams[T scanner.QualifiedName] struct {
	Type T
}

type ScanWant struct {
	Err   error
	Mocks []*mok.Mock
}

func RunScanTestCase[T scanner.QualifiedName](t *testing.T, tc ScanTestCase[T]) {
	err := scanner.Scan(tc.Setup.Config, tc.Setup.Name, tc.Setup.Op, tc.Setup.Tops)
	asserterror.EqualError(t, err, tc.Want.Err)
	asserterror.EqualDeep(t, mok.CheckCalls(tc.Want.Mocks), mok.EmptyInfomap)
}
