package shunt


import (
	"errors"
)


var (
	errNilDB   = errors.New("Internal Error: Nil DB")
	errNilRows = errors.New("Internal Error: Nil Rows")
)
