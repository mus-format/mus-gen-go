package types

import "errors"

var (
	ErrZeroLen        = errors.New("len must be non-zero")
	ErrNegativeNum    = errors.New("num must be non-negative")
	ErrEmptyStr       = errors.New("str must not be empty")
	ErrNotValidStruct = errors.New("struct is not valid")
)

func ValidateInt(num ValidInt) error {
	if num < 0 {
		return ErrNegativeNum
	}
	return nil
}

func ValidateLen(l int) error {
	if l == 0 {
		return ErrZeroLen
	}
	return nil
}

func ValidateNum(num int) error {
	if num < 0 {
		return ErrNegativeNum
	}
	return nil
}

func ValidateStr(str string) error {
	if str == "" {
		return ErrEmptyStr
	}
	return nil
}

func ValidateStruct(s ValidStruct) error {
	if s.Num == 10 && s.Str == "hello" {
		return ErrNotValidStruct
	}
	return nil
}

func ValidateComplexMapValue(map[*[]int][][]float32) error {
	return nil
}
