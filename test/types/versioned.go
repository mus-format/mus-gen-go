package types

import (
	"strconv"

	com "github.com/mus-format/common-go"
)

const (
	Ver1DTM com.DTM = iota + 1
	Ver2DTM
)

type Versioned Ver2

type Ver2 string

type Ver1 int

func MigrateVer1(v Ver1) Versioned {
	return Versioned(strconv.Itoa(int(v)))
}
