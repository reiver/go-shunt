package shunt


import (
	"github.com/reiver/go-shunt/dbsqldriver"


	"database/sql"
)


// Rows conceptually "converts" an Iterator to a *sql.Rows.
//
// Iterator and *sql.Rows have the exact same interface.
//
// This is useful for creating "middleware" for "database/sql", in situations where
// there is an expectation that *sql.Rows is used (and not an interface that fits
// *sql.Rows).
//
// This can also be useful for composition of "middleware", in situations where
// there is an expectation that *sql.Rows is used (and not an interface that fits
// *sql.Rows).
func Rows(iterator Iterator) (*sql.Rows, error) {

	shunted, err := shuntdbsqldriver.Shunt(iterator)
	if nil != err {
		return nil, err
	}
	defer shunted.Close()


	db, err := sql.Open(shuntdbsqldriver.Name, shunted.DataSourceName())
	if nil != err {
		return nil, err
	}
	if nil == db {
		return nil, errNilDB
	}
	defer db.Close()


	if err := db.Ping(); nil != err {
		return nil, err
	}


	rows, err := db.Query(``)
	if nil != err {
		return nil, err
	}
	if nil == rows {
		return nil, errNilRows
	}


	return rows, nil
}


// MustRows is like the Rows func, but does not return an error;
// instead it panic()s if there is an error.
//
// For example:
//
//      rows := refs.MustRows(iterator) // This could panic()!
func MustRows(iterator Iterator) *sql.Rows {
	rows, err := Rows(iterator)
	if nil != err {
		panic(err)
	}

	return rows
}
