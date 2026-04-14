package types

import "strconv"

type VersionedRegister Ver4

type Ver4 string

type Ver3 int

func MigrateVer3(v Ver3) VersionedRegister {
	return VersionedRegister(strconv.Itoa(int(v)))
}
