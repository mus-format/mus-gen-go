package builder

import (
	"errors"

	tpops "github.com/mus-format/mus-gen-go/options/type"
	"github.com/mus-format/mus-gen-go/scanner"
	"github.com/mus-format/mus-gen-go/spec"
	"github.com/mus-format/mus-gen-go/typename"
)

var ErrStop = errors.New("stop")

func NewMakeAnonOp(converter TypeNameConverter) *MakeAnonDataOp {
	m := map[spec.AnonSerName]spec.AnonType{}
	return &MakeAnonDataOp{CollectAnonOp: NewCollectAnonOp(m, converter)}
}

type MakeAnonDataOp struct {
	*CollectAnonOp
	anonType spec.AnonType
	ok       bool
}

func (o *MakeAnonDataOp) ProcessType(t scanner.TypeInfo[typename.FullName],
	tops tpops.Options) (err error) {
	if err = o.CollectAnonOp.ProcessType(t, tops); err != nil {
		return
	}
	if anonType := o.CollectAnonOp.FirstAnonType(); anonType != nil {
		o.anonType = *anonType
		o.ok = true
	}
	return ErrStop
}

func (o *MakeAnonDataOp) AnonType() (anonType spec.AnonType, ok bool) {
	return o.anonType, o.ok
}
