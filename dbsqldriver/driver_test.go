package shuntdbsqldriver


import (
	"database/sql"

	"testing"
)


func TestDatabaseSqlOpenFail(t *testing.T) {

	tests := []struct{
		DataSourceName string
	}{
		{
			DataSourceName: "",
		},



		{
			DataSourceName: "apple",
		},
		{
			DataSourceName: "banana",
		},
		{
			DataSourceName: "cherry",
		},
	}

	const numMoreTests = 50
	for i:=0; i<numMoreTests; i++ {
		randomString := generateKey() // We just want a random string here.

		test := struct{
			DataSourceName string
		}{
			DataSourceName: randomString,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {
		db, err := sql.Open(Name, test.DataSourceName)
		if nil != err {
			t.Errorf("For test #%d, after called sql.Open(%q, %q), did not expect an error but got one: %v", testNumber, Name, test.DataSourceName, err)
			return
		}
		if nil == db {
			t.Errorf("For test #%d, after called sql.Open(%q, %q), expected db to not be nil but was: %v", testNumber, Name, test.DataSourceName, db)
			return
		}

		// Need to call Ping() to make it so the driver's Open() method actually gets called.
		if err := db.Ping(); nil == err {
			t.Errorf("For test #%d, after called sql.DB.Ping(), expected an error but did not get one: %v", testNumber, err)
			return
		}
	}
}


func TestDatabaseSqlOpen(t *testing.T) {


	const numTests = 50
	for testNumber := 0; testNumber < numTests; testNumber++ {

		randomString := generateKey() // We just want a random string here.

		iterator := internalEmptyIterator{ Cols: []string{"does","not","matter"} }

		if err := insertIterator(randomString, iterator); nil != err {
			t.Errorf("For test #%d, after called insertIterator(%q, iterator), did not expect an error but got one: %v", testNumber, randomString, err)
			return
		}


		db, err := sql.Open(Name, randomString)
		if nil != err {
			t.Errorf("For test #%d, after called sql.Open(%q, %q), did not expect an error but got one: %v", testNumber, Name, randomString, err)
			return
		}
		if nil == db {
			t.Errorf("For test #%d, after called sql.Open(%q, %q), expected db to not be nil but was: %v", testNumber, Name, randomString, db)
			return
		}


		// Need to call Ping() to make it so the driver's Open() method actually gets called.
		if err := db.Ping(); nil != err {
			t.Errorf("After called sql.DB.Ping(), did not expect an error but actually got one: %v", err)
			return
		}


		deleteIterator(randomString)
	}
}
