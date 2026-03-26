package musgen

import (
	"reflect"

	// introps "github.com/mus-format/musgen-go/options/interface"

	tpopts "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/spec"
	"github.com/mus-format/mus-gen-go/typename"
)

type TypeBuilder interface {
	BuildDefinedType(t reflect.Type, tops tpopts.Options) (spec.DefinedType, error)
	BuildAnonType(name typename.FullName, topts tpopts.Options) (
		spec.AnonType, bool, error)
	// BuildStructData(t reflect.Type, sops structopts.Options) (d data.TypeData,
	// 	err error)
	// BuildInterfaceData(t reflect.Type, iops introps.Options) (d data.TypeData,
	// 	err error)
	// BuildDTSData(t reflect.Type) (d data.TypeData, err error)
	// BuildTimeData(t reflect.Type, tops *tpopts.Options) (d data.TypeData,
	// 	err error)
	// ToFullName(t reflect.Type) typename.FullName
	// ToRelName(fname typename.FullName) typename.RelName
}
