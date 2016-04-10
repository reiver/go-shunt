package shuntdbsqldriver


import (
	"database/sql/driver"
)


type internalTx struct {
	conn *internalConn
}


func newTx(conn *internalConn) (driver.Tx, error) {
	if nil == conn {
		return nil, errNilConn
	}

	tx := internalTx{
		conn:conn,
	}

	return &tx, nil
}


func (tx *internalTx) Commit() error {
	if nil == tx {
		return errNilReceiver
	}

	return errNotImplementedCommit
}


func (tx *internalTx) Rollback() error {
	if nil == tx {
		return errNilReceiver
	}

	return errNotImplementedRollback
}
