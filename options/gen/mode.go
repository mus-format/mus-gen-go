package genopts

type Mode int

const (
	ModeSafe Mode = iota
	ModeUnsafe
	ModeNotUnsafe
)
