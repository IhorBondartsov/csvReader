package entity

import (
	"errors"
)

type CustomError int

const (
	ErrInvalidNumber CustomError = iota
	ErrDataMalformed
	ErrInvalidID
)

func (s CustomError) Error() error {
	return [...]error{
		errors.New("ErrInvalidNumber"),
		errors.New("ErrDataMalformed"),
		errors.New("ErrInvalidID"),
	}[s]
}
