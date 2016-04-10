package shuntdbsqldriver


import (
	"database/sql/driver"
)


type internalStmt struct{
	conn     *internalConn
	query     string
}


func newStmt(conn *internalConn, query string) (driver.Stmt, error) {

	if nil == conn {
		return nil, errNilConn
	}

	stmt := internalStmt{
		conn:conn,
		query:query,
	}

	return &stmt, nil
}


func (stmt *internalStmt) Close() error {

	return nil
}


func (stmt *internalStmt) NumInput() int {

	return 0
}


func (stmt *internalStmt) Exec(args []driver.Value) (driver.Result, error) {
//@TODO: Is driver.ErrSkip the correct error to return?
	return nil, driver.ErrSkip
}


func (stmt *internalStmt) Query(args []driver.Value) (driver.Rows, error) {
	if nil == stmt {
		return nil, errNilReceiver
	}

	return newRows(stmt, args)
}
