package shuntdbsqldriver


import (
	"database/sql/driver"
)


type internalConn struct {
	name     string
	iterator Iterator
}


func newConn(name string) (driver.Conn, error) {

	iterator := getIterator(name)
	if nil == iterator {
		return nil, errNotFound
	}

	conn := internalConn{
		name:name,
		iterator:iterator,
	}

	return &conn, nil
}


func (conn *internalConn) Close() error {
	return nil
}


func (conn *internalConn) Begin() (driver.Tx, error) {
	if nil == conn {
		return nil, errNilReceiver
	}

	return newTx(conn)
}


func (conn *internalConn) Prepare(query string) (driver.Stmt, error) {
	if nil == conn {
		return nil, errNilReceiver
	}

	return newStmt(conn, query)
}
