package shuntdbsqldriver


import (
	"database/sql"

	"testing"
)


func TestDatabaseSqlDbPrepare(t *testing.T) {

	randomString := generateKey() // We just want a random string here.

	iterator := internalEmptyIterator{ Cols: []string{"does","not","matter"} }

	if err := insertIterator(randomString, iterator); nil != err {
		t.Errorf("After called insertIterator(%q, iterator), did not expect an error but got one: %v", randomString, err)
		return
	}



	db, err := sql.Open(Name, randomString)
	if nil != err {
		t.Errorf("After called sql.Open(%q, %q), did not expect an error but got one: %v", Name, randomString, err)
		return
	}
	if nil == db {
		t.Errorf("After called sql.Open(%q, %q), expected db to not be nil, but was: %v", Name, randomString, db)
		return
	}

	// Need to call Ping() to make it so the driver's Open() method actually gets called.
	if err := db.Ping(); nil != err {
		t.Errorf("After called sql.DB.Ping(), did not expect an error, but got one: %v", err)
		return
	}



	deleteIterator(randomString)



	query := ``
	stmt, err := db.Prepare(query)
	if nil != err {
		t.Errorf("After called sql.DB.Prepare(%q), did not expect an error, but got one: %v", query, err)
		return
	}
	if nil == stmt {
		t.Errorf("After called sql.DB.Prepare(%q), expected stmt to not be nil, but was: %v", query, stmt)
		return
	}
}
