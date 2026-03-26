package types

type Impl3 struct {
	Num int
}

func (i Impl3) MarshalTypedMUS(bs []byte) (n int) {
	return Impl3TypedMUS.Marshal(i, bs)
}

func (i Impl3) SizeTypedMUS() (size int) {
	return Impl3TypedMUS.Size(i)
}

type Impl4 string

func (i Impl4) MarshalTypedMUS(bs []byte) (n int) {
	return Impl4TypedMUS.Marshal(i, bs)
}

func (i Impl4) SizeTypedMUS() (size int) {
	return Impl4TypedMUS.Size(i)
}

type Impl5 struct {
	Num int
}

func (i Impl5) MarshalTypedMUS(bs []byte) (n int) {
	return Impl5TypedMUS.Marshal(i, bs)
}

func (i Impl5) SizeTypedMUS() (size int) {
	return Impl5TypedMUS.Size(i)
}

type Impl6 string

func (i Impl6) MarshalTypedMUS(bs []byte) (n int) {
	return Impl6TypedMUS.Marshal(i, bs)
}

func (i Impl6) SizeTypedMUS() (size int) {
	return Impl6TypedMUS.Size(i)
}

type InterfaceRegister any

type MarshallerInterfaceRegister any
