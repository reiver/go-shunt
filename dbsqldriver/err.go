package shuntdbsqldriver


import (
	"errors"
)


var (
	errAlreadyExists  = errors.New("Already Exists")
	errNilConn        = errors.New("Nil Conn")
	errNilIterator    = errors.New("Nil Iterator")
	errNilReceiver    = errors.New("Nil Receiver")
	errNilStmt        = errors.New("Nil Stmt")
	errNotFound       = errors.New("Not Found")
	errTooManyTries   = errors.New("Internal Error")

	errNotImplementedCommit   = errors.New("driver.Tx.Commit Not Implemented")
	errNotImplementedRollback = errors.New("driver.Tx.Rollback Not Implemented")
)
