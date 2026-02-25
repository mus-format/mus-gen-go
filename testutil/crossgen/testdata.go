package testutil

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/musgen-go/testutil/crossgen/pkg"
)

const MyStructDTM com.DTM = iota + 1

type MyMap map[int]pkg.MyInterface

type MyStructWithCrossgen struct {
	Field1 pkg.MyInterface
}
