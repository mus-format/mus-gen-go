package musgen

import (
	genopts "github.com/mus-format/mus-gen-go/options/gen"
	tpopts "github.com/mus-format/mus-gen-go/options/type"
)

const (
	VarintPkg   = "varint"
	RawPkg      = "raw"
	OrdPkg      = "ord"
	UnsafePkg   = "unsafe"
	DefTimeSer  = "TimeUnixUTC"
	TimeSerRule = "time_ser"
)

func safeModeRules(tops tpopts.Options) map[string]string {
	var (
		m       = make(map[string]string)
		numPkg  = VarintPkg
		timeSer = DefTimeSer
	)
	if tops.NumEncoding != tpopts.NumEncodingUndefined {
		numPkg = tops.NumEncoding.Package()
	}
	if tops.TimeUnit != tpopts.TimeUnitUndefined {
		timeSer = tops.TimeUnit.Ser()
	}
	m["int"] = numPkg
	m["int64"] = numPkg
	m["int32"] = numPkg
	m["int16"] = numPkg
	m["uint"] = numPkg
	m["uint64"] = numPkg
	m["uint32"] = numPkg
	m["uint16"] = numPkg

	m["int8"] = RawPkg
	m["uint8"] = RawPkg
	m["float64"] = RawPkg
	m["float32"] = RawPkg
	m["byte"] = RawPkg
	m["time.Time"] = RawPkg

	m["bool"] = OrdPkg
	m["string"] = OrdPkg
	m["[]byte"] = OrdPkg
	m["[]uint8"] = OrdPkg

	m[TimeSerRule] = timeSer
	return m
}

func unsafeModeRules(tops tpopts.Options, gops genopts.Options) map[string]string {
	var (
		m            = make(map[string]string)
		numPkg       = UnsafePkg
		byteSlicePkg = UnsafePkg
		timeSer      = DefTimeSer
	)
	if tops.NumEncoding != tpopts.NumEncodingUndefined {
		numPkg = tops.NumEncoding.Package()
	}
	if tops.TimeUnit != tpopts.TimeUnitUndefined {
		timeSer = tops.TimeUnit.Ser()
	}
	if gops.Stream {
		byteSlicePkg = OrdPkg
	}
	m["int"] = numPkg
	m["int64"] = numPkg
	m["int32"] = numPkg
	m["int16"] = numPkg
	m["uint"] = numPkg
	m["uint64"] = numPkg
	m["uint32"] = numPkg
	m["uint16"] = numPkg

	m["int8"] = UnsafePkg
	m["uint8"] = UnsafePkg
	m["float64"] = UnsafePkg
	m["float32"] = UnsafePkg
	m["time.Time"] = UnsafePkg
	m["byte"] = UnsafePkg
	m["bool"] = UnsafePkg
	m["string"] = UnsafePkg

	m["[]byte"] = byteSlicePkg
	m["[]uint8"] = byteSlicePkg

	m[TimeSerRule] = timeSer
	return m
}

func notUnsafeModeRules(tops tpopts.Options, gops genopts.Options) map[string]string {
	m := unsafeModeRules(tops, gops)
	m["string"] = OrdPkg
	return m
}
